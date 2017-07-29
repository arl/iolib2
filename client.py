#! /usr/bin/env python2

# from http://learning-0mq-with-pyzmq.readthedocs.io/en/latest/pyzmq/patterns/client_server.html
import zmq
import sys

def main():
    port = "5556"
    if len(sys.argv) > 1:
        port =  sys.argv[1]
        int(port)

    if len(sys.argv) > 2:
        port1 =  sys.argv[2]
        int(port1)

    context = zmq.Context()
    print "Connecting to server..."
    socket = context.socket(zmq.REQ)
    socket.connect ("tcp://localhost:%s" % port)
    if len(sys.argv) > 2:
        socket.connect ("tcp://localhost:%s" % port1)


    cmds = [
            "SET-PORT|file://name=/tmp/file.hpgl",
            "WRITE|coucou\n",
            "WRITE|coucou2\n",
            "SEND|",
        ]

    for cmd in cmds:
        # request
        print "Sending command ", cmd,"..."
        ret = socket.send (cmd)

        # reply
        msg = socket.recv()
        print "Received reply ", cmd, "[", msg, "]"

if __name__ == "__main__":
    main()
