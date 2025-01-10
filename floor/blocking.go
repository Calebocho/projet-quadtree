package floor

import "gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"

// Blocking retourne, étant donnée la position du personnage,
// un tableau de booléen indiquant si les cases au dessus (0),
// à droite (1), au dessous (2) et à gauche (3) du personnage
// sont bloquantes.
func (f Floor) Blocking(characterXPos, characterYPos, camXPos, camYPos int) (blocking [4]bool) {

	relativeXPos := characterXPos - camXPos + configuration.Global.ScreenCenterTileX
	relativeYPos := characterYPos - camYPos + configuration.Global.ScreenCenterTileY

	blocking[0] = relativeYPos <= 0 || f.isCellBlocking(relativeXPos, relativeYPos-1)
	blocking[1] = relativeXPos >= configuration.Global.NumTileX-1 || f.isCellBlocking(relativeXPos+1, relativeYPos)
	blocking[2] = relativeYPos >= configuration.Global.NumTileY-1 || f.isCellBlocking(relativeXPos, relativeYPos+1)
	blocking[3] = relativeXPos <= 0 || f.isCellBlocking(relativeXPos-1, relativeYPos)
	return blocking
}

func (f Floor) isCellBlocking(relative_x, relative_y int) bool {
	cell := f.content[relative_y][relative_x]

	// Il n'est logique de marcher ni dans le vide, ni sur l'eau, ni sur un mur
	return cell == -1 || (configuration.Global.Blocking && (cell == 4) || (cell == 2))
}
