package main

import (
	"fmt"
)

func InputField(data []string) Field {
	var err error
	var field Field

	var size int
	size = StrToInt(data[0])
	data = data[1:]
	var p Point
	points := make([]Point, 0, size)
	for i := 0; i < size; i++ {
		_, err = fmt.Sscan(data[i], &p.X, &p.Y)
		if err != nil {
			panic(err)
		}
		points = append(points, p)
	}

	// ground lines
	field.Ground = make(Lines, 0, size-1)
	var ln Line
	for i := 0; i < size-1; i++ {
		ln = Line{points[i], points[i+1]}
		if ln.From.Y == ln.To.Y {
			field.Flat = ln
		}
		field.Ground = append(field.Ground, ln)
	}

	return field
}

func InputLander(data []string) Lander {
	var err error

	var lander Lander
	_, err = fmt.Sscan(
		data[0],
		&lander.Pos.X,
		&lander.Pos.Y,
		&lander.Speed.X,
		&lander.Speed.Y,
		&lander.Fuel,
		&lander.Angle,
		&lander.Power,
	)
	if err != nil {
		panic(err)
	}

	return lander
}
