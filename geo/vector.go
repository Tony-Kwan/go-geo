package geo

import (
	. "math"
)

var (
	nNorth = &nVector{x: 0, y: 0, z: 1}
)

type nVector struct {
	x, y, z float64
}

func newNE(lngDeg, latDeg float64) *nVector {
	lng, lat := ToRadians(lngDeg), ToRadians(latDeg)
	sinLng, cosLng := Sin(lng), Cos(lng)
	sinLat, cosLat := Sin(lat), Cos(lat)
	return &nVector{x: cosLat * cosLng, y: cosLat * sinLng, z: sinLat}
}

func newNEWithPoint(point *Point) *nVector {
	return newNE(point.X(), point.Y())
}

func (v *nVector) toPoint() *Point {
	return NewPoint(
		ToDegrees(Atan2(v.y, v.x)),
		ToDegrees(Atan2(v.z, Sqrt(v.x*v.x+v.y*v.y))),
		GeoCtx,
	)
}

func (v *nVector) unit() *nVector {
	n := v.norm()
	return &nVector{v.x / n, v.y / n, v.z / n}
}

func (v *nVector) norm() float64 {
	return Sqrt(v.x*v.x + v.y*v.y + v.z*v.z)
}

func (v *nVector) plus(u *nVector) *nVector {
	return &nVector{
		x: v.x + u.x,
		y: v.y + u.y,
		z: v.z + u.z,
	}
}

func (v *nVector) times(t float64) *nVector {
	return &nVector{
		x: t * v.x,
		y: t * v.y,
		z: t * v.z,
	}
}

func (v *nVector) crossProduct(u *nVector) *nVector {
	return &nVector{
		x: v.y*u.z - v.z*u.y,
		y: -(v.x*u.z - v.z*u.x),
		z: v.x*u.y - v.y*u.x,
	}
}

func (v *nVector) dotProduct(u *nVector) float64 {
	return v.x*u.x + v.y*u.y + v.z*u.z
}

func (v *nVector) angleTo(u, n *nVector) float64 {
	sign := sign(v.crossProduct(u).dotProduct(n))
	return Atan2(v.crossProduct(u).norm()*float64(sign), v.dotProduct(u))
}

func (p *Point) greatCircle(bearingDeg float64) *nVector {
	lng, lat := ToRadians(p.X()), ToRadians(p.Y())
	bearing := ToRadians(bearingDeg)
	return &nVector{
		x: Sin(lng)*Cos(bearing) - Sin(lat)*Cos(lng)*Sin(bearing),
		y: -Cos(lng)*Cos(bearing) - Sin(lat)*Sin(lng)*Sin(bearing),
		z: Cos(lat) * Sin(bearing),
	}
}

//======================================================================================================================

type vector2 struct {
	x, y float64
}

func (v *vector2) cross(u *vector2) float64 {
	return v.x*u.y - u.x*v.y
}
