package main

// A set of helper methods for outputting debug information
// to the Stderr stream in text or JSON format.

import (
	"encoding/json"
	"fmt"
	"os"
)

var debug = true
var debugOutput = os.Stderr

func asText(a ...any) {
	if !debug {
		return
	}
	fmt.Fprintln(debugOutput, a...)
}

func asJson(a any) {
	if !debug {
		return
	}
	b, _ := json.Marshal(a)
	asText(string(b))
}

func asJsonPretty(a any) {
	if !debug {
		return
	}
	b, _ := json.MarshalIndent(a, ``, `  `)
	asText(string(b))
}

func u(a ...any) {
}
