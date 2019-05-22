package geo

import (
	"fmt"
	"math"
	"strconv"
)

type Triangle struct {
	AbstractShape

	a *Point
	b *Point
	c *Point
}

func NewTriangle(a, b, c *Point, ctx GeoContext) *Triangle {
	tri := &Triangle{a: a, b: b, c: c}
	tri.ctx = ctx
	return tri
}

func (tri *Triangle) GetArea() float64 {
	return tri.GetContext().GetCalculator().Area(tri)
}

func (tri *Triangle) IsDisjoint(other *Triangle) bool {
	a1, b1, c1 := pointHash(tri.a), pointHash(tri.b), pointHash(tri.c)
	a2, b2, c2 := pointHash(other.a), pointHash(other.b), pointHash(other.c)
	return len(map[uint64]bool{a1: true, b1: true, c1: true, a2: true, b2: true, c2: true}) != 4
}

func (tri *Triangle) String() string {
	return fmt.Sprintf(
		"POLYGON((%s %s, %s %s, %s %s, %s %s))",
		strconv.FormatFloat(tri.a.X(), 'f', -1, 64), strconv.FormatFloat(tri.a.Y(), 'f', -1, 64),
		strconv.FormatFloat(tri.b.X(), 'f', -1, 64), strconv.FormatFloat(tri.b.Y(), 'f', -1, 64),
		strconv.FormatFloat(tri.c.X(), 'f', -1, 64), strconv.FormatFloat(tri.c.Y(), 'f', -1, 64),
		strconv.FormatFloat(tri.a.X(), 'f', -1, 64), strconv.FormatFloat(tri.a.Y(), 'f', -1, 64),
	)
}

func (tri *Triangle) contains(point *Point) bool {
	v0 := NewVector2(tri.a, tri.c)
	v1 := NewVector2(tri.a, tri.b)
	v2 := NewVector2(tri.a, point)
	dot00 := v0.dot(v0)
	dot01 := v0.dot(v1)
	dot02 := v0.dot(v2)
	dot11 := v1.dot(v1)
	dot12 := v1.dot(v2)
	invDenom := 1.0 / (dot00*dot11 - dot01*dot01)
	u := (dot11*dot02 - dot01*dot12) * invDenom
	v := (dot00*dot12 - dot01*dot02) * invDenom
	return (u >= 0) && (v >= 0) && (u+v < 1)
}

func pointHash(p *Point) uint64 {
	return math.Float64bits(p.x) ^ math.Float64bits(p.y)
}
