package main

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/jacobsa/go-serial/serial"
)

type serialPort struct {
	port io.ReadWriteCloser
}

func newSerialPort() port {
	return &serialPort{}
}

func (sp *serialPort) name() string {
	return "serial"
}

func (sp *serialPort) set(cfg cfgDict) error {
	log.Debugf("serialPort.set(%v)", cfg)
	if sp.port != nil {
		return fmt.Errorf("serial port error, can't set an initialized port")
	}

	options, err := parseSerialPortOptions(cfg)
	if err != nil {
		return fmt.Errorf("serial port error, config parsing error, %s", err)
	}

	options.MinimumReadSize = 0
	options.InterCharacterTimeout = 500
	sp.port, err = serial.Open(options)
	if err != nil {
		return fmt.Errorf("serial port error, can't open port, %s", err)
	}
	return nil
}

func (sp *serialPort) reset() error {
	if sp.port != nil {
		sp.port.Close()
	}
	return nil
}

func (sp *serialPort) write(buf []byte) error {
	if sp.port != nil {
		_, err := sp.port.Write(buf)
		if err != nil {
			return fmt.Errorf("serial port error, write error, %s", err)
		}
	}
	return nil
}

func (sp *serialPort) enumerate() ([]portEntry, error) {
	// nothing to enumerate
	return []portEntry{}, nil
}

func parseSerialPortOptions(cfg cfgDict) (options serial.OpenOptions, err error) {

	var (
		parms, port string
		intVal      int
		baudRate    uint
		dataBits    uint
		stopBits    uint
		parity      serial.ParityMode
		ok          bool
	)

	if port, ok = cfg["port"]; !ok {
		err = fmt.Errorf("no 'port' parameter")
		return
	}
	if parms, ok = cfg["parms"]; !ok {
		err = fmt.Errorf("no 'parms' parameter")
		return
	}

	optVals := strings.Split(parms, ",")
	if len(optVals) != 4 {
		err = fmt.Errorf("incomplete 'parms' \"%v\"", parms)
		return
	}

	// parse serial port options
	intVal, err = strconv.Atoi(optVals[0])
	if err != nil {
		err = fmt.Errorf("malformed 'bauds rate' \"%v\", %v", optVals[0], err)
		return
	}
	baudRate = uint(intVal)
	switch optVals[1] {
	case "n":
		parity = serial.PARITY_NONE
	case "o":
		parity = serial.PARITY_ODD
	case "e":
		parity = serial.PARITY_EVEN
	default:
		err = fmt.Errorf("malformed 'parity mode' '%v'", optVals[1])
		return
	}
	intVal, err = strconv.Atoi(optVals[2])
	if err != nil {
		err = fmt.Errorf("malformed 'data bits' \"%v\", %v", optVals[2], err)
		return
	}
	dataBits = uint(intVal)

	intVal, err = strconv.Atoi(optVals[3])
	if err != nil {
		err = fmt.Errorf("malformed 'stop bits' \"%v\", %v", optVals[3], err)
		return
	}
	stopBits = uint(intVal)

	// Set up options.
	return serial.OpenOptions{
		PortName:   port,
		BaudRate:   baudRate,
		ParityMode: parity,
		DataBits:   dataBits,
		StopBits:   stopBits,
	}, nil
}
