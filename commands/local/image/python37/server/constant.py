# -*- coding: utf-8 -*-


HTTP_REQUEST_ID = 'HTTP_X_FC_REQUEST_ID'
HTTP_FUNCTION_ID = 'HTTP_X_FC_FUNCTION_ID'
HTTP_FUNCTION_NAME = 'HTTP_X_FC_FUNCTION_NAME'
HTTP_FUNCTION_HANDLER = 'HTTP_X_FC_FUNCTION_HANDLER'
HTTP_FUNCTION_MEMORY = 'HTTP_X_FC_FUNCTION_MEMORY'
HTTP_FUNCTION_TIMEOUT = 'HTTP_X_FC_FUNCTION_TIMEOUT'
HTTP_FUNCTION_ENV = 'HTTP_X_FC_FUNCTION_ENV'
HTTP_FUNCTION_VERSION = 'HTTP_X_FC_FUNCTION_VERSION'
HTTP_FUNCTION_MD5 = 'HTTP_X_FC_FUNCTION_MD5'
HTTP_FUNCTION_CODE_PATH = 'HTTP_X_FC_FUNCTION_CODEPATH'
HTTP_CONTEXT_LOG_SET = "HTTP_X_FC_CONTEXT_LOGSET"
HTTP_CONTEXT_LOG_TOPIC = "HTTP_X_FC_CONTEXT_LOGTOPIC"

STD_OUT_RETURN = "StdoutReturn"
STD_ERR_RETURN = "StderrReturn"
FUNCTION_RETURN = "FunctionReturn"
MEM_USAGE = "MemUsage"
TIME_USAGE = "TimeUsage"

HTTP_CONTENT_LENGTH = 'CONTENT_LENGTH'
HTTP_REQUEST_METHOD = 'REQUEST_METHOD'
HTTP_PATH_INFO = 'PATH_INFO'
HTTP_WSGI_INPUT = 'wsgi.input'
METHOD_POST = 'POST'
INVOKE_PATH = '/invoke'

WSGI_HANDLER = "HANDLER"
WSGI_MD5 = "CHECKSUM"
WSGI_CODEPATH = "CODEPATH"

# log
SERVER_LOG_PATH = 'FC_SERVER_LOG_PATH'  # Server app log path.
SERVER_LOG_PATH_VALUE = '/tmp'  # Server app log path.
LOG_TAIL_START_PREFIX_INVOKE = 'FC Invoke Start RequestId: ' # Start of invoke log tail mark
LOG_TAIL_END_PREFIX_INVOKE = 'FC Invoke End RequestId: '   # End of invoke log tail mark

# status code
STATUS_OK = 200
STATUS_ERR = 500
STATUS_USER_ERR = 400

# return length
MAX_LENGTH = -1024*100
