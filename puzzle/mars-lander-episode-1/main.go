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

	pid := &PID{
		Kp: 1.85,
		Ki: 0.04,
		Kd: 16.00,
	}

	isControl := false
	expSpeed := -MaxVSpeed + (MaxPower-G)*4

	for {
		turnData := ReadTurn(scanner)
		UpdateLander(turnData[0], lander)

		if !isControl && lander.VSpeed < expSpeed {
			isControl = true
		}
		if !isControl {
			continue
		}

		control := pid.Control(expSpeed, lander.VSpeed)
		if control <= 0 {
			if lander.Power > 0 {
				lander.Power = lander.Power - 1
			}
		} else {
			if lander.Power < MaxPower {
				lander.Power = lander.Power + 1
			}
		}

		exec(0, lander.Power)
	}
}
