# -*- coding: utf-8 -*-
import sys
import os
import logging
import constant
from logging.handlers import TimedRotatingFileHandler

# global variables
fc_logger = None
global_std_out = None
global_std_err = None


class SelfStdout:
    def __init__(self):
        self.buff = ''
        self.__console__ = sys.stdout

    def write(self, output_stream):
        self.buff += output_stream


class SelfStderr:
    def __init__(self):
        self.buff = ''
        self.__console__ = sys.stderr

    def write(self, output_stream):
        self.buff += output_stream


def get_self_std_out():
    global global_std_out
    return global_std_out


def get_self_std_err():
    global global_std_err
    return global_std_err


def init_logging():
    global global_std_out
    global global_std_err
    global_std_out = SelfStdout()
    sys.stdout = global_std_out
    global_std_err = SelfStderr()
    sys.stderr = global_std_err


def get_fc_logger():
    global fc_logger
    return fc_logger


def init_fc_logger():
    """
    print log to file
    :return:
    """
    global fc_logger

    log_file = "{}/log_{}.log".format(constant.SERVER_LOG_PATH_VALUE, os.popen("id -u").read().replace("\n", ""))
    formatter = logging.Formatter('%(asctime)s.%(msecs)d [%(levelname)s] %(message)s', '%Y-%m-%d %H:%M:%S')
    file_time_handler = TimedRotatingFileHandler(log_file, when="d", interval=1)
    file_time_handler.setFormatter(formatter)

    fc_logger = logging.getLogger("fc_logger")
    fc_logger.setLevel(logging.INFO)
    fc_logger.addHandler(file_time_handler)
    return





