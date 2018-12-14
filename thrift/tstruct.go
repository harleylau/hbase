package thrift

/**
 * Helper class that encapsulates struct metadata.
 *
 */
type TStruct interface {
	TFieldContainer
	TStructName() string
	ThriftName() string
	TStructFields() TFieldContainer
	String() string
	AttributeFromFieldId(fieldId int) interface{}
	AttributeFromFieldName(fieldName string) interface{}
}

type tStruct struct {
	TFieldContainer
	name string
}

func NewTStructEmpty(name string) TStruct {
	return &tStruct{
		name:            name,
		TFieldContainer: NewTFieldContainer(make([]TField, 0, 0)),
	}
}

func NewTStruct(name string, fields []TField) TStruct {
	return &tStruct{
		name:            name,
		TFieldContainer: NewTFieldContainer(fields),
	}
}

func (p *tStruct) TStructName() string {
	return p.name
}

func (p *tStruct) ThriftName() string {
	return p.name
}

func (p *tStruct) TStructFields() TFieldContainer {
	return p.TFieldContainer
}

func (p *tStruct) String() string {
	return p.name
}

func (p *tStruct) AttributeFromFieldId(fieldId int) interface{} {
	return nil
}

func (p *tStruct) AttributeFromFieldName(fieldName string) interface{} {
	return p.AttributeFromFieldId(p.FieldIdFromFieldName(fieldName))
}

var ANONYMOUS_STRUCT TStruct

func init() {
	ANONYMOUS_STRUCT = NewTStructEmpty("")
}
