package geo

import (
	"errors"
	. "math"
)

var (
	nNorth = &vector3{x: 0, y: 0, z: 1}
)

type VectorCalculator struct {
}

type vector3 struct {
	x, y, z float64
}

type vector2 struct {
	x, y float64
}

func newNE(lngDeg, latDeg float64) *vector3 {
	lng, lat := ToRadians(lngDeg), ToRadians(latDeg)
	sinLng, cosLng := Sin(lng), Cos(lng)
	sinLat, cosLat := Sin(lat), Cos(lat)
	return &vector3{x: cosLat * cosLng, y: cosLat * sinLng, z: sinLat}
}

func newNEWithPoint(point *Point) *vector3 {
	return newNE(point.X(), point.Y())
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

func (v *vector3) times(t float64) *vector3 {
	return &vector3{
		x: t * v.x,
		y: t * v.y,
		z: t * v.z,
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

func (v *vector3) angleTo(u, n *vector3) float64 {
	sign := sign(v.crossProduct(u).dotProduct(n))
	return Atan2(v.crossProduct(u).norm()*float64(sign), v.dotProduct(u))
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

func (v *vector2) cross(u *vector2) float64 {
	return v.x*u.y - u.x*v.y
}

//======================================================================================================================

func (VectorCalculator) meanPosition(points ...*Point) *Point {
	nM := &vector3{}
	for _, point := range points {
		nE := newNEWithPoint(point)
		nM.x += nE.x
		nM.y += nE.y
		nM.z += nE.z
	}
	return newPoint(nM.unit())
}

func (VectorCalculator) Distance(from, to *Point) float64 {
	nFrom, nTo := newNE(from.X(), from.Y()), newNE(to.X(), to.Y())
	return Atan2(nFrom.crossProduct(nTo).norm(), nFrom.dotProduct(nTo))
}

func (vc *VectorCalculator) DistanceXY(fromX, fromY, toX, toY float64) float64 {
	return vc.Distance(NewPoint(fromX, fromY, nil), NewPoint(toX, toY, nil))
}

func (vc *VectorCalculator) Mid(from, to *Point, ctx GeoContext) *Point {
	if from.Equals(to) {
		return from.clone().(*Point)
	}
	return vc.meanPosition(from, to)
}

func (VectorCalculator) Bearing(from, to *Point) float64 {
	nFrom, nTo := newNE(from.X(), from.Y()), newNE(to.X(), to.Y())
	c1, c2 := nFrom.crossProduct(nTo), nFrom.crossProduct(nNorth)
	bearing := c1.angleTo(c2, nFrom)
	return Mod(ToDegrees(bearing)+360., 360.)
}

func (VectorCalculator) PointOnBearing(from *Point, distRad, bearingDeg float64, ctx GeoContext) *Point {
	nFrom := newNEWithPoint(from)
	bearing := ToRadians(bearingDeg)
	de := nNorth.crossProduct(nFrom).unit()
	dn := nFrom.crossProduct(de)
	deSin := de.times(Sin(bearing))
	dnCos := dn.times(Cos(bearing))
	d := dnCos.plus(deSin)
	x := nFrom.times(Cos(distRad))
	y := d.times(Sin(distRad))
	return newPoint(x.plus(y))
}

func (VectorCalculator) Area(s Shape) float64 {
	sphereCalc := &SphereCalculator{}
	return sphereCalc.Area(s)
}

func (vc *VectorCalculator) MinCoverCircle(points ...*Point) (*Circle, error) {
	n := len(points)
	if n == 0 {
		return nil, errors.New("empty points")
	}

	ps := make([]*Point, n)
	copy(ps, points)
	//rand.Shuffle(n, func(i, j int) {
	//	ps[i], ps[j] = ps[j], ps[i]
	//})
	c, r := ps[0], 0.
	var err error
	for i := 1; i < n; i++ {
		if vc.Distance(ps[i], c) > r {
			c, r = ps[i], 0.
			for j := 0; j < i; j++ {
				if vc.Distance(ps[j], c) > r {
					c = vc.Mid(ps[i], ps[j], nil)
					r = vc.Distance(c, ps[j])
					for k := 0; k < j; k++ {
						if vc.Distance(ps[k], c) > r {
							c, err = vc.Circumcenter(ps[i], ps[j], ps[k])
							if err != nil {
								return nil, err
							}
							r = Max(vc.Distance(c, ps[k]), Max(vc.Distance(c, ps[i]), vc.Distance(c, ps[j])))
						}
					}
				}
			}
		}
	}
	return NewCircle(c.X(), c.Y(), r, nil), nil
}

func (vc *VectorCalculator) Circumcenter(p1, p2, p3 *Point) (*Point, error) {
	//v := &vector2{x: p2.X() - p1.X(), y: p2.Y() - p1.Y()}
	//u := &vector2{x: p3.X() - p1.X(), y: p3.Y() - p1.Y()}
	//cross := v.cross(u)
	//if Abs(cross) < E12 {
	//	return nil, fmt.Errorf("3点共线: p1=%v, p2=%v, p3=%v", p1, p2, p3)
	//} else if cross < 0 { //CW
	//} else { //CCW
	//}

	p12Mid, p23Mid := vc.Mid(p1, p2, nil), vc.Mid(p2, p3, nil)
	crs12Mid2 := vc.Bearing(p12Mid, p2)
	crs23Mid3 := vc.Bearing(p23Mid, p3)
	p4, err := vc.IntersectionOfTwoPath(
		vc.PointOnBearing(p12Mid, ToRadians(90-E6), crs12Mid2-90, nil),
		vc.PointOnBearing(p12Mid, ToRadians(90-E6), crs12Mid2+90, nil),
		vc.PointOnBearing(p23Mid, ToRadians(90-E6), crs23Mid3-90, nil),
		vc.PointOnBearing(p23Mid, ToRadians(90-E6), crs23Mid3+90, nil),
	)
	if err != nil {
		return nil, err
	}
	return p4, nil
}

//func (VectorCalculator) Intersection(pa *Point, brng1 float64, pb *Point, brng2 float64) (*Point, error) {
//	p1, p2 := newNE(pa.X(), pa.Y()), newNE(pb.X(), pb.Y())
//	c1, c2 := pa.greatCircle(brng1), pb.greatCircle(brng2)
//	i1, i2 := c1.crossProduct(c2), c2.crossProduct(c1)
//	dir1 := sign(c1.crossProduct(p1).dotProduct(i1))
//	dir2 := sign(c2.crossProduct(p2).dotProduct(i1))
//	switch dir1 + dir2 {
//	case 2:
//		return newPoint(i1), nil
//	case -2:
//		return newPoint(i2), nil
//	case 0:
//		if p1.plus(p2).dotProduct(i1) > 0 {
//			return newPoint(i2), nil
//		} else {
//			return newPoint(i1), nil
//		}
//	default:
//		return nil, fmt.Errorf("program should not run here: pa=%v, brng1=%f, pb=%v, brng2=%f", pa, brng1, pb, brng2)
//	}
//}

func (VectorCalculator) IntersectionOfTwoPath(pa1, pa2, pb1, pb2 *Point) (*Point, error) {
	p1, p2 := newNE(pa1.X(), pa1.Y()), newNE(pb1.X(), pb1.Y())
	p11, p22 := newNE(pa2.X(), pa2.Y()), newNE(pb2.X(), pb2.Y())
	c1, c2 := p1.crossProduct(p11), p2.crossProduct(p22)
	i1, i2 := c1.crossProduct(c2), c2.crossProduct(c1)
	mid := p1.plus(p2).plus(p11).plus(p22)
	if mid.dotProduct(i1) > 0 {
		return newPoint(i1), nil
	}
	return newPoint(i2), nil
}
