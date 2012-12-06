package ge2d

import (
	"github.com/0xe2-0x9a-0x9b/Go-SDL/sdl"
)

type Sprite struct {
	// gid		uint
	surface		*sdl.Surface
	srcrect		*sdl.Rect
	destrect	*sdl.Rect
	// DestX		int16
	// DestY		int16
	// DestWidth	uint16
	// DestHeight	uint16
	// SrcX		int16
	// SrcY		int16
	// SrcWidth	uint16
	// SrcHeight	uint16
}

func NewSprite() *Sprite {
	return &Sprite { nil, nil, nil }
}

func (this *Sprite) GetSurface() *sdl.Surface {
	return this.surface
}

func (this *Sprite) SetSurface(surface *sdl.Surface) {
	this.surface = surface
}

// func (this *Sprite) GetGid() uint {
// 	return this.gid
// }

// func (this *Sprite) SetGid(gid uint) {
// 	this.gid = gid
// }

func (this *Sprite) GetSrcRect() *sdl.Rect {
	return this.srcrect
}

func (this *Sprite) SetSrcRect(rect *sdl.Rect) {
	this.srcrect = rect
}

func (this *Sprite) GetDestRect() *sdl.Rect {
	return this.destrect
}

func (this *Sprite) SetDestRect(rect *sdl.Rect) {
	this.destrect = rect
}
