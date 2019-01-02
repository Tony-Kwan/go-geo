package geo

import (
	"errors"
	. "math"
)

type VectorCalculator struct {
}

type vector3 struct {
	x, y, z float64
}

func newNE(lngDeg, latDeg float64) *vector3 {
	lng, lat := ToRadians(lngDeg), ToRadians(latDeg)
	sinLng, cosLng := Sin(lng), Cos(lng)
	sinLat, cosLat := Sin(lat), Cos(lat)
	return &vector3{x: cosLat * cosLng, y: cosLat * sinLng, z: sinLat}
}

func newPoint(nE *vector3) *Point {
	return NewPoint(
		ToDegrees(Atan2(nE.y, nE.x)),
		ToDegrees(Atan2(nE.z, Sqrt(nE.x*nE.x+nE.y*nE.y))),
		GeoCtx,
	)
}

func sign(a float64) int {
	switch {
	case a < 0:
		return -1
	case a > 0:
		return +1
	}
	return 0
}

func (v *vector3) unit() *vector3 {
	n := v.norm()
	return &vector3{v.x / n, v.y / n, v.z / n}
}

func (v *vector3) norm() float64 {
	return Sqrt(v.x*v.x + v.y*v.y + v.z*v.z)
}

func (v *vector3) plus(u *vector3) *vector3 {
	return &vector3{
		x: v.x + u.x,
		y: v.y + u.y,
		z: v.z + u.z,
	}
}

func (v *vector3) crossProduct(u *vector3) (p *vector3) {
	p = &vector3{
		x: v.y*u.z - v.z*u.y,
		y: -(v.x*u.z - v.z*u.x),
		z: v.x*u.y - v.y*u.x,
	}
	return
}

func (v *vector3) dotProduct(u *vector3) float64 {
	return v.x*u.x + v.y*u.y + v.z*u.z
}

func (p *Point) greatCircle(bearingDeg float64) *vector3 {
	lng, lat := ToRadians(p.X()), ToRadians(p.Y())
	bearing := ToRadians(bearingDeg)
	return &vector3{
		x: Sin(lng)*Cos(bearing) - Sin(lat)*Cos(lng)*Sin(bearing),
		y: -Cos(lng)*Cos(bearing) - Sin(lat)*Sin(lng)*Sin(bearing),
		z: Cos(lat) * Sin(bearing),
	}
}

//======================================================================================================================

func (VectorCalculator) meanPosition(nEs ...*vector3) *vector3 {
	nM := &vector3{}
	for _, nE := range nEs {
		nM.x += nE.x
		nM.y += nE.y
		nM.z += nE.z
	}
	return nM.unit()
}

func (vc *VectorCalculator) Distance(from, to *Point) float64 {
	nFrom, nTo := newNE(from.X(), from.Y()), newNE(to.X(), to.Y())
	return Atan2(nFrom.crossProduct(nTo).norm(), nFrom.dotProduct(nTo))
}

func (vc *VectorCalculator) DistanceXY(fromX, fromY, toX, toY float64) float64 {
	return vc.Distance(NewPoint(fromX, fromY, nil), NewPoint(toX, toY, nil))
}

func (VectorCalculator) Bearing(from, to *Point) float64 {
	return 0 //TODO:impl
}

func (VectorCalculator) PointOnBearing(from *Point, distDeg, bearingDeg float64, ctx GeoContext) *Point {
	return nil //TODO:impl
}

func (VectorCalculator) Area(s Shape) float64 {
	sphereCalc := &SphereCalculator{}
	return sphereCalc.Area(s)
}

func (VectorCalculator) Intersection(pa *Point, brng1 float64, pb *Point, brng2 float64) (*Point, error) {
	p1, p2 := newNE(pa.X(), pa.Y()), newNE(pb.X(), pb.Y())
	c1, c2 := pa.greatCircle(brng1), pb.greatCircle(brng2)
	i1, i2 := c1.crossProduct(c2), c2.crossProduct(c1)
	dir1 := sign(c1.crossProduct(p1).dotProduct(i1))
	dir2 := sign(c2.crossProduct(p2).dotProduct(i1))
	switch dir1 + dir2 {
	case 2:
		return newPoint(i1), nil
	case -2:
		return newPoint(i2), nil
	case 0:
		if p1.plus(p2).dotProduct(i1) > 0 {
			return newPoint(i2), nil
		} else {
			return newPoint(i1), nil
		}
	default:
		return nil, errors.New("program should not run here")
	}
}
