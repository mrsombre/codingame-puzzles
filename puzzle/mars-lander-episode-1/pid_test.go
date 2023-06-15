package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPID_Control(t *testing.T) {
	lander := debugLander()
	pid := &PID{
		Kp: 1.850,
		Ki: 0.038,
		Kd: 17.718,
	}

	const it = 100
	expSpeed := -MaxVSpeed + (MaxPower-G)*4

	isControl := false
	for i := 0; i < it; i++ {
		lander.VSpeed += -G + float64(lander.Power)

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

		fmt.Printf("vspeed: %f, power: %d\n", lander.VSpeed, lander.Power)
	}

	assert.Greater(t, lander.VSpeed, -float64(MaxVSpeed))
}
