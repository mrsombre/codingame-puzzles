package main

// A set of helper methods for type conversions.

import (
	"math"
	"strconv"
)

func StrToInt(s string) int {
	result, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return result
}

func IsFloatsEqual(a, b, delta float64) bool {
	return math.Abs(a-b) <= delta
}
