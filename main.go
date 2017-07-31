package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	log "github.com/Sirupsen/logrus"
	zmq "github.com/pebbe/zmq4"
)

// from http://learning-0mq-with-pyzmq.readthedocs.io/en/latest/pyzmq/patterns/client_server.html

func main() {
	log.StandardLogger().SetLevel(log.DebugLevel)

	if len(os.Args) != 2 {
		log.Fatalln("Usage iolib2.exe PORTNUMBER")
	}

	var (
		port int
		err  error
		ctx  *zmq.Context
		sck  *zmq.Socket
	)
	handler := newPortHandler()

	handler.registerPort("file", newFilePort)
	handler.registerPort("serial", newSerialPort)

	port, err = strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalln("invalid port,", err)
	}

	ctx, err = zmq.NewContext()
	if err != nil {
		log.Fatalln("error creating context,", err)
	}

	sck, err = ctx.NewSocket(zmq.REP)
	if err != nil {
		log.Fatalln("error creating socket,", err)
	}

	if err = sck.Bind(fmt.Sprintf("tcp://*:%v", port)); err != nil {
		log.Fatalln("error binding socket,", err)
	}

	for {
		// wait for next request from client
		msg, err := sck.Recv(zmq.DONTWAIT)
		if err == nil {
			log.Printf("received request \"%s\"", msg)
			err = handler.handleMessage(msg)
			errString := newErrorString(err)
			sck.Send(errString, zmq.DONTWAIT)
		}
		time.Sleep(2 * time.Millisecond)
	}
}

func newErrorString(err error) string {
	var s string
	if err == nil {
		s = "0|"
	} else {
		s = fmt.Sprintf("1|%v", err)
	}
	return s
}
