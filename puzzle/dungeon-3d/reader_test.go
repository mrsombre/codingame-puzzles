package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInputGame = `
2 3 3
8

..#
.##
A##

#.#
#.S
#..
`

func TestReadGame(t *testing.T) {
	input := strings.TrimSpace(testInputGame)

	r := strings.NewReader(input)
	expected := strings.Split(input, "\n")

	scanner := NewScanner(r)
	actual := ReadGame(scanner)

	assert.Equal(t, expected, actual)
}
