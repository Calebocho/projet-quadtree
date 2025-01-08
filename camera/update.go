package camera

import "gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"

// Update met à jour la position de la caméra à chaque pas
// de temps, c'est-à-dire tous les 1/60 secondes.
func (c *Camera) Update(characterPosX, characterPosY, floor_width, floor_height int) {

	switch configuration.Global.CameraMode {
	case Static:
		c.updateStatic()
	case FollowCharacter:
		c.updateFollowCharacter(characterPosX, characterPosY, floor_width, floor_height)
	}
}

// updateStatic est la mise-à-jour d'une caméra qui reste
// toujours à la position (0,0). Cette fonction ne fait donc
// rien.
func (c *Camera) updateStatic() {}

// updateFollowCharacter est la mise-à-jour d'une caméra qui
// suit toujours le personnage. Elle prend en paramètres deux
// entiers qui indiquent les coordonnées du personnage et place
// la caméra au même endroit.
func (c *Camera) updateFollowCharacter(characterPosX, characterPosY, floor_width, floor_height int) {
	if !configuration.Global.CameraBlockedByEdge {
		c.X = characterPosX
		c.Y = characterPosY
	}

	leftViewEdgeX := characterPosX - configuration.Global.ScreenCenterTileX
	rightViewEdgeX := leftViewEdgeX + configuration.Global.NumTileX
	if leftViewEdgeX >= 0 && rightViewEdgeX <= floor_width {
		c.X = characterPosX
	}

	topViewEdgeY := characterPosY - configuration.Global.ScreenCenterTileY
	bottomViewEdgeY := topViewEdgeY + configuration.Global.NumTileY
	if topViewEdgeY >= 0 && bottomViewEdgeY <= floor_height {
		c.Y = characterPosY
	}
}
