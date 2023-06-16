package geometry

import (
	"math"
)

// MovingDistance formula: D = v*t + 1/2*a*t^2
func MovingDistance(speed, acceleration, time float64) float64 {
	return (speed * time) + (0.5 * acceleration * time * time)
}

// MovingTime formula: t = (-u ± √(u^2 + 2as)) / a
func MovingTime(distance, speed, acceleration float64) float64 {
	if acceleration == 0 {
		return distance / speed
	}
	uSquared := speed * speed
	term := 2 * acceleration * distance
	squareRoot := math.Sqrt(uSquared + term)
	time := (speed + squareRoot) / acceleration
	return time
}
