package test

import (
	"fmt"
	"github.com/Tony-Kwan/go-geo/geo"
	"github.com/Tony-Kwan/go-geo/geo/io/wkt"
	"testing"
)

func Benchmark_MinCoverCircle(b *testing.B) {
	wktStr := "POLYGON((120.584017 30.334264,120.616301 30.273781,120.526286 30.244407,120.501884 30.284578,120.584017 30.334264,120.345366 30.212234,120.368149 30.177084,120.284766 30.127091,120.248699 30.187233,120.345366 30.212234,120.52629261821812 30.24436668295212,120.52697582270933 30.233299334035987,120.52382297492878 30.216926182050383,120.50816882313957 30.192957338856885,120.483794 30.177527,120.43219775825095 30.16072212797702,120.40038343656032 30.16166599384457,120.37713335485847 30.17080612177303,120.368149 30.177084,120.34537517739905 30.212194490559547,120.34438558172529 30.217674992600042,120.34778835792177 30.24504085732033,120.36426475505374 30.26859901749025,120.382862 30.280802,120.416186504205 30.295798257235646,120.44058500229838 30.30241947134751,120.4661354368617 30.301431620999843,120.49500712737346 30.289765916116554,120.501884 30.284578,120.52629261821812 30.24436668295212,120.5953138486885 30.39368525224143,120.68281747356784 30.230453059776597,120.27156347260748 30.06603842681942,120.1842060444106 30.229270619443145,120.5953138486885 30.39368525224143,120.574023 30.350935,120.606332 30.287108,120.514625 30.261084,120.491845 30.30125,120.574023 30.350935,120.335386 30.230601,120.358116 30.190422,120.274747 30.140427,120.242044 30.200586,120.335386 30.230601,120.51463306543144 30.26104406230654,120.51572370882518 30.250001097220096,120.51317486382023 30.23354956586078,120.50876208686653 30.223142956821338,120.49840831046319 30.209162540888467,120.477138 30.194198,120.42572487592481 30.17352591501829,120.39388785681989 30.17355089354241,120.36504062974964 30.185198748272395,120.358116 30.190422,120.33539312139733 30.23056081396441,120.33498290762775 30.247155269232408,120.33749194435569 30.258024365139146,120.35283562834141 30.28214751503115,120.376206 30.297473,120.42752707781861 30.314916441792395,120.45307374576718 30.316036178574258,120.48310678971662 30.30680009209954,120.491845 30.30125,120.51463306543144 30.26104406230654,120.58865815180187 30.408691824196655,120.67618595808533 30.24546393634518,120.26488532399638 30.0810311963203,120.17750385884608 30.244259084330878,120.58865815180187 30.408691824196655))"
	var polygon = wkt.MustPolygon(wkt.WktReader{}.Read(wktStr))
	for i := 0; i < b.N; i++ {
		circle, err := calc.MinCoverCircle(polygon.Shell...)
		if err != nil {
			b.Error(err)
			return
		}
		fmt.Println(circle.GetRadius() * geo.EarthRadius)
	}
}

func Test_MinCoverCircle(t *testing.T) {
	wktStr := "POLYGON((120.584017 30.334264,120.616301 30.273781,120.526286 30.244407,120.501884 30.284578,120.584017 30.334264,120.345366 30.212234,120.368149 30.177084,120.284766 30.127091,120.248699 30.187233,120.345366 30.212234,120.52629261821812 30.24436668295212,120.52697582270933 30.233299334035987,120.52382297492878 30.216926182050383,120.50816882313957 30.192957338856885,120.483794 30.177527,120.43219775825095 30.16072212797702,120.40038343656032 30.16166599384457,120.37713335485847 30.17080612177303,120.368149 30.177084,120.34537517739905 30.212194490559547,120.34438558172529 30.217674992600042,120.34778835792177 30.24504085732033,120.36426475505374 30.26859901749025,120.382862 30.280802,120.416186504205 30.295798257235646,120.44058500229838 30.30241947134751,120.4661354368617 30.301431620999843,120.49500712737346 30.289765916116554,120.501884 30.284578,120.52629261821812 30.24436668295212,120.5953138486885 30.39368525224143,120.68281747356784 30.230453059776597,120.27156347260748 30.06603842681942,120.1842060444106 30.229270619443145,120.5953138486885 30.39368525224143,120.574023 30.350935,120.606332 30.287108,120.514625 30.261084,120.491845 30.30125,120.574023 30.350935,120.335386 30.230601,120.358116 30.190422,120.274747 30.140427,120.242044 30.200586,120.335386 30.230601,120.51463306543144 30.26104406230654,120.51572370882518 30.250001097220096,120.51317486382023 30.23354956586078,120.50876208686653 30.223142956821338,120.49840831046319 30.209162540888467,120.477138 30.194198,120.42572487592481 30.17352591501829,120.39388785681989 30.17355089354241,120.36504062974964 30.185198748272395,120.358116 30.190422,120.33539312139733 30.23056081396441,120.33498290762775 30.247155269232408,120.33749194435569 30.258024365139146,120.35283562834141 30.28214751503115,120.376206 30.297473,120.42752707781861 30.314916441792395,120.45307374576718 30.316036178574258,120.48310678971662 30.30680009209954,120.491845 30.30125,120.51463306543144 30.26104406230654,120.58865815180187 30.408691824196655,120.67618595808533 30.24546393634518,120.26488532399638 30.0810311963203,120.17750385884608 30.244259084330878,120.58865815180187 30.408691824196655))"
	var polygon = wkt.MustPolygon(wkt.WktReader{}.Read(wktStr))
	circle, err := calc.MinCoverCircle(polygon.Shell...)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(circle.GetRadius() * geo.EarthRadius)
	//cirPolygon := circle.ToPolygon(32)
	//wkt := fmt.Sprintf("GEOMETRYCOLLECTION(%s,%s)", polygon.String(), cirPolygon.String())
	//t.Log(wkt)
	//clipboard.WriteAll(wkt)
}

func Test_Circumcenter(t *testing.T) {
	p1, p2, p3 := geo.NewPoint(120.58865815180188, 30.408691824196666, nil), geo.NewPoint(120.67618595808533, 30.24546393634516, nil), geo.NewPoint(120.68281747356784, 30.230453059776593, nil)
	polygon := geo.NewPolygon([]geo.Point{p1, p2, p3})
	t.Log(polygon.String())
	t.Log(calc.Bearing(p1, p2))
	t.Log(calc.Bearing(p1, p3))
	t.Log(calc.Bearing(p2, p3))

	c, err := calc.Circumcenter(p1, p2, p3)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(c)
	t.Log(calc.Distance(c, p1) * geo.EarthRadius)
	t.Log(calc.Distance(c, p2) * geo.EarthRadius)
	t.Log(calc.Distance(c, p3) * geo.EarthRadius)

	circle, _ := calc.MinCoverCircle(p2, p1, p3)
	t.Log(circle.GetRadius() * geo.EarthRadius)
}

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
//func TestVectorCalculator_PointOnBearing(t *testing.T) {
//	info.Println(p1, "->", p2)
//	dist := vectorCalc.Distance(p1, p2)
//	bearingDeg := vectorCalc.Bearing(p1, p2)
//	p3 := vectorCalc.PointOnBearing(p1, dist, bearingDeg, GeoCtx)
//	info.Println(dist * EarthRadius)
//	info.Println(bearingDeg)
//	info.Println(p3)
//}
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
