package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
)

// Update met à jour les données du jeu à chaque 1/60 de seconde.
// Il faut bien faire attention à l'ordre des mises-à-jour car elles
// dépendent les unes des autres (par exemple, pour le moment, la
// mise-à-jour de la caméra dépend de celle du personnage et la définition
// du terrain dépend de celle de la caméra).
func (g *Game) Update() error {

	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		configuration.Global.DebugMode = !configuration.Global.DebugMode
	}

	useTeleporters := configuration.Global.Teleporters
	if useTeleporters && inpututil.IsKeyJustPressed(ebiten.KeyT) {
		teleporter_x, teleporter_y := g.character.GetPosInFront(1)
		g.floor.SetNewTeleporter(teleporter_x, teleporter_y)
	}

	justMoved := g.character.Update(g.floor.Blocking(g.character.X, g.character.Y, g.camera.X, g.camera.Y))
	if useTeleporters && justMoved {
		teleportedPos := g.floor.Teleport(g.character.X, g.character.Y)
		if teleportedPos != nil {
			g.character.X = teleportedPos.X
			g.character.Y = teleportedPos.Y
		}
	}

	g.camera.Update(g.character.X, g.character.Y, g.floor.GetWidth(), g.floor.GetHeight())
	g.floor.Update(g.camera.X, g.camera.Y)

	return nil
}
