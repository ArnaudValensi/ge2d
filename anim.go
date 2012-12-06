// TODO: free

package ge2d

import (
	"log"
	"github.com/kr/pretty"
	// "github.com/0xe2-0x9a-0x9b/Go-SDL/sdl"
)

type Anim struct {
	name		string
	frequency	uint
	sprites		map[uint]RessourceImage
	nbImage		uint
	// sprite		[]*sdl.Surface
	// nbCall		uint
	// set		*SpriteSet
	// gset		*SpriteSetCollection
}

func NewAnim(name string, frequency uint) *Anim {
	// set := NewSpriteSet("Test.png", 101, 171)
	// gset := NewSpriteSetCollection()
	// gset.LoadSpriteSet("Test.png", 101, 171)
	// gset.LoadSpriteSet("set_windrider.png", 80, 78)
	// return &Anim {name, frequency, make([]*sdl.Surface, nbSprites), set, gset}
	return &Anim {
		name,
		frequency,
		make(map[uint]RessourceImage),
		0,
	}
}

// func (this *Anim) Free() {

// }

func (this *Anim) GetName() string {
	return this.name
}

func (this *Anim) GetFrequency() uint {
	return this.frequency
}

// TODO: ne pas exit
func (this *Anim) GetResource(n uint) *RessourceImage {
	pretty.Printf("SSS: %# v\n", this.sprites)
	if ressource, exist := this.sprites[n]; exist {
		return &ressource
	}
	log.Fatal("[Anim] GetSprite(): n is out of bound") 
	return nil
}

// func (this *Anim) GetSprite2(n uint) (*sdl.Surface, *sdl.Rect, error) {
// 	// return this.set.GetSprite(n % 9)
// 	return this.gset.GetSprite(n % 17)
// }

func (this *Anim) GetNumberSprite() uint {
	return this.nbImage
}

func (this *Anim) AddSprite(spriteSetName string, spriteNum uint) {
	this.sprites[this.nbImage] = RessourceImage { spriteSetName, spriteNum }
	this.nbImage++
}

// func (this *Anim) AddSprite(file string, i int) {
// 	sprite := sdl.Load(file)
// 	this.gset.LoadSpriteSet(file, uint(sprite.W), uint(sprite.H))
// 	sprite.Free()
// 	this.nbImage++
// }


// func (this *Anim) AddSpriteSet(file string, elementWidth uint, elementHeight uint) {
// 	this.gset.LoadSpriteSet(file, elementWidth, elementHeight)
// }

// func (this *Anim) GetNextImage() *sdl.Surface {
// 	image := this.Sprite[this.nbCall / this.Frequency]
// 	var nbImage uint = uint(len(this.Sprite))
// 	this.nbCall = (this.nbCall + 1) % (nbImage * this.Frequency)
// 	return image
// }
