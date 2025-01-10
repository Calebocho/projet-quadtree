package floor

import (
	"log"
	"os"

	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
)

func (f *Floor) SaveFloor(fileName string) {
	fullContent := [][]int{}

	switch configuration.Global.FloorKind {
	case GridFloor:
		log.Fatal("infinite grid floor saving not possible")
	case FromFileFloor:
		fullContent = f.fullContent
	case QuadTreeFloor:
		fullContent = make([][]int, f.quadtreeContent.GetHeight())
		for y := 0; y < len(fullContent); y++ {
			fullContent[y] = make([]int, f.quadtreeContent.GetWidth())
		}

		f.quadtreeContent.GetContent(0, 0, fullContent)
	}

	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(fullContent); i++ {
		for j := 0; j < len(fullContent[i]); j++ {
			if _, err := file.Write([]byte{byte(fullContent[i][j]) + byte('0')}); err != nil {
				file.Close()
				log.Fatal(err)
			}
		}

		if _, err := file.Write([]byte{byte('\n')}); err != nil {
			file.Close()
			log.Fatal(err)
		}
	}
	if err := file.Close(); err != nil {
		log.Fatal(err)
	}
}
