package test

import (
	"github.com/Tony-Kwan/go-geo/geo"
	"testing"
)

var (
	p1 = geo.NewPoint(111.434, -7.603, geo.GeoCtx)
	p2 = geo.NewPoint(111.435683, -7.626217, geo.GeoCtx)
)

//
//func TestSphereCalculator_Distance(t *testing.T) {
//	dist := GeoCtx.GetCalculator().Distance(p1, p2)
//	info.Println(dist, dist*EarthRadius)
//}
//
func TestSphereCalculator_Bearing(t *testing.T) {
	calc := &geo.SphereCalculator{}
	t.Log(calc.Bearing(p1, p2))
}

func TestSphereCalculator_PointOnBearing(t *testing.T) {
	t.Log(p1, "->", p2)
	calc := &geo.SphereCalculator{}
	bearingDeg := calc.Bearing(p1, p2)
	p3 := calc.PointOnBearing(p2, 3000.0/geo.EarthRadius, bearingDeg, geo.GeoCtx)
	p4 := calc.PointOnBearing(p3, 600/geo.EarthRadius, bearingDeg-90, geo.GeoCtx)
	t.Log(bearingDeg)
	t.Log(p3)
	t.Log(p4)
}

//
//func TestSphereCalculator_Mid(t *testing.T) {
//	info.Println(p1, "->", p2)
//	calc := &SphereCalculator{}
//	info.Println(calc.Mid(p1, p2, GeoCtx))
//}
//
//func TestSphereCalculator_Circle_Area(t *testing.T) {
//	circle := NewCircle(0, 0, 90, GeoCtx)
//	info.Println(sphereCalc.Area(circle) * EarthRadius2 * 2)
//}
//
//func TestSphereCalculator_Rectangle_Area(t *testing.T) {
//	rect := NewRectangle(0, 180, 0, 90, GeoCtx)
//	info.Println(sphereCalc.Area(rect) * EarthRadius2 * 4)
//}
//
//func TestSphereCalculator_Triangle_Area(t *testing.T) {
//	tris := []*Triangle{
//		NewTriangle(
//			NewPoint(0, 0, GeoCtx),
//			NewPoint(90, 0, GeoCtx),
//			NewPoint(0, 90, GeoCtx),
//			GeoCtx,
//		),
//		NewTriangle(
//			NewPoint(0, 0, GeoCtx),
//			NewPoint(90, 0, GeoCtx),
//			NewPoint(45, 45, GeoCtx),
//			GeoCtx,
//		),
//		NewTriangle(
//			NewPoint(0, 0, GeoCtx),
//			NewPoint(30, 30, GeoCtx),
//			NewPoint(45, 45, GeoCtx),
//			GeoCtx,
//		),
//		NewTriangle(
//			NewPoint(90, 0, GeoCtx),
//			NewPoint(-80, 0, GeoCtx),
//			NewPoint(40, 40, GeoCtx),
//			GeoCtx,
//		),
//	}
//	for _, tri := range tris {
//		info.Println(sphereCalc.Area(tri) * EarthRadius2)
//	}
//}
