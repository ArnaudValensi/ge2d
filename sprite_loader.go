//TODO: ApriteLoader: free
//TODO: Rename SpriteLoader into RessoursesManager or something like

package ge2d

import (
	"log"
	"os"
	"encoding/xml"
	"github.com/0xe2-0x9a-0x9b/Go-SDL/sdl"
	"strconv"
)

type SpriteLoader struct {
	animMap			map[string]*Anim
	mapSprite		map[string]*SpriteSet
	data			XmlSpriteFormat
	singleImageCount	uint
}

type XmlImage struct {
	Path		string		`xml:"path,attr"`
	Tileset		string		`xml:"tileset,attr"`
	Sprite_num	uint		`xml:"sprite_num,attr"`
}

type XmlTileset struct {
	// FirstGid	string		`xml:"firstgid,attr"`
	Name		string		`xml:"name,attr"`
	TileWidth	uint		`xml:"tilewidth,attr"`
	TileHeight	uint		`xml:"tileheight,attr"`
	Source		string		`xml:"source,attr"`
}

type XmlAnim struct {
	Name		string		`xml:"name,attr"`
	Frequency	uint		`xml:"frequency,attr"`
	Images		[]XmlImage	`xml:"image"`
}

type XmlSpriteFormat struct {
	XMLName		xml.Name	`xml:"sprite"`
	Tilesets	[]XmlTileset	`xml:"tileset"`
	Anims		[]XmlAnim	`xml:"anim"`
}

type RessourceImage struct {
	spriteSetName	string
	imageNum	uint
}

func NewSpriteLoader() *SpriteLoader {
	return &SpriteLoader {
		make(map[string]*Anim),
		make(map[string]*SpriteSet),
		XmlSpriteFormat {},
		0,
	}
}

func (this *SpriteLoader) loadTilesets() {
	for _, tileset := range this.data.Tilesets {
		if tileset.Name == "" {
			log.Fatal("[SpriteLoader] gspr: 'tileset' must contain a 'name'")
		}
		if tileset.TileWidth == 0 || tileset.TileHeight == 0 {
			log.Fatal("[SpriteLoader] gspr: 'tileset' must contain",
				" a positive 'tilewidth' and 'tileheight'")
		}
		if tileset.Source == "" {
			log.Fatal("[SpriteLoader] gspr: 'tileset' must contain a 'source'")
		}
		this.mapSprite[tileset.Name] = NewSpriteSet(
			tileset.Source, 
			tileset.TileWidth,
			tileset.TileHeight,
		)
		log.Printf("[SpriteLoader] TileSet loaded: %s\n", tileset.Name)
	}
}

func (this *SpriteLoader) loadAnims() {
	for _, anim := range this.data.Anims {
		if anim.Name == "" {
			log.Fatal("[SpriteLoader] gspr: 'anim' must contain",
				" a 'name'")
		}
		if anim.Frequency == 0 {
			log.Fatal("[SpriteLoader] gspr: 'anim' must contain",
				" a 'frequency'")
		}
		newAnim := NewAnim(anim.Name, anim.Frequency)
		for _, image := range anim.Images {
			if image.Path == "" && image.Tileset == "" {
				log.Fatal("[SpriteLoader] gspr: 'anim>image' must contain",
					" a 'path' or a 'tileset'")
			}
			if image.Path != "" && image.Tileset != "" {
				log.Fatal("[SpriteLoader] gspr: 'anim>image' must not contain",
					" either 'path' or 'tileset'")
			}
			if image.Tileset != "" {
				newAnim.AddSprite(image.Tileset, image.Sprite_num)
			} else {
				sprite := sdl.Load(image.Path)
				imageName := strconv.Itoa(int(this.singleImageCount))
				this.mapSprite[imageName] = NewSpriteSet(
					image.Path, 
					uint(sprite.W),
					uint(sprite.H),
					)
				sprite.Free()
				newAnim.AddSprite(imageName, 0)
				this.singleImageCount++
			}
		}
		this.animMap[anim.Name] = newAnim
	}
}

func (this *SpriteLoader) Load(gsprFile string) {
	xmlFile, err := os.Open(gsprFile)
	if err != nil {
		log.Fatal("[SpriteLoader] Error opening file:", err)
		return
	}
	defer xmlFile.Close()
	
	xml.NewDecoder(xmlFile).Decode(&this.data)
	if this.data.XMLName.Local != "sprite" {
		log.Fatal(
			"[SpriteLoader] \"", 
			gsprFile, 
			"\" file does not content <sprite></sprite>")
	}

	// pretty.Printf("tmx: %# v\n", this.data)

	this.loadTilesets()
	this.loadAnims()
}

func (this *SpriteLoader) GetAnimMap() map[string]*Anim {
	return this.animMap
}

func (this *SpriteLoader) GetAnim(name string) *Anim {
	if anim, exist := this.animMap[name]; exist {
		return anim
	}
	log.Fatal("[SpriteLoader] GetAnim: unknown anim")
	return nil
}

func (this *SpriteLoader) GetSprite(res *RessourceImage) (*sdl.Surface, *sdl.Rect, error) {
	if spriteset, exist := this.mapSprite[res.spriteSetName]; exist {
		return spriteset.GetSprite(res.imageNum)
	}
	log.Fatal("[SpriteLoader] GetSprite: spriteset name does not exist")
	return nil, nil, nil
}
