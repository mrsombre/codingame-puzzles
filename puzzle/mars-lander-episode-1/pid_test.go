package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPID_Control(t *testing.T) {
	lander := &Lander{}
	pid := &PID{
		Kp:          3.0,
		Ki:          2.0,
		Kd:          10.0,
		IntegralMax: 1.0,
		IntegralMin: -1.0,
	}

	const it = 100
	expSpeed := -MaxVSpeed + 1

	for i := 0; i < it; i++ {
		lander.VSpeed += -G + float64(lander.Power)
		if lander.VSpeed > expSpeed+10 {
			continue
		}

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

		fmt.Printf("control: %.2f | vspeed: %f | power: %d\n", control, lander.VSpeed, lander.Power)
	}

	assert.Greater(t, lander.VSpeed, -MaxVSpeed)
}
