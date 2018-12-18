package hbase

import (
	"errors"
	"fmt"

	"github.com/J-J-J/hbase/thrift" // will replace it later
)

const (
	stateDefault = iota // default close
	stateOpen           // open
)

// protocol
const (
	TBinaryProtocol = iota
	TCompactProtocol
)

// transport
const (
	TSocket = iota
)

func newProtocolFactory(protocol int, trans thrift.TTransport) (thrift.TProtocolFactory, error) {
	switch protocol {
	case TBinaryProtocol:
		return thrift.NewTBinaryProtocol(trans, false, true), nil
	}
	return nil, errors.New(fmt.Sprint("invalid protocol:", protocol))
}
