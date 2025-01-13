package quadtree

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// GetContent remplit le tableau contentHolder (qui représente
// un terrain dont la case le plus en haut à gauche a pour coordonnées
// (topLeftX, topLeftY)) à partir du qadtree q.
func (q Quadtree) GetContent(topLeftX, topLeftY int, contentHolder [][]int) {
	// Il peut arriver que la zone à afficher (contentHolder) soit plus grande
	// que celle que le quadtree gère, ou que les topLeftX/Y soient négatifs:
	// dans ce cas là il est plus simple de laisser seulement la gestion des
	// zones qui existent dans le quadtree à fillContentFromQuadtreeNode, donc
	// on gère les zones vides (avec -1) ici
	contentWidth := len(contentHolder[0])
	contentHeight := len(contentHolder)
	for line := 0; line < contentHeight; line++ {
		for column := 0; column < contentWidth; column++ {
			contentHolder[line][column] = -1
		}
	}

	visible_width := min(q.width - abs(topLeftX), contentWidth)
	visible_height := min(q.height - abs(topLeftY), contentHeight)
	if q.width <= contentWidth {
		visible_width = q.width
	}
	if q.height <= contentHeight {
		visible_height = q.height
	}

	var contentX, contentY int
	if topLeftX < 0 {
		contentX = -topLeftX
		topLeftX = 0
	}
	if topLeftY < 0 {
		contentY = -topLeftY
		topLeftY = 0
	}

	fillContentFromNode(q.root, topLeftX, topLeftY, contentX, contentY, visible_width, visible_height, contentHolder)
}

// Remplit la zone de contentHolder correspondant au rectangle de taille width * height
// positionné avec son coin haut gauche en (contentX, contentY) et récupéré
// à partir du coin haut gauche "topLeft" en (topLeftX, topLeftY) du noeud de quadtree.
func fillContentFromNode(node *node, topLeftX, topLeftY, contentX, contentY, width, height int, contentHolder [][]int) {
	if node == nil {
		return
	}

	if node.isLeaf {
		for y := node.topLeftY; y < node.topLeftY + node.height; y++ {
			contentHolderY := contentY + y - topLeftY
			if contentHolderY >= len(contentHolder) || y >= topLeftY + height {
				break
			}
			if y < topLeftY {
				continue
			}

			for x := node.topLeftX; x < node.topLeftX + node.width; x++ {
				contentHolderX := contentX + x - topLeftX
				if contentHolderX >= len(contentHolder[contentHolderY]) || x >= topLeftX + width {
					break
				}
				if x < topLeftX {
					continue
				}
				contentHolder[contentHolderY][contentHolderX] = node.content
			}
		}
		return
	}

	fillContentFromNode(node.topLeftNode, topLeftX, topLeftY, contentX, contentY, width, height, contentHolder)
	fillContentFromNode(node.topRightNode, topLeftX, topLeftY, contentX, contentY, width, height, contentHolder)
	fillContentFromNode(node.bottomLeftNode, topLeftX, topLeftY, contentX, contentY, width, height, contentHolder)
	fillContentFromNode(node.bottomRightNode, topLeftX, topLeftY, contentX, contentY, width, height, contentHolder)
}