package ge2d

import "fmt"

type RenderComponent struct {

}

func (this *RenderComponent) handleMessage() {
	fmt.Print("[RenderComponent] handleMessage")
}
