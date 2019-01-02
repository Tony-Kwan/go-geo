package geo

import (
	"github.com/paulsmith/gogeos/geos"
)

var (
	unitRadius = 1.0
)

type Shape interface {
	ToGeos() (*geos.Geometry, error)

	GetArea() float64

	GetCenter() *Point

	GetContext() GeoContext

	clone() Shape
}

type AbstractShape struct {
	ctx GeoContext
}

func (*AbstractShape) ToGeos() (*geos.Geometry, error) {
	return nil, ErrUnsupportedOperation
}

func (*AbstractShape) GetArea() float64 {
	panic(ErrUnsupportedOperation)
	return 0
}

func (*AbstractShape) GetCenter() *Point {
	panic(ErrUnsupportedOperation)
	return nil
}

func (s *AbstractShape) GetContext() GeoContext {
	return s.ctx
}

func (AbstractShape) clone() Shape {
	panic(ErrUnsupportedOperation)
	return nil
}
