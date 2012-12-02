package ge2d

import (
	"github.com/0xe2-0x9a-0x9b/Go-SDL/sdl"
	"errors"
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
func NewSpriteSet (file string, elemWidth uint, elemHeight uint) *SpriteSet {
	surface := sdl.Load(file)
	nbElem := (uint(surface.W) / elemWidth) * (uint(surface.H) / elemHeight)
	elemList := make([]Vector2d, nbElem)
	return  &SpriteSet { surface, elemWidth, elemHeight, elemList }
}

// Get a sprite (an elem). It begin by 0
func (this *SpriteSet) GetSprite (id uint) (*sdl.Surface, int, int, error) {
	if id >= uint(len(this.elemList)) {
		err := errors.New("id for request sprite is out of bound")
		return this.surface, this.elemList[id].X, this.elemList[id].Y, err
	}
	return this.surface, this.elemList[id].X, this.elemList[id].Y, nil
}

// Return the number of sprites
func (this *SpriteSet) GetNumberSprites() uint {
	return (uint(this.surface.W) / this.elemWidth) * (uint(this.surface.H) / this.elemHeight)
}