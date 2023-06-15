package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func debugGrid() *Grid {
	data := DataImport(strings.TrimSpace(debugInputGame))
	return NewGridFromInput(data)
}

func TestGridCells(t *testing.T) {
	grid := debugGrid()

	assert.Len(t, grid.Cells, 6)

	assert.Equal(t, 0, grid.Cells[0].X)
	assert.Equal(t, 100, grid.Cells[0].Y)

	assert.Equal(t, 6999, grid.Cells[5].X)
	assert.Equal(t, 1000, grid.Cells[5].Y)
}

func TestGridFlat(t *testing.T) {
	grid := debugGrid()

	assert.Equal(t, grid.Cells[2], grid.Flat[0])
	assert.Equal(t, grid.Cells[3], grid.Flat[1])
}
