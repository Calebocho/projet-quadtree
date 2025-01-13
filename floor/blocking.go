package floor

import "gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"

// Blocking retourne, étant donnée la position du personnage,
// un tableau de booléen indiquant si les cases au dessus (0),
// à droite (1), au dessous (2) et à gauche (3) du personnage
// sont bloquantes.
func (f Floor) Blocking(characterXPos, characterYPos, camXPos, camYPos int) (blocking [4]bool) {

	relativeXPos := characterXPos - camXPos + configuration.Global.ScreenCenterTileX
	relativeYPos := characterYPos - camYPos + configuration.Global.ScreenCenterTileY

	blocking[0] = relativeYPos <= 0 || IsCellBlocking(f.content[relativeYPos-1][relativeXPos])
	blocking[1] = relativeXPos >= configuration.Global.NumTileX-1 || IsCellBlocking(f.content[relativeYPos][relativeXPos+1])
	blocking[2] = relativeYPos >= configuration.Global.NumTileY-1 || IsCellBlocking(f.content[relativeYPos+1][relativeXPos])
	blocking[3] = relativeXPos <= 0 || IsCellBlocking(f.content[relativeYPos][relativeXPos-1])
	return blocking
}

// Vérifie si la case "cell" est un terrain qui n'est pas bloquant pour les déplacements
// du personnage.
func IsCellBlocking(cell int) bool {
	// Il n'est logique de marcher ni dans le vide, ni sur l'eau, ni sur un mur
	return cell == -1 || (configuration.Global.Blocking && (cell == 4) || (cell == 2))
}
