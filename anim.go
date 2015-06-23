// TODO: Anim: free

package ge2d

import (
	"log"
)

type Anim struct {
	name		string
	frequency	uint
	sprites		map[uint]RessourceImage
	nbImage		uint
}

func NewAnim(name string, frequency uint) *Anim {
	return &Anim {
		name,
		frequency,
		make(map[uint]RessourceImage),
		0,
	}
}

func (this *Anim) GetName() string {
	return this.name
}

func (this *Anim) GetFrequency() uint {
	return this.frequency
}

// TODO: ne pas exit
func (this *Anim) GetResource(n uint) *RessourceImage {
	if ressource, exist := this.sprites[n]; exist {
		return &ressource
	}
	log.Fatal("[Anim] GetSprite(): n is out of bound") 
	return nil
}

func (this *Anim) GetNumberSprite() uint {
	return this.nbImage
}

func (this *Anim) AddSprite(spriteSetName string, spriteNum uint) {
	this.sprites[this.nbImage] = RessourceImage { spriteSetName, spriteNum }
	this.nbImage++
}
