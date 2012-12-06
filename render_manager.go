package ge2d

import (
	"fmt"
	"github.com/0xe2-0x9a-0x9b/Go-SDL/sdl"
	"log"
	//"math"
	//"os"
	//"strings"
	//"time"
)

type RenderManager struct {
	screen		*sdl.Surface
	image		*sdl.Surface
	animMap		map[string]*Anim
	// spriteSlice	[]*sdl.Surface
}

func NewRenderManager() *RenderManager {
	this := &RenderManager {nil, nil, make(map[string]*Anim)}
	this.Init()
	return this
}

func (this *RenderManager) Init() {
	if sdl.Init(sdl.INIT_EVERYTHING) != 0 {
		log.Fatal(sdl.GetError())
	}

	this.screen = sdl.SetVideoMode(640, 480, 32, sdl.RESIZABLE)

	if this.screen == nil {
		log.Fatal(sdl.GetError())
	}

	var video_info = sdl.GetVideoInfo()

	println("HW_available = ", video_info.HW_available)
	println("WM_available = ", video_info.WM_available)
	println("Video_mem = ", video_info.Video_mem, "kb")

	sdl.EnableUNICODE(1)

	sdl.WM_SetCaption("Go-SDL SDL Test", "")

	// image := sdl.Load(resourcePath + "/test.png")

	// if image == nil {
	// 	log.Fatal(sdl.GetError())
	// }

	// sdl.WM_SetIcon(image, nil)

	// running := true

	if sdl.GetKeyName(270) != "[+]" {
		log.Fatal("GetKeyName broken")
	}

	spriteLoader := NewSpriteLoader()
	spriteLoader.Load("test.gspr")
	this.animMap = spriteLoader.GetAnimMap()

	// DEBUG
	// this.DebugRessources()
}

// TODO: Sprite are not free (sdl)
func (this *RenderManager) Quit() {
	// image.Free()
	sdl.Quit()
}

var i int16 = 0
func (this *RenderManager) Update(sceneManager *SceneManager) {
	this.screen.FillRect(nil, 0x302019)
	this.BrowseNode(sceneManager.GetRootNode())
	this.screen.Flip()

	// this.screen.FillRect(nil, 0x302019)
	// this.screen.Blit(&sdl.Rect{i, 0, 0, 0}, this.image, nil)
	// i++
	// this.screen.Flip()
}

// func (this *RenderManager) DebugRessources() {
// 	this.image = sdl.Load(resourcePath + "/test.png")

// 	if this.image == nil {
// 		log.Fatal(sdl.GetError())
// 	}

// 	// sdl.WM_SetIcon(image, nil)
// }

func (this *RenderManager) CreateRenderComponent(sprite string) *RenderComponent {
	if _, exist := this.animMap[sprite]; !exist {
		log.Fatal("[RenderManager] CreateRenderComponent(): Unknown sprite")
	}
	renderComponent := NewRenderComponent(this, sprite)
	return renderComponent
}

func (this *RenderManager) BrowseNode(node INode) {
	objectMap := node.GetObjectMap()
	for _, object := range objectMap {
		renderComponent := 
			object.GetComposantByTypeName("render").(*RenderComponent)
		if renderComponent != nil {
			this.Blit(renderComponent, node)
			// this.Blit(renderComponent.GetNextImage(), node.GetPosition())
		}
	}
	for _, node := range node.GetChildMap() {
		this.BrowseNode(node)
	}
}

func (this *RenderManager) Blit(component *RenderComponent, node INode) {
	anim := this.animMap[component.GetAnimation()]
	imageCount := component.GetImageCount()

	log.Printf("imageCount: %d, image: %d\n", *imageCount, *imageCount / anim.GetFrequency())

	// TODO: handle err
	image, srcrect, err := anim.GetSprite2(*imageCount / anim.GetFrequency())
	if err != nil {
		log.Fatal("[RenderManager] Blit(): GetSprite2 error")
	}
	nbImage := anim.GetNumberSprite()
	*imageCount = (*imageCount + 1) % (nbImage * anim.GetFrequency())

	position := node.GetPosition()

	if image == nil {
		log.Fatal("[RenderManager] Blit(): nil surface")
	}
	
	// this.screen.Blit(
	// 	&sdl.Rect{
	// 	int16(position.X), int16(position.Y), 
	// 	uint16(image.W), uint16(image.H)},
	// 	image, nil)

	fmt.Printf("size: %d, %d -- pos: %d, %d\n", 
		srcrect.W, srcrect.H, srcrect.X, srcrect.Y)

	this.screen.Blit(
		&sdl.Rect{
		int16(position.X), int16(position.Y), 
		uint16(image.W), uint16(image.H)},
		image,
		srcrect,
	)

}

func (this *RenderManager) LoadSprite() {
	
}