package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBFS(t *testing.T) {
	grid := testGrid()

	actual := BFS(grid)
	assert.Equal(t, "6", actual)
}
