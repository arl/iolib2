#! /usr/bin/env python2

import zmq

import sys
import argparse


def connect_and_send(port, url, hpgl):
    context = zmq.Context()
    print "Connecting to server..."
    socket = context.socket(zmq.REQ)
    socket.connect("tcp://localhost:%s" % port)

    cmds = [
        "SET-PORT|{url}".format(**locals()),
        "WRITE|{hpgl}\n".format(**locals()),
        "SEND|",
        "RESET|",
    ]

    for cmd in cmds:
        # request
        print "Sending command ", cmd, "..."
        ret = socket.send(cmd)

        # reply
        msg = socket.recv()
        print "Received reply ", cmd, "[", msg, "]"
        ret_val, ret_str = msg.split("|")
        if ret_val != '0':
            print "Error, %v: %v".format(ret_val, ret_str)
            break


def main():
    usage = """ {} -p PORT -u URL -f FILE
Sends commands to plotter via iolib2.exe.

With no FILE, read HPGL commands from standard input.

Some example plotter URL strings are:
    - file://name=/tmp/file.hpgl
    - serial://port=/dev/pts/8;parms=9600,n,8,1
    - net://ip=localhost;port=9100
""".format(__file__)

    parser = argparse.ArgumentParser(usage=usage)
    parser.add_argument('-p', '--port', type=int, required=True,
                        help='communicate with iolib2.exe on PORT')
    parser.add_argument('-u', '--url', required=True,
                        help='send commands to plotter at URL')
    parser.add_argument('-f', '--file',
                        help="read HPGL commands from FILE")
    args = parser.parse_args()
    if not args.file:
        hpgl = sys.stdin.readlines()
    else:
        with open(args.file, "rb") as f:
            hpgl = f.readlines()
    connect_and_send(args.port, args.url, ''.join(hpgl))


if __name__ == "__main__":
    main()
