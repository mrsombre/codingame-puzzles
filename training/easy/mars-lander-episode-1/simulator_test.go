package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimulate(t *testing.T) {
	tests := []struct {
		name     string
		game     func() Game
		commands Commands
		want     Lander
	}{
		{
			name: `landed`,
			game: func() Game {
				field := InputField(DataImport(debugField))
				lander := Lander{
					Pos:   Point{2000, 110},
					Speed: Point{10, -39},
					Angle: 0,
					Power: 4,
					Fuel:  100,
				}
				return Game{field, lander}
			},
			commands: Commands{MoveCommand{0, 4}, MoveCommand{0, 4}},
			want: Lander{
				Pos:   Point{2002.573, 100},
				Speed: Point{10, -38.711},
				Angle: 0,
				Power: 4,
				Fuel:  96,
			},
		},
		{
			name: `lost`,
			game: func() Game {
				field := InputField(DataImport(debugField))
				lander := Lander{
					Pos:   Point{70, 2475},
					Speed: Point{-71, 5},
					Angle: 15,
					Power: 4,
					Fuel:  226,
				}
				return Game{field, lander}
			},
			commands: Commands{MoveCommand{15, 4}, MoveCommand{15, 4}},
			want: Lander{
				Pos:   Point{-1.51, 2480},
				Speed: Point{-72, 5.152},
				Angle: 15,
				Power: 4,
				Fuel:  222,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			game := tc.game()
			got := Simulate(game, tc.commands)

			lander := got.Lander
			assert.InDelta(t, tc.want.Pos.X, lander.Pos.X, 0.1)
			assert.InDelta(t, tc.want.Pos.Y, lander.Pos.Y, 0.1)
			assert.InDelta(t, tc.want.Speed.X, lander.Speed.X, 0.1)
			assert.InDelta(t, tc.want.Speed.Y, lander.Speed.Y, 0.1)
			assert.InDelta(t, tc.want.Angle, lander.Angle, 0.1)
			assert.InDelta(t, tc.want.Power, lander.Power, 0.1)
			assert.InDelta(t, tc.want.Fuel, lander.Fuel, 0.1)
		})
	}
}

func TestSimulateFall(t *testing.T) {
	game := testGame()
	commands := make(Commands, 0, 100)
	for i := 0; i < 100; i++ {
		commands = append(commands, MoveCommand{0, 0})
	}
	game = Simulate(game, commands)

	lander := game.Lander
	assert.InDelta(t, 2500, lander.Pos.X, 0.1)
	assert.InDelta(t, 100, lander.Pos.Y, 0.1)
	assert.InDelta(t, 0, lander.Speed.X, 0.1)
	assert.InDelta(t, -133.59, lander.Speed.Y, 0.1)
	assert.InDelta(t, 500, lander.Fuel, 0.1)
}

func TestSimulateCollision(t *testing.T) {
	game := testGame()
	commands := make(Commands, 0, 100)
	power := 0.0
	for i := 0; i < 100; i++ {
		power = math.Min(3, power+1.0)
		commands = append(commands, MoveCommand{-15, power})
	}
	game = Simulate(game, commands)

	lander := game.Lander
	assert.InDelta(t, 3944, lander.Pos.X, 0.1)
	assert.InDelta(t, 760.763, lander.Pos.Y, 0.1)
	assert.InDelta(t, 47.363, lander.Speed.X, 0.1)
	assert.InDelta(t, -53.317, lander.Speed.Y, 0.1)
	assert.InDelta(t, 317, lander.Fuel, 0.1)
}

// Benchmarks

func BenchmarkSimulate(b *testing.B) {
	field := InputField(DataImport(debugField))
	lander := InputLander(DataImport(debugLander))
	game := Game{field, lander}
	chromosome := RandChromosome()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Simulate(game, chromosome.Genes)
	}
}
