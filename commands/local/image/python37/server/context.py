# -*- coding: utf-8 -*-

import json


class FunctionMeta:
    def __init__(self, invoked_function_id, function_name, function_version, function_handler, memory_size, timeout):
        self.invoked_function_id = invoked_function_id
        self.function_name = function_name
        self.function_version = function_version
        self.function_handler = function_handler
        self.memory_size = memory_size
        self.timeout = timeout

    def to_dict(self):
        return {
            "id": self.invoked_function_id,
            "name": self.function_name,
            "version": self.function_version,
            "handler": self.function_handler,
            "memory": self.memory_size,
            "timeout": self.timeout,
        }


class FCContext:
    def __init__(self, request_id, function_meta, log_set, log_topic):
        self.request_id = request_id
        self.function = function_meta
        self.log_set = log_set
        self.log_topic = log_topic

    def to_json(self):
        return json.dumps({"requestId": self.request_id,
                           "function": self.function.to_dict(),
                           "logSet": self.log_set,
                           "logTopic": self.log_topic,
                           })
