package geo

type CartesianCalculator struct {
}

func (cc *CartesianCalculator) Distance(from, to *Point) float64 {
	return 0
}

func (cc *CartesianCalculator) DistanceXY(fromX, fromY, toX, toY float64) float64 {
	return 0
}

func (cc *CartesianCalculator) Bearing(from, to *Point) float64 {
	return 0
}

func (cc *CartesianCalculator) PointOnBearing(from *Point, distRad, bearingDeg float64, ctx GeoContext) *Point {
	return nil
}

func (cc *CartesianCalculator) Area(s Shape) float64 {
	switch s.(type) {
	case *Rectangle:
		return cc.areaOfRectangle(s.(*Rectangle))
	default:
		erro.Printf("unsupported shape type: %+v\n", s)
		return -1
	}
}

func (cc *CartesianCalculator) areaOfRectangle(rect *Rectangle) float64 {
	return rect.GetWidth() * rect.GetHeight()
}
