package ge2d

import (
	"github.com/0xe2-0x9a-0x9b/Go-SDL/sdl"
	// "errors"
)

// Sprite is used to load all sprite set and access to sprites with a
// global id (gid)
type Sprite struct {
	spriteMap	map[uint]*spriteSetPair
	// It used to quickly know where is the requested sprite in the spriteMap 
	globalIndex	map[uint]uint
}

type spriteSetPair struct {
	spriteSet	*SpriteSet
	firstGid	uint
}

func NewSprite() *Sprite {
	return &Sprite { make(map[uint]*spriteSetPair), make(map[uint]uint) }
}

func (this *Sprite) LoadSpriteSet(file string, elemWidth uint, elemHeight uint) {
	set := NewSpriteSet(file, elemWidth, elemHeight)
	spriteMapIndex := uint(len(this.spriteMap))
	pair := &spriteSetPair{ set, spriteMapIndex }
	this.spriteMap[spriteMapIndex] = pair

	globalIndexFirstFree := uint(len(this.globalIndex))
	nbSprite := set.GetNumberSprites()
	for i := 0; uint(i) < nbSprite; i++ {
		this.globalIndex[globalIndexFirstFree + uint(i)] = spriteMapIndex
	}
}

// Get sprite by the global id
func (this *Sprite) GetSprite(gid uint) (*sdl.Surface, int, int, error) {
	pair := this.spriteMap[this.globalIndex[gid]]
	spriteId := gid - pair.firstGid
	return pair.spriteSet.GetSprite(spriteId)
}