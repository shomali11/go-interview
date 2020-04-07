package grids

import (
	"fmt"
	"strconv"

	"github.com/shomali11/go-interview/algorithms/astar"
)

// New creates a 2D grid of nodes
func New(rows, columns int) *GridWorld {
	grid := make([][]*GridNode, rows)
	world := GridWorld{grid: grid}

	for i := range grid {
		grid[i] = make([]*GridNode, columns)
	}

	for row := 0; row < rows; row++ {
		for column := 0; column < columns; column++ {
			grid[row][column] = &GridNode{X: row, Y: column, Cost: 1, world: &world}
		}
	}
	return &world
}

// GridWorld grid world
type GridWorld struct {
	grid [][]*GridNode
}

// Get retrieves a Grid Node
func (n *GridWorld) Get(i int, j int) *GridNode {
	return n.grid[i][j]
}

// IsValid checks whether coordinates are within the grid
func (n *GridWorld) IsValid(i int, j int) bool {
	return i >= 0 && i < len(n.grid) && j >= 0 && j < len(n.grid[0])
}

// PrintGrid prints a Grid
func (n *GridWorld) PrintGrid(nodes []astar.Node) {
	rows := len(n.grid)
	if rows == 0 {
		return
	}

	columns := len(n.grid[0])

	printGrid := make([][]string, rows)
	for i := range printGrid {
		printGrid[i] = make([]string, columns)
	}

	for i := range printGrid {
		for j := range printGrid[0] {
			node := n.Get(i, j)
			if node.IsObstacle {
				printGrid[i][j] = "X"
			} else {
				printGrid[i][j] = strconv.FormatFloat(node.Cost, 'f', 0, 64)
			}
		}
	}

	for _, value := range nodes {
		node := value.(*GridNode)
		printGrid[node.X][node.Y] = "*"
	}

	for i := range printGrid {
		fmt.Println(printGrid[i])
	}
}
