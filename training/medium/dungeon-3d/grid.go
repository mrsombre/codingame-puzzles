package main

import (
	"fmt"
)

type Cell struct {
	ID int

	X, Y, Z    int
	Directions []*Cell

	IsWall bool
}

type Grid struct {
	Levels int
	Width  int
	Height int

	Cells [][][]*Cell

	Start, End *Cell
}

func NewGridFromInput(input []string) *Grid {
	grid := &Grid{}

	fmt.Sscan(input[0], &grid.Levels, &grid.Height, &grid.Width)

	id := 0
	ln := 2
	grid.Cells = make([][][]*Cell, grid.Levels)
	for z := 0; z < grid.Levels; z++ {
		grid.Cells[z] = make([][]*Cell, grid.Height)
		ln++
		for y := 0; y < grid.Height; y++ {
			grid.Cells[z][y] = make([]*Cell, grid.Width)
			x := 0
			for _, c := range input[ln] {
				cell := &Cell{
					ID:     id,
					Z:      z,
					Y:      y,
					X:      x,
					IsWall: c == '#',
				}

				switch c {
				case 'S':
					grid.End = cell
				case 'A':
					grid.Start = cell
				}

				grid.Cells[z][y][x] = cell
				id++
				x++
			}
			ln++
		}
	}

	for z := 0; z < grid.Levels; z++ {
		for y := 0; y < grid.Height; y++ {
			for x := 0; x < grid.Width; x++ {
				cell := grid.Cells[z][y][x]
				if cell.IsWall {
					continue
				}

				// left
				if x > 0 {
					cell.Directions = append(cell.Directions, grid.Cells[z][y][x-1])
				}
				// right
				if x < grid.Width-1 {
					cell.Directions = append(cell.Directions, grid.Cells[z][y][x+1])
				}
				// top
				if y > 0 {
					cell.Directions = append(cell.Directions, grid.Cells[z][y-1][x])
				}
				// bottom
				if y < grid.Height-1 {
					cell.Directions = append(cell.Directions, grid.Cells[z][y+1][x])
				}

				// up
				if z > 0 {
					cell.Directions = append(cell.Directions, grid.Cells[z-1][y][x])
				}
				// down
				if z < grid.Levels-1 {
					cell.Directions = append(cell.Directions, grid.Cells[z+1][y][x])
				}
			}
		}
	}

	return grid
}
