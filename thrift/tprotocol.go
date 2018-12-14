package thrift

import "encoding/base64"

// VERSION
const (
	VERSION_MASK = 0xffff0000
	VERSION_1    = 0x80010000
)

// TProtocolFactory factory
type TProtocolFactory interface {
	GetProtocol(trans TTransport) TProtocol
}

// TProtocol thrift protocol
type TProtocol interface {
	WriteMessageBegin(name string, typeId TMessageType, seqid int32) TProtocolException
	WriteMessageEnd() TProtocolException
	WriteStructBegin(name string) TProtocolException
	WriteStructEnd() TProtocolException
	WriteFieldBegin(name string, typeId TType, id int16) TProtocolException
	WriteFieldEnd() TProtocolException
	WriteFieldStop() TProtocolException
	WriteMapBegin(keyType TType, valueType TType, size int) TProtocolException
	WriteMapEnd() TProtocolException
	WriteListBegin(elemType TType, size int) TProtocolException
	WriteListEnd() TProtocolException
	WriteSetBegin(elemType TType, size int) TProtocolException
	WriteSetEnd() TProtocolException
	WriteBool(value bool) TProtocolException
	WriteByte(value int8) TProtocolException
	WriteI16(value int16) TProtocolException
	WriteI32(value int32) TProtocolException
	WriteI64(value int64) TProtocolException
	WriteDouble(value float64) TProtocolException
	WriteString(value string) TProtocolException
	WriteBinary(value []byte) TProtocolException

	ReadMessageBegin() (name string, typeId TMessageType, seqid int32, err TProtocolException)
	ReadMessageEnd() TProtocolException
	ReadStructBegin() (name string, err TProtocolException)
	ReadStructEnd() TProtocolException
	ReadFieldBegin() (name string, typeId TType, id int16, err TProtocolException)
	ReadFieldEnd() TProtocolException
	ReadMapBegin() (keyType TType, valueType TType, size int, err TProtocolException)
	ReadMapEnd() TProtocolException
	ReadListBegin() (elemType TType, size int, err TProtocolException)
	ReadListEnd() TProtocolException
	ReadSetBegin() (elemType TType, size int, err TProtocolException)
	ReadSetEnd() TProtocolException
	ReadBool() (value bool, err TProtocolException)
	ReadByte() (value int8, err TProtocolException)
	ReadI16() (value int16, err TProtocolException)
	ReadI32() (value int32, err TProtocolException)
	ReadI64() (value int64, err TProtocolException)
	ReadDouble() (value float64, err TProtocolException)
	ReadString() (value string, err TProtocolException)
	ReadBinary() (value []byte, err TProtocolException)

	Skip(fieldType TType) (err TProtocolException)
	Flush() (err TProtocolException)

	Transport() TTransport
}

/**
 * The maximum recursive depth the skip() function will traverse before
 * throwing a TException.
 */
var (
	MaxSkipDepth = 1<<31 - 1
)

/**
 * Specifies the maximum recursive depth that the skip function will
 * traverse before throwing a TException.  This is a global setting, so
 * any call to skip in this JVM will enforce this value.
 *
 * @param depth  the maximum recursive depth.  A value of 2 would allow
 *    the skip function to skip a structure or collection with basic children,
 *    but it would not permit skipping a struct that had a field containing
 *    a child struct.  A value of 1 would only allow skipping of simple
 *    types and empty structs/collections.
 */
func SetMaxSkipDepth(depth int) {
	MaxSkipDepth = depth
}

/**
 * Skips over the next data element from the provided input TProtocol object.
 *
 * @param prot  the protocol object to read from
 * @param type  the next value will be intepreted as this TType value.
 */
func SkipDefaultDepth(prot TProtocol, typeId TType) (err TProtocolException) {
	return Skip(prot, typeId, MaxSkipDepth)
}

/**
 * Skips over the next data element from the provided input TProtocol object.
 *
 * @param prot  the protocol object to read from
 * @param type  the next value will be intepreted as this TType value.
 * @param maxDepth  this function will only skip complex objects to this
 *   recursive depth, to prevent Java stack overflow.
 */
