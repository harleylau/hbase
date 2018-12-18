package thrift

// TMessage Helper class that encapsulates struct metadata.
type TMessage interface {
	Name() string
	TypeID() TMessageType
	SeqID() int
	Equals(other TMessage) bool
}
type tMessage struct {
	name   string
	typeID TMessageType
	seqid  int
}

// EMPTY MESSAGE
var (
	EMPTYMESSAGE TMessage
)

func init() {
	EMPTYMESSAGE = NewTMessage("", STOP, 0)
}

// NewTMessage NewTMessage
func NewTMessage(n string, t TMessageType, s int) TMessage {
	return &tMessage{name: n, typeID: t, seqid: s}
}

func (p *tMessage) Name() string {
	return p.name
}

func (p *tMessage) TypeID() TMessageType {
	return p.typeID
}

func (p *tMessage) SeqID() int {
	return p.seqid
}

func (p *tMessage) String() string {
	return "<TMessage name:'" + p.name + "' type: " + string(p.typeID) + " seqid:" + string(p.seqid) + ">"
}

func (p *tMessage) Equals(other TMessage) bool {
	return p.name == other.Name() && p.typeID == other.TypeID() && p.seqid == other.SeqID()
}

// TMessageType TMessageType
type TMessageType int32

// TMessageType
const (
	INVALID   TMessageType = 0
	CALL      TMessageType = 1
	REPLY     TMessageType = 2
	EXCEPTION TMessageType = 3
	ONEWAY    TMessageType = 4
)
