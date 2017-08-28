package main

// the unique iolib2 error type
type iolib2Error struct {
	code    errorCode
	wrapped error
}

// TODO: remove, probably useless?
func newIOLib2Error(code errorCode, wrapped error) iolib2Error {
	return iolib2Error{
		code:    code,
		wrapped: wrapped,
	}
}

func (e iolib2Error) Error() string {
	return ``
}

type errorCode int

const (
	success    errorCode = iota
	portError            // illogical operation (eg write on an uninitialized port)
	argError             // arguments that are invalid, malformed, unexpected
	writeError           // error happened during a write operation
	osError              // permission error, file not found, busy port, etc.

	WRITE_ERROR
	ARGS_ERROR
	URI_ERROR
	FILE_ERROR
	PERMISSION_ERROR
	NODEVS_ERROR
	NETWORK_INVALID_SOCKET_ERROR
	NETWORK_SOCKET_ERROR
	NETWORK_WRITE_ERROR
	NETWORK_INCOMPLETE_SEND_ERROR
	UNKNOWN_ERROR
)
