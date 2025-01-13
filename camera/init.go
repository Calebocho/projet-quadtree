package camera

// Init met en place une caméra.
func (c *Camera) Init(characterPosX, characterPosY, xShift, yShift int) {
	c.X = characterPosX
	c.Y = characterPosY
	c.XShift = xShift
	c.YShift = yShift
}
