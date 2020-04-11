package grids

import (
	"image/color"
	"log"
	"strconv"

	"github.com/golang/freetype/truetype"
	"github.com/shomali11/go-interview/algorithms/astar"
	"github.com/shomali11/gridder"
	"golang.org/x/image/font/gofont/goregular"
)

func createImage(rows int, columns int, start *GridNode, goal *GridNode, path []astar.Node, world *GridWorld) {
	imageConfig := gridder.ImageConfig{
		Width:  2000,
		Height: 2000,
		Name:   "output.png",
	}
	gridConfig := gridder.GridConfig{
		Rows:              rows,
		Columns:           rows,
		MarginWidth:       32,
		LineStrokeWidth:   2,
		BorderStrokeWidth: 20,
	}

	grid, err := gridder.New(imageConfig, gridConfig)
	if err != nil {
		log.Fatal(err)
	}

	font, err := truetype.Parse(goregular.TTF)
	if err != nil {
		log.Fatal(err)
	}

	fontFace := truetype.NewFace(font, &truetype.Options{Size: 48})
	lineConfig := gridder.PathConfig{Dashes: 10}

	for row := 0; row < rows; row++ {
		for column := 0; column < columns; column++ {
			node := world.Get(row, column)
			if node.IsObstacle {
				grid.PaintCell(row, column, color.NRGBA{R: 0, G: 0, B: 0, A: 255 / 2})
				grid.DrawString(row, column, "Block", fontFace)
			} else {
				cost := strconv.FormatFloat(node.Cost, 'f', 0, 64)
				grid.DrawString(row, column, cost, fontFace)
			}
		}
	}

	currentX := start.X
	currentY := start.Y
	grid.DrawCircle(start.X, start.Y, gridder.CircleConfig{Color: color.NRGBA{R: 255 / 2, G: 0, B: 0, A: 255 / 2}, Radius: 60})
	for i := 1; i < len(path)-1; i++ {
		node := path[i].(*GridNode)
		grid.DrawPath(currentX, currentY, node.X, node.Y, lineConfig)
		currentX = node.X
		currentY = node.Y
	}
	grid.DrawPath(currentX, currentY, goal.X, goal.Y, lineConfig)
	grid.DrawCircle(goal.X, goal.Y, gridder.CircleConfig{Color: color.NRGBA{R: 0, G: 255 / 2, B: 0, A: 255 / 2}, Radius: 60})

	grid.SavePNG()
}
