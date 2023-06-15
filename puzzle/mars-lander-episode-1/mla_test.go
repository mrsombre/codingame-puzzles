package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func debugMla() *Mla {
	grid := debugGrid()
	lander := debugLander()
	return NewMla(grid, lander)
}

func TestMlaControl(t *testing.T) {
	mla := debugMla()

	s := 1
	for {
		command := mla.Control()
		mla.lander.Power = command.power

		mla.lander.VSpeed -= G - float64(mla.lander.Power)
		distance := calculateDistance(mla.lander.VSpeed, G, 1)
		mla.lander.Y += int(math.Round(distance))

		groundHeight := mla.GetGroundHeight()
		if groundHeight < 0 {
			break
		}
		s++
	}

	assert.Less(t, -mla.lander.VSpeed, float64(MaxVSpeed))
}

func TestMla_GetYFromX(t *testing.T) {
	mla := debugMla()

	assert.Equal(t, 100, mla.GetYFromX(2500))
}

func TestMla_GetGroundHeight(t *testing.T) {
	mla := debugMla()

	assert.Equal(t, 2400, mla.GetGroundHeight())
}
