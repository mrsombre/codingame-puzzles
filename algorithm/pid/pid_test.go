package pid

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

const (
	G = 3.711

	MaxSpeed = 40.0
	MaxPower = 4
)

type Lander struct {
	VSpeed float64
	Power  int
}

func TestPID_Control(t *testing.T) {
	lander := &Lander{
		VSpeed: 0,
		Power:  0,
	}
	pid := &PID{Kp: 4.0, Ki: 3.0, Kd: 3.0, IntegralMax: 10, IntegralMin: -10}

	const it = 100
	expSpeed := -MaxSpeed + 1

	jsonData := struct {
		Speed []float64 `json:"speed"`
		Power []int     `json:"power"`
	}{
		Speed: make([]float64, 0, it),
		Power: make([]int, 0, it),
	}

	for i := 0; i < it; i++ {
		lander.VSpeed += -G + float64(lander.Power)

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

		jsonData.Speed = append(jsonData.Speed, lander.VSpeed)
		jsonData.Power = append(jsonData.Power, lander.Power)

		fmt.Printf("control: %.2f | vspeed: %f | power: %d\n", control, lander.VSpeed, lander.Power)
	}

	// write points to file json
	// Convert the points slice to JSON
	json, err := json.Marshal(jsonData)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// Write the JSON data to file
	err = ioutil.WriteFile("points.json", json, 0644)
	if err != nil {
		panic(err)
	}
}
