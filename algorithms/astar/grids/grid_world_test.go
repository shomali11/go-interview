package grids

import (
	"fmt"
	"testing"

	"github.com/shomali11/go-interview/algorithms/astar"
)

func TestSearch(t *testing.T) {
	columns := 9
	rows := 9
	world := New(rows, columns)

	start := world.Get(0, 0)
	goal := world.Get(rows-1, columns-1)

	world.Get(1, 2).IsObstacle = true
	world.Get(2, 1).IsObstacle = true
	world.Get(3, 0).IsObstacle = true
	world.Get(0, 2).Cost = 3
	world.Get(0, 3).Cost = 3

	aStar := astar.New()
	path, distance := aStar.Search(start, goal)

	fmt.Println("Distance:", distance)
	fmt.Println()

	world.PrintGrid(path)

	createImage(rows, columns, start, goal, path, world)
}
