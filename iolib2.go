package main

import (
	"bytes"
	"fmt"
	"strings"

	log "github.com/Sirupsen/logrus"
)

type portHandler struct {
	registry map[string]func() port // port registry
	curPort  port                   // the current output port
	buf      bytes.Buffer           // transmit buffer
}

func newPortHandler() *portHandler {
	return &portHandler{
		registry: map[string]func() port{},
	}
}

func (ph *portHandler) registerPort(name string, fn func() port) {
	ph.registry[name] = fn
}

func (ph *portHandler) handleMessage(msg string) error {
	cmdParm := strings.Split(msg, "|")
	if len(cmdParm) != 2 {
		return fmt.Errorf("malformed request, \"%s\"", msg)
	}
	cmd, parm := cmdParm[0], cmdParm[1]

	log.Debugf("received %s with param %s", cmd, parm)

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
	// check port is currently uninitialized
	if ph.curPort != nil {
		return fmt.Errorf("port has already initialized (curPort:%s)", ph.curPort.name())
	}

	// parse URL
	vals := strings.Split(urlStr, "://")
	if len(vals) != 2 {
		return fmt.Errorf("malformed URL \"%s\"", urlStr)
	}

	// URL scheme is the port type
	portType, params := vals[0], vals[1]
	newPort, ok := ph.registry[portType]
	if !ok {
		return fmt.Errorf("unknown port type %s for \"%s\" URL", portType, urlStr)
	}

	// parse paramaters
	m, err := toMap(params)
	if err != nil {
		return fmt.Errorf("malformed URL parameters, %v", err)
	}

	// create and initialize the port
	ph.curPort = newPort()
	return ph.curPort.set(cfgDict(m))
}

func (ph *portHandler) reset() error {
	if ph.curPort != nil {
		log.Infof("reset current port, \"%v\"", ph.curPort.name())
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

func toMap(s string) (map[string]string, error) {
	m := map[string]string{}

	for _, v := range strings.Split(s, ";") {
		fmt.Println(v)
		s2 := strings.Split(v, "=")
		if len(s2) != 2 {
			return m, fmt.Errorf("malformed parameter \"%v\"", v)
		}
		param, val := s2[0], s2[1]
		m[param] = val
	}
	return m, nil
}
