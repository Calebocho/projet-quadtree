package floor

import "log"

// Vérifie si le point (x, y) correspond au téléporteur numéro "teleporter_num".
// Il y a seulement deux téléporteurs, liés entre eux, donc teleporter_num <= 1
func (f Floor) isOnTeleporter(teleporter_num int, x, y int) bool {
	if teleporter_num > len(f.teleporters) {
		log.Fatal("Il y a seulement", len(f.teleporters), "téléporteurs mais celui d'indice", teleporter_num, "est demandé")
	}

	teleporter := f.teleporters[teleporter_num]
	return teleporter != nil && teleporter.X == x && teleporter.Y == y
}

// Vérifie si le point (x, y) correspond à un des téléporteurs associés à ce sol.
func (f Floor) IsOnAnyTeleporter(x, y int) bool {
	return f.isOnTeleporter(0, x, y) || f.isOnTeleporter(1, x, y)
}

// Vérifie si le point (x, y) est bien contenu par la surface du sol.
func (f Floor) IsInside(x, y int) bool {
	return x >= 0 && y >= 0 && x < f.GetWidth() && y < f.GetHeight()
}

// Retourne la position du téléporteur sur lequel se trouvant en (x, y), ou
// nil s'il n'y a pas de téléporteur à cette position.
func (f Floor) Teleport(x, y int) *Pos {
	if f.isOnTeleporter(0, x, y) {
		return f.teleporters[1]
	} else if f.isOnTeleporter(1, x, y) {
		return f.teleporters[0]
	} else {
		return nil
	}
}

// S'assure que l'on ne se trouve pas sur un téléporteur existant et que la position
// demandée pour le nouveau téléporteur est bien valide pour que le personnage
// puisse y marcher, puis crée un nouveau téléporteur à la place du plus ancien
// créé, en échangeant les téléporteurs existants préalablement de manière à ce
// que les téléporteurs successifs les plus récents soient conservés.
//
// Retourne true si le téléporteur a pu être placé, false sinon.
func (f *Floor) SetNewTeleporter(teleporter_x, teleporter_y int) bool {
	if f.IsOnAnyTeleporter(teleporter_x, teleporter_y) ||
		!f.IsInside(teleporter_x, teleporter_y) ||
		IsCellBlocking(f.content[teleporter_y - f.topLeftY][teleporter_x - f.topLeftX]) {
		return false
	}

	teleporterPos := Pos {X: teleporter_x, Y: teleporter_y}
	if f.teleporters[0] == nil {
		f.teleporters[0] = &teleporterPos
	} else {
		if f.teleporters[1] == nil {
			f.teleporters[1] = &teleporterPos
		} else {
			// On remplace le portail le plus ancien par le plus récent et on crée un nouveau
			// portail pour remplacer l'ancien, à la position actuelle
			f.teleporters[0], f.teleporters[1] = f.teleporters[1], &teleporterPos
		}
	}

	return true
}