package floor

import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/quadtree"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
)

type Pos struct {
	X, Y int
}

// Floor représente les données du terrain. Pour le moment
// aucun champs n'est exporté.
//
//   - content : partie du terrain qui doit être affichée à l'écran
//   - fullContent : totalité du terrain (utilisé seulement avec le type
//     d'affichage du terrain "fromFileFloor")
//   - quadTreeContent : totalité du terrain sous forme de quadtree (utilisé
//     avec le type d'affichage du terrain "quadtreeFloor")
type Floor struct {
	content         [][]int
	fullContent     [][]int
	quadtreeContent quadtree.Quadtree
	topLeftX, topLeftY    int

	frameCounter          int
	currentAnimationFrame int

	teleporters           [2]*Pos
}

// types d'affichage du terrain disponibles
const (
	GridFloor int = iota
	FromFileFloor
	QuadTreeFloor
)

// GetHeight retourne la hauteur (en cases) du terrain
// à partir du tableau fullContent, en supposant que
// ce tableau représente un terrain rectangulaire
func (f Floor) GetHeight() (height int) {
	switch configuration.Global.FloorKind {
	case FromFileFloor:
		return len(f.fullContent)
	case QuadTreeFloor:
		return f.quadtreeContent.GetHeight()
	}
	return
}

// GetWidth retourne la largeur (en cases) du terrain
// à partir du tableau fullContent, en supposant que
// ce tableau représente un terrain rectangulaire
func (f Floor) GetWidth() (width int) {
	switch configuration.Global.FloorKind {
	case FromFileFloor:
		if len(f.fullContent) > 0 {
			width = len(f.fullContent[0])
		}
		return
	case QuadTreeFloor:
		return f.quadtreeContent.GetWidth()
	}
	return
}