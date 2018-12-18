package thrift

import (
	"io"
	"log"
	"os"
	"strconv"
)

// TTransport Generic class that encapsulates the I/O layer. This is basically a thin
// wrapper around the combined functionality of Java input/output streams.
type TTransport interface {
	/**
	 * Queries whether the transport is open.
	 *
	 * @return True if the transport is open.
	 */
	IsOpen() bool

	/**
	 * Opens the transport for reading/writing.
	 *
	 * @returns TTransportException if the transport could not be opened
	 */
	Open() (err error)

	/**
	 * Closes the transport.
	 */
	Close() (err error)

	/**
	 * Reads up to len bytes into buffer buf, starting att offset off.
	 *
	 * @param buf Array to read into
	 * @param off Index to start reading at
	 * @param len Maximum number of bytes to read
	 * @return The number of bytes actually read
	 * @return TTransportException if there was an error reading data
	 */
	Read(buf []byte) (n int, err error)

	/**
	 * Guarantees that all of len bytes are actually read off the transport.
	 *
	 * @param buf Array to read into
	 * @param off Index to start reading at
	 * @param len Maximum number of bytes to read
	 * @return The number of bytes actually read, which must be equal to len
	 * @return TTransportException if there was an error reading data
	 */
	ReadAll(buf []byte) (n int, err error)

	/**
	 * Writes the buffer to the output
	 *
	 * @param buf The output data buffer
	 * @return Number of bytes written
	 * @return TTransportException if an error occurs writing data
	 */
	Write(buf []byte) (n int, err error)

	/**
	 * Flush any pending data out of a transport buffer.
	 *
	 * @return TTransportException if there was an error writing out data.
	 */
	Flush() (err error)

	/**
	 * Is there more data to be read?
	 *
	 * @return True if the remote side is still alive and feeding us
	 */
	Peek() bool
}

// ReadAllTransport Guarantees that all of len bytes are actually read off the transport.
// @param buf Array to read into
// @param off Index to start reading at
// @param len Maximum number of bytes to read
// @return The number of bytes actually read, which must be equal to len
// @return TTransportException if there was an error reading data
func ReadAllTransport(p TTransport, buf []byte) (n int, err error) {
	size := len(buf)
	for n < size {
		ret, err := p.Read(buf[n:])
		if ret <= 0 {
			if err != nil {
				err = NewTTransportExceptionDefaultString("Cannot read. Remote side has closed. Tried to read " + strconv.Itoa(size) + " bytes, but only got " + strconv.Itoa(n) + " bytes.")
			}
			return ret, err
		}
		n += ret
	}
	return n, err
}

var (
	LOGGER *log.Logger
)

func init() {
	LOGGER = log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile)
}

type TTransportException interface {
	TException
	TypeId() int
}

const (
	UNKNOWN_TRANSPORT_EXCEPTION = 0
	NOT_OPEN                    = 1
	ALREADY_OPEN                = 2
	TIMED_OUT                   = 3
	END_OF_FILE                 = 4
)

type tTransportException struct {
	typeId  int
	message string
}

func (p *tTransportException) TypeId() int {
	return p.typeId
}

func (p *tTransportException) Error() string {
	return p.message
}

func NewTTransportExceptionDefault() TTransportException {
	return NewTTransportExceptionDefaultType(UNKNOWN_TRANSPORT_EXCEPTION)
}

func NewTTransportExceptionDefaultType(t int) TTransportException {
	return NewTTransportException(t, "")
}

func NewTTransportExceptionDefaultString(m string) TTransportException {
	return NewTTransportException(UNKNOWN_TRANSPORT_EXCEPTION, m)
}

func NewTTransportException(t int, m string) TTransportException {
	return &tTransportException{typeId: t, message: m}
}

func NewTTransportExceptionFromOsError(e error) TTransportException {
	if e == nil {
		return nil
	}
	t, ok := e.(TTransportException)
	if ok {
		return t
	}
	if e == io.EOF {
		return NewTTransportException(END_OF_FILE, e.Error())
	}
	return NewTTransportExceptionDefaultString(e.Error())
}
