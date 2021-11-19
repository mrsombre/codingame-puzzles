package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDistance(t *testing.T) {
	var a, b point

	a = point{0, 0}
	b = point{0, 5}
	assert.EqualValues(t, 5, distance(a, b))

	a = point{0, 5}
	b = point{0, 0}
	assert.EqualValues(t, 5, distance(a, b))

	a = point{0, 0}
	b = point{4, 4}
	assert.EqualValues(t, 5, distance(a, b))

	a = point{8, 8}
	b = point{4, 4}
	assert.EqualValues(t, 5, distance(a, b))
}
