package test

//var (
//	p1 = NewPoint(-94.57202911376953, 28.61737745885398, GeoCtx)
//	p2 = NewPoint(-82.85270690917969, 45.463742648776844, GeoCtx)
//)
//
//func TestSphereCalculator_Distance(t *testing.T) {
//	dist := GeoCtx.GetCalculator().Distance(p1, p2)
//	info.Println(dist, dist*EarthRadius)
//}
//
//func TestSphereCalculator_Bearing(t *testing.T) {
//	bearingRad := GeoCtx.GetCalculator().Bearing(p1, p2)
//	bearingDeg := ToDegrees(bearingRad)
//	info.Println(bearingRad)
//	info.Println(bearingDeg)
//}
//
//func TestSphereCalculator_PointOnBearing(t *testing.T) {
//	info.Println(p1, "->", p2)
//	calc := &SphereCalculator{}
//	dist := calc.Distance(p1, p2)
//	bearingDeg := calc.Bearing(p1, p2)
//	p3 := calc.PointOnBearing(p1, dist, bearingDeg, GeoCtx)
//	info.Println(dist * EarthRadius)
//	info.Println(bearingDeg)
//	info.Println(p3)
//}
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
