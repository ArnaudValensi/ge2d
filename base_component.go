package ge2d

import "fmt"

type IComponent interface {
	HandleMessage(message IMessage)
}

type BaseComponent struct {

}

func (this *BaseComponent) HandleMessage(message IMessage) {
	fmt.Print("[BaseComponent] handleMessage")
}

// func Test() {
// 	o := BaseComponent{}
// 	p := RenderComponent{}
// 	o = p
// }