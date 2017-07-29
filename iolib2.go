package main

import (
	"bytes"
	"fmt"
	"log"
	"net/url"
	"strings"
)

type portHandler struct {
	curPort port         // the current output port
	buf     bytes.Buffer // transmit buffer
}

func (ph *portHandler) handleMessage(msg string) error {
	cmdParm := strings.Split(msg, "|")
	if len(cmdParm) != 2 {
		return fmt.Errorf("malformed request, \"%s\"", msg)
	}
	cmd, parm := cmdParm[0], cmdParm[1]

	log.Printf("received %s with param %s", cmd, parm)

	switch cmd {
	case "SET-PORT":
		return ph.setPort(parm)
	case "RESET":
		return ph.reset()
	case "WRITE":
		return ph.write([]byte(parm))
	case "SEND":
		return ph.send()
	}
	return fmt.Errorf("invalid command, \"%s\"", cmd)
}

func (ph *portHandler) setPort(urlStr string) error {
	u, err := url.Parse(urlStr)
	if err != nil {
		return err
	}
	log.Print(u)

	// TODO continuer ici

	return nil
}

func (ph *portHandler) reset() error {
	if ph.curPort != nil {
		return ph.curPort.reset()
	}
	// does nothing if port is not configured
	return nil
}

func (ph *portHandler) write(buf []byte) error {
	if ph.curPort != nil {
		ph.buf.Write(buf)
		return nil
	}
	return fmt.Errorf("writing on a unintialized port")
}

func (ph *portHandler) send() error {
	if ph.curPort != nil {
		return ph.curPort.write(ph.buf.Bytes())
	}
	return fmt.Errorf("sending to a unintialized port")
}
