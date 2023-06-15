package main

type PID struct {
	Kp, Ki, Kd             float64
	LastError, IntegralSum float64
}

func (pid *PID) Control(desiredSpeed, actualSpeed float64) float64 {
	pe := desiredSpeed - actualSpeed

	// Proportional
	proportional := pid.Kp * pe

	// Integral
	pid.IntegralSum += pe
	integral := pid.Ki * pid.IntegralSum

	// Derivative
	derivative := pid.Kd * (pe - pid.LastError)
	pid.LastError = pe

	return proportional + integral + derivative
}
