package main

import (
	"fmt"
)

type Command interface {
	String() string
}

type Commands []Command

type MoveCommand struct {
	Angle float64
	Power float64
}

func (c MoveCommand) String() string {
	return fmt.Sprintf("%.f %.f", c.Angle, c.Power)
}

func ExecuteCommand(commands Commands) {
	for _, command := range commands {
		fmt.Println(command)
	}
}
