package geo

import (
	"errors"
	"fmt"
	. "math"
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

func (vc *VectorCalculator) Circumcenter(pa, pb, pc *Point) (*Point, error) {
	p12Mid, p23Mid := vc.Mid(pa, pb, nil), vc.Mid(pb, pc, nil)
	bearing12Mid2 := vc.Bearing(p12Mid, pb)
	bearing23Mid3 := vc.Bearing(p23Mid, pc)
	c1, c2 := p12Mid.greatCircle(bearing12Mid2+90), p23Mid.greatCircle(bearing23Mid3+90)
	i1, i2, err := vc.intersectionOfTwoGreatCircle(c1, c2)
	if err != nil {
		return nil, err
	}
	nMid := newNEWithPoint(vc.meanPosition(pa, pb, pc))
	if nMid.dot(i1) > 0 {
		return i1.toPoint(), nil
	}
	return i2.toPoint(), nil
}

func (vc *VectorCalculator) intersectionOfTwoGreatCircle(c1, c2 *vector3) (*vector3, *vector3, error) {
	if c1.ApproxEqual(c2) || c1.mul(-1).ApproxEqual(c2) {
		return nil, nil, fmt.Errorf("infinite solutions: %v, %v", c1, c2)
	}
	i := c1.cross(c2)
	return i, i.mul(-1), nil
}

func (vc *VectorCalculator) IntersectionOfTwoGreatCircle(pa *Point, bearingDegA float64, pb *Point, bearingDegB float64) (*Point, *Point, error) {
	c1, c2 := pa.greatCircle(bearingDegA), pb.greatCircle(bearingDegB)
	i1, i2, err := vc.intersectionOfTwoGreatCircle(c1, c2)
	if err != nil {
		return nil, nil, err
	}
	return i1.toPoint(), i2.toPoint(), nil
}

func (vc *VectorCalculator) IntersectionOfTwoPath(pa1, pa2, pb1, pb2 *Point) (*Point, error) {
	na1, nb1 := newNE(pa1.X(), pa1.Y()), newNE(pb1.X(), pb1.Y())
	na2, nb2 := newNE(pa2.X(), pa2.Y()), newNE(pb2.X(), pb2.Y())
	c1, c2 := na1.cross(na2), nb1.cross(nb2)
	i1, i2, err := vc.intersectionOfTwoGreatCircle(c1, c2)
	if err != nil {
		return nil, err
	}
	mid := newNEWithPoint(vc.meanPosition(pa1, pa2, pb1, pb2))
	if mid.dot(i1) > 0 { //select nearest intersection of mid of all points
		return i1.toPoint(), nil
	}
	return i2.toPoint(), nil
}

func (vc *VectorCalculator) Triangulation(pa *Point, bearingDegA float64, pb *Point, bearingDegB float64) (*Point, error) {
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

	i1, i2, err := vc.intersectionOfTwoGreatCircle(c1, c2)
	if err != nil {
		return nil, err
	}
	dir1 := sign(c1.cross(na).dot(i1))
	dir2 := sign(c2.cross(nb).dot(i2))
	switch dir1 + dir2 {
	case 2:
		return i1.toPoint(), nil
	case -2:
		return i2.toPoint(), nil
	case 0:
		if na.add(nb).dot(i1) > 0 {
			return i1.toPoint(), nil
		}
		return i2.toPoint(), nil
	default:
		return nil, fmt.Errorf("program should not run here: dir=%d, dir2=%d", dir1, dir2)
	}
}
