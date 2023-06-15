package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var debugInputGame = `H4sIAAAAAAAA/4pWMlPSUTJQMDQwUNJRMjQwMFAwhTBNDWCixgZwpimYCVFhZmlpCRI2UIoFBAAA//91mcCRRgAAAA==`

/**
6
0 100
1000 500
1500 100
3000 100
5000 1500
6999 1000
*/

var debugInputTurn = `H4sIAAAAAAAA/4pWMjI1MFAAEyAIpZViAQEAAP//bKzxWhkAAAA=`

/**
2500 2500 0 0 500 0 0
*/

func TestInputGame(t *testing.T) {
	data := DataImport(strings.TrimSpace(debugInputGame))
	assert.Equal(t, 7, len(data))
}

func TestInputTurn(t *testing.T) {
	data := DataImport(strings.TrimSpace(debugInputTurn))
	assert.Equal(t, 1, len(data))
}
