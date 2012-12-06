package ge2d

import (
	"github.com/kyleconroy/go-tmx/tmx"
	"os"
	"log"
	"github.com/kr/pretty"
	"github.com/0xe2-0x9a-0x9b/Go-SDL/sdl"
)

type TmxTileMap struct {
	renderManager	*RenderManager
	spriteBuffer	Sprite
	tmxMap		*tmx.Map
	width		uint
	height		uint
	tileWidth	uint
	tileHeight	uint
}

func NewTmxTileMap(renderManager *RenderManager) *TmxTileMap {
	return &TmxTileMap { renderManager, Sprite {nil, nil, nil}, nil, 0, 0, 0, 0 }
}

// TODO: voir la place du defer
func (this *TmxTileMap) Load(file string) {
	chk := func(err error) {
		if err != nil {
			log.Fatal("[TmxTileMap] Load(): ", err)
		}
	}

	fd, err := os.Open(file)
	chk(err)
	defer fd.Close()
	
	this.tmxMap, err = tmx.Read(fd)
	chk(err)
	
	pretty.Printf("tmx: %# v\n", this.tmxMap)
	// log.Printf("================\ntmx: %+v\n", m.Version)

	for _, tileset := range this.tmxMap.Tilesets {
		this.renderManager.LoadMapSpriteSet(
			tileset.Image.Source, 
			uint(tileset.TileWidth), 
			uint(tileset.TileHeight),
			)
	}

	this.width = uint(this.tmxMap.Width)
	this.height = uint(this.tmxMap.Height)
	this.tileWidth = uint(this.tmxMap.TileWidth)
	this.tileHeight = uint(this.tmxMap.TileHeight)
}

func (this *TmxTileMap) blitLayer() {
	
}

func (this *TmxTileMap) BlitMap() {
	for _, layer := range this.tmxMap.Layers {
		for n, tile := range layer.DecodedTiles {
			if tile.Tileset != nil {
				tilenum := uint(tile.Tileset.FirstGID) + uint(tile.ID) - 1
				surface, srcrect, err := this.renderManager.GetSprite(tilenum)
				if err != nil {
					log.Fatal("[TmxTileMap] BlitMap(): ", err)
				}

				x := n % int(this.width) * int(this.tileWidth)
				y := n / int(this.width) * int(this.tileHeight) - int(srcrect.H) + this.tmxMap.TileHeight
				log.Printf("n: %d, x: %d, y: %d\n", n, x, y)


				destrect := &sdl.Rect{ int16(x), int16(y), srcrect.W, srcrect.H }
				this.spriteBuffer.SetSurface(surface)
				this.spriteBuffer.SetSrcRect(srcrect)
				this.spriteBuffer.SetDestRect(destrect)

				this.renderManager.BlitSprite(&this.spriteBuffer)
			}
		}
	}
}