func Skip(self TProtocol, fieldType TType, maxDepth int) (err TProtocolException) {
	switch fieldType {
	case STOP:
		return
	case BOOL:
		_, err = self.ReadBool()
		return
	case BYTE:
		_, err = self.ReadByte()
		return
	case I16:
		_, err = self.ReadI16()
		return
	case I32:
		_, err = self.ReadI32()
		return
	case I64:
		_, err = self.ReadI64()
		return
	case DOUBLE:
		_, err = self.ReadDouble()
		return
	case STRING:
		_, err = self.ReadString()
		return
	case STRUCT:
		{
			_, err = self.ReadStructBegin()
			if err != nil {
				return
			}
			for {
				_, typeId, _, _ := self.ReadFieldBegin()
				if typeId == STOP {
					break
				}
				Skip(self, typeId, maxDepth-1)
				self.ReadFieldEnd()
			}
			return self.ReadStructEnd()
		}
	case MAP:
		{
			keyType, valueType, l, err := self.ReadMapBegin()
			if err != nil {
				return err
			}
			size := int(l)
			for i := 0; i < size; i++ {
				Skip(self, keyType, maxDepth-1)
				self.Skip(valueType)
			}
			return self.ReadMapEnd()
		}
	case SET:
		{
			elemType, l, err := self.ReadSetBegin()
			if err != nil {
				return err
			}
			size := int(l)
			for i := 0; i < size; i++ {
				Skip(self, elemType, maxDepth-1)
			}
			return self.ReadSetEnd()
		}
	case LIST:
		{
			elemType, l, err := self.ReadListBegin()
			if err != nil {
				return err
			}
			size := int(l)
			for i := 0; i < size; i++ {
				Skip(self, elemType, maxDepth-1)
			}
			return self.ReadListEnd()
		}
	}
	return nil
}

// TProtocolException
type TProtocolException interface {
	TException
	TypeId() int
}

const (
	UNKNOWN_PROTOCOL_EXCEPTION = 0
	INVALID_DATA               = 1
	NEGATIVE_SIZE              = 2
	SIZE_LIMIT                 = 3
	BAD_VERSION                = 4
	NOT_IMPLEMENTED            = 5
)

type tProtocolException struct {
	typeId  int
	message string
}

func (p *tProtocolException) TypeId() int {
	return p.typeId
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
	return &tProtocolException{typeId: t, message: m}
}

func NewTProtocolExceptionReadField(fieldId int, fieldName string, structName string, e TProtocolException) TProtocolException {
	t := e.TypeId()
	if t == UNKNOWN_PROTOCOL_EXCEPTION {
		t = INVALID_DATA
	}
	return NewTProtocolException(t, "Unable to read field "+string(fieldId)+" ("+fieldName+") in "+structName+" due to: "+e.Error())
}

func NewTProtocolExceptionWriteField(fieldId int, fieldName string, structName string, e TProtocolException) TProtocolException {
	t := e.TypeId()
	if t == UNKNOWN_PROTOCOL_EXCEPTION {
		t = INVALID_DATA
	}
	return NewTProtocolException(t, "Unable to write field "+string(fieldId)+" ("+fieldName+") in "+structName+" due to: "+e.Error())
}

func NewTProtocolExceptionReadStruct(structName string, e TProtocolException) TProtocolException {
	t := e.TypeId()
	if t == UNKNOWN_PROTOCOL_EXCEPTION {
		t = INVALID_DATA
	}
	return NewTProtocolException(t, "Unable to read struct "+structName+" due to: "+e.Error())
}

func NewTProtocolExceptionWriteStruct(structName string, e TProtocolException) TProtocolException {
	t := e.TypeId()
	if t == UNKNOWN_PROTOCOL_EXCEPTION {
		t = INVALID_DATA
	}
	return NewTProtocolException(t, "Unable to write struct "+structName+" due to: "+e.Error())
}

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
