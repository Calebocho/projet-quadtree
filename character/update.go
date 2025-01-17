package character

import (
	"math/rand"
	"log"

	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/particle"

	"github.com/hajimehoshi/ebiten/v2"
)

// Update met à jour la position du personnage, son orientation
// et son étape d'animation (si nécessaire) à chaque pas
// de temps, c'est-à-dire tous les 1/60 secondes.
//
// Retourne true si l'on vient de finir un déplacement au dernier appel à Update()
func (c *Character) Update(blocking [4]bool, currentCell int) (justMoved bool, newParticle *particle.Particle) {
	justMoved = false

	if !c.moving {
		if ebiten.IsKeyPressed(ebiten.KeyRight) {
			c.orientation = orientedRight
			if !blocking[1] {
				c.xInc = 1
				c.moving = true
			}
		} else if ebiten.IsKeyPressed(ebiten.KeyLeft) {
			c.orientation = orientedLeft
			if !blocking[3] {
				c.xInc = -1
				c.moving = true
			}
		} else if ebiten.IsKeyPressed(ebiten.KeyUp) {
			c.orientation = orientedUp
			if !blocking[0] {
				c.yInc = -1
				c.moving = true
			}
		} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
			c.orientation = orientedDown
			if !blocking[2] {
				c.yInc = 1
				c.moving = true
			}
		}
	} else {
		if configuration.Global.Particles {
			newParticle = c.placeWalkingParticle(currentCell)
		}

		c.animationFrameCount++
		if c.animationFrameCount >= configuration.Global.NumFramePerCharacterAnimImage {
			c.animationFrameCount = 0
			shiftStep := configuration.Global.TileSize / configuration.Global.NumCharacterAnimImages
			c.shift += shiftStep
			c.animationStep = -c.animationStep
			if c.shift > configuration.Global.TileSize-shiftStep {
				c.shift = 0
				c.moving = false
				c.X += c.xInc
				c.Y += c.yInc
				c.xInc = 0
				c.yInc = 0
				c.placedParticle = false
				justMoved = true
				return
			}
		}
	}
	return
}

func (c *Character) placeWalkingParticle(currentCell int) (newParticle *particle.Particle) {
	orientationDeg := 0
	if currentCell != 4 {
		switch c.orientation {
		case orientedDown:
			orientationDeg = 180
		case orientedLeft:
			orientationDeg = 270
		case orientedRight:
			orientationDeg = 90
		case orientedUp:
			orientationDeg = 0
		default:
			log.Fatal("orientation invalide du personnage dans character.Update()")
		}
	}

	particleType := particle.GetParticleTypeWhenWalkingOnCell(currentCell)

	if !c.placedParticle && rand.Intn(particleType.GetWalkingAppearanceFrequency()) == 0 {
		xShift, yShift := c.GetShift()
		p := particle.NewParticle(particleType, float64(orientationDeg), c.X, c.Y, xShift, yShift)
		newParticle = &p
		c.placedParticle = particleType.IsAloneOnCell()
	}

	return
}