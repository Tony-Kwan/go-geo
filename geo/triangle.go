package geo

import (
	"fmt"
	"math"
	"strconv"
)

type Triangle struct {
	AbstractShape

	A *Point
	B *Point
	C *Point
}

func NewTriangle(a, b, c *Point, ctx GeoContext) *Triangle {
	tri := &Triangle{A: a, B: b, C: c}
	tri.ctx = ctx
	return tri
}

func (tri *Triangle) GetArea() float64 {
	return tri.GetContext().GetCalculator().Area(tri)
}

func (tri *Triangle) IsConnected(other *Triangle) bool {
	a1, b1, c1 := pointHash(tri.A), pointHash(tri.B), pointHash(tri.C)
	a2, b2, c2 := pointHash(other.A), pointHash(other.B), pointHash(other.C)
	return len(map[uint64]bool{a1: true, b1: true, c1: true, a2: true, b2: true, c2: true}) == 4 //TODO: check if triangle contain another point of triangle
}

func (tri Triangle) ToPolygon() Polygon {
	shell := LinearRing{
		*tri.A,
		*tri.B,
		*tri.C,
		*tri.A,
	}
	return *NewPolygon(shell)
}

func (tri *Triangle) String() string {
	return fmt.Sprintf(
		"POLYGON((%s %s, %s %s, %s %s, %s %s))",
		strconv.FormatFloat(tri.A.X(), 'f', -1, 64), strconv.FormatFloat(tri.A.Y(), 'f', -1, 64),
		strconv.FormatFloat(tri.B.X(), 'f', -1, 64), strconv.FormatFloat(tri.B.Y(), 'f', -1, 64),
		strconv.FormatFloat(tri.C.X(), 'f', -1, 64), strconv.FormatFloat(tri.C.Y(), 'f', -1, 64),
		strconv.FormatFloat(tri.A.X(), 'f', -1, 64), strconv.FormatFloat(tri.A.Y(), 'f', -1, 64),
	)
}

func (tri *Triangle) contains(point *Point) bool {
	switch tri.GetContext().(type) {
	case *SpatialContext:
		v1, v2, v3 := newNEWithPoint(tri.A), newNEWithPoint(tri.B), newNEWithPoint(tri.C)
		v0 := newNEWithPoint(point)
		return v1.cross(v2).dot(v0) >= 0 && v2.cross(v3).dot(v0) >= 0 && v3.cross(v1).dot(v0) >= 0
	default:
		v0 := NewVector2(tri.A, tri.C)
		v1 := NewVector2(tri.A, tri.B)
		v2 := NewVector2(tri.A, point)
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
}

func pointHash(p *Point) uint64 {
	return math.Float64bits(p.x) ^ math.Float64bits(p.y)
}
