package test

import (
	"github.com/Tony-Kwan/go-geo/geo"
	"testing"
)

//func TestVectorCalculator_LngLat2nE(t *testing.T) {
//	fmt.Println(newNE(0, 90))
//}
//
//func TestVectorCalculator_nE2Lnglat(t *testing.T) {
//	ll := newNE(0, 90).toPoint()
//	fmt.Println(ll)
//}
//
//func TestVectorCalculator_meanPosition(t *testing.T) {
//	points := []*Point{
//		NewPoint(-81.9465923309326, 36.309868813086695, nil),
//		NewPoint(-84.01296615600586, 33.51764054105411, nil),
//	}
//	m := vectorCalc.MeanPosition(points...)
//	fmt.Println(m)
//	for _, nE := range points {
//		fmt.Println(vectorCalc.Distance(m, nE) * EarthRadius)
//	}
//}
//
//func TestVectorCalculator_Mid(t *testing.T) {
//	from := NewPoint(-81.9465923309326, 36.309868813086695, nil)
//	to := NewPoint(-84.01296615600586, 33.51764054105411, nil)
//	fmt.Println(vectorCalc.Mid(from, to, GeoCtx))
//	fmt.Println(sphereCalc.Mid(from, to, GeoCtx))
//}
//

func TestVectorCalculator_Bearing(t *testing.T) {
	calc := &geo.VectorCalculator{}
	t.Log(calc.Bearing(p1, p2))
}

func TestVectorCalculator_PointOnBearing(t *testing.T) {
	t.Log(p1, "->", p2)
	vectorCalc := &geo.VectorCalculator{}
	bearingDeg := vectorCalc.Bearing(p1, p2)
	p3 := vectorCalc.PointOnBearing(p2, 3000.0/geo.EarthRadius, bearingDeg, geo.GeoCtx)
	p4 := vectorCalc.PointOnBearing(p3, 600/geo.EarthRadius, bearingDeg-90, geo.GeoCtx)
	t.Log(bearingDeg)
	t.Log(p3)
	t.Log(p4)
}

