package floor

func (f Floor) isOnTeleporter(teleporter_num int, x, y int) bool {
	teleporter := f.teleporters[teleporter_num]
	return teleporter != nil && teleporter.X == x && teleporter.Y == y
}

func (f Floor) IsOnAnyTeleporter(x, y int) bool {
	return f.isOnTeleporter(0, x, y) || f.isOnTeleporter(1, x, y)
}

func (f Floor) IsInside(x, y int) bool {
	return x >= 0 && y >= 0 && x < f.GetWidth() && y < f.GetHeight()
}

func (f Floor) Teleport(x, y int) *Pos {
	if f.isOnTeleporter(0, x, y) {
		return f.teleporters[1]
	} else if f.isOnTeleporter(1, x, y) {
		return f.teleporters[0]
	} else {
		return nil
	}
}

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