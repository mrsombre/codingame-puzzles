package pid

import (
	"math"
)

type PID struct {
	Kp float64
	Ki float64
	Kd float64

	LastError   float64
	IntegralSum float64

	IntegralMax float64 // max value for integral sum
	IntegralMin float64 // min value for integral sum
}

func (pid *PID) Control(referenceSignal, actualSignal float64) float64 {
	currentError := referenceSignal - actualSignal

	// Proportional
	proportional := pid.Kp * currentError

	// Integral
	pid.IntegralSum += currentError
	// Limit integral sum to avoid windup
	if pid.IntegralMax != 0 {
		pid.IntegralSum = math.Min(pid.IntegralSum, pid.IntegralMax)
	}
	if pid.IntegralMin != 0 {
		pid.IntegralSum = math.Max(pid.IntegralSum, pid.IntegralMin)
	}
	integral := pid.Ki * pid.IntegralSum

	// Derivative
	derivative := pid.Kd * (currentError - pid.LastError)

	pid.LastError = currentError

	return proportional + integral + derivative
}
