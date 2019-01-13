package geo

import (
	"fmt"
	"math"
	"strconv"
)

type Point struct {
	AbstractShape

	x float64
	y float64
}

func NewPoint(x, y float64, ctx GeoContext) *Point {
	p := &Point{x: x, y: y}
	p.ctx = ctx
	return p
}

func (p *Point) X() float64 {
	return p.x
}

func (p *Point) Y() float64 {
	return p.y
}

func (p *Point) Reset(x, y float64) {
	p.x, p.y = x, y
}

func (p *Point) GetCenter() *Point {
	return p.clone().(*Point)
}

func (*Point) GetArea() float64 {
	return 0
}

func (p *Point) clone() Shape {
	return NewPoint(p.x, p.y, p.ctx)
}

func (p *Point) String() string {
	return fmt.Sprintf(
		"POINT (%s %s)",
		strconv.FormatFloat(p.x, 'f', -1, 64),
		strconv.FormatFloat(p.y, 'f', -1, 64),
	)
}

func (p *Point) Equals(other *Point) bool {
	if p == other {
		return true
	}
	return p.x == other.x && p.y == other.y && p.ctx == other.ctx
}

func (p *Point) ApproxEqual(other *Point) bool {
	if p == nil || other == nil {
		return false
	}
	const eps = E10
	return math.Abs(p.x-other.x) < eps && math.Abs(p.y-other.y) < eps
}
