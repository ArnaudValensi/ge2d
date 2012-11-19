package ge2d

type MessageType uint

const (
	MSG_NONE		MessageType = iota
	MSG_SET_POSITION
	MSG_GET_POSITION
)

type IMessage interface {
	GetDestinationObjectId() uint
	setDestinationObjectId(id uint)
	getMessageType() MessageType
}

type BaseMessage struct {
	destinationObjectId uint
	messageType MessageType
}

func NewBaseMessage(destinationObjectId uint, messageType MessageType) *BaseMessage {
	return &BaseMessage {destinationObjectId, messageType}
}

func (this *BaseMessage) GetDestinationObjectId() uint {
	return this.destinationObjectId
}

func (this *BaseMessage) setDestinationObjectId(id uint) {
	this.destinationObjectId = id
}

func (this *BaseMessage) getMessageType() MessageType {
	return this.messageType
}
