package geo

import (
	"errors"
	. "math"
	"math/rand"
)

type VectorCalculator struct {
}

func (VectorCalculator) meanPosition(points ...*Point) *Point {
	nM := &vector3{}
	for _, point := range points {
		nE := newNEWithPoint(point)
		nM.x += nE.x
		nM.y += nE.y
		nM.z += nE.z
	}
	return nM.unit().toPoint()
}

func (VectorCalculator) Distance(from, to *Point) float64 {
	nFrom, nTo := newNE(from.X(), from.Y()), newNE(to.X(), to.Y())
	return Atan2(nFrom.cross(nTo).norm(), nFrom.dot(nTo))
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
	c1, c2 := nFrom.cross(nTo), nFrom.cross(nNorth)
	bearing := c1.angleTo(c2, nFrom)
	return Mod(ToDegrees(bearing)+360., 360.)
}

func (VectorCalculator) PointOnBearing(from *Point, distRad, bearingDeg float64, ctx GeoContext) *Point {
	nFrom := newNEWithPoint(from)
	bearing := ToRadians(bearingDeg)
	de := nNorth.cross(nFrom).unit()
	dn := nFrom.cross(de)
	deSin := de.mul(Sin(bearing))
	dnCos := dn.mul(Cos(bearing))
	d := dnCos.add(deSin)
	x := nFrom.mul(Cos(distRad))
	y := d.mul(Sin(distRad))
	return x.add(y).toPoint()
}

func (VectorCalculator) Area(s Shape) float64 {
	return sphereCalc.Area(s)
}

func (vc *VectorCalculator) MinCoverCircle(points ...*Point) (*Circle, error) {
	n := len(points)
	if n == 0 {
		return nil, errors.New("empty points")
	}

	ps := make([]*Point, n)
	copy(ps, points)
	rand.Shuffle(n, func(i, j int) {
		ps[i], ps[j] = ps[j], ps[i]
	})
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
	p12Mid, p23Mid := vc.Mid(p1, p2, nil), vc.Mid(p2, p3, nil)
	crs12Mid2 := vc.Bearing(p12Mid, p2)
	crs23Mid3 := vc.Bearing(p23Mid, p3)
	p4, err := vc.IntersectionOfTwoPath( //TODO: calc with clever way
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

func (vc *VectorCalculator) IntersectionOfTwoPath(pa1, pa2, pb1, pb2 *Point) (*Point, error) {
	p1, p2 := newNE(pa1.X(), pa1.Y()), newNE(pb1.X(), pb1.Y())
	p11, p22 := newNE(pa2.X(), pa2.Y()), newNE(pb2.X(), pb2.Y())
	c1, c2 := p1.cross(p11), p2.cross(p22)
	i1, i2 := c1.cross(c2), c2.cross(c1)
	mid := p1.add(p2).add(p11).add(p22)
	if mid.dot(i1) > 0 {
		return i1.toPoint(), nil
	}
	return i2.toPoint(), nil
}

func (vc *VectorCalculator) Intersection(pa *Point, bearingDegA float64, pb *Point, bearingDegB float64) (*Point, error) {
	na, nb := newNEWithPoint(pa), newNEWithPoint(pb)
	bearingA, bearingB := ToRadians(bearingDegA), ToRadians(bearingDegB)

	dae := nNorth.cross(na).unit()
	dan := na.cross(dae)
	da := dan.mul(Cos(bearingA)).add(dae.mul(Sin(bearingA)))
	c1 := na.cross(da)

	dbe := nNorth.cross(nb).unit()
	dbn := nb.cross(dbe)
	db := dbn.mul(Cos(bearingB)).add(dbe.mul(Sin(bearingB)))
	c2 := nb.cross(db)

	nc := c1.cross(c2)
	return nc.toPoint(), nil
}
