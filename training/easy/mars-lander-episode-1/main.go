package main

import (
	"bufio"
	"math/rand"
	"os"
	"runtime"
	"sort"
	"time"
)

var rnd *rand.Rand

func init() {
	runtime.GOMAXPROCS(1)
	rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
	debug = true
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	dataGame := ReadGame(scanner)
	asText(DataExport(dataGame))
	field := InputField(dataGame)

	dataStep := ReadStep(scanner)
	asText(DataExport(dataStep))
	lander := InputLander(dataStep)

	game := Game{field, lander}
	e := NewEvolution(game)
	e.Evolve(100)

	chromosomes := e.Population.Chromosomes
	sort.Slice(chromosomes, func(i, j int) bool {
		return chromosomes[i].Fitness > chromosomes[j].Fitness
	})

	best := chromosomes[0]
	commands := best.Genes

	ExecuteCommand(Commands{commands[0]})
	commands = commands[1:]

	for {
		dataStep = ReadStep(scanner)
		asText(DataExport(dataStep))
		lander = InputLander(dataStep)

		ExecuteCommand(Commands{commands[0]})
		commands = commands[1:]
	}
}