//
//func TestVectorCalculator_Distance(t *testing.T) {
//	p1 := NewPoint(-83.85829925537108, 37.351328227794866, nil)
//	p2 := NewPoint(-86.93447113037108, 32.600915527883345, nil)
//	dis1 := vectorCalc.Distance(p1, p2)
//	dis2 := sphereCalc.Distance(p1, p2)
//	fmt.Println(dis1, dis2)
//	fmt.Println(dis1*EarthRadius, dis2*EarthRadius)
//	fmt.Println((dis1 - dis2) * EarthRadius)
//}
//
//func TestVectorCalculator_IntersectionOfTwoPath(t *testing.T) {
//	p1 := NewPoint(113.021085, 23.292487, nil)
//	p2 := NewPoint(113.212321, 23.253895, nil)
//	p3 := NewPoint(113.212321, 23.253895+E10, nil)
//
//	p4, err := vectorCalc.IntersectionOfTwoPath(p1, p2, p1, p3)
//	if err != nil {
//		t.Error(err)
//		return
//	}
//	fmt.Println(p4)
//}
//
//func TestVectorCalculator_Triangulation(t *testing.T) {
//	p1 := NewPoint(113.021085, 23.292487, nil)
//	p2 := NewPoint(113.212321, 23.253895, nil)
//	p3 := NewPoint(113.121248, 22.873807, nil)
//	bearing12 := vectorCalc.Bearing(p1, p2)
//	bearing23 := vectorCalc.Bearing(p2, p3)
//	bearing32 := vectorCalc.Bearing(p3, p2)
//	tests := []struct {
//		points   [2]*Point
//		bearings [2]float64
//		want     *Point
//	}{
//		{
//			points:   [2]*Point{p1, p2},
//			bearings: [2]float64{bearing12, bearing23},
//			want:     p2,
//		},
//		{
//			points:   [2]*Point{p1, p3},
//			bearings: [2]float64{bearing12, bearing32},
//			want:     p2,
//		},
//		{
//			points:   [2]*Point{NewPoint(0, 0, nil), NewPoint(1, 0, nil)},
//			bearings: [2]float64{0, 90},
//			want:     NewPoint(0, 0, nil),
//		},
//	}
//	for _, test := range tests {
//		p4, err := vectorCalc.Triangulation(test.points[0], test.bearings[0], test.points[1], test.bearings[1])
//		if err != nil {
//			t.Error(err)
//			continue
//		} else if !p4.ApproxEqual(test.want) {
//			t.Errorf("expect=%v, found=%v", test.want, p4)
//		}
//	}
//}
//
//func TestVectorCalculator_Circumcenter(t *testing.T) {
//	p1 := NewPoint(113.021085, 23.292487, nil)
//	p2 := NewPoint(113.212321, 23.253895, nil)
//	p3 := NewPoint(113.121248, 22.873807, nil)
//	c, err := vectorCalc.Circumcenter(p1, p2, p3)
//	if err != nil {
//		t.Error(err)
//		return
//	}
//	fmt.Println(c)
//
//	fmt.Println(vectorCalc.Distance(c, p1) * EarthRadius)
//	fmt.Println(vectorCalc.Distance(c, p2) * EarthRadius)
//	fmt.Println(vectorCalc.Distance(c, p3) * EarthRadius)
//}
//
//func TestVectorCalculator_MinCoverCircle(t *testing.T) {
//	points := []*Point{
//		//NewPoint(7.098206748478991, 46.658250853066654, nil),
//		//NewPoint(7.096818204096314, 46.6255295921, nil),
//		//NewPoint(7.081215882059562, 46.62424586456617, nil),
//		//NewPoint(7.068605669086119, 46.655815315282965, nil),
//		//NewPoint(7.098206748478991, 46.658250853066654, nil),
//
//		//NewPoint(7.107510838765815, 46.56415356342374, nil),
//		//NewPoint(7.120093200561858, 46.53258287949015, nil),
//		//NewPoint(7.090541763528792, 46.530148356963075, nil),
//		//NewPoint(7.091926209321755, 46.56286965110933, nil),
//		//NewPoint(7.107510838765815, 46.56415356342374, nil),
//
//		NewPoint(3.4187225890205233, 46.20547150536718, nil),
//		NewPoint(3.4039554760716264, 46.13243249643665, nil),
//		NewPoint(3.3885327467298407, 46.13392817417585, nil),
//		NewPoint(3.403279258370794, 46.20696665022393, nil),
//		NewPoint(3.4187225890205233, 46.20547150536718, nil),
//	}
//	circle, err := vectorCalc.MinCoverCircle(points...)
//	if err != nil {
//		t.Error(err)
//		return
//	}
//	//fmt.Println(circle)
//	fmt.Println(circle.radius * EarthRadius)
//}
//
//func TestVectorCalculator_MinCoverCircle_2(t *testing.T) {
//	for i := 0; i < 1000; i++ {
//		TestVectorCalculator_MinCoverCircle(t)
//	}
//}
//
//func TestVectorCalculator_AreaOfPolygon(t *testing.T) {
//	polygon := &Polygon{shell: LinearRing{
//		*NewPoint(-1.6991043090820295, 43.32317958748908, nil),
//		*NewPoint(3.936882019042976, 42.161367532874806, nil),
//		*NewPoint(3.581542968750002, 43.11953018614551, nil),
//		*NewPoint(6.507682800292972, 42.94360648947236, nil),
//		*NewPoint(7.775917053222664, 43.726700289570346, nil),
//		*NewPoint(7.122917175292971, 44.847836846819575, nil),
//		*NewPoint(6.737022399902351, 46.3104186157758, nil),
//		*NewPoint(7.2142410278320375, 47.52856107277904, nil),
//		*NewPoint(8.028602600097656, 49.073640995014955, nil),
//		*NewPoint(2.7785110473632812, 51.27093716704937, nil),
//		*NewPoint(1.4405822753906263, 50.82654161609699, nil),
//		*NewPoint(1.5284729003906241, 49.996484644867365, nil),
//		*NewPoint(-1.6918945312499958, 49.65629401936761, nil),
//		*NewPoint(-1.718673706054683, 48.69005384564221, nil),
//		*NewPoint(-5.080490112304683, 48.803471664383096, nil),
//		*NewPoint(-4.904708862304683, 47.97912122349871, nil),
//		*NewPoint(-2.723236083984373, 47.16450869129574, nil),
//		*NewPoint(-1.6956710815429656, 46.10775536585561, nil),
//		*NewPoint(-1.722450256347653, 44.31844432059691, nil),
//		*NewPoint(-1.6991043090820295, 43.32317958748908, nil),
//	}}
//	t.Log(vectorCalc.areaOfPolygon(polygon) * EarthRadius2)
//}
//func TestVectorCalculator_Triangle_Area(t *testing.T) {
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
//		NewTriangle(
//			NewPoint(0, 90, GeoCtx),
//			NewPoint(0, -90, GeoCtx),
//			NewPoint(1, 0, GeoCtx),
//			GeoCtx,
//		),
//	}
//	for _, tri := range tris {
//		info.Println(vectorCalc.Area(tri) * EarthRadius2)
//	}
//}
