package main

import (
	"fmt"
	"math"
)

const (
	// World
	width     = 7000
	height    = 3000
	g         = 3.711
	flatWidth = 1000

	// Rules
	angleMax  = 90
	angleStep = 15
	powerMax  = 4
	powerStep = 1

	// Successful landing
	vSpeedMax = 40
	hSpeedMax = 20
)

type Field struct {
	Ground Lines
	Flat   Line
}

type Lander struct {
	Pos   Point
	Speed Point

	Angle float64
	Power float64
	Fuel  float64
}

func (l Lander) Move(angle, power float64) Lander {
	thrust := MovingVector(angle, power)
	thrust.Y -= g

	l.Pos.X += MovingDistance(l.Speed.X, thrust.X, 1)
	l.Pos.Y += MovingDistance(l.Speed.Y, thrust.Y, 1)
	l.Speed.X += thrust.X
	l.Speed.Y += thrust.Y

	l.Angle = angle
	l.Power = power
	l.Fuel -= l.Power

	return l
}

func (l Lander) String() string {
	return fmt.Sprintf("Pos: %s, Speed: %s, Angle: %.f, Power: %.f, Fuel: %.f", l.Pos, l.Speed, l.Angle, l.Power, l.Fuel)
}

type Game struct {
	Field  Field
	Lander Lander
}

func (g Game) IsLanderInBound() bool {
	return g.Lander.Pos.IsInBound(width, height)
}

func (g Game) IsLanderBetweenFlat() bool {
	return g.Lander.Pos.IsInXRange(g.Field.Flat.From.X, g.Field.Flat.To.X)
}

func (g Game) IsLanderOnFlat() bool {
	return g.Lander.Pos.Y == g.Field.Flat.To.Y
}

func (g Game) IsLanderLanded() bool {
	return g.IsLanderBetweenFlat() && g.IsLanderOnFlat()
}

func (g Game) IsLanderAngleOk() bool {
	return g.Lander.Angle == 0
}

func (g Game) IsLanderSpeedOk() bool {
	return math.Abs(g.Lander.Speed.Y) <= vSpeedMax && math.Abs(g.Lander.Speed.X) <= hSpeedMax
}

func (g Game) IsLanderLandedOk() bool {
	return g.IsLanderLanded() && g.IsLanderAngleOk() && g.IsLanderSpeedOk()
}
