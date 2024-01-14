package main

import (
	"fmt"
	"math"
)

func Simulate(game Game, commands Commands) Game {
	lander := game.Lander

	var move Line
	for _, command := range commands {
		switch c := command.(type) {
		case MoveCommand:
			game.Lander = lander.Move(c.Angle, c.Power)

			// check rules
			if math.Abs(game.Lander.Angle) > angleMax {
				panic(fmt.Sprintf("invalid angle: %s", game.Lander))
			}
			if math.Abs(game.Lander.Angle-lander.Angle) > angleStep {
				panic(fmt.Sprintf("invalid angle step: %s", game.Lander))
			}
			if game.Lander.Power < 0 || game.Lander.Power > powerMax {
				panic(fmt.Sprintf("invalid power: %s", game.Lander))
			}
			if math.Abs(game.Lander.Power-lander.Power) > powerStep {
				panic(fmt.Sprintf("invalid power step: %s", game.Lander))
			}

			// lander left the field
			if !game.IsLanderInBound() {
				return game
			}

			// collision
			move = Line{lander.Pos, game.Lander.Pos}
			for _, line := range game.Field.Ground {
				if lander.Pos.IsInXRange(line.From.X, line.To.X) {
					if p, ok := move.SegmentsIntersection(line); ok {
						game.Lander.Pos = p
						return game
					}
				}
			}

			// next iteration
			lander = game.Lander

		default:
			panic(fmt.Sprintf("invalid command: %v", command))
		}
	}

	return game
}
