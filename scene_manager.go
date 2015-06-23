package ge2d

import (
	"fmt"
	"log"
)

type SceneManager struct {
	objectMap	map[uint]*Object
	rootNode	INode
}

func NewSceneManager() *SceneManager {
	this := &SceneManager {make(map[uint]*Object), nil}
	this.rootNode = NewBaseNode("root", nil, this, Vector2d {0, 0})
	return this
}

// Used by node classes to register an object to message system
func (this *SceneManager) AddObject(object *Object) {
	if _, exist := this.objectMap[object.GetId()]; exist {
		log.Printf(
			"[Warning] [SceneManager] AddObject(): object(id: %d) %s",
			object.GetId(),
			"already attached\n")
	}
	this.objectMap[object.GetId()] = object
}

func (this *SceneManager) GetRootNode() INode {
	return this.rootNode
}

func (this *SceneManager) HandleMessage(message IMessage) {
	fmt.Println("[SceneManager] handleMessage: ", message)
	fmt.Println("[SceneManager] destination object: ", message.GetDestinationObjectId())

	if object, exist := this.objectMap[message.GetDestinationObjectId()]; exist {
		object.HandleMessage(message)
		return
	}
	log.Fatal("[SceneManager] HandleMessage(): Unknown destination object id")
}

// Load tile map in tmx format. You can use http://www.mapeditor.org/ 
func (this *SceneManager) LoadTileMap(file string) {
	
}
