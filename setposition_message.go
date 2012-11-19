package ge2d

type SetPositionMessage struct {
	BaseMessage
	position Vector2d
}

func NewSetPositionMessage(destination uint, position Vector2d) *SetPositionMessage {
	return &SetPositionMessage {
		BaseMessage{destination, MSG_GET_POSITION}, 
		position}
}
