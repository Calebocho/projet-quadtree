package game

import (
	"log"

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

	currentViewPosX := g.character.X - g.camera.X + configuration.Global.ScreenCenterTileX
	currentViewPosY := g.character.Y - g.camera.Y + configuration.Global.ScreenCenterTileY
	posInsideCameraView, currentCell := g.floor.GetCameraViewCell(currentViewPosX, currentViewPosY)
	if !posInsideCameraView {
		log.Fatal("code cassé, ne devrait pas trouvé une position hors de la vue de la caméra ici")
	}

	justMoved, newParticle := g.character.Update(g.floor.Blocking(g.character.X, g.character.Y, g.camera.X, g.camera.Y), currentCell)
	if useTeleporters && justMoved {
		teleportedPos := g.floor.Teleport(g.character.X, g.character.Y)
		if teleportedPos != nil {
			g.character.X = teleportedPos.X
			g.character.Y = teleportedPos.Y
		}
	}
	if newParticle != nil {
		particleAdded := false
		for i, particle := range(g.particles) {
			if particle.Disappeared() {
				g.particles[i] = *newParticle
				particleAdded = true
				break
			}
		}
		if !particleAdded {
			g.particles = append(g.particles, *newParticle)
		}
	}

	for i := 0; i < len(g.particles); i++ {
		(&g.particles[i]).Update()
	}

	xShift, yShift := g.character.GetShift()
	g.camera.Update(g.character.X, g.character.Y, xShift, yShift, g.floor.GetWidth(), g.floor.GetHeight())
	g.floor.Update(g.camera.X, g.camera.Y)

	if configuration.Global.Zoomable {
		if inpututil.IsKeyJustPressed(ebiten.KeyNumpadSubtract) {
			configuration.Global.NumTileX += 1
			configuration.Global.NumTileY += 1
			configuration.SetComputedFields()
			g.floor.UpdateCameraViewSize()
			g.floor.Update(g.camera.X, g.camera.Y)
			g.camera.Update(g.character.X, g.character.Y, g.camera.XShift, g.camera.YShift, g.floor.GetWidth(), g.floor.GetHeight())
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyNumpadAdd) && configuration.Global.NumTileX > 3 && configuration.Global.NumTileY > 3 {
			configuration.Global.NumTileX -= 1
			configuration.Global.NumTileY -= 1
			configuration.SetComputedFields()
			g.floor.UpdateCameraViewSize()
			g.floor.Update(g.camera.X, g.camera.Y)
			g.camera.Update(g.character.X, g.character.Y, g.camera.XShift, g.camera.YShift, g.floor.GetWidth(), g.floor.GetHeight())
		}
	}

	return nil
}
