package ge2d

import (
	"fmt"
	"github.com/0xe2-0x9a-0x9b/Go-SDL/sdl"
	"log"
	//"math"
	"os"
	"strings"
	"time"
)

var resourcePath string

func Run() {
	{
		GOPATH := os.Getenv("GOPATH")
		if GOPATH == "" {
			log.Fatal("No such environment variable: GOPATH")
		}
		for _, gopath := range strings.Split(GOPATH, ":") {
			a := gopath + "/src/binpix/gosdl_test"
			_, err := os.Stat(a)
			if err == nil {
				resourcePath = a
				break
			}
		}
		if resourcePath == "" {
			log.Fatal("Failed to find resource directory")
		}
	}

	renderManager := NewRenderManager()
	renderManager.Init()
	// var i int16 = 0
	// image := sdl.Load(resourcePath + "/test.png")

	// if image == nil {
	// 	log.Fatal(sdl.GetError())
	// }

	// sdl.WM_SetIcon(image, nil)

	running := true

	if sdl.GetKeyName(270) != "[+]" {
		log.Fatal("GetKeyName broken")
	}

	ticker := time.NewTicker(time.Second / 50) // 50 Hz
	for running {
		select {
		case <-ticker.C:
			// screen.FillRect(nil, 0x302019)
			// screen.Blit(&sdl.Rect{i, 0, 0, 0}, image, nil)
			// i++
			// // var p Point
			// // sdl.GetMouseState(&p.x, &p.y)
			// // worm_in <- p

			// screen.Flip()
			renderManager.Update()
		case _event := <-sdl.Events:
			switch e := _event.(type) {
			case sdl.QuitEvent:
				running = false

			case sdl.KeyboardEvent:
				println("")
				println(e.Keysym.Sym, ": ", sdl.GetKeyName(sdl.Key(e.Keysym.Sym)))

				if e.Keysym.Sym == sdl.K_ESCAPE {
					running = false
				}

				fmt.Printf("%04x ", e.Type)

				for i := 0; i < len(e.Pad0); i++ {
					fmt.Printf("%02x ", e.Pad0[i])
				}
				println()

				fmt.Printf("Type: %02x Which: %02x State: %02x Pad: %02x\n", e.Type, e.Which, e.State, e.Pad0[0])
				fmt.Printf("Scancode: %02x Sym: %08x Mod: %04x Unicode: %04x\n", e.Keysym.Scancode, e.Keysym.Sym, e.Keysym.Mod, e.Keysym.Unicode)
				// case sdl.ResizeEvent:
				// 	println("resize screen ", e.W, e.H)
				
				// 	screen = sdl.SetVideoMode(int(e.W), int(e.H), 32, sdl.RESIZABLE)
				
				// 	if screen == nil {
				// 		log.Fatal(sdl.GetError())
				// 	}
				// }
			}
		}
		// image.Free()
	}
}
