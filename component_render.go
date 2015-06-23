package ge2d

import (
	"fmt"
)

type RenderComponent struct {
	BaseComponent
	renderManager	*RenderManager
	// animation	string
	animation	*Anim
	imageCount	uint
}

func NewRenderComponent(renderManager *RenderManager, animation *Anim) *RenderComponent {
	return &RenderComponent {
		BaseComponent {"render"}, 
		renderManager, 
		animation, 
		0}
}

func (this *RenderComponent) GetAnimation() *Anim {
	return this.animation
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
