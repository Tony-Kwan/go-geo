package geo

import (
	"fmt"
	. "math"
	"strconv"
)

var (
	zero   = vector3{}
	nNorth = vector3{x: 0, y: 0, z: 1}
	nSouth = vector3{x: 0, y: 0, z: -1}
)

//vector represents a point in R3
type vector3 struct {
	x, y, z float64
}

func newNE(lngDeg, latDeg float64) vector3 {
	lng, lat := ToRadians(lngDeg), ToRadians(latDeg)
	sinLng, cosLng := Sin(lng), Cos(lng)
	sinLat, cosLat := Sin(lat), Cos(lat)
	return vector3{x: cosLat * cosLng, y: cosLat * sinLng, z: sinLat}
}

func newNEWithPoint(point Point) vector3 {
	return newNE(point.X(), point.Y())
}

func (v vector3) toPoint() Point {
	return NewPoint(
		ToDegrees(Atan2(v.y, v.x)),
		ToDegrees(Atan2(v.z, Sqrt(v.x*v.x+v.y*v.y))),
		GeoCtx,
	)
}

func (v vector3) unit() vector3 {
	n := v.norm()
	return vector3{v.x / n, v.y / n, v.z / n}
}

func (v vector3) norm() float64 {
	return Sqrt(v.dot(v))
}

func (v vector3) add(u vector3) vector3 {
	return vector3{
		x: v.x + u.x,
		y: v.y + u.y,
		z: v.z + u.z,
	}
}

func (v vector3) mul(t float64) vector3 {
	return vector3{
		x: t * v.x,
		y: t * v.y,
		z: t * v.z,
	}
}

func (v vector3) cross(u vector3) vector3 {
	return vector3{
		x: v.y*u.z - v.z*u.y,
		y: -(v.x*u.z - v.z*u.x),
		z: v.x*u.y - v.y*u.x,
	}
}

func (v vector3) dot(u vector3) float64 {
	return v.x*u.x + v.y*u.y + v.z*u.z
}

func (v vector3) angleTo(u, n vector3) float64 {
	sin := v.cross(u).norm() * float64(sign(v.cross(u).dot(n)))
	cos := v.dot(u)
	return Atan2(sin, cos)
}

func (v vector3) ApproxEqual(u vector3) bool {
	const eps = E12
	return Abs(v.x-u.x) < eps && Abs(v.y-u.y) < eps && Abs(v.z-u.z) < eps
}

func (v vector3) String() string {
	return fmt.Sprintf(
		"[%s %s %s]",
		strconv.FormatFloat(v.x, 'f', -1, 64),
		strconv.FormatFloat(v.y, 'f', -1, 64),
		strconv.FormatFloat(v.z, 'f', -1, 64),
	)
}

func (p Point) greatCircle(bearingDeg float64) vector3 {
	lng, lat := ToRadians(p.X()), ToRadians(p.Y())
	bearing := ToRadians(bearingDeg)
	return vector3{
		x: Sin(lng)*Cos(bearing) - Sin(lat)*Cos(lng)*Sin(bearing),
		y: -Cos(lng)*Cos(bearing) - Sin(lat)*Sin(lng)*Sin(bearing),
		z: Cos(lat) * Sin(bearing),
	}
}

//======================================================================================================================

//vector represents a point in R2
type vector2 struct {
	x, y float64
}

func newVector2(p0, p1 Point) vector2 {
	return vector2{x: p1.x - p0.x, y: p1.y - p0.y}
}

func (v vector2) cross(u vector2) float64 {
	return v.x*u.y - u.x*v.y
}

func (v vector2) dot(u vector2) float64 {
	return v.x*u.x + v.y*u.y
}
