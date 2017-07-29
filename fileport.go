package main

import (
	"fmt"
	"os"
)

type filePort struct {
	fn string // filename
}

func newFilePort() *filePort {
	return &filePort{}
}

func (fp *filePort) Set(cfg cfgDict) error {
	if name, ok := cfg["name"]; !ok {
		return fmt.Errorf("fileport error: no name parameter")
	} else {
		fp.fn = name
	}
	return nil
}

func (fp *filePort) Reset() error {
	// nothing to do
	return nil
}

func (fp *filePort) Write(buf []byte) error {
	// create file or open if in order to append to it if it exists
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
