package ge2d

import (
	"github.com/0xe2-0x9a-0x9b/Go-SDL/sdl"
	"log"
	// "fmt"
	//"math"
	// "os"
	//"strings"
	//"time"
)

type RenderManager struct {
	screen			*sdl.Surface
	image			*sdl.Surface
	// animMap			map[string]*Anim
	tmx			*TmxTileMap
	mapSprite		*SpriteSetCollection
	// spriteSlice		[]*sdl.Surface
	backgroundColor		uint32
	ressourceManager	*SpriteLoader
}

func NewRenderManager() *RenderManager {
	mapSprite := NewSpriteSetCollection()
	this := &RenderManager {
		nil, 
		nil, 
		// make(map[string]*Anim), 
		nil, 
		mapSprite,
		0x0,
		NewSpriteLoader(),
	}
	tmx := NewTmxTileMap(this)
	this.tmx = tmx
	this.Init()
	return this
}

func (this *RenderManager) Init() {
	this.tmx.Load("./test_cute.tmx")
	if sdl.Init(sdl.INIT_EVERYTHING) != 0 {
		log.Fatal(sdl.GetError())
	}
	this.screen = sdl.SetVideoMode(707, 600, 32, sdl.RESIZABLE)
	if this.screen == nil {
		log.Fatal(sdl.GetError())
	}
	var video_info = sdl.GetVideoInfo()
	println("HW_available = ", video_info.HW_available)
	println("WM_available = ", video_info.WM_available)
	println("Video_mem = ", video_info.Video_mem, "kb")
	sdl.EnableUNICODE(1)
	sdl.WM_SetCaption("Go-SDL SDL Test", "")
	if sdl.GetKeyName(270) != "[+]" {
		log.Fatal("GetKeyName broken")
	}
	this.ressourceManager = NewSpriteLoader()
	this.ressourceManager.Load("test_tileset.gspr")
	// this.animMap = ressourceManager.GetAnimMap()
}

func (this *RenderManager) LoadMapSpriteSet(file string, elemWidth uint, elemHeight uint) {
	this.mapSprite.LoadSpriteSet(file, elemWidth, elemHeight)
}

// TODO: Sprite are not free (sdl)
func (this *RenderManager) Quit() {
	// image.Free()
	sdl.Quit()
}

var i int16 = 0
func (this *RenderManager) Update(sceneManager *SceneManager) {
	// if this.backgroundColor != 0 {
	this.screen.FillRect(nil, this.backgroundColor)
	// }
	// this.screen.FillRect(nil, 0x302019)
	this.BlitMap()
	this.BrowseNode(sceneManager.GetRootNode())
	this.screen.Flip()

	// this.screen.FillRect(nil, 0x302019)
	// this.screen.Blit(&sdl.Rect{i, 0, 0, 0}, this.image, nil)
	// i++
	// this.screen.Flip()
}

func (this *RenderManager) CreateRenderComponent(animName string) *RenderComponent {
	anim := this.ressourceManager.GetAnim(animName)
	renderComponent := NewRenderComponent(this, anim)
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

func (this *RenderManager) BlitMap() {
	this.tmx.BlitMap()
}

func (this *RenderManager) Blit(component *RenderComponent, node INode) {
	// anim := this.animMap[component.GetAnimation()]
	// imageCount := component.GetImageCount()
	
	anim := component.GetAnimation()
	imageCount := component.GetImageCount()

	log.Printf("imageCount: %d, image: %d\n", *imageCount, *imageCount / anim.GetFrequency())

	ressource := anim.GetResource(*imageCount / anim.GetFrequency())
	image, srcrect, err := this.ressourceManager.GetSprite(ressource)
	if err != nil {
		log.Fatal("[RenderManager] Blit(): ", err)
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

	// fmt.Printf("size: %d, %d -- pos: %d, %d\n", 
	// 	srcrect.W, srcrect.H, srcrect.X, srcrect.Y)

	this.screen.Blit(
		&sdl.Rect{
		int16(position.X), int16(position.Y), 
		uint16(image.W), uint16(image.H)},
		image,
		srcrect,
	)

}

func (this *RenderManager) BlitSprite(sprite *Sprite) {
	this.screen.Blit(
		sprite.GetDestRect(),
		sprite.GetSurface(),
		sprite.GetSrcRect(),
	)

}

// func (this *RenderManager) GetSpriteBuffer() *Sprite {
// 	return this.spriteBuffer
// }

func (this *RenderManager) GetSprite(gid uint) (*sdl.Surface, *sdl.Rect, error) {
	return this.mapSprite.GetSprite(gid)
}

//TODO: maybe to move in Scene
func (this *RenderManager) GetBackgroundColor() uint32 {
	return this.backgroundColor
}

func (this *RenderManager) SetBackgroundColor(backgroundColor uint32) {
	this.backgroundColor = backgroundColor
}