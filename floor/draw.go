package floor

import (
	"image"

	"gitlab.univ-nantes.fr/jezequel-l/quadtree/assets"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"

	"github.com/hajimehoshi/ebiten/v2"
)

// Draw affiche dans une image (en général, celle qui représente l'écran),
// la partie du sol qui est visible (qui doit avoir été calculée avec Get avant).
func (f Floor) Draw(screen *ebiten.Image, xShift, yShift int) {
	for y := range f.content {
		for x := range f.content[y] {
			if f.content[y][x] != -1 {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(x*configuration.Global.TileSize - xShift), float64(y*configuration.Global.TileSize - yShift))

				shiftX := f.content[y][x] * configuration.Global.TileSize
				shiftY := f.currentAnimationFrame * configuration.Global.TileSize

				screen.DrawImage(assets.FloorImage.SubImage(
					image.Rect(shiftX, shiftY, shiftX+configuration.Global.TileSize, shiftY+configuration.Global.TileSize),
				).(*ebiten.Image), op)

				if f.IsOnAnyTeleporter(f.topLeftX + x, f.topLeftY + y) {
					screen.DrawImage(assets.TeleporterImage, op)
				}
			}
		}
	}
}