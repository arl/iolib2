# -*- coding: utf-8 -*-
from __future__ import print_function, division

import sys
import os
import logging
import subprocess

import zmq


iolib2_com = None


class IOLibError(IOError):
    def __init__(self, msg, errno, config):
        self.msg = msg
        self.errno = -1  # TODO: undefined for now
        self.config = config

    def __str__(self):
        return "iolib2 error({}), cfg={} : {}".format(self.errno, self.config, self.msg)


class IOLib2Exe(object):
    if not hasattr(sys, 'frozen'):
        exe_path = os.path.join(os.path.dirname(__file__), "iolib2.exe")
    else:
        exe_path = os.path.join(os.path.dirname(sys.argv[0]), "iolib2.exe")

    port_number = None

    def __init__(self, port):
        self.log = logging.getLogger(__name__)
        if self.port_number is None:
            self.port_number = port
        else:
            self.log.exception("IOLIb2.exe port must not be more than once")
            raise RuntimeError("IOLIb2.exe port must not be more than once")
        self.proc = None

    def start(self):
        try:
            args = [self.exe_path, str(self.port_number)]
            self.proc = subprocess.Popen(args, close_fds=True)
            if self.proc.poll() is not None:
                # process has returned prematurely
                return False
        except OSError as exc:
            self.log.exception("Couldn't start {}: {}".format(args, exc))
            return False
        return True

    def enumerate(self):
        pass

    def stop(self):
        if self.proc.poll() is None:
            # process is still alive
            self.proc.terminate()


class IOLib2Com(object):

    def __init__(self, port_number):
        self.log = logging.getLogger(__name__)
        self.port_number = port_number
        self.addr = "tcp://localhost:{}".format(self.port_number)
        self.context = None
        self.socket = None

    def connect(self):
        if self.context is None:
            self.context = zmq.Context()
            if self.socket is None:
                self.socket = self.context.socket(zmq.REQ)
                self.log.info("IOlib2 connecting to {}".format(self.addr))
                self.socket.connect(self.addr)

    def disconnect(self):
        # import ipdb
        # ipdb.set_trace()
        if self.socket is not None:
            self.socket.disconnect(self.addr)
            if not self.socket.closed:
                self.socket.close()
            self.socket = None
        if self.context is not None:
            self.context.term()
            self.conext = None

    def set_port(self, device_url):
        self.__send_request("SET-PORT|{}".format(device_url))
        return self.__wait_for_reply()

    def reset_port(self):
        self.__send_request("RESET|")
        return self.__wait_for_reply()

    def write(self, s):
        self.__send_request("WRITE|{}".format(s))
        return self.__wait_for_reply()

    def send(self):
        self.__send_request("SEND|")
        return self.__wait_for_reply()

    def __send_request(self, cmd):
        if self.context and self.socket and not self.socket.closed:
            self.log.debug("Sending command \"\"".format(cmd))
            self.socket.send(cmd)

    def __wait_for_reply(self):
        if self.context and self.socket and not self.socket.closed:
            self.log.debug("Waiting for reply...")
            reply = self.socket.recv()
            self.log.debug("Received {}".format(reply))
            ret_val, ret_str = reply.split("|")
            if ret_val != '0':
                self.log.error("Error, {}: {}".format(ret_val, ret_str))
                return int(ret_val), ret_str
        return 0, None
