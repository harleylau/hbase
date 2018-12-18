package thrift

// VERSION
const (
	VERSIONMASK = 0xffff0000
	VERSION1    = 0x80010000
)

// TProtocolFactory factory
type TProtocolFactory interface {
	GetProtocol(trans TTransport) TProtocol
}

// TProtocol thrift protocol
type TProtocol interface {
	WriteMessageBegin(name string, typeID TMessageType, seqid int32) TProtocolException
	WriteMessageEnd() TProtocolException
	WriteStructBegin(name string) TProtocolException
	WriteStructEnd() TProtocolException
	WriteFieldBegin(name string, typeID TType, id int16) TProtocolException
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

	ReadMessageBegin() (name string, typeID TMessageType, seqid int32, err TProtocolException)
	ReadMessageEnd() TProtocolException
	ReadStructBegin() (name string, err TProtocolException)
	ReadStructEnd() TProtocolException
	ReadFieldBegin() (name string, typeID TType, id int16, err TProtocolException)
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

// MaxSkipDepth
var (
	MaxSkipDepth = 1<<31 - 1
)

// SetMaxSkipDepth SetMaxSkipDepth
func SetMaxSkipDepth(depth int) {
	MaxSkipDepth = depth
}

// SkipDefaultDepth SkipDefaultDepth
func SkipDefaultDepth(prot TProtocol, typeID TType) (err TProtocolException) {
	return Skip(prot, typeID, MaxSkipDepth)
}

// Skip Skip
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
				_, typeID, _, _ := self.ReadFieldBegin()
				if typeID == STOP {
					break
				}
				Skip(self, typeID, maxDepth-1)
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
