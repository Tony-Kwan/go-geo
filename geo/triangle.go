package geo

import (
	"fmt"
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

func (tri *Triangle) String() string {
	return fmt.Sprintf(
		"POLYGON ((%s %s, %s %s, %s %s, %s %s))",
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
