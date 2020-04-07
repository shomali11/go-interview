package grids

import (
	"math"

	"github.com/shomali11/go-interview/algorithms/astar"
)

// GridNode grid node
type GridNode struct {
	X          int
	Y          int
	Cost       float64
	IsObstacle bool
	world      *GridWorld
}

// EstimateDistanceToGoal get distance to the goal
func (n *GridNode) EstimateDistanceToGoal(goal astar.Node) float64 {
	return distance(n, goal.(*GridNode))
}

// ActualDistanceToNeighbor get distance to a neighbor
func (n *GridNode) ActualDistanceToNeighbor(node astar.Node) float64 {
	if isDiagonal(n, node.(*GridNode)) {
		return math.Sqrt2 * n.Cost
	}
	return n.Cost
}

// GetNeighbors get neighbors
func (n *GridNode) GetNeighbors() []astar.Node {
	nodes := make([]astar.Node, 0)

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			// If Same Object ...
			if i == 0 && j == 0 {
				// Skip ...
				continue
			}

			if !n.world.IsValid(n.X+i, n.Y+j) {
				continue
			}

			node := n.world.Get(n.X+i, n.Y+j)
			if node.IsObstacle {
				continue
			}

			if n.isBehindObstacles(node) {
				continue
			}
			nodes = append(nodes, node)
		}
	}
	return nodes
}

func (n *GridNode) isBehindObstacles(node *GridNode) bool {
	dx := node.X - n.X
	dy := node.Y - n.Y

	i := node.X
	j := node.Y

	if dx == 1 && dy == 1 { // Lower Right
		if n.world.IsValid(i, j-1) && n.world.IsValid(i-1, j) {
			right := n.world.Get(i, j-1)
			bottom := n.world.Get(i-1, j)

			if right.IsObstacle && bottom.IsObstacle {
				return true
			}
		}
	} else if dx == -1 && dy == -1 { // Upper Left
		if n.world.IsValid(i+1, j) && n.world.IsValid(i, j+1) {
			top := n.world.Get(i+1, j)
			left := n.world.Get(i, j+1)

			if top.IsObstacle && left.IsObstacle {
				return true
			}
		}
	} else if dx == 1 && dy == -1 { // Upper Right
		if n.world.IsValid(i-1, j) && n.world.IsValid(i, j+1) {
			top := n.world.Get(i-1, j)
			right := n.world.Get(i, j+1)

			if top.IsObstacle && right.IsObstacle {
				return true
			}
		}
	} else if dx == -1 && dy == 1 { // Lower Left
		if n.world.IsValid(i, j-1) && n.world.IsValid(i+1, j) {
			left := n.world.Get(i, j-1)
			bottom := n.world.Get(i+1, j)

			if left.IsObstacle && bottom.IsObstacle {
				return true
			}
		}
	}
	return false
}

func isDiagonal(n *GridNode, node *GridNode) bool {
	dx := node.X - n.X
	dy := node.Y - n.Y

	if dx == 1 && dy == 1 { // Lower Right
		return true
	} else if dx == -1 && dy == -1 { // Upper Left
		return true
	} else if dx == 1 && dy == -1 { // Upper Right
		return true
	} else if dx == -1 && dy == 1 { // Lower Left
		return true
	}
	return false
}

func distance(current *GridNode, other *GridNode) float64 {
	currentX := float64(current.X)
	currentY := float64(current.Y)
	otherX := float64(other.X)
	otherY := float64(other.Y)
	return math.Sqrt(math.Pow(currentX-otherX, 2) + math.Pow(currentY-otherY, 2))
}
