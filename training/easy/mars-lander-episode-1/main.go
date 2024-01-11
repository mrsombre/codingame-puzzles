package main

import (
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

func exec(a ...any) {
	fmt.Println(a...)
}

func main() {
	scanner := NewScanner(os.Stdin)
	_ = ReadGame(scanner)

	lander := &Lander{}
	pid := &PID{Kp: 4.0, Ki: 3.0, Kd: 3.0, IntegralMax: 10.0, IntegralMin: -10.0}

	expSpeed := -MaxVSpeed + 1

	for {
		turnData := ReadTurn(scanner)
		UpdateLander(turnData[0], lander)

		control := pid.Control(expSpeed, lander.VSpeed)
		if control <= 0 {
			if lander.Power > 0 {
				lander.Power = lander.Power - TurnPower
			}
		} else {
			if lander.Power < MaxPower {
				lander.Power = lander.Power + TurnPower
			}
		}

		exec(0, lander.Power)
	}
}
