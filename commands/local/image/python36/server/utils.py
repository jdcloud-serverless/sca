# -*- coding: utf-8 -*-

"""
JSON serializer for objects not serializable by default json code
 like bytes, datetime, decimal.Decimal....
"""


import sys
import traceback
import json
import constant


def make_json(obj):
    return json.dumps(obj, indent=4, sort_keys=True)


# ca response error function
def make_error_handler(status, ex):
    exc_type, exc_value, exc_traceback = sys.exc_info()
    tb = traceback.format_exception_only(exc_type, exc_value)
    trace_list = []
    for idx, trace_line in enumerate(tb):
        trace_line_lst = trace_line.split('\n')
        for item in trace_line_lst:
            r_item = item.replace('^','').strip()
            if not r_item:
                continue
            if idx != 0:
                trace_list.append(r_item)
            else:
                # python stack trace file and line info in the first
                # like '  File "/code/hello_world.py", line 3\n', ' , need split by ,
                for content in r_item.split(','):
                    r_content = content.strip()
                    if r_content:
                        trace_list.append(r_content)
    msg = make_json(make_error_info(ex, trace_list))

    def result(*args):
        return status, msg

    return result


# error_info is an array
def make_error_info(ex, trace_list):
    result = {}
    err_msg = str(ex)
    result['error_message'] = err_msg
    result['error_type'] = ex.__class__.__name__
    if trace_list:
        result['error_trace'] = trace_list
    return result


def load_handler_error_handler(e, modname):
    if isinstance(e, ImportError):
        return make_error_handler(constant.STATUS_USER_ERR, ImportError("Import module '{}' error".format(modname)))
    elif isinstance(e, SyntaxError):
        return make_error_handler(constant.STATUS_USER_ERR, SyntaxError("Syntax error in module '{}'".format(modname)))
    elif isinstance(e, ValueError):
        return make_error_handler(constant.STATUS_USER_ERR, ValueError("Invalid handler '{}'".format(modname)))
    else:
        return make_error_handler(constant.STATUS_ERR, Exception("Load handler error: '%s'" % str(e)))


def format_trace_list(trace_list):
    def format_trace_item(trace_item):
        first_split = trace_item.split('\n')
        result = []
        for item in first_split[0].split(','):
            result.append(item.strip())

        result.append(first_split[1].strip())
        return result

    for index in range(len(trace_list)):
        trace_list[index] = format_trace_item(trace_list[index])
    return trace_list


def get_error_message(ex, tb, rm_first_trace=False):
    """
    param rm_first_trace: remove first trace in trace stack
    :param ex: 
    :param tb: 
    :param rm_first_trace: 
    :return: 
    """
    trace = traceback.format_tb(tb)
    if rm_first_trace:
        trace = trace[1:]
    trace_list = format_trace_list(trace)
    return make_json(make_error_info(ex, trace_list))
