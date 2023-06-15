package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInputGame = `
2
0 1500
3500 500
`

var testInputTurn = `
5000 2500 -50 0 1000 90 0
`

func TestReadGame(t *testing.T) {
	input := strings.TrimSpace(testInputGame)

	r := strings.NewReader(input)
	expected := strings.Split(input, "\n")

	scanner := NewScanner(r)
	actual := ReadGame(scanner)

	assert.Equal(t, expected, actual)
}

func TestReadTurn(t *testing.T) {
	input := strings.TrimSpace(testInputTurn)

	r := strings.NewReader(input)
	expected := strings.Split(input, "\n")

	scanner := NewScanner(r)
	actual := ReadTurn(scanner)

	assert.Equal(t, expected, actual)
}
