package ge2d

import (
	"container/list"
)

type INode interface {
	GetParentNode() INode
	SetName(name string)
	GetName() string
	GetId() uint
	CreateNamedChild(name string, relativePosition Vector2d)
	CreateChild(relativePosition Vector2d)
	addChild(child INode)
	RemoveNamedChild(name string)
	RemoveChild(child INode)
	// AttachObject(object *Object)
	// DetachObject(object *Object)
	// DetachObjectById(id uint)
	SetPosition(position Vector2d)
	GetPosition() Vector2d
	AddComponent(component *BaseComponent)
	HandleMessage(message IMessage)
}

type BaseNode struct {
	id uint
	name string
	parent INode
	position Vector2d
	nodeMap map[uint]INode
	componentList *list.List
	// objectMap map[uint]*Object
}

var lastNodeId uint = 0

// position is relative to the parent
func NewBaseNode(name string, parent INode, position Vector2d) *BaseNode {
	lastNodeId++
	return &BaseNode {
		lastNodeId,
		name, 
		nil, 
		position, 
		make(map[uint]INode), 
		list.New()}
		// make(map[uint]*Object)}
}

func (this *BaseNode) GetParentNode() INode {
	return this.parent
}

func (this *BaseNode) SetName(name string) {
	this.name = name
}

func (this *BaseNode) GetName() string {
	return this.name
}

func (this *BaseNode) GetId() uint {
	return this.id
}

func (this *BaseNode) CreateNamedChild(name string, relativePosition Vector2d) {
	newNode := NewBaseNode(name, this, relativePosition)
	this.nodeMap[newNode.GetId()] = newNode
}

func (this *BaseNode) CreateChild(relativePosition Vector2d) {
	this.CreateNamedChild("", relativePosition)
}

func (this *BaseNode) addChild(child INode) {
	this.nodeMap[child.GetId()] = child
}

func (this *BaseNode) RemoveNamedChild(name string) {
	for key, value := range this.nodeMap {
		if value.GetName() == name {
			delete(this.nodeMap, key)
		}
	}
}

func (this *BaseNode) RemoveChild(child INode) {
	for key, value := range this.nodeMap {
		if value == child {
			delete(this.nodeMap, key)
		}
	}
}

// func (this *BaseNode) AttachObject(object *Object) {
// 	this.objectMap[object.GetId()] = object
// }

// func (this *BaseNode) DetachObject(object *Object) {
// 	for key, value := range this.objectMap {
// 		if value == object {
// 			delete(this.objectMap, key)
// 		}
// 	}
// }

// func (this *BaseNode) DetachObjectById(id uint) {
// 	delete(this.objectMap, id)
// }

func (this *BaseNode) SetPosition(position Vector2d) {
	this.position = position
}

func (this *BaseNode) GetPosition() Vector2d {
	return this.position
}

func (this *BaseNode) AddComponent(component *BaseComponent) {
	this.componentList.PushBack(component)
}

func (this *BaseNode) HandleMessage(message IMessage) {
	for e := this.componentList.Front(); e != nil; e = e.Next() {
		e.Value.(IComponent).HandleMessage(message)
	}
}