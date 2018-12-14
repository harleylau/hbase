package hbase

import (
	"bytes"

	"github.com/J-J-J/hbase/Hbase"
)

// Error return standard Hbase Error
type Error struct {
	IOErr  *Hbase.IOError         // IOError
	ArgErr *Hbase.IllegalArgument // IllegalArgument
	Err    error                  // error

}

func newError(io *Hbase.IOError, arg *Hbase.IllegalArgument, err error) *Error {
	return &Error{
		IOErr:  io,
		ArgErr: arg,
		Err:    err,
	}
}

// String return error string.
func (e *Error) String() string {
	if e == nil {
		return "<nil>"
	}
	var b bytes.Buffer
	if e.IOErr != nil {
		b.WriteString("IOError:")
		b.WriteString(e.IOErr.Message)
		b.WriteString(";")
	}
	if e.ArgErr != nil {
		b.WriteString("ArgumentError:")
		b.WriteString(e.ArgErr.Message)
		b.WriteString(";")
	}
	if e.Err != nil {
		b.WriteString("Error:")
		b.WriteString(e.Err.Error())
		b.WriteString(";")
	}
	return b.String()
}

// Error is implement of error interface.
func (e *Error) Error() string {
	return e.String()
}

func checkError(io *Hbase.IOError, err error) error {
	if io != nil || err != nil {
		return newError(io, nil, err)
	}
	return nil
}

func checkHbaseArgError(io *Hbase.IOError, arg *Hbase.IllegalArgument, err error) error {
	if io != nil || arg != nil || err != nil {
		return newError(io, arg, err)
	}
	return nil
}
