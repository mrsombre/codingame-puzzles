package geometry

type Coord struct {
	X, Y int
}

func (cr Coord) Equal(cmp Coord) bool {
	return cr.X == cmp.X && cr.Y == cmp.Y
}
