package thrift

/**
 * Helper class that encapsulates struct metadata.
 *
 */
type TMessage interface {
	Name() string
	TypeId() TMessageType
	SeqId() int
	Equals(other TMessage) bool
}
type tMessage struct {
	name   string
	typeId TMessageType
	seqid  int
}

func NewTMessageDefault() TMessage {
	return NewTMessage("", STOP, 0)
}

func NewTMessage(n string, t TMessageType, s int) TMessage {
	return &tMessage{name: n, typeId: t, seqid: s}
}

func (p *tMessage) Name() string {
	return p.name
}

func (p *tMessage) TypeId() TMessageType {
	return p.typeId
}

func (p *tMessage) SeqId() int {
	return p.seqid
}

func (p *tMessage) String() string {
	return "<TMessage name:'" + p.name + "' type: " + string(p.typeId) + " seqid:" + string(p.seqid) + ">"
}

func (p *tMessage) Equals(other TMessage) bool {
	return p.name == other.Name() && p.typeId == other.TypeId() && p.seqid == other.SeqId()
}

var EMPTY_MESSAGE TMessage

func init() {
	EMPTY_MESSAGE = NewTMessageDefault()
}

type TMessageType int32

const (
	INVALID_TMESSAGE_TYPE TMessageType = 0
	CALL                  TMessageType = 1
	REPLY                 TMessageType = 2
	EXCEPTION             TMessageType = 3
	ONEWAY                TMessageType = 4
)