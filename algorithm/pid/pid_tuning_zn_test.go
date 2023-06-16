package pid

import (
	"fmt"
	"testing"
)

// This test is used to tune the PID controller
// https://en.wikipedia.org/wiki/Ziegler%E2%80%93Nichols_method
// https://github.com/jackw01/arduino-pid-autotuner/blob/master/pidautotuner.cpp
func TestPID_ControlTuning(t *testing.T) {
	lander := &Lander{VSpeed: 0, Power: 0}
	pid := &PID{
		Kp: 0.0,
		Ki: 0.0,
		Kd: 0.0,
	}

	expSpeed := -float64(MaxSpeed)

	lastVSpeed := 0.0
	isControl := false
	cnt := 0
	ci := 0

	for i := 0; i < 1000; i++ {
		lastVSpeed = lander.VSpeed
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

		if lastVSpeed > expSpeed && lander.VSpeed < expSpeed || lastVSpeed < expSpeed && lander.VSpeed > expSpeed {
			if cnt > 10 { // We have a stable oscillation
				Ku := pid.Kp
				Pu := float64(ci) / float64(cnt)
				pid.Kp = 0.6 * Ku
				pid.Ki = 2 * pid.Kp / Pu
				pid.Kd = pid.Kp * Pu / 8
				cnt = 0
				ci = 0
				continue
			}
			pid.Kp += 0.1
			cnt++
		}

		ci++
	}

	fmt.Printf("Tuning Done: Kp = %f, Ki = %f, Kd = %f\n", pid.Kp, pid.Ki, pid.Kd)
}
