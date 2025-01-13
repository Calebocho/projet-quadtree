package camera

import "gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"

// Update met à jour la position de la caméra à chaque pas
// de temps, c'est-à-dire tous les 1/60 secondes.
func (c *Camera) Update(characterPosX, characterPosY, xShift, yShift, floor_width, floor_height int) {

	switch configuration.Global.CameraMode {
	case Static:
		c.updateStatic()
	case FollowCharacter:
		c.updateFollowCharacter(characterPosX, characterPosY, xShift, yShift, floor_width, floor_height)
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
func (c *Camera) updateFollowCharacter(characterPosX, characterPosY, xShift, yShift, floor_width, floor_height int){
	c.OutsideXEdges = false
	c.OutsideYEdges = false

	if !configuration.Global.CameraBlockedByEdge {
		c.X = characterPosX
		c.Y = characterPosY
		c.XShift = xShift
		c.YShift = yShift
		return
	}

	leftViewEdgeX := characterPosX - configuration.Global.ScreenCenterTileX
	rightViewEdgeX := leftViewEdgeX + configuration.Global.NumTileX
	if leftViewEdgeX < 0 {
		c.X = configuration.Global.ScreenCenterTileX
		c.OutsideXEdges = true
	} else if rightViewEdgeX <= floor_width {
		if leftViewEdgeX == 0 {
			c.OutsideXEdges = xShift < 0
		} else if rightViewEdgeX == floor_width {
			c.OutsideXEdges = xShift > 0
		}
		c.X = characterPosX
	} else {
		c.X = floor_width - (configuration.Global.NumTileX - configuration.Global.ScreenCenterTileX)
		c.OutsideXEdges = true
	}
	c.XShift = xShift

	topViewEdgeY := characterPosY - configuration.Global.ScreenCenterTileY
	bottomViewEdgeY := topViewEdgeY + configuration.Global.NumTileY
	if topViewEdgeY < 0 {
		c.Y = configuration.Global.ScreenCenterTileY
		c.OutsideYEdges = true
	} else if bottomViewEdgeY <= floor_height {
		if topViewEdgeY == 0 {
			c.OutsideYEdges = yShift < 0
		} else if bottomViewEdgeY == floor_height {
			c.OutsideYEdges = yShift > 0
		}
		c.Y = characterPosY
	} else {
		c.Y = floor_height - (configuration.Global.NumTileY - configuration.Global.ScreenCenterTileY)
		c.OutsideYEdges = true
	}
	c.YShift = yShift

}
