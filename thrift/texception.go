package thrift

import (
	"encoding/base64"
	"errors"
)

/**
 * Generic exception class for Thrift.
 *
 */

type TException interface {
	Error() string
}

type tException struct {
	message string
}

func (p *tException) Error() string {
	return p.message
}

func NewTException(m string) TException {
	return &tException{message: m}
}

func NewTExceptionFromOsError(e error) TException {
	if e == nil {
		return nil
	}
	t, ok := e.(TException)
	if ok {
		return t
	}
	return NewTException(e.Error())
}

const (
	UNKNOWN_APPLICATION_EXCEPTION  = 0
	UNKNOWN_METHOD                 = 1
	INVALID_MESSAGE_TYPE_EXCEPTION = 2
	WRONG_METHOD_NAME              = 3
	BAD_SEQUENCE_ID                = 4
	MISSING_RESULT                 = 5
	INTERNAL_ERROR                 = 6
	PROTOCOL_ERROR                 = 7
)

/**
 * Application level exception
 *
 */
type TApplicationException interface {
	TException
	TypeID() int32
	Read(iprot TProtocol) (TApplicationException, error)
	Write(oprot TProtocol) error
}

type tApplicationException struct {
	TException
	type_ int32
}

func NewTApplicationExceptionDefault() TApplicationException {
	return NewTApplicationException(UNKNOWN_APPLICATION_EXCEPTION, "UNKNOWN")
}

func NewTApplicationExceptionType(type_ int32) TApplicationException {
	return NewTApplicationException(type_, "UNKNOWN")
}

func NewTApplicationException(type_ int32, message string) TApplicationException {
	return &tApplicationException{TException: NewTException(message), type_: type_}
}

func NewTApplicationExceptionMessage(message string) TApplicationException {
	return NewTApplicationException(UNKNOWN_APPLICATION_EXCEPTION, message)
}

func (p *tApplicationException) TypeID() int32 {
	return p.type_
}

func (p *tApplicationException) Read(iprot TProtocol) (error TApplicationException, err error) {
	_, err = iprot.ReadStructBegin()

	// this shouldn't be needed
	er := errors.New("empty")
	if er == nil {
		return
	}
	if err != nil {
		return
	}

	message := ""
	type_ := int32(UNKNOWN_APPLICATION_EXCEPTION)

	for {
		_, ttype, id, er := iprot.ReadFieldBegin()
		if er != nil {
			return nil, er
		}
		if ttype == STOP {
			break
		}
		switch id {
		case 1:
			if ttype == STRING {
				message, err = iprot.ReadString()
				if err != nil {
					return
				}
			} else {
				err = SkipDefaultDepth(iprot, ttype)
				if err != nil {
					return
				}
			}
			break
		case 2:
			if ttype == I32 {
				type_, err = iprot.ReadI32()
				if err != nil {
					return
				}
			} else {
				err = SkipDefaultDepth(iprot, ttype)
				if err != nil {
					return
				}
			}
			break
		default:
			err = SkipDefaultDepth(iprot, ttype)
			if err != nil {
				return
			}
			break
		}
		err = iprot.ReadFieldEnd()
		if err != nil {
			return
		}
	}
	err = iprot.ReadStructEnd()
	error = NewTApplicationException(type_, message)
	return
}

func (p *tApplicationException) Write(oprot TProtocol) (err error) {
	err = oprot.WriteStructBegin("TApplicationException")
	if len(p.Error()) > 0 {
		err = oprot.WriteFieldBegin("message", STRING, 1)
		if err != nil {
			return
		}
		err = oprot.WriteString(p.Error())
		if err != nil {
			return
		}
		err = oprot.WriteFieldEnd()
		if err != nil {
			return
		}
	}
	err = oprot.WriteFieldBegin("type", I32, 2)
	if err != nil {
		return
	}
	err = oprot.WriteI32(p.type_)
	if err != nil {
		return
	}
	err = oprot.WriteFieldEnd()
	if err != nil {
		return
	}
	err = oprot.WriteFieldStop()
	if err != nil {
		return
	}
	err = oprot.WriteStructEnd()
	return
}

// TProtocolException TProtocolException
type TProtocolException interface {
	TException
	TypeID() int
}

// exception type
const (
	UNKNOWN_PROTOCOL_EXCEPTION = 0
	INVALID_DATA               = 1
	NEGATIVE_SIZE              = 2
	SIZE_LIMIT                 = 3
	BAD_VERSION                = 4
	NOT_IMPLEMENTED            = 5
)

type tProtocolException struct {
	typeID  int
	message string
}

func (p *tProtocolException) TypeID() int {
	return p.typeID
}

func (p *tProtocolException) String() string {
	return p.message
}

func (p *tProtocolException) Error() string {
	return p.message
}

func NewTProtocolExceptionDefault() TProtocolException {
	return NewTProtocolExceptionDefaultType(UNKNOWN_PROTOCOL_EXCEPTION)
}

func NewTProtocolExceptionDefaultType(t int) TProtocolException {
	return NewTProtocolException(t, "")
}

func NewTProtocolExceptionDefaultString(m string) TProtocolException {
	return NewTProtocolException(UNKNOWN_PROTOCOL_EXCEPTION, m)
}

func NewTProtocolException(t int, m string) TProtocolException {
	return &tProtocolException{typeID: t, message: m}
}

func NewTProtocolExceptionReadField(fieldId int, fieldName string, structName string, e TProtocolException) TProtocolException {
	t := e.TypeID()
	if t == UNKNOWN_PROTOCOL_EXCEPTION {
		t = INVALID_DATA
	}
	return NewTProtocolException(t, "Unable to read field "+string(fieldId)+" ("+fieldName+") in "+structName+" due to: "+e.Error())
}

func NewTProtocolExceptionWriteField(fieldId int, fieldName string, structName string, e TProtocolException) TProtocolException {
	t := e.TypeID()
	if t == UNKNOWN_PROTOCOL_EXCEPTION {
		t = INVALID_DATA
	}
	return NewTProtocolException(t, "Unable to write field "+string(fieldId)+" ("+fieldName+") in "+structName+" due to: "+e.Error())
}

func NewTProtocolExceptionReadStruct(structName string, e TProtocolException) TProtocolException {
	t := e.TypeID()
	if t == UNKNOWN_PROTOCOL_EXCEPTION {
		t = INVALID_DATA
	}
	return NewTProtocolException(t, "Unable to read struct "+structName+" due to: "+e.Error())
}

func NewTProtocolExceptionWriteStruct(structName string, e TProtocolException) TProtocolException {
	t := e.TypeID()
	if t == UNKNOWN_PROTOCOL_EXCEPTION {
		t = INVALID_DATA
	}
	return NewTProtocolException(t, "Unable to write struct "+structName+" due to: "+e.Error())
}

// NewTProtocolExceptionFromOsError NewTProtocolExceptionFromOsError
func NewTProtocolExceptionFromOsError(e error) TProtocolException {
	if e == nil {
		return nil
	}
	if t, ok := e.(TProtocolException); ok {
		return t
	}
	if te, ok := e.(TTransportException); ok {
		return NewTProtocolExceptionFromTransportException(te)
	}
	if _, ok := e.(base64.CorruptInputError); ok {
		return NewTProtocolException(INVALID_DATA, e.Error())
	}
	return NewTProtocolExceptionDefaultString(e.Error())
}

func NewTProtocolExceptionFromTransportException(e TTransportException) TProtocolException {
	if e == nil {
		return nil
	}
	if t, ok := e.(TProtocolException); ok {
		return t
	}
	return NewTProtocolExceptionDefaultString(e.Error())
}
