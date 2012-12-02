package ge2d

import (
	"fmt"
	"log"
)

type SceneManager struct {
<<<<<<< HEAD
	rootNode	INode
	nodeMap		map[uint]INode
}

func NewSceneManager() *SceneManager {
	rootNode := NewBaseNode("root", nil, Vector2d {0, 0})
	nodeMap := make(map[uint]INode)
	nodeMap[rootNode.GetId()] = rootNode
	return &SceneManager {rootNode, nodeMap}
}

func (this *SceneManager) AddNode(node INode) {
	this.nodeMap[node.GetId()] = node
=======
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
>>>>>>> obj
}

func (this *SceneManager) GetRootNode() INode {
	return this.rootNode
}

func (this *SceneManager) HandleMessage(message IMessage) {
	fmt.Println("[SceneManager] handleMessage: ", message)
	fmt.Println("[SceneManager] destination object: ", message.GetDestinationObjectId())

<<<<<<< HEAD
	this.nodeMap[message.GetDestinationObjectId()].HandleMessage(message)
=======
	if object, exist := this.objectMap[message.GetDestinationObjectId()]; exist {
		object.HandleMessage(message)
		return
	}
	log.Fatal("[SceneManager] HandleMessage(): Unknown destination object id")
>>>>>>> obj
}

// Load tile map in tmx format. You can use http://www.mapeditor.org/ 
func (this *SceneManager) LoadTileMap(file string) {
	
}
