package particle

import (
	"image"

	"gitlab.univ-nantes.fr/jezequel-l/quadtree/assets"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"

	"github.com/hajimehoshi/ebiten/v2"
)

func (particle Particle) Draw(screen *ebiten.Image, viewTopLeftX, viewTopLeftY, viewShiftX, viewShiftY int) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM = particle.transformMatrix
	op.GeoM.Translate(
		float64((particle.X - viewTopLeftX) * configuration.Global.TileSize + particle.ShiftX - viewShiftX),
		float64((particle.Y - viewTopLeftY) * configuration.Global.TileSize + particle.ShiftY - viewShiftY),
	)

	shiftX := int(particle.particleType) * configuration.Global.TileSize
	shiftY := particle.currentAnimationFrame * configuration.Global.TileSize
	screen.DrawImage(assets.ParticlesImage.SubImage(
		image.Rect(shiftX, shiftY, shiftX + configuration.Global.TileSize, shiftY + configuration.Global.TileSize),
	).(*ebiten.Image), op)
}