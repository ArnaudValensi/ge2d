package ge2d

import "fmt"

type IComponent interface {
	HandleMessage(message IMessage)
	GetTypeName() string
}

type BaseComponent struct {
	typeName string
}

func (this *BaseComponent) HandleMessage(message IMessage) {
	fmt.Print("[BaseComponent] handleMessage")
}

func (this *BaseComponent) GetTypeName() string {
	return this.typeName
}

// func Test() {
// 	o := BaseComponent{}
// 	p := RenderComponent{}
// 	o = p
// }