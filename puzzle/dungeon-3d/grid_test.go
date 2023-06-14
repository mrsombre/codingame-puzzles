package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var debugInput = `H4sIAAAAAAAA/4pWMlUwVjBR0lEyMlDSUVLSUdLT01NW0lFSVlbWg1DBUHFlPT0UcWWYuDJEvR4OcZg0VNwRbj6cigUEAAD///bULIiGAAAA`

var testGrid = func() *Grid {
	r := strings.NewReader(strings.TrimSpace(testInputGame))
	s := NewScanner(r)
	data := ReadGame(s)
	return NewGridFromInput(data)
}

func TestNewGridFromInput(t *testing.T) {
	grid := testGrid()

	assert.Len(t, grid.Cells, 2)

	assert.Same(t, grid.Start, grid.Cells[0][2][0])
	assert.Same(t, grid.End, grid.Cells[1][1][2])
}

func TestGrid_CellSides(t *testing.T) {
	grid := testGrid()

	tests := []struct {
		name     string
		cell     *Cell
		expected []*Cell
	}{
		{
			name: `start`,
			cell: grid.Cells[0][2][0],
			expected: []*Cell{
				grid.Cells[0][2][1],
				grid.Cells[0][1][0],
				grid.Cells[1][2][0],
			},
		},
		{
			name: `bottom`,
			cell: grid.Cells[0][0][1],
			expected: []*Cell{
				grid.Cells[0][0][0],
				grid.Cells[0][0][2],
				grid.Cells[0][1][1],
				grid.Cells[1][0][1],
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.cell.Directions)
		})
	}
}

func TestDebug(t *testing.T) {
	data := DataImport(debugInput)
	grid := NewGridFromInput(data)

	asJsonPretty(grid)
}
