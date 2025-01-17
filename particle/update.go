package particle

import "log"

func (particle *Particle) Update() {
	if particle.framesLeft > 0 {
		// Indicé avec un ParticleType
		framesPerAnimationFrame := []int {-1, 6, 10}
		animationFramesCount := []int {1, 8, 2}

		if int(particle.particleType) >= len(framesPerAnimationFrame) || int(particle.particleType) >= len(animationFramesCount) {
			log.Fatal("les données relatives aux frames d'animation de la particule ", particle.particleType, " doivent être ajoutés dans particle.Update()")
		}

		particle.framesLeft--

		isParticleAnimated := animationFramesCount[particle.particleType] > 1
		if isParticleAnimated && particle.framesLeft % framesPerAnimationFrame[particle.particleType] == 0 {
			particle.currentAnimationFrame++
			// On cycle de nouveau sur les mêmes frame d'animation pendant
			// toute la durée d'affichage framesLeft de la particule.
			if particle.currentAnimationFrame >= animationFramesCount[particle.particleType] {
				particle.currentAnimationFrame = 0
				if particle.oneShot {
					particle.framesLeft = 0
				}
			}
		}
	}
}