package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZedTurnsToHuman(t *testing.T) {
	var h human
	zombie := zombie{}

	h = human{crd: point{200, 200}}
	assert.EqualValues(t, 1, zombie.turnsToHuman(h))

	h = human{crd: point{1000, 0}}
	assert.EqualValues(t, 3, zombie.turnsToHuman(h))
}

func TestZedClosestHuman(t *testing.T) {
	zombie := zombie{humanDistance: []float64{1000, 500}}

	cd, cid := zombie.closestHuman()
	assert.EqualValues(t, 500, cd)
	assert.EqualValues(t, 1, cid)
}
