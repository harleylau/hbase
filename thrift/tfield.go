package thrift

import (
	"sort"
)

// TField Helper class that encapsulates field metadata.
type TField interface {
	Name() string
	TypeID() TType
	ID() int
	String() string
}

type tField struct {
	name   string
	typeID TType
	id     int
}

// ANONYMOUS FIELD
var (
	ANONYMOUSFIELD TField
)

func init() {
	ANONYMOUSFIELD = NewTField("", STOP, 0)
}

// NewTField NewTField
func NewTField(n string, t TType, i int) TField {
	return &tField{name: n, typeID: t, id: i}
}

func (p *tField) Name() string {
	return p.name
}

func (p *tField) TypeID() TType {
	return p.typeID
}

func (p *tField) ID() int {
	return p.id
}

func (p *tField) String() string {
	return "<TField name:'" + p.name + "' type:" + string(p.typeID) + " field-id:" + string(p.id) + ">"
}

type tFieldArray []TField

func (p tFieldArray) Len() int {
	return len(p)
}

func (p tFieldArray) Less(i, j int) bool {
	return p[i].ID() < p[j].ID()
}

func (p tFieldArray) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// TFieldContainer TFieldContainer
type TFieldContainer interface {
	FieldNameFromFieldID(id int) string
	FieldIDFromFieldName(name string) int
	FieldFromFieldID(id int) TField
	FieldFromFieldName(name string) TField
	At(i int) TField
}

type tFieldContainer struct {
	fields         []TField
	nameToFieldMap map[string]TField
	idToFieldMap   map[int]TField
}

// NewTFieldContainer NewTFieldContainer
func NewTFieldContainer(fields []TField) TFieldContainer {
	var (
		sortedFields   = make([]TField, len(fields))
		nameToFieldMap = make(map[string]TField)
		idToFieldMap   = make(map[int]TField)
	)
	for i, field := range fields {
		sortedFields[i] = field
		idToFieldMap[field.ID()] = field
		if field.Name() != "" {
			nameToFieldMap[field.Name()] = field
		}
	}
	sort.Sort(tFieldArray(sortedFields))
	return &tFieldContainer{
		fields:         fields,
		nameToFieldMap: nameToFieldMap,
		idToFieldMap:   idToFieldMap,
	}
}

func (p *tFieldContainer) FieldNameFromFieldID(id int) string {
	if field, ok := p.idToFieldMap[id]; ok {
		return field.Name()
	}
	return ""
}

func (p *tFieldContainer) FieldIDFromFieldName(name string) int {
	if field, ok := p.nameToFieldMap[name]; ok {
		return field.ID()
	}
	return -1
}

func (p *tFieldContainer) FieldFromFieldID(id int) TField {
	if field, ok := p.idToFieldMap[id]; ok {
		return field
	}
	return ANONYMOUSFIELD
}

func (p *tFieldContainer) FieldFromFieldName(name string) TField {
	if field, ok := p.nameToFieldMap[name]; ok {
		return field
	}
	return ANONYMOUSFIELD
}

func (p *tFieldContainer) Len() int {
	return len(p.fields)
}

func (p *tFieldContainer) At(i int) TField {
	return p.FieldFromFieldID(i)
}

func (p *tFieldContainer) iterate(c chan<- TField) {
	for _, v := range p.fields {
		c <- v
	}
	close(c)
}
