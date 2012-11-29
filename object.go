package ge2d

import (
	"container/list"
)


type Object struct {
	id uint
	componentList *list.List
	position Vector2d
}

func NewObject(id uint) *Object {
	return &Object {
		id, 
		list.New(), 
		Vector2d {0, 0},
	}
}

func (this *Object) AddComponent(component IComponent) {
	this.componentList.PushBack(component)
}

func (this *Object) GetId() uint {
	return this.id
}

func (this *Object) HandleMessage(message IMessage) {
	for e := this.componentList.Front(); e != nil; e = e.Next() {
		e.Value.(IComponent).HandleMessage(message)
	}
}

func (this *Object) GetComposantByTypeName(typeName string) IComponent {
	for e := this.componentList.Front(); e != nil; e = e.Next() {
		if e.Value.(IComponent).GetTypeName() == typeName {
			return e.Value.(IComponent)
		}
	}
	return nil
}
