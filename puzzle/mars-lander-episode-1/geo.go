package main

import (
	"math"
)

func calculateDistance(speed, acceleration, time float64) float64 {
	return (speed * time) + (0.5 * acceleration * time * time)
}

// formula: t = (-u ± √(u^2 + 2as)) / a
func calculateTime(distance, speed, acceleration float64) float64 {
	uSquared := speed * speed
	term := 2 * acceleration * distance
	squareRoot := math.Sqrt(uSquared + term)
	time := (speed + squareRoot) / acceleration
	return time
}
