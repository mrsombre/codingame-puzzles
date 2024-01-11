package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

func console(a ...any) {
	fmt.Fprintln(os.Stderr, a...)
}

func exec(a ...any) {
	fmt.Println(a...)
}

func asJsonPretty(a any) {
	b, _ := json.MarshalIndent(a, ``, `  `)
	console(string(b))
}

func intToStr(x int) string {
	return strconv.Itoa(x)
}

func main() {
	scanner := NewScanner(os.Stdin)
	data := ReadGame(scanner)

	grid := NewGridFromInput(data)
	answer := BFS(grid)

	exec(answer)
}
