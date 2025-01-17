package particle

import (
	"math"
	"log"

	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"

	"github.com/hajimehoshi/ebiten/v2"
)

type ParticleType int

const (
	FootstepsParticle = iota
	WaterSplashingParticle
	WalkThroughGrass
)

type Particle struct {
	particleType                 ParticleType
	framesLeft                   int
	currentAnimationFrame        int
	transformMatrix              ebiten.GeoM
	oneShot                      bool
	X, Y, ShiftX, ShiftY         int
}

// Crée une particule d'un certain type et orientée (sens trigonométrique, 0° à droite)
// placée sur la case (X, Y) avec un décallage supplémentaire de ShiftX et ShiftY
// pixels par rapport à la case centrée.
func NewParticle(particleType ParticleType, rotationDeg float64, X, Y, ShiftX, ShiftY int) Particle {
	durationSeconds := 0.
	oneShot := true
	switch particleType {
	case FootstepsParticle:
		durationSeconds = 2.
	case WaterSplashingParticle:
		durationSeconds = 1.
	case WalkThroughGrass:
		durationSeconds = 3.
	default:
		log.Fatal("cas manquant à traiter dans particle.NewParticle(): ", particleType)
	}

	transformMatrix := ebiten.GeoM {}
	halfTileSize := float64(configuration.Global.TileSize) / 2.
	transformMatrix.Translate(-halfTileSize, -halfTileSize)
	rotationRad := 2. * math.Pi * rotationDeg / 360.
	transformMatrix.Rotate(rotationRad)
	transformMatrix.Translate(halfTileSize, halfTileSize)
	return Particle {
		particleType: particleType,
		framesLeft: int(durationSeconds * 60.),
		currentAnimationFrame: 0,
		transformMatrix: transformMatrix,
		oneShot: oneShot,
		X: X,
		Y: Y,
		ShiftX: ShiftX,
		ShiftY: ShiftY,
	}
}

func (particle Particle) Disappeared() bool {
	return particle.framesLeft == 0
}

// La "fréquence" à laquelle cette particule apparaît lorsque le personnage marche.
// La fréquence réelle est de 1/freq par frame de Update() (1/60 de seconde).
func (particleType ParticleType) GetWalkingAppearanceFrequency() (frequency int) {
	switch particleType {
	case FootstepsParticle:
		frequency = 60
	case WaterSplashingParticle:
		frequency = 30
	case WalkThroughGrass:
		frequency = 5
	default:
		log.Fatal("cas manquant à traiter pour particleType.GetWalkingAppearanceFrequency(): ", particleType)
	}
	return
}

func (particleType ParticleType) IsAloneOnCell() (isAloneOnCell bool) {
	switch particleType {
	case FootstepsParticle:
		isAloneOnCell = true
	case WaterSplashingParticle:
		isAloneOnCell = true
	case WalkThroughGrass:
		isAloneOnCell = false
	default:
		log.Fatal("cas manquant à traiter pour particleType.GetWalkingAppearanceFrequency(): ", particleType)
	}
	return
}

func GetParticleTypeWhenWalkingOnCell(cell int) (particleType ParticleType) {
	switch cell {
	case 0:
		particleType = WalkThroughGrass
	case 4:
		particleType = WaterSplashingParticle
	case 1, 2, 3:
		particleType = FootstepsParticle
	default:
		log.Fatal("cas manquant à traiter dans particle.GetParticleTypeWhenWalkingOnCell(): ", particleType)
	}
	return
}