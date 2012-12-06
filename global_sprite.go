package ge2d

import (
	"github.com/0xe2-0x9a-0x9b/Go-SDL/sdl"
)

// SpriteSetCollection is used to load all sprite set and access to sprites with a
// global id (gid)
type SpriteSetCollection struct {
	spriteMap	map[uint]*spriteSetPair
	// It used to quickly know where is the requested sprite in the spriteMap 
	globalIndex	map[uint]uint
}

type spriteSetPair struct {
	spriteSet	*SpriteSet
	firstGid	uint
}

func NewSpriteSetCollection() *SpriteSetCollection {
	return &SpriteSetCollection { make(map[uint]*spriteSetPair), make(map[uint]uint) }
}

func (this *SpriteSetCollection) LoadSpriteSet(file string, elemWidth uint, elemHeight uint) {
	set := NewSpriteSet(file, elemWidth, elemHeight)
	spriteMapIndex := uint(len(this.spriteMap))
	globalIndexFirstFree := uint(len(this.globalIndex))
	pair := &spriteSetPair{ set, globalIndexFirstFree }
	this.spriteMap[spriteMapIndex] = pair

	nbSprite := set.GetNumberSprites()
	for i := 0; uint(i) < nbSprite; i++ {
		this.globalIndex[globalIndexFirstFree + uint(i)] = spriteMapIndex
	}
}

func (this *SpriteSetCollection) GetSprite(gid uint) (*sdl.Surface, *sdl.Rect, error) {
	pair := this.spriteMap[this.globalIndex[gid]]
	spriteId := gid - pair.firstGid
	return pair.spriteSet.GetSprite(spriteId)
}