package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAshTurnsToHuman(t *testing.T) {
	var h human
	ash := hero{point{0, 0}}

	h = human{crd: point{1500, 0}}
	assert.EqualValues(t, 0, ash.safeTurnsToHuman(h))

	h = human{crd: point{2500, 0}}
	assert.EqualValues(t, 1, ash.safeTurnsToHuman(h))

	h = human{crd: point{3500, 0}}
	assert.EqualValues(t, 2, ash.safeTurnsToHuman(h))
}
