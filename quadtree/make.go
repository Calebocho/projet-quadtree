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
	isEmpty := width == 0 || height == 0
	if isEmpty || (width == 1 && height == 1) {
		content := floorContent[y][x]

		return &node {
			topLeftX: x,
			topLeftY: y,

			width: width,
			height: height,

			content: content,
			isLeaf: true,

			topLeftNode: nil,
			topRightNode: nil,
			bottomLeftNode: nil,
			bottomRightNode: nil,
		}
	}

	halfWidth := width / 2
	halfHeight := height / 2
	new_node := node {
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

	first_content := new_node.topLeftNode.content
	if new_node.topLeftNode.isLeaf &&
		new_node.topRightNode.isLeaf && new_node.topRightNode.content == first_content &&
		new_node.bottomLeftNode.isLeaf && new_node.bottomLeftNode.content == first_content &&
		new_node.bottomRightNode.isLeaf && new_node.bottomRightNode.content == first_content {
		return &node {
			topLeftX: x,
			topLeftY: y,

			width: width,
			height: height,

			content: first_content,
			isLeaf: true,

			topLeftNode: nil,
			topRightNode: nil,
			bottomLeftNode: nil,
			bottomRightNode: nil,
		}
	} else {
		return &new_node
	}
}