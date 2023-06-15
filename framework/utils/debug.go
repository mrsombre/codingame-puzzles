package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

const debug = true

func console(a ...any) {
	if !debug {
		return
	}
	fmt.Fprintln(os.Stderr, a...)
}

func asJson(a any) {
	if !debug {
		return
	}
	b, _ := json.Marshal(a)
	console(string(b))
}

func asJsonPretty(a any) {
	if !debug {
		return
	}
	b, _ := json.MarshalIndent(a, ``, `  `)
	console(string(b))
}
