package ge2d

import "fmt"

type Scene struct {
	objectMap map[uint]*Object
}

func NewScene() *Scene {
	return &Scene {make(map[uint]*Object)}
}

func (this *Scene) AddObject(object *Object) {
	this.objectMap[object.GetId()] = object
}

func (this *Scene) HandleMessage(message IMessage) {
	fmt.Print("[Scene] handleMessage: ", message)
	fmt.Print("[Scene] destination object: ", message.GetDestinationObjectId())

	this.objectMap[message.GetDestinationObjectId()].HandleMessage(message)
}
