package main

import (
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
)

type filePort struct {
	fn string // filename
}

func newFilePort() port {
	return &filePort{}
}

func (fp *filePort) name() string {
	return "file"
}

func (fp *filePort) set(cfg cfgDict) error {
	log.Debugf("fileport.set(%v)", cfg)
	if name, ok := cfg["name"]; !ok {
		return fmt.Errorf("fileport error: no name parameter")
	} else {
		fp.fn = name
		log.Infof("fileport, set filename to %v", fp.fn)
	}
	return nil
}

func (fp *filePort) reset() error {
	// nothing to do
	return nil
}

func (fp *filePort) write(buf []byte) error {
	// create file or open it in order to append to it if it exists
	f, err := os.OpenFile(fp.fn, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(buf)
	if err != nil {
		return err
	}
	f.Sync()
	return nil
}

func (fp *filePort) enumerate() ([]portEntry, error) {
	// nothing to enumerate
	return []portEntry{}, nil
}
