package iolib2

import (
	"fmt"
	"net"

	log "github.com/Sirupsen/logrus"
)

type TcpPort struct {
	conn net.Conn
}

func NewtcpPort() Port {
	return &TcpPort{}
}

func (sp *TcpPort) Name() string {
	return "net"
}

func (sp *TcpPort) Set(cfg cfgDict) error {
	log.Debugf("tcpPort.set(%v)", cfg)
	if sp.conn != nil {
		return fmt.Errorf("tcp port, can't set an initialized port")
	}

	var (
		ip, port string
		ok       bool
		err      error
	)
	if ip, ok = cfg["ip"]; !ok {
		return fmt.Errorf("tcp port, no 'ip' parameter")
	}
	if port, ok = cfg["port"]; !ok {
		return fmt.Errorf("tcp port, no 'port' parameter")
	}

	sp.conn, err = net.Dial("tcp", fmt.Sprintf("%s:%s", ip, port))
	if err != nil {
		return fmt.Errorf("tcp port error, can't open connection, %s", err)
	}
	return nil
}

func (sp *TcpPort) Reset() error {
	if sp.conn != nil {
		return sp.conn.Close()
	}
	return nil
}

func (sp *TcpPort) Write(buf []byte) error {
	if sp.conn != nil {
		var ntot int
		for {
			n, err := sp.conn.Write(buf)
			ntot += n
			if err != nil {
				return fmt.Errorf("tcp port error, write error, %s", err)
			}
			if n == ntot {
				break
			}
		}
	}
	return nil
}

func (sp *TcpPort) Enumerate() ([]portEntry, error) {
	// nothing to enumerate
	return []portEntry{}, nil
}
