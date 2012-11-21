package ge2d

import (
	"github.com/0xe2-0x9a-0x9b/Go-SDL/sdl"
)
type Anim struct {
	Name		string
	Frequency	uint
	Sprite		[]*sdl.Surface
}

func NewAnim(name string, frequency uint, nbSprites int) *Anim {
	return &Anim {name, frequency, make([]*sdl.Surface, nbSprites)}
}

func (this *Anim) Free() {
	for _, image := range this.Sprite {
		image.Free()
	}

}