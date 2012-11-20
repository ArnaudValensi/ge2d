package ge2d

import (
//"fmt"
"github.com/0xe2-0x9a-0x9b/Go-SDL/sdl"
"log"
//"math"
//"os"
//"strings"
//"time"
)

type RenderManager struct {
	screen *sdl.Surface
	image *sdl.Surface
}

func NewRenderManager() *RenderManager {
	return &RenderManager {nil, nil}
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

	// DEBUG
	this.DebugRessources()
}


func (this *RenderManager) Quit() {
	// image.Free()
	sdl.Quit()
}

var i int16 = 0
func (this *RenderManager) Update() {
	this.screen.FillRect(nil, 0x302019)
	this.screen.Blit(&sdl.Rect{i, 0, 0, 0}, this.image, nil)
	i++
	this.screen.Flip()
}

func (this *RenderManager) DebugRessources() {
	this.image = sdl.Load(resourcePath + "/test.png")

	if this.image == nil {
		log.Fatal(sdl.GetError())
	}

	// sdl.WM_SetIcon(image, nil)
}
