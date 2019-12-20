import os
import sys
import json
import imp
import resource
import base64
import psutil
import signal
import time

import constant
from log import init_logging, get_self_std_out, get_self_std_err, init_fc_logger, get_fc_logger
from context import FCContext, FunctionMeta
from utils import make_json, make_error_handler, load_handler_error_handler, get_error_message

global_request_id = None
global_handlers = dict()
return_dict = dict()


class TimeoutError(AssertionError):
    """Thrown when a timeout occurs in the `timeout` context manager."""
    def __init__(self, value="Timed Out"):
        self.value = value

    def __str__(self):
        return repr(self.value)


def set_request_id(request_id):
    global global_request_id
    global_request_id = request_id


def get_request_id():
    global global_request_id
    return global_request_id


def send_error(err_code, start_response):
    header_code = str(err_code) + " NOK"
    start_response(header_code, [('Content-type', 'application/json')])
    return


def send_success(start_response):
    start_response('200 OK', [('Content-type', 'application/json')])
    return


def set_fc_env(fc_env_raw):
    """
    set function env
    :param fc_env_raw:
    :return:
    """
    fc_env = json.loads(fc_env_raw)
    for k in fc_env:
        os.environ[k] = fc_env[k]


def gen_fc_context(environ):
    """
    generate function context
    :param environ:
    :return:
    """
    fc_id = environ[constant.HTTP_FUNCTION_ID]
    name = environ[constant.HTTP_FUNCTION_NAME]
    version = environ[constant.HTTP_FUNCTION_VERSION]
    handler = environ[constant.HTTP_FUNCTION_HANDLER]
    memory = int(environ[constant.HTTP_FUNCTION_MEMORY])
    timeout = int(environ[constant.HTTP_FUNCTION_TIMEOUT])
    meta = FunctionMeta(fc_id, name, version, handler, memory, timeout)

    # log_set = environ[constant.HTTP_CONTEXT_LOG_SET]
    # log_topic = environ[constant.HTTP_CONTEXT_LOG_TOPIC]
    log_set = None
    log_topic = None
    ctx = FCContext(get_request_id(), meta, log_set, log_topic)

    return ctx


def gen_return_dict(mem_info, msg, time_usage):
    return_dict[constant.STD_OUT_RETURN] = get_self_std_out().buff[constant.MAX_LENGTH:]
    return_dict[constant.STD_ERR_RETURN] = get_self_std_err().buff[constant.MAX_LENGTH:]
    return_dict[constant.FUNCTION_RETURN] = msg
    return_dict[constant.MEM_USAGE] = round(float(resource.getrusage(resource.RUSAGE_SELF).ru_maxrss) / 1024 - float(
        mem_info.rss) / 1024 / 1024, 2)
    return_dict[constant.TIME_USAGE] = time_usage
    return make_json(return_dict)


def parse_params(environ):
    """
    parse input parameters
    :param environ:
    :return:
    """
    try:
        length = int(environ[constant.HTTP_CONTENT_LENGTH], 0)
    except ValueError:
        length = 0
    evt = environ[constant.HTTP_WSGI_INPUT].read(length) if length else b""

    ctx = gen_fc_context(environ)
    handler = environ[constant.HTTP_FUNCTION_HANDLER]

    return evt, ctx, handler


def load_handler(handler, md5, pathname):
    global global_handlers
    if md5 in global_handlers:
        return True, global_handlers[md5]

    global_handlers.clear()

    try:
        (modname, func_name) = handler.rsplit(".", 1)
    except ValueError as e:
        return False, load_handler_error_handler(e, handler)

    file_handle, desc = None, None
    try:
        if os.path.islink(pathname):
            pathname = os.readlink(pathname)
        pos = modname.rfind("/")
        if pos != -1:
            path_suffix, segments = modname[:pos], modname[pos+1:]
            pathname = os.path.join(pathname, path_suffix)
            if segments:
                segments = segments.split(".")
        else:
            segments = modname.split(".")
        for segment in segments:
            if pathname:
                pathname = [pathname]
            file_handle, pathname, desc = imp.find_module(segment, pathname)

        m = imp.load_module(modname, file_handle, pathname, desc)
    except Exception as load_ex:
        return False, load_handler_error_handler(load_ex, modname)
    finally:
        if file_handle is not None:
            file_handle.close()

    try:
        request_handler = getattr(m, func_name)
    except AttributeError:
        return False, make_error_handler(constant.STATUS_ERR, Exception("'{}' not find in module '{}'".format(func_name, modname)))

    global_handlers[md5] = request_handler
    return True, request_handler


