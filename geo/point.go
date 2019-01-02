package geo

import (
	"fmt"
	"github.com/paulsmith/gogeos/geos"
)

type Point struct {
	AbstractShape

	coord geos.Coord
}

func NewPoint(x, y float64, ctx GeoContext) *Point {
	p := &Point{coord: geos.NewCoord(x, y)}
	p.ctx = ctx
	return p
}

func (p *Point) X() float64 {
	return p.coord.X
}

func (p *Point) Y() float64 {
	return p.coord.Y
}

func (p *Point) Reset(x, y float64) {
	p.coord.X, p.coord.Y = x, y
}

func (p *Point) ToGeos() (*geos.Geometry, error) {
	return geos.NewPoint(geos.NewCoord(p.coord.X, p.coord.Y))
}

func (p *Point) GetCenter() *Point {
	return p.clone().(*Point)
}

func (*Point) GetArea() float64 {
	return 0
}

func (p *Point) clone() Shape {
	return NewPoint(p.coord.X, p.coord.Y, p.ctx)
}

func (p *Point) String() string {
	return fmt.Sprintf("POINT (%v %v)", p.coord.X, p.coord.Y)
}

func (p *Point) Equals(other *Point) bool {
	if p == other {
		return true
	}
	return false //TODO: impl this func
}
