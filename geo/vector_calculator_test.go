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
	m := vectorCalc.meanPosition(points...)
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

func TestVectorCalculator_Intersection(t *testing.T) {
	p1 := NewPoint(113.021085, 23.292487, nil)
	p2 := NewPoint(113.212321, 23.253895, nil)
	p3 := NewPoint(113.121248, 22.873807, nil)
	p12mid := vectorCalc.Mid(p1, p2, nil)
	p23mid := vectorCalc.Mid(p2, p3, nil)
	crs12mid4 := vectorCalc.Bearing(p12mid, p2) + 90
	crs23mid4 := vectorCalc.Bearing(p23mid, p3) + 90
	fmt.Println(p12mid, crs12mid4)
	fmt.Println(p23mid, crs23mid4)
	p4, err := vectorCalc.Intersection(p12mid, crs12mid4, p23mid, crs23mid4)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(p4)

	fmt.Println(vectorCalc.Distance(p4, p1) * EarthRadius)
	fmt.Println(vectorCalc.Distance(p4, p2) * EarthRadius)
	fmt.Println(vectorCalc.Distance(p4, p3) * EarthRadius)
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
		NewPoint(113.021085, 23.292487, nil),
		NewPoint(113.212321, 23.253895, nil),
		NewPoint(113.121248, 22.873807, nil),
		NewPoint(112.930559, 22.91243, nil),
		NewPoint(113.021085, 23.292487, nil),
	}
	circle, err := vectorCalc.MinCoverCircle(points...)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(circle)
	fmt.Println(circle.radius * EarthRadius)
}
