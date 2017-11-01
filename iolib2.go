package iolib2

import (
	"bytes"
	"fmt"
	"strings"

	log "github.com/Sirupsen/logrus"
)

type PortHandler struct {
	registry map[string]func() Port // port registry
	curPort  Port                   // the current output port
	buf      bytes.Buffer           // transmit buffer
}

func NewPortHandler() *PortHandler {
	return &PortHandler{
		registry: map[string]func() Port{},
	}
}

func (ph *PortHandler) RegisterPort(name string, fn func() Port) {
	ph.registry[name] = fn
}

func (ph *PortHandler) HandleMessage(msg string) error {

	//# TODO: continue here

	cmdParm := strings.Split(msg, "|")
	if len(cmdParm) != 2 {
		return fmt.Errorf("malformed request, \"%s\"", msg)
	}
	cmd, parm := cmdParm[0], cmdParm[1]

	log.Debugf("received %s with param %s", cmd, parm)

	switch cmd {
	case "SET-PORT":
		return ph.SetPort(parm)
	case "RESET":
		return ph.Reset()
	case "WRITE":
		return ph.Write([]byte(parm))
	case "SEND":
		return ph.Send()
	}
	return fmt.Errorf("unknown command, \"%s\"", cmd)
}

func (ph *PortHandler) SetPort(urlStr string) error {
	// check port is currently uninitialized
	if ph.curPort != nil {
		return fmt.Errorf("port has already initialized (curPort:%s)", ph.curPort.Name())
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
	return ph.curPort.Set(cfgDict(m))
}

func (ph *PortHandler) Reset() error {
	if ph.curPort != nil {
		log.Infof("reset current port, \"%v\"", ph.curPort.Name())
		return ph.curPort.Reset()
	}
	// does nothing if port is not configured
	return nil
}

func (ph *PortHandler) Write(buf []byte) error {
	if ph.curPort != nil {
		ph.buf.Write(buf)
		return nil
	}
	return fmt.Errorf("writing on a unintialized port")
}

func (ph *PortHandler) Send() error {
	if ph.curPort != nil {
		return ph.curPort.Write(ph.buf.Bytes())
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
