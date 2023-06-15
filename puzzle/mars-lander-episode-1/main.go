package main

import (
	"encoding/json"
	"fmt"
	"os"
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

func main() {
	scanner := NewScanner(os.Stdin)

	gameData := ReadGame(scanner)
	// console(DataExport(gameData))

	grid := NewGridFromInput(gameData)
	lander := &Lander{}
	mla := NewMla(grid, lander)

	for {
		turnData := ReadTurn(scanner)
		// console(DataExport(turnData))

		UpdateLander(turnData[0], lander)

		command := mla.Control()
		exec(0, command.power)
	}
}
