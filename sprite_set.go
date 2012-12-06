// TODO: error if the surface is not divisable by tilesize
// TODO: free

package ge2d

import (
	"github.com/0xe2-0x9a-0x9b/Go-SDL/sdl"
	"errors"
	"github.com/kr/pretty"
	"log"
	// "fmt"
)

// SpriteSet is used to load a sprite set (an image which contain many 
// subimages/sprites).
type SpriteSet struct {
	surface		*sdl.Surface
	elemWidth	uint
	elemHeight	uint
	elemList	[]Vector2d
}

// Load a file as sprite set. elemWidth/elemHeight are the size of each
// element in the sprite set
func NewSpriteSet(file string, elemWidth uint, elemHeight uint) *SpriteSet {
	surface := sdl.Load(file)
	if surface == nil {
		log.Fatal("[SpriteSet] NewSpriteSet(): unable to load resource")
	}
	nbElemWidth := int(surface.W) / int(elemWidth)
	nbElemHeight := int(surface.H) / int(elemHeight)
	nbElem := nbElemWidth * nbElemHeight
	elemList := make([]Vector2d, nbElem)

	pretty.Printf("=====: %# v, nbElem %d, %d\n", elemList, nbElem, elemWidth)

	var x, y int = 0, 0
	for i := 0; i < nbElem; i++ {
		elemList[i] = Vector2d { x * int(elemWidth), y * int(elemHeight) }
		x = (x + 1) % nbElemWidth
		if x == 0 {
			y++
		}
	}

	return  &SpriteSet { surface, elemWidth, elemHeight, elemList }
}

// Get a sprite (an elem). It begin by 0
func (this *SpriteSet) GetSprite(id uint) (*sdl.Surface, *sdl.Rect, error) {
	if id >= uint(len(this.elemList)) {
		err := errors.New("id for requested sprite is out of bound")
		return nil, nil, err
	}
	rect := &sdl.Rect {
		int16(this.elemList[id].X), int16(this.elemList[id].Y), 
		uint16(this.elemWidth), uint16(this.elemHeight)}
	return this.surface, rect, nil
}

// Return the number of sprites
func (this *SpriteSet) GetNumberSprites() uint {
	return (uint(this.surface.W) / this.elemWidth) * (uint(this.surface.H) / this.elemHeight)
}

func (this *SpriteSet) Debug() {
	pretty.Printf("tmx: %# v\n", this.elemList)
}
