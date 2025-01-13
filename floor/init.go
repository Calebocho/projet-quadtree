package floor

import (
	"bufio"
	"log"
	"math/rand"
	"os"

	"gitlab.univ-nantes.fr/jezequel-l/quadtree/assets"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/quadtree"
)

// Init initialise les structures de données internes de f.
func (f *Floor) Init() {
	f.UpdateCameraViewSize()

	switch configuration.Global.FloorKind {
	case FromFileFloor:
		f.fullContent = readFloorFromFile(configuration.Global.FloorFile)
	case QuadTreeFloor:
		//Extension 1
		if configuration.Global.RandomFloor {
			tileWidth := assets.FloorImage.Bounds().Dx() / configuration.Global.TileSize
			floor := make([][]int, configuration.Global.RandomTileY)
			for i := 0; i < len(floor); i++ {
				floor[i] = make([]int, configuration.Global.RandomTileX)
				for j := 0; j < len(floor[i]); j++ {
					floor[i][j] = rand.Intn(tileWidth)
				}
			}
			f.quadtreeContent = quadtree.MakeFromArray(floor)
		} else {
			f.quadtreeContent = quadtree.MakeFromArray(readFloorFromFile(configuration.Global.FloorFile))
		}
	}
}

func (f *Floor) UpdateCameraViewSize() {
	f.content = make([][]int, configuration.Global.NumTileY)
	for y := 0; y < len(f.content); y++ {
		f.content[y] = make([]int, configuration.Global.NumTileX)
	}
}

// lecture du contenu d'un fichier représentant un terrain
// pour le stocker dans un tableau
func readFloorFromFile(fileName string) (floorContent [][]int) {
	var file *os.File
	var err error
	file, err = os.Open(fileName)
	if err != nil {
		log.Fatal("erreur lors de l'ouverture du fichier de sol", fileName, err)
	}
	var scanner *bufio.Scanner = bufio.NewScanner(file)

	for lineNumber := 0; scanner.Scan(); lineNumber++ {
		line := scanner.Text()
		var floorLine []int = make([]int, len(line))
		for column, c := range line {
			floorCell := int(c) - int('0')
			if floorCell < 0 || floorCell > 9 {
				log.Fatal("caractère invalide", c, lineNumber, column)
			}
			floorLine[column] = floorCell
			if err != nil {
				log.Fatal("erreur lors de la conversion caractère vers nombre dans le fichier",
					fileName, "à la ligne", lineNumber,
					column, "ième caractère", err)
			}
		}
		if lineNumber != 0 && len(floorLine) != len(floorContent[len(floorContent)-1]) {
			log.Fatal("le terrain n'est pas rectangulaire à partir de la ligne", lineNumber)
		}
		floorContent = append(floorContent, floorLine)
	}

	err = scanner.Err()
	if err != nil {
		log.Fatal("erreur lors de la lecture du fichier", fileName, err)
	}

	err = file.Close()
	if err != nil {
		log.Fatal("erreur lors de la fermeture du fichier", fileName, err)
	}

	return
}
