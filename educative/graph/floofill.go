package graph

import "errors"

var deltaRow = [...]int{-1, 0, 1, 0}
var deltaCol = [...]int{0, 1, 0, -1}

type coordinate struct {
	R int
	C int
}

func newCoordinate(R, C int) *coordinate {
	return &coordinate{R, C}
}

func (c *coordinate) isWellWithinBoundary(numRows, numCols int) bool {
	return (0 <= c.R && c.R < numRows) && (0 <= c.C && c.C < numCols)
}

type coordinateQueue struct {
	Elements []coordinate
}

func getCoordinateQueue() *coordinateQueue {
	return &coordinateQueue{
		Elements: make([]coordinate, 0),
	}
}

func (q *coordinateQueue) enqueue(c coordinate) {
	q.Elements = append(q.Elements, c)
}

func (q *coordinateQueue) size() int {
	return len(q.Elements)
}

func (q *coordinateQueue) isEmpty() bool {
	return q.size() == 0
}

func (q *coordinateQueue) dequeue() (coordinate, error) {
	if !q.isEmpty() {
		element := q.Elements[0]
		q.Elements = q.Elements[1:len(q.Elements)]
		return element, nil
	}

	return coordinate{}, errors.New("empty queue")
}

func bfs(image [][]int, root coordinate, replacementColor int, numRows, numCols int) {
	q := getCoordinateQueue()

	q.enqueue(root)

	visited := make([][]bool, 0)

	for i := 0; i < numRows; i++ {
		visited[i] = make([]bool, numCols)
	}

	rootColor := image[root.R][root.C]
	//replace root color
	image[root.R][root.C] = replacementColor
	visited[root.R][root.C] = true

	for !q.isEmpty() {
		node, error := q.dequeue()
		if error == nil {
			neighbors := GetNeighborsWithSameColor(image, node, rootColor, numRows, numCols)

			for _, neighbor := range neighbors {
				if visited[neighbor.R][neighbor.C] {
					continue
				} else {
					image[neighbor.R][neighbor.C] = replacementColor
					q.enqueue(neighbor)
					visited[neighbor.R][neighbor.C] = true
				}
			}
		}
	}

}

func GetNeighborsWithSameColor(image [][]int, node coordinate, rootColor int, numRows, numCols int) []coordinate {
	neighbors := make([]coordinate, 0)

	for i := range deltaRow {
		neighborRow := node.C + deltaRow[i]
		neighborCol := node.R + deltaCol[i]

		neighborCell := coordinate{neighborRow, neighborCol}

		if neighborCell.isWellWithinBoundary(numRows, numCols) && image[neighborRow][neighborCol] == rootColor {
			neighbors = append(neighbors, neighborCell)
		}
	}
	return neighbors
}
