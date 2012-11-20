package ge2d

import (
	"fmt"
	"log"
	"os"
	"encoding/xml"
)

type SpriteLoader struct {

}

type XmlImage struct {
	Path		string		`xml:"path,attr"`
	Time		uint		`xml:"time,attr"`
}

type XmlSpriteFormat struct {
	XMLName		xml.Name	`xml:"sprite"`
	Images		[]XmlImage	`xml:"image"`
}

func NewSpriteLoader() *SpriteLoader {
	return &SpriteLoader {}
}

func (this *SpriteLoader) Load(gsprFile string) {
	xmlFile, err := os.Open(gsprFile)
	if err != nil {
		log.Fatal("Error opening file:", err)
		return
	}
	defer xmlFile.Close()
	
	var q XmlSpriteFormat
	xml.NewDecoder(xmlFile).Decode(&q)
	// xml.Unmarshal(xmlFile, &q)
	fmt.Println(q)
	// for _, episode := range q.EpisodeList {
	// 	fmt.Printf("\t%s\n", episode)
	// }
}