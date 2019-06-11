package geo

import (
	. "math"
)

type SphereCalculator struct {
}

func (calc *SphereCalculator) Distance(from, to *Point) float64 {
	return calc.DistanceXY(from.X(), from.Y(), to.X(), to.Y())
}

// haversine formula
// return: radians
func (calc *SphereCalculator) DistanceXY(fromX, fromY, toX, toY float64) float64 {
	lng1, lat1 := ToRadians(fromX), ToRadians(fromY)
	lng2, lat2 := ToRadians(toX), ToRadians(toY)
	if lng1 == lng2 && lat1 == lat2 {
		return 0
	}
	hSinLng, hSinLat := Sin((lng1-lng2)/2), Sin((lat1-lat2)/2)
	h := hSinLat*hSinLat + Cos(lat1)*Cos(lat2)*hSinLng*hSinLng
	if h > 1 {
		h = 1
	}
	return 2 * Atan2(Sqrt(h), Sqrt(1-h))
}

func (calc *SphereCalculator) Bearing(from, to *Point) float64 {
	lng1, lat1 := ToRadians(from.X()), ToRadians(from.Y())
	lng2, lat2 := ToRadians(to.X()), ToRadians(to.Y())
	dLng := lng2 - lng1
	x := Sin(dLng) * Cos(lat2)
	y := Cos(lat1)*Sin(lat2) - Sin(lat1)*Cos(lat2)*Cos(dLng)
	return ToDegrees(Mod(Pi/2-Atan2(y, x)+2*Pi, 2*Pi))
}

func (calc *SphereCalculator) PointOnBearing(from *Point, distRad, bearingDeg float64, ctx GeoContext) *Point {
	lng, lat := ToRadians(from.X()), ToRadians(from.Y())
	bearingRad := ToRadians(bearingDeg)
	endLat := Asin(Sin(lat)*Cos(distRad) + Cos(lat)*Sin(distRad)*Cos(bearingRad))
	endLng := lng + Atan2(Sin(bearingRad)*Sin(distRad)*Cos(lat), Cos(distRad)-Sin(lat)*Sin(endLat))
	return NewPoint((Mod(ToDegrees(endLng)+540, 360))-180, ToDegrees(endLat), ctx)
}

func (SphereCalculator) MeanPosition(points ...*Point) *Point {
	meanPoint := NewPoint(0, 0, nil)
	for _, p := range points {
		meanPoint.x += p.x
		meanPoint.y += p.y
	}
	meanPoint.x /= float64(len(points))
	meanPoint.y /= float64(len(points))
	return meanPoint
}

func (calc *SphereCalculator) Mid(from, to *Point, ctx GeoContext) *Point {
	if from.Equals(to) {
		return from.clone().(*Point)
	}
	lng1, lat1 := ToRadians(from.X()), ToRadians(from.Y())
	lng2, lat2 := ToRadians(to.X()), ToRadians(to.Y())
	dLng := lng2 - lng1
	bx, by := Cos(lat2)*Cos(dLng), Cos(lat2)*Sin(dLng)
	midLat := Atan2(Sin(lat1)+Sin(lat2), Sqrt(Pow(Cos(lat1)+bx, 2)+Pow(by, 2)))
	midLng := lng1 + Atan2(by, Cos(lat1)+bx)
	return NewPoint(Mod(ToDegrees(midLng)+540, 360)-180, ToDegrees(midLat), ctx)
}

func (calc *SphereCalculator) Area(s Shape) float64 {
	switch s.(type) {
	case *Circle:
		return calc.areaOfCircle(s.(*Circle))
	//case *Triangle:
	//	return calc.areaOfTriangle(s.(*Triangle))
	case *Rectangle:
		return calc.areaOfRectangle(s.(*Rectangle))
	default:
		erro.Printf("unsupported shape type: %+v\n", s)
		return -1
	}
}

func (calc *SphereCalculator) areaOfCircle(circle *Circle) float64 {
	lat := ToRadians(90 - circle.radius)
	return 2 * Pi * (1 - Sin(lat))
}

func (calc *SphereCalculator) areaOfRectangle(rect *Rectangle) float64 {
	w := rect.maxX - rect.minX
	if w < 0 {
		w += 360.
	}
	minLat, maxLat := ToRadians(rect.minY), ToRadians(rect.maxY)
	return Pi / 180. * Abs(Sin(minLat)-Sin(maxLat)) * w
}

func (calc *SphereCalculator) areaOfTriangle(tri *Triangle) float64 {
	//A, B, C := calc.Distance(tri.B, tri.C), calc.Distance(tri.A, tri.C), calc.Distance(tri.A, tri.B)
	//A := Acos((Cos(A) - Cos(B)*Cos(C)) / Sin(B) / Sin(C))
	//B := Acos((Cos(B) - Cos(A)*Cos(C)) / Sin(A) / Sin(C))
	//C := Acos((Cos(C) - Cos(A)*Cos(B)) / Sin(A) / Sin(B))
	//fmt.Println(ToDegrees(A), ToDegrees(B), ToDegrees(C))
	//return A + B + C - Pi
	return 0
}
