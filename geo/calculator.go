package geo

type Calculator interface {
	Distance(from, to *Point) float64

	DistanceXY(fromX, fromY, toX, toY float64) float64

	Bearing(from, to *Point) float64

	PointOnBearing(from *Point, distRad, bearingDeg float64, ctx GeoContext) *Point

	Area(s Shape) float64
}
