package ge2d

import "log"

type INode interface {
	GetParentNode() INode
	SetName(name string)
	GetName() string
	GetId() uint
	CreateNamedChild(name string, relativePosition Vector2d)
	CreateChild(relativePosition Vector2d)
	AddChild(child INode)
	RemoveNamedChild(name string)
	RemoveChild(child INode)
	SetPosition(position Vector2d)
	GetPosition() Vector2d
	AttachObject(object *Object)
	DetachObject(object *Object)
	DetachObjectById(id uint)
	GetChildMap() map[uint]INode
	GetObjectMap() map[uint]*Object
}

type BaseNode struct {
	id uint
	name string
	parent INode
	sceneManager *SceneManager
	position Vector2d
	childMap map[uint]INode
	objectMap map[uint]*Object
}

var lastNodeId uint = 0

// position is relative to the parent
func NewBaseNode(
	name string, 
	parent INode, 
	sceneManager *SceneManager, 
	position Vector2d) *BaseNode {

	lastNodeId++
	return &BaseNode {
		lastNodeId,
		name, 
		parent,
		sceneManager,
		position, 
		make(map[uint]INode), 
		make(map[uint]*Object)}
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
	newNode := NewBaseNode(name, this, this.sceneManager, relativePosition)
	this.childMap[newNode.GetId()] = newNode
}

func (this *BaseNode) CreateChild(relativePosition Vector2d) {
	this.CreateNamedChild("", relativePosition)
}

func (this *BaseNode) AddChild(child INode) {
	this.childMap[child.GetId()] = child
}

func (this *BaseNode) RemoveNamedChild(name string) {
	for key, value := range this.childMap {
		if value.GetName() == name {
			delete(this.childMap, key)
		}
	}
}

func (this *BaseNode) RemoveChild(child INode) {
	for key, value := range this.childMap {
		if value == child {
			delete(this.childMap, key)
		}
	}
}

func (this *BaseNode) SetPosition(position Vector2d) {
	this.position = position
}

func (this *BaseNode) GetPosition() Vector2d {
	return this.position
}

func (this *BaseNode) AttachObject(object *Object) {
	if _, exist := this.objectMap[object.GetId()]; exist {
		log.Printf(
			"[Warning] [BaseNode] AttachObject(): object(id: %d) %s",
			object.GetId(),
			"already attached to node\n")
	}
	this.objectMap[object.GetId()] = object
	this.sceneManager.AddObject(object)
}

func (this *BaseNode) DetachObject(object *Object) {
	for key, value := range this.objectMap {
		if value == object {
			delete(this.objectMap, key)
		}
	}
}

func (this *BaseNode) DetachObjectById(id uint) {
	delete(this.objectMap, id)
}


func (this *BaseNode) GetChildMap() map[uint]INode {
	return this.childMap
}

func (this *BaseNode) GetObjectMap() map[uint]*Object {
	return this.objectMap
}
