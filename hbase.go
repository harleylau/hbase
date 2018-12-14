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
	TFramedTransport = iota
	TSocket
)

func newProtocolFactory(protocol int) (thrift.TProtocolFactory, error) {
	switch protocol {
	case TBinaryProtocol:
		return thrift.NewTBinaryProtocolFactoryDefault(), nil
	case TCompactProtocol:
		return thrift.NewTCompactProtocolFactory(), nil
	}
	return nil, errors.New(fmt.Sprint("invalid protocol:", protocol))
}
