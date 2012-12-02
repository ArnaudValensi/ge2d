package ge2d

import (
	"github.com/0xe2-0x9a-0x9b/Go-SDL/sdl"
)

// TODO: private attributes
type Anim struct {
	name		string
	frequency	uint
	sprite		[]*sdl.Surface
	// nbCall		uint
}

func NewAnim(name string, frequency uint, nbSprites int) *Anim {
	return &Anim {name, frequency, make([]*sdl.Surface, nbSprites)}
}

func (this *Anim) Free() {
	for _, image := range this.sprite {
		image.Free()
	}
}

func (this *Anim) GetName() string {
	return this.name
}

func (this *Anim) GetFrequency() uint {
	return this.frequency
}

func (this *Anim) GetSprite(n uint) *sdl.Surface {
	return this.sprite[n]
}

func (this *Anim) GetNunberSprite() uint {
	return uint(len(this.sprite))
}

func (this *Anim) AddSprite(file string, i int) {
	this.sprite[i] = sdl.Load(file)
}

// func (this *Anim) GetNextImage() *sdl.Surface {
// 	image := this.Sprite[this.nbCall / this.Frequency]
// 	var nbImage uint = uint(len(this.Sprite))
// 	this.nbCall = (this.nbCall + 1) % (nbImage * this.Frequency)
// 	return image
// }
