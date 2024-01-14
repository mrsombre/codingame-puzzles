package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var inputGame = []string{
	"4",
	"0 1000",
	"1000 500",
	"2000 500",
	"3000 1000",
}

func TestInputField(t *testing.T) {
	game := InputField(inputGame)

	assert.Equal(t, 3, len(game.Ground))
	assert.Equal(t, Line{Point{0, 1000}, Point{1000, 500}}, game.Ground[0])
	assert.Equal(t, Line{Point{1000, 500}, Point{2000, 500}}, game.Ground[1])
	assert.Equal(t, Line{Point{2000, 500}, Point{3000, 1000}}, game.Ground[2])
	assert.Equal(t, Line{Point{1000, 500}, Point{2000, 500}}, game.Flat)
}

var inputStep = []string{
	"2500 2700 21 32 550 11 1",
}

func TestInputLander(t *testing.T) {
	lander := InputLander(inputStep)

	assert.Equal(t, Point{2500, 2700}, lander.Pos)
	assert.Equal(t, Point{21, 32}, lander.Speed)
	assert.EqualValues(t, 11, lander.Angle)
	assert.EqualValues(t, 1, lander.Power)
	assert.EqualValues(t, 550, lander.Fuel)
}
