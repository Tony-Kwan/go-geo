package geo

import (
	"fmt"
	"testing"
)

var (
	calc       VectorCalculator
	sphereCalc SphereCalculator
)

func TestVectorCalculator_LngLat2nE(t *testing.T) {
	fmt.Println(newNE(0, 90))
}

func TestVectorCalculator_nE2Lnglat(t *testing.T) {
	ll := newPoint(newNE(0, 90))
	fmt.Println(ll)
}

func TestVectorCalculator_meanPosition(t *testing.T) {
	points := []*vector3{
		newNE(-81.9465923309326, 36.309868813086695),
		newNE(-84.01296615600586, 33.51764054105411),
	}
	m := newPoint(calc.meanPosition(points...))
	fmt.Println(m)
	for _, nE := range points {
		fmt.Println(calc.Distance(m, newPoint(nE)) * EarthRadius)
	}
}

func TestVectorCalculator_Distance(t *testing.T) {
	p1 := NewPoint(-83.85829925537108, 37.351328227794866, nil)
	p2 := NewPoint(-86.93447113037108, 32.600915527883345, nil)
	dis1 := calc.Distance(p1, p2)
	dis2 := sphereCalc.Distance(p1, p2)
	fmt.Println(dis1, dis2)
	fmt.Println(dis1*EarthRadius, dis2*EarthRadius)
	fmt.Println((dis1 - dis2) * EarthRadius)
}

func TestVectorCalculator_Intersection(t *testing.T) {
	p1 := NewPoint(113.6679634862422, 22.720836440519747, nil)
	p2 := NewPoint(113.82518535746098, 22.504796985005584, nil)
	p3 := NewPoint(114.0053646777264, 22.60559468990624, nil)
	p12mid := sphereCalc.Mid(p1, p2, nil)
	p23mid := sphereCalc.Mid(p2, p3, nil)
	crs12mid4 := sphereCalc.Bearing(p12mid, p2) - 90
	crs23mid4 := sphereCalc.Bearing(p23mid, p3) - 90
	fmt.Println(p12mid, crs12mid4)
	fmt.Println(p23mid, crs23mid4)
	p4, err := calc.Intersection(p12mid, crs12mid4, p23mid, crs23mid4)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(p4)

	fmt.Println(calc.Distance(p4, p1) * EarthRadius)
	fmt.Println(calc.Distance(p4, p2) * EarthRadius)
	fmt.Println(calc.Distance(p4, p3) * EarthRadius)
}
