package ge2d

import "fmt"

type SceneManager struct {
	objectMap map[uint]*Object
}

func NewSceneManager() *SceneManager {
	return &SceneManager {make(map[uint]*Object)}
}

func (this *SceneManager) AddObject(object *Object) {
	this.objectMap[object.GetId()] = object
}

func (this *SceneManager) HandleMessage(message IMessage) {
	fmt.Print("[SceneManager] handleMessage: ", message)
	fmt.Print("[SceneManager] destination object: ", message.GetDestinationObjectId())

	this.objectMap[message.GetDestinationObjectId()].HandleMessage(message)
}
