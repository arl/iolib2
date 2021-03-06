package iolib2

import (
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
)

type ParallelPort struct {
	fn string // filename
}

func NewParallelPort() Port {
	return &ParallelPort{}
}

func (fp *ParallelPort) Name() string {
	return "parallel"
}

func (fp *ParallelPort) Set(cfg cfgDict) error {
	panic("not implemented")
	panic("to not have problems due to slow operations on port opening of the parallel port, we should, contatrily to the standard file port, keep the file pointer opened")

	log.Debugf("parallel.set(%v)", cfg)
	name, ok := cfg["name"]
	if !ok {
		return fmt.Errorf("parallel error: no name parameter")
	}
	fp.fn = name
	log.Infof("parallel, set filename to %v", fp.fn)
	return nil
}

func (fp *ParallelPort) Reset() error {
	// nothing to do
	return nil
}

func (fp *ParallelPort) Write(buf []byte) error {
	// create file or open it in order to append to it if it exists
	f, err := os.OpenFile(fp.fn, os.O_RDWR, 0644)
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

func (fp *ParallelPort) Enumerate() ([]portEntry, error) {
	// nothing to enumerate
	return []portEntry{}, nil
}
