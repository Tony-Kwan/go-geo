package geo

import (
	"fmt"
	"testing"
)

func TestVectorCalculator_LngLat2nE(t *testing.T) {
	fmt.Println(newNE(0, 90))
}

func TestVectorCalculator_nE2Lnglat(t *testing.T) {
	ll := newNE(0, 90).toPoint()
	fmt.Println(ll)
}

func TestVectorCalculator_meanPosition(t *testing.T) {
	points := []*Point{
		NewPoint(-81.9465923309326, 36.309868813086695, nil),
		NewPoint(-84.01296615600586, 33.51764054105411, nil),
	}
	m := vectorCalc.MeanPosition(points...)
	fmt.Println(m)
	for _, nE := range points {
		fmt.Println(vectorCalc.Distance(m, nE) * EarthRadius)
	}
}

func TestVectorCalculator_Mid(t *testing.T) {
	from := NewPoint(-81.9465923309326, 36.309868813086695, nil)
	to := NewPoint(-84.01296615600586, 33.51764054105411, nil)
	fmt.Println(vectorCalc.Mid(from, to, GeoCtx))
	fmt.Println(sphereCalc.Mid(from, to, GeoCtx))
}

func TestVectorCalculator_PointOnBearing(t *testing.T) {
	info.Println(p1, "->", p2)
	dist := vectorCalc.Distance(p1, p2)
	bearingDeg := vectorCalc.Bearing(p1, p2)
	p3 := vectorCalc.PointOnBearing(p1, dist, bearingDeg, GeoCtx)
	info.Println(dist * EarthRadius)
	info.Println(bearingDeg)
	info.Println(p3)
}

func TestVectorCalculator_Distance(t *testing.T) {
	p1 := NewPoint(-83.85829925537108, 37.351328227794866, nil)
	p2 := NewPoint(-86.93447113037108, 32.600915527883345, nil)
	dis1 := vectorCalc.Distance(p1, p2)
	dis2 := sphereCalc.Distance(p1, p2)
	fmt.Println(dis1, dis2)
	fmt.Println(dis1*EarthRadius, dis2*EarthRadius)
	fmt.Println((dis1 - dis2) * EarthRadius)
}

func TestVectorCalculator_IntersectionOfTwoPath(t *testing.T) {
	p1 := NewPoint(113.021085, 23.292487, nil)
	p2 := NewPoint(113.212321, 23.253895, nil)
	p3 := NewPoint(113.212321, 23.253895+E10, nil)

	p4, err := vectorCalc.IntersectionOfTwoPath(p1, p2, p1, p3)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(p4)
}

func TestVectorCalculator_Triangulation(t *testing.T) {
	p1 := NewPoint(113.021085, 23.292487, nil)
	p2 := NewPoint(113.212321, 23.253895, nil)
	p3 := NewPoint(113.121248, 22.873807, nil)
	bearing12 := vectorCalc.Bearing(p1, p2)
	bearing23 := vectorCalc.Bearing(p2, p3)
	bearing32 := vectorCalc.Bearing(p3, p2)
	tests := []struct {
		points   [2]*Point
		bearings [2]float64
		want     *Point
	}{
		{
			points:   [2]*Point{p1, p2},
			bearings: [2]float64{bearing12, bearing23},
			want:     p2,
		},
		{
			points:   [2]*Point{p1, p3},
			bearings: [2]float64{bearing12, bearing32},
			want:     p2,
		},
		{
			points:   [2]*Point{NewPoint(0, 0, nil), NewPoint(1, 0, nil)},
			bearings: [2]float64{0, 90},
			want:     NewPoint(0, 0, nil),
		},
	}
	for _, test := range tests {
		p4, err := vectorCalc.Triangulation(test.points[0], test.bearings[0], test.points[1], test.bearings[1])
		if err != nil {
			t.Error(err)
			continue
		} else if !p4.ApproxEqual(test.want) {
			t.Errorf("expect=%v, found=%v", test.want, p4)
		}
	}
}

func TestVectorCalculator_Circumcenter(t *testing.T) {
	p1 := NewPoint(113.021085, 23.292487, nil)
	p2 := NewPoint(113.212321, 23.253895, nil)
	p3 := NewPoint(113.121248, 22.873807, nil)
	c, err := vectorCalc.Circumcenter(p1, p2, p3)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(c)

	fmt.Println(vectorCalc.Distance(c, p1) * EarthRadius)
	fmt.Println(vectorCalc.Distance(c, p2) * EarthRadius)
	fmt.Println(vectorCalc.Distance(c, p3) * EarthRadius)
}

func TestVectorCalculator_MinCoverCircle(t *testing.T) {
	points := []*Point{
		//NewPoint(7.098206748478991, 46.658250853066654, nil),
		//NewPoint(7.096818204096314, 46.6255295921, nil),
		//NewPoint(7.081215882059562, 46.62424586456617, nil),
		//NewPoint(7.068605669086119, 46.655815315282965, nil),
		//NewPoint(7.098206748478991, 46.658250853066654, nil),

		//NewPoint(7.107510838765815, 46.56415356342374, nil),
		//NewPoint(7.120093200561858, 46.53258287949015, nil),
		//NewPoint(7.090541763528792, 46.530148356963075, nil),
		//NewPoint(7.091926209321755, 46.56286965110933, nil),
		//NewPoint(7.107510838765815, 46.56415356342374, nil),

		NewPoint(3.4187225890205233, 46.20547150536718, nil),
		NewPoint(3.4039554760716264, 46.13243249643665, nil),
		NewPoint(3.3885327467298407, 46.13392817417585, nil),
		NewPoint(3.403279258370794, 46.20696665022393, nil),
		NewPoint(3.4187225890205233, 46.20547150536718, nil),
	}
	circle, err := vectorCalc.MinCoverCircle(points...)
	if err != nil {
		t.Error(err)
		return
	}
	//fmt.Println(circle)
	fmt.Println(circle.radius * EarthRadius)
}

func TestVectorCalculator_MinCoverCircle_2(t *testing.T) {
	for i := 0; i < 1000; i++ {
		TestVectorCalculator_MinCoverCircle(t)
	}
}

func TestVectorCalculator_AreaOfTriangle(t *testing.T) {
	tri := &Triangle{
		a: NewPoint(0, 0, nil),
		b: NewPoint(0, 89, nil),
		c: NewPoint(90, 45, nil),
	}
	t.Log(vectorCalc.areaOfTriangle(tri) * EarthRadius2)
}

func TestVectorCalculator_AreaOfPolygon(t *testing.T) {
	polygon := &Polygon{shell: LinearRing{
		*NewPoint(-92.38059997558592, 45.38157243512828, nil),
		*NewPoint(-90.47378540039062, 40.582670638095294, nil),
		*NewPoint(-81.50894165039061, 39.77397788285171, nil),
		*NewPoint(-87.19505310058592, 37.1594957106433, nil),
		*NewPoint(-86.93138122558594, 32.01972036197235, nil),
		*NewPoint(-91.06224060058594, 35.909908145897035, nil),
		*NewPoint(-95.95733642578125, 31.974007590177635, nil),
		*NewPoint(-94.81475830078125, 37.60580020781012, nil),
		*NewPoint(-102.54913330078125, 38.29882852868994, nil),
		*NewPoint(-93.93585205078125, 40.27140563877154, nil),
		*NewPoint(-92.38059997558592, 45.38157243512828, nil),
	}}
	t.Log(vectorCalc.areaOfPolygon(polygon))
}
