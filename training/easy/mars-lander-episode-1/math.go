package main

import (
	"math"
	"sort"
)

func Median(values []float64) float64 {
	sort.Float64s(values)

	middle := len(values) / 2
	if len(values)%2 == 0 {
		return (values[middle-1] + values[middle]) / 2
	}

	return values[middle]
}

func ExpectedExpDiff(expected, current, decay float64) float64 {
	diff := math.Abs(expected - current)
	if diff == 0 {
		return 1.0
	}
	if diff > math.Abs(expected) {
		return 0.0
	}

	return math.Exp(-diff / (expected / decay))
}
