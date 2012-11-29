package ge2d

import (
	"github.com/0xe2-0x9a-0x9b/Go-SDL/sdl"
)

// TODO: private attributes
type Anim struct {
	Name		string
	Frequency	uint
	Sprite		[]*sdl.Surface
	// nbCall		uint
}

func NewAnim(name string, frequency uint, nbSprites int) *Anim {
	return &Anim {name, frequency, make([]*sdl.Surface, nbSprites)}
}

func (this *Anim) Free() {
	for _, image := range this.Sprite {
		image.Free()
	}

}

// func (this *Anim) GetNextImage() *sdl.Surface {
// 	image := this.Sprite[this.nbCall / this.Frequency]
// 	var nbImage uint = uint(len(this.Sprite))
// 	this.nbCall = (this.nbCall + 1) % (nbImage * this.Frequency)
// 	return image
// }
