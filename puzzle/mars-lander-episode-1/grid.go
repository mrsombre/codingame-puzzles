package main

import (
	"fmt"
)

type Cell struct {
	X, Y int
}

type Grid struct {
	Cells []*Cell

	Flat []*Cell
}

func NewGridFromInput(data []string) *Grid {
	var size int
	fmt.Sscan(data[0], &size)

	grid := &Grid{}
	grid.Cells = make([]*Cell, 0, size)
	for i := 1; i < len(data); i++ {
		cell := &Cell{}
		fmt.Sscan(data[i], &cell.X, &cell.Y)
		grid.Cells = append(grid.Cells, cell)
	}

	findFlat(grid)

	return grid
}

func findFlat(grid *Grid) {
	var prev *Cell
	for _, cell := range grid.Cells {
		if prev != nil {
			if prev.Y == cell.Y && cell.X-prev.X >= 1000 {
				grid.Flat = append(grid.Flat, prev)
				grid.Flat = append(grid.Flat, cell)
			}
		}
		prev = cell
	}
}