def execute_handler(request_handler, evt, ctx, timeout):
    try:
        start_time = time.time()
        data = base64.b64decode(evt.decode())
        event = json.loads(data)
        signal.alarm(int(timeout))
        result = request_handler(event, ctx)
    except Exception as user_ex:
        end_time = time.time()
        exc_info = sys.exc_info()
        message = get_error_message(user_ex, exc_info[2], True)
        if isinstance(user_ex, MemoryError):
            message = "Out of Memory"
        elif isinstance(user_ex, TimeoutError):
            message = "Time out"
        return False, message, round(end_time-start_time, 2)
    else:
        end_time = time.time()
        signal.alarm(0)
        if not isinstance(result, (str, bytes)):
            try:
                return True, make_json(result), round(end_time-start_time, 2)
            except TypeError:
                return True, str(result), round(end_time-start_time, 2)
        return True, str(result), round(end_time-start_time, 2)


def execute_invoke(environ):
    try:
        # get_fc_logger().info(constant.LOG_TAIL_START_PREFIX_INVOKE + get_request_id())

        # validate request
        if get_request_id() is None:
            return constant.STATUS_ERR, "request id is none"
        if environ[constant.HTTP_REQUEST_METHOD] != constant.METHOD_POST:
            return constant.STATUS_ERR, "http method should be post"
        if environ[constant.HTTP_PATH_INFO] != constant.INVOKE_PATH:
            return constant.STATUS_ERR, "path should be /invoke"

        evt, ctx, handler = parse_params(environ)

        # load user handler
        md5 = environ[constant.HTTP_FUNCTION_MD5]
        code_path = environ[constant.HTTP_FUNCTION_CODE_PATH]
        load_success, request_handler = load_handler(handler, md5, code_path)
        if not load_success:
            status, message = request_handler()
            process_mem_info = psutil.Process(os.getpid()).memory_info()
            return status, gen_return_dict(process_mem_info, message, 0)

        # limit memory
        process_mem_info = psutil.Process(os.getpid()).memory_info()
        limit_as = int(process_mem_info.vms) + int(ctx.function.memory_size) * 1024 * 1024
        soft, hard = resource.getrlimit(resource.RLIMIT_AS)
        resource.setrlimit(resource.RLIMIT_AS, (limit_as, hard))

        # execute handler
        execute_success, resp_msg, time_usage = execute_handler(request_handler, evt, ctx, environ[constant.HTTP_FUNCTION_TIMEOUT])
        if not execute_success:
            ret_code = constant.STATUS_USER_ERR
        else:
            ret_code = constant.STATUS_OK
        
        return ret_code, gen_return_dict(process_mem_info, resp_msg, time_usage)
    except Exception as ex:
        exc_info = sys.exc_info()
        ret = get_error_message(ex, exc_info[2])
        process_mem_info = psutil.Process(os.getpid()).memory_info()
        return constant.STATUS_ERR, gen_return_dict(process_mem_info, ret, 0)
    finally:
        pass
        # get_fc_logger().info(constant.LOG_TAIL_END_PREFIX_INVOKE + get_request_id())


def init_env():
    # os.environ = dict()
    os.environ['PATH'] = '/usr/local/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin'
    os.environ['PYTHON_VERSION'] = '2.7.16'
    os.environ['_'] = '/usr/local/bin/python2.7'
    os.environ['PWD'] = '/tmp'
    os.environ['HOME'] = '/tmp'
    os.environ[constant.SERVER_LOG_PATH] = constant.SERVER_LOG_PATH_VALUE


def timeout_handler(sig, frame):
    raise TimeoutError("time out")


def register_timeout_handler():
    signal.signal(signal.SIGALRM, timeout_handler)


def application(environ, start_response):
    """
    main function
    :param environ:
    :param start_response:
    :return:
    """
    imp.reload(sys)
    set_fc_env(environ[constant.HTTP_FUNCTION_ENV])
    init_logging()
    # init_fc_logger()
    sys.path.insert(0, environ[constant.HTTP_FUNCTION_CODE_PATH])

    # register timeout handler
    register_timeout_handler()

    # get request id
    set_request_id(environ[constant.HTTP_REQUEST_ID])

    # change process group
    pgid = os.getpgid(os.getpid())
    new_pgid = os.getpid()
    os.setpgid(os.getpid(), new_pgid)

    status, ret = execute_invoke(environ)
    if isinstance(ret, str):
        ret = ret.encode("utf-8")

    try:
        # recover process group
        os.setpgid(os.getpid(), pgid)
        # clean subprocess
        os.killpg(new_pgid, signal.SIGKILL)
    except Exception:
        pass

    while 1:
        try:
            os.waitpid(-1, os.WNOHANG)
        except OSError:
            break

    if status == constant.STATUS_OK:
        send_success(start_response)
    else:
        send_error(status, start_response)

    return [ret]


# Init environment
init_env()
# load function
g_handler = os.environ[constant.WSGI_HANDLER]
g_md5 = os.environ[constant.WSGI_MD5]
g_code_path = os.environ[constant.WSGI_CODEPATH]
load_handler(g_handler, g_md5, g_code_path)
