package main

type Command struct {
	rotate int
	power  int
}

type Mla struct {
	grid   *Grid
	lander *Lander
}

func NewMla(grid *Grid, lander *Lander) *Mla {
	return &Mla{
		grid:   grid,
		lander: lander,
	}
}

func (m *Mla) Control() *Command {
	var rotate, power int

	dv := MaxVSpeed - 1 - -m.lander.VSpeed
	if dv < 0 {
		power = min(MaxPower, m.lander.Power+1)
	} else if m.lander.Power > 0 {
		power = max(0, m.lander.Power-1)
	}

	return &Command{
		rotate: rotate,
		power:  power,
	}
}

func (m *Mla) GetYFromX(x int) int {
	var left, right *Cell

	for _, cell := range m.grid.Cells {
		if cell.X < x {
			if left == nil || cell.X > left.X {
				left = cell
			}
		}
		if cell.X > x {
			if right == nil || cell.X < right.X {
				right = cell
			}
		}
	}

	return left.Y + (x-left.X)*(right.Y-left.Y)/(right.X-left.X)
}

func (m *Mla) GetGroundHeight() int {
	return m.lander.Y - m.GetYFromX(m.lander.X)
}
