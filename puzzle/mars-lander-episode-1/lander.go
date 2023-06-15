package main

import (
	"fmt"
)

const (
	G = 3.711

	MaxHSpeed = 20
	MaxVSpeed = 40

	TurnRotate = 15
	TurnPower  = 1

	MaxPower = 4
)

type Lander struct {
	X, Y                int
	HSpeed, VSpeed      float64
	Fuel, Rotate, Power int
}

func UpdateLander(data string, lander *Lander) {
	fmt.Sscan(data, &lander.X, &lander.Y, &lander.HSpeed, &lander.VSpeed, &lander.Fuel, &lander.Rotate, &lander.Power)
}
