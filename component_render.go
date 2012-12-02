package ge2d

import (
	"fmt"
)

type RenderComponent struct {
<<<<<<< HEAD
	
=======
	BaseComponent
	renderManager	*RenderManager
	animation	string
	imageCount	uint
}

func NewRenderComponent(renderManager *RenderManager, animation string) *RenderComponent {
	return &RenderComponent {
		BaseComponent {"render"}, 
		renderManager, 
		animation, 
		0}
}

func (this *RenderComponent) GetAnimation() string {
	return this.animation
>>>>>>> obj
}

func (this *RenderComponent) HandleMessage(message IMessage) {
	fmt.Print("[RenderComponent] handleMessage")
}

func (this *RenderComponent) GetImageCount() *uint {
	return &this.imageCount
}

func (this *RenderComponent) SetImageCount(number uint) {
	this.imageCount = number
}

// func (this *RenderComponent) GetNextImage() *sdl.Surface {
// 	return this.animation.GetNextImage()
// }
