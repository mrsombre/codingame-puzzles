package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func debugLander() *Lander {
	data := DataImport(strings.TrimSpace(debugInputTurn))
	lander := &Lander{}
	UpdateLander(data[0], lander)
	return lander
}

func TestUpdateLander(t *testing.T) {
	lander := debugLander()

	assert.Equal(t, 2500, lander.X)
	assert.Equal(t, 2500, lander.Y)

	assert.Equal(t, 0.0, lander.HSpeed)
	assert.Equal(t, 0.0, lander.VSpeed)

	assert.Equal(t, 500, lander.Fuel)
	assert.Equal(t, 0, lander.Rotate)
}
