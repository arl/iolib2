package main

type cfgDict map[string]string

// port is the interface implemented by various kinds of plotters.
type port interface {
	// set configures the writer with given configuration dict.
	set(cfg cfgDict) error

	// reset closes the writer and performs eventual cleanup.
	reset() error

	// write writes a buffer s to the writer.
	write(buf []byte) error

	// enumerate enumerates the available ports of the machine and their
	// configuration string.
	enumerate() ([]portEntry, error)

	// name returns the port name
	name() string
}

// portEntry represents a port entry.
type portEntry struct{ name, cfg string }
