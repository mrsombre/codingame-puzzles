package pid

const (
	G = 3.711

	MaxSpeed = 40
	MaxPower = 4
)

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

type Lander struct {
	VSpeed float64
	Power  int
}
