package main

import (
	"fmt"
	"math"
)

const (
	up    = 0
	left  = 90
	right = -90
)

type Point struct {
	X, Y float64
}

func (p Point) IsInXRange(from, to float64) bool {
	return p.X >= from && p.X <= to
}

func (p Point) IsInBound(width, height float64) bool {
	return p.X >= 0 && p.X < width && p.Y >= 0 && p.Y < height
}

func (p Point) Distance(t Point) float64 {
	return math.Sqrt(math.Pow(p.X-t.X, 2) + math.Pow(p.Y-t.Y, 2))
}

func (p Point) Sub(t Point) Point {
	return Point{p.X - t.X, p.Y - t.Y}
}

func (p Point) DotProduct(t Point) float64 {
	return p.X*t.X + p.Y*t.Y
}

func (p Point) CrossProduct(t Point) float64 {
	return p.X*t.Y - p.Y*t.X
}

func (p Point) Normalize(length float64) Point {
	return Point{p.X / length, p.Y / length}
}

func (p Point) String() string {
	return fmt.Sprintf("[X:%.f,Y:%.f]", p.X, p.Y)
}

func NewPoint(x, y float64) Point {
	return Point{
		X: math.Round(x),
		Y: math.Round(y),
	}
}

type Line struct {
	From, To Point
}

func (ln Line) Length() float64 {
	return ln.From.Distance(ln.To)
}

func (ln Line) Vector() Point {
	return ln.To.Sub(ln.From)
}

func closestPoint(line Line, point Point, isSegment bool) Point {
	nv := line.Vector().Normalize(line.Length())
	dp := point.Sub(line.From).DotProduct(nv)

	if isSegment {
		if dp <= 0 {
			return line.From
		}
		if dp >= line.Length() {
			return line.To
		}
	}

	return NewPoint(
		line.From.X+nv.X*dp,
		line.From.Y+nv.Y*dp,
	)
}

func (ln Line) ClosestPointOnSegment(t Point) Point {
	return closestPoint(ln, t, true)
}

func linesIntersection(lineA, lineB Line, isSegmentA, isSegmentB bool) (Point, bool) {
	av := lineA.Vector()
	bv := lineB.Vector()

	vcp := av.CrossProduct(bv)
	if vcp == 0 {
		return Point{}, false
	}

	sv := lineB.From.Sub(lineA.From)
	acp := sv.CrossProduct(av)
	bcp := sv.CrossProduct(bv)
	t := bcp / vcp
	u := acp / vcp

	if isSegmentA && (t < 0 || t > 1) {
		return Point{}, false
	}
	if isSegmentB && (u < 0 || u > 1) {
		return Point{}, false
	}

	return Point{
		X: lineA.From.X + t*av.X,
		Y: lineA.From.Y + t*av.Y,
	}, true
}

func (ln Line) SegmentsIntersection(t Line) (Point, bool) {
	return linesIntersection(ln, t, true, true)
}

func (ln Line) Rect() Rect {
	return NewRectangle(ln.From.X, ln.To.X, ln.From.Y, ln.To.Y)
}

func (ln Line) String() string {
	return fmt.Sprintf("(%s->%s)", ln.From, ln.To)
}

type Lines []Line

type Rect struct {
	Xf, Xt, Yf, Yt float64
}

func (r Rect) String() string {
	return fmt.Sprintf("[X:%.f>%.f,Y:%.f>%.f]", r.Xf, r.Xt, r.Yf, r.Yt)
}

func NewRectangle(xf, xt, yf, yt float64) Rect {
	if xf > xt {
		xf, xt = xt, xf
	}
	if yf > yt {
		yf, yt = yt, yf
	}

	return Rect{
		Xf: xf,
		Xt: xt,
		Yf: yf,
		Yt: yt,
	}
}

type Rects []Rect

// MovingDistance calculates the distance traveled by an object.
// formula: s = ut + 1/2at^2
func MovingDistance(speed, acceleration, time float64) float64 {
	return (speed * time) + (0.5 * acceleration * time * time)
}

// MovingVector calculates the vector of a moving object with static angle coordinate system.
func MovingVector(angle, power float64) Point {
	rad := angle * (math.Pi / 180)

	return Point{
		X: -power * math.Sin(rad),
		Y: power * math.Cos(rad),
	}
}
