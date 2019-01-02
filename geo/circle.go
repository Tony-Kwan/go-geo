package geo

import (
	"fmt"
	"github.com/paulsmith/gogeos/geos"
	"strconv"
)

type Circle struct {
	AbstractShape

	center *Point
	radius float64 //if circle is a spherical cap, unit of radius is radian
}

func NewCircle(x, y, radiusDeg float64, ctx GeoContext) *Circle {
	c := &Circle{center: NewPoint(x, y, ctx), radius: radiusDeg}
	c.ctx = ctx
	return c
}

func (c *Circle) ToGeos() (*geos.Geometry, error) {
	return nil, ErrUnsupportedOperation
}

func (c *Circle) GetArea() float64 {
	return c.ctx.GetCalculator().Area(c)
}

func (c *Circle) GetCenter() *Point {
	return c.center.clone().(*Point)
}

func (c *Circle) clone() Shape {
	return NewCircle(c.center.X(), c.center.Y(), c.radius, c.ctx)
}

func (c *Circle) String() string {
	return fmt.Sprintf(
		"CIRCLE((%s, %s), %s)",
		strconv.FormatFloat(c.center.X(), 'f', -1, 64),
		strconv.FormatFloat(c.center.Y(), 'f', -1, 64),
		strconv.FormatFloat(c.radius, 'f', -1, 64),
	)
}
