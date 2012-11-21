package ge2d

import (
	// "fmt"
	"log"
	"os"
	"encoding/xml"
	"github.com/0xe2-0x9a-0x9b/Go-SDL/sdl"
)

type SpriteLoader struct {
	animMap		map[string]*Anim
}

type XmlImage struct {
	Path		string		`xml:"path,attr"`
	// Time		uint		`xml:"time,attr"`
}

type XmlAnim struct {
	Name		string		`xml:"name,attr"`
	Frequency	uint		`xml:"frequency,attr"`
	Images		[]XmlImage	`xml:"image"`
}

type XmlSpriteFormat struct {
	XMLName		xml.Name	`xml:"sprite"`
	Anim		[]XmlAnim	`xml:"anim"`
	// Name		string		`xml:"name,attr"`
	// Images	[]XmlImage	`xml:"image"`
}

func NewSpriteLoader() *SpriteLoader {
	return &SpriteLoader {make(map[string]*Anim)}
}

func (this *SpriteLoader) Load(gsprFile string) {
	xmlFile, err := os.Open(gsprFile)
	if err != nil {
		log.Fatal("[SpriteLoader] Error opening file:", err)
		return
	}
	defer xmlFile.Close()
	
	var q XmlSpriteFormat
	xml.NewDecoder(xmlFile).Decode(&q)

	if q.XMLName.Local != "sprite" {
		log.Fatal(
			"[SpriteLoader] \"", 
			gsprFile, 
			"\" file does not content <sprite></sprite>")
	}

	for _, anim := range q.Anim {
		if _, exist := this.animMap[anim.Name]; exist {
			log.Fatal("[SpriteLoader] Sprite with same name already loaded:", anim.Name)
		}
		this.animMap[anim.Name] = 
			NewAnim(anim.Name, anim.Frequency, len(anim.Images))
		for i, image := range anim.Images {
			this.animMap[anim.Name].Sprite[i] = sdl.Load(image.Path)
			if this.animMap[anim.Name].Sprite[i] == nil {
				log.Fatal("[SpriteLoader] ", sdl.GetError())
			}
		}
	}	
	// fmt.Println(this.animMap)
	// fmt.Println(this.animMap["walk_face"])
}

func (this *SpriteLoader) GetAnimMap() map[string]*Anim {
	return this.animMap
}