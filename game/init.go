package game

import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
)

// Init initialise les données d'un jeu. Il faut bien
// faire attention à l'ordre des initialisation car elles
// pourraient dépendre les unes des autres.
func (g *Game) Init() {
	g.floor.Init()
	if configuration.Global.SaveFloor {
		g.floor.SaveFloor("../floor-files/saved-floor")
	}
	g.character.Init(g.floor.GetWidth(), g.floor.GetHeight())
	g.camera.Init(g.character.X, g.character.Y, 0, 0)
}
