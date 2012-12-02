package ge2d

import (
	"github.com/0xe2-0x9a-0x9b/Go-SDL/sdl"
)

type Anim struct {
	name		string
	frequency	uint
	sprite		[]*sdl.Surface
	// nbCall		uint
	set		*SpriteSet
}

func NewAnim(name string, frequency uint, nbSprites int) *Anim {
	set := NewSpriteSet("Test.png", 101, 171)
	return &Anim {name, frequency, make([]*sdl.Surface, nbSprites), set}
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

func (this *Anim) GetSprite2(n uint) (*sdl.Surface, *sdl.Rect, error) {
	return this.set.GetSprite(n % 9)
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
