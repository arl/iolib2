package iolib2

import (
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
)

type FilePort struct {
	fn string // filename
}

func NewFilePort() Port {
	return &FilePort{}
}

func (fp *FilePort) Name() string {
	return "file"
}

func (fp *FilePort) Set(cfg cfgDict) error {
	log.Debugf("fileport.set(%v)", cfg)
	name, ok := cfg["name"]
	if !ok {
		return fmt.Errorf("fileport error: no name parameter")
	}
	fp.fn = name
	log.Infof("fileport, set filename to %v", fp.fn)
	return nil
}

func (fp *FilePort) Reset() error {
	// nothing to do
	return nil
}

func (fp *FilePort) Write(buf []byte) error {
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

func (fp *FilePort) Enumerate() ([]portEntry, error) {
	// nothing to enumerate
	return []portEntry{}, nil
}
