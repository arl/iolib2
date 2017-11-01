package iolib2

type cfgDict map[string]string

// port is the interface implemented by various kinds of plotters.
type Port interface {
	// Set configures the writer with given configuration dict.
	Set(cfg cfgDict) error

	// Reset closes the writer and performs eventual cleanup.
	Reset() error

	// Write writes a buffer s to the writer.
	Write(buf []byte) error

	// Enumerate enumerates the available ports of the machine and their
	// configuration string.
	Enumerate() ([]portEntry, error)

	// Name returns the port name
	Name() string
}

// portEntry represents a port entry.
type portEntry struct{ name, cfg string }
