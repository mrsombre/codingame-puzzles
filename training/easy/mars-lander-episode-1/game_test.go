package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

var debugField = `H4sIAAAAAAAA/4pWMlPSUTJQMDQwUNJRMjQwMFAwhTBNDWCixgZwpimYCVFhZmlpCRI2UIoFBAAA//91mcCRRgAAAA==`
var debugLander = `H4sIAAAAAAAA/4pWMjI1MFAAEyAIpZViAQEAAP//bKzxWhkAAAA=`
var debugStep = `H4sIAAAAAAAA/4pWMjI1MFAwMjYzUTBQ0DUyVzCxNFAwUDBUigUEAAD//3t0J7wbAAAA`

func testGame() Game {
	return Game{
		Field:  InputField(DataImport(debugField)),
		Lander: InputLander(DataImport(debugLander)),
	}
}

func TestLander_Move(t *testing.T) {
	tests := []struct {
		name   string
		lander Lander
		angle  float64
		power  float64
		want   Lander
	}{
		{
			name: `zero`,
			lander: Lander{
				Pos: Point{2500, 2500},
			},
			want: Lander{
				Pos:   Point{2500, 2498.1445},
				Speed: Point{0, -g},
			},
		},
		{
			name: `fall`,
			lander: Lander{
				Pos:   Point{2500, 2498.1445},
				Speed: Point{0, -g},
			},
			want: Lander{
				Pos:   Point{2500, 2492.578},
				Speed: Point{0, -2 * g},
			},
		},
		{
			name: `left turn`,
			lander: Lander{
				Pos:  Point{2500, 2500},
				Fuel: 500,
			},
			angle: left,
			power: 4,
			want: Lander{
				Pos:   Point{2498, 2498.1445},
				Speed: Point{-4, -g},
				Fuel:  496,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			move := tc.lander.Move(tc.angle, tc.power)
			assert.InDelta(t, tc.want.Pos.X, move.Pos.X, 0.1)
			assert.InDelta(t, tc.want.Pos.Y, move.Pos.Y, 0.1)
			assert.InDelta(t, tc.want.Speed.X, move.Speed.X, 0.1)
			assert.InDelta(t, tc.want.Speed.Y, move.Speed.Y, 0.1)
			assert.EqualValues(t, tc.want.Fuel, move.Fuel)
		})
	}
}

func TestEvolutionMiddle(t *testing.T) {
	field := InputField(DataImport(debugField))
	lander := InputLander(DataImport(debugStep))

	e := NewEvolution(Game{
		Field:  field,
		Lander: lander,
	})
	e.Evolve(100)

	chromosomes := e.Population.Chromosomes
	assert.Len(t, chromosomes, populationSize)

	sort.Slice(chromosomes, func(i, j int) bool {
		return chromosomes[i].Fitness > chromosomes[j].Fitness
	})

	best := chromosomes[0]
	r := Simulate(e.Game, best.Genes)

	asJsonPretty(r.Lander)
}

// Benchmarks

var (
	GlobalLander Lander
)
