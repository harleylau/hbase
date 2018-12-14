package thrift

// TType type
type TType int8

// TType
const (
	STOP   = 0
	VOID   = 1
	BOOL   = 2
	BYTE   = 3
	DOUBLE = 4
	I16    = 6
	I32    = 8
	I64    = 10
	STRING = 11
	STRUCT = 12
	MAP    = 13
	SET    = 14
	LIST   = 15
	ENUM   = 16
)

// ThriftTypeID type id
func (p TType) ThriftTypeID() int8 {
	return int8(p)
}

func (p TType) String() string {
	switch p {
	case STOP:
		return "STOP"
	case VOID:
		return "VOID"
	case BOOL:
		return "BOOL"
	case BYTE:
		return "BYTE"
	case DOUBLE:
		return "DOUBLE"
	case I16:
		return "I16"
	case I32:
		return "I32"
	case I64:
		return "I64"
	case STRING:
		return "STRING"
	case STRUCT:
		return "STRUCT"
	case MAP:
		return "MAP"
	case SET:
		return "SET"
	case LIST:
		return "LIST"
	case ENUM:
		return "ENUM"
	}
	return "Unknown"
}
