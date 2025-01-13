package quadtree

// MakeFromArray construit un quadtree représentant un terrain
// étant donné un tableau représentant ce terrain.
func MakeFromArray(floorContent [][]int) (q Quadtree) {
	if len(floorContent) == 0 || len(floorContent[0]) == 0 {
		return
	}
	q.height = len(floorContent)
	q.width = len(floorContent[0])
	q.root = makeQuadtreeNodeFromArea(floorContent, 0, 0, q.width, q.height)

	return
}

// Retourne un nœud de quadtree à partir de la zone contenue dans floorContent
// définie par son coin haut gauche en (x, y) (pour le floorContent et les positions
// topLeftX/Y du nœud de quadtree) et sa taille width * height.
func makeQuadtreeNodeFromArea(floorContent [][]int, x, y, width, height int) *node {
	if width == 0 || height == 0 {
		return &node {
			topLeftX: x,
			topLeftY: y,

			width: width,
			height: height,

			content: 0,
			isLeaf: true,

			topLeftNode: nil,
			topRightNode: nil,
			bottomLeftNode: nil,
			bottomRightNode: nil,
		}
	}

	isContentUniform := true
	previousContent := floorContent[y][x]
	for column := x; column < x+width && isContentUniform; column++ {
		for line := y; line < y+height; line++ {
			if floorContent[line][column] != previousContent {
				isContentUniform = false
				break
			}
			previousContent = floorContent[line][column]
		}
	}

	if isContentUniform {
		return &node {
			topLeftX: x,
			topLeftY: y,

			width: width,
			height: height,

			content: floorContent[y][x],
			isLeaf: true,

			topLeftNode: nil,
			topRightNode: nil,
			bottomLeftNode: nil,
			bottomRightNode: nil,
		}
	}

	halfWidth := width / 2
	halfHeight := height / 2
	return &node {
		topLeftX: x,
		topLeftY: y,

		width: width,
		height: height,

		content: -1,
		isLeaf: false,

		topLeftNode: makeQuadtreeNodeFromArea(floorContent, x, y, halfWidth, halfHeight),
		topRightNode: makeQuadtreeNodeFromArea(floorContent, x + halfWidth, y, width - halfWidth, halfHeight),
		bottomLeftNode: makeQuadtreeNodeFromArea(floorContent, x, y + halfHeight, halfWidth, height - halfHeight),
		bottomRightNode: makeQuadtreeNodeFromArea(floorContent, x + halfWidth, y + halfHeight, width - halfWidth, height - halfHeight),
	}
}