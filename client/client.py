#! /usr/bin/env python2

import zmq

import sys
import argparse

from iolib2 import IOLib2Com

def connect_and_send(port, url, data):
    iolib2_com = IOLib2Com(port)
    print "Connecting to daemon..."
    iolib2_com.connect()

    ret = iolib2_com.set_port(url)
    print('set_port ->', ret)

    ret = iolib2_com.write(data)
    print('write ->', ret)

    ret = iolib2_com.send()
    print('send ->', ret)

    ret = iolib2_com.reset_port()
    print('reset_port ->', ret)

    iolib2_com.disconnect()


def main():
    usage = """ {} -p PORT -u URL -f FILE
Sends data to URL, via iolib2 daemon, listening on PORT

With no FILE, read data from standard input.

Some example URL strings are:
    - file://name=/tmp/file.hpgl
    - serial://port=/dev/pts/8;parms=9600,n,8,1
    - net://ip=localhost;port=9100
""".format(__file__)

    parser = argparse.ArgumentParser(usage=usage)
    parser.add_argument('-p', '--port', type=int, required=True,
                        help='communicate with iolib2 daemon on PORT')
    parser.add_argument('-u', '--url', required=True,
                        help='send commands to device at URL')
    parser.add_argument('-f', '--file',
                        help="read data FILE")
    args = parser.parse_args()
    if not args.file:
        hpgl = sys.stdin.readlines()
    else:
        with open(args.file, "rb") as f:
            hpgl = f.readlines()
    connect_and_send(args.port, args.url, ''.join(hpgl))


if __name__ == "__main__":
    main()
