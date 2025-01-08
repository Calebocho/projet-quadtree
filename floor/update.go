package floor

import (
	"log"

	"gitlab.univ-nantes.fr/jezequel-l/quadtree/assets"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
)

// Update se charge de stocker dans la structure interne (un tableau)
// de f une représentation de la partie visible du terrain à partir
// des coordonnées absolues de la case sur laquelle se situe la
// caméra.
//
// On aurait pu se passer de cette fonction et tout faire dans Draw.
// Mais cela permet de découpler le calcul de l'affichage.
func (f *Floor) Update(camXPos, camYPos int) {
	topLeftX := camXPos - configuration.Global.ScreenCenterTileX
	topLeftY := camYPos - configuration.Global.ScreenCenterTileY
	f.topLeftX, f.topLeftY = topLeftX, topLeftY
	switch configuration.Global.FloorKind {
	case GridFloor:
		f.updateGridFloor(topLeftX, topLeftY)
	case FromFileFloor:
		f.updateFromFileFloor(topLeftX, topLeftY)
	case QuadTreeFloor:
		f.updateQuadtreeFloor(topLeftX, topLeftY)
	}

	if configuration.Global.AnimateFloor {
		f.frameCounter++
		if f.frameCounter < configuration.Global.NumFramePerFloorAnimImage {
			return
		}
		f.frameCounter = 0

		floorImageHeight := assets.FloorImage.Bounds().Dy()
		if floorImageHeight % configuration.Global.TileSize != 0 {
			log.Fatal("floor image height is not divisible by the tile size!")
		}
		frameCount := floorImageHeight / configuration.Global.TileSize
		f.currentAnimationFrame++
		if f.currentAnimationFrame >= frameCount {
			f.currentAnimationFrame = 0
		}
	}
}

// le sol est un quadrillage de tuiles d'herbe et de tuiles de désert
func (f *Floor) updateGridFloor(topLeftX, topLeftY int) {
	for y := 0; y < len(f.content); y++ {
		for x := 0; x < len(f.content[y]); x++ {
			absX := topLeftX
			if absX < 0 {
				absX = -absX
			}
			absY := topLeftY
			if absY < 0 {
				absY = -absY
			}
			f.content[y][x] = ((x + absX%2) + (y + absY%2)) % 2
		}
	}
}

// le sol est récupéré depuis un tableau, qui a été lu dans un fichier
//
// la version actuelle recopie fullContent dans content, ce qui n'est pas
// le comportement attendu dans le rendu du projet
func (f *Floor) updateFromFileFloor(topLeftX, topLeftY int) {
	for y := 0; y < len(f.content); y++ {
		yAbsolu := y + topLeftY
		for x := 0; x < len(f.content[y]); x++ {
			xAbsolu := x + topLeftX
			if yAbsolu >= 0 && yAbsolu < len(f.fullContent) &&
				xAbsolu >= 0 && xAbsolu < len(f.fullContent[yAbsolu]) {
				f.content[y][x] = f.fullContent[yAbsolu][xAbsolu]
			} else {
				f.content[y][x] = -1
			}
		}
	}
}

// le sol est récupéré depuis un quadtree, qui a été lu dans un fichier
func (f *Floor) updateQuadtreeFloor(topLeftX, topLeftY int) {
	f.quadtreeContent.GetContent(topLeftX, topLeftY, f.content)
}
