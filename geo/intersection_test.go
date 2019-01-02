package geo

import (
	"errors"
	"fmt"
	. "math"
	"testing"
)

func Test1(t *testing.T) {
	calc := GeoCtx.GetCalculator().(*SphereCalculator)
	p1 = NewPoint(113.80877344009748, 22.52216933506496, GeoCtx)
	p13 := NewPoint(113.92542765142109, 22.606924488995247, GeoCtx)
	p2 = NewPoint(113.90269032162367, 22.538125022580147, GeoCtx)
	p23 := NewPoint(113.83335981639193, 22.585522743725548, GeoCtx)

	crs13 := calc.Bearing(p1, p13)
	crs23 := calc.Bearing(p2, p23)
	//fmt.Println("crs13:", ToDegrees(crs13))
	//fmt.Println("crs23:", ToDegrees(crs23))
	f(p1, p2, crs13, crs23)
}

func Test2(t *testing.T) {
	calc := GeoCtx.GetCalculator().(*SphereCalculator)
	p1 = NewPoint(113.6679634862422, 22.720836440519747, GeoCtx)
	p2 = NewPoint(113.82518535746098, 22.504796985005584, GeoCtx)
	p3 := NewPoint(114.0053646777264, 22.60559468990624, GeoCtx)
	p12mid := calc.Mid(p1, p2, GeoCtx)
	p23mid := calc.Mid(p2, p3, GeoCtx)
	crs12mid4 := calc.Bearing(p1, p2) - Pi/2
	crs23mid4 := calc.Bearing(p2, p3) - Pi/2
	p4, err := f(p12mid, p23mid, crs12mid4, crs23mid4)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("POINT(%v %v)\n", p4.X(), p4.Y())
	fmt.Println("dis14:", calc.Distance(p1, p4)*EarthRadius)
	fmt.Println("dis24:", calc.Distance(p2, p4)*EarthRadius)
	fmt.Println("dis34:", calc.Distance(p3, p4)*EarthRadius)

	lng5 := (p1.X() + p2.X() + p3.X()) / 3.
	lat5 := (p1.Y() + p2.Y() + p3.Y()) / 3.
	p5 := NewPoint(lng5, lat5, GeoCtx)
	fmt.Printf("POINT(%v %v)\n", p5.X(), p5.Y())
	fmt.Println("dis15:", calc.Distance(p1, p5)*EarthRadius)
	fmt.Println("dis25:", calc.Distance(p2, p5)*EarthRadius)
	fmt.Println("dis35:", calc.Distance(p3, p5)*EarthRadius)
}

func f(p1, p2 *Point, crs13, crs23 float64) (*Point, error) {
	lon1, lat1 := ToRadians(p1.X()), ToRadians(p1.Y())
	lon2, lat2 := ToRadians(p2.X()), ToRadians(p2.Y())

	var crs12, crs21 float64
	dst12 := 2. * Asin(Sqrt(Pow(Sin((lat1-lat2)/2.), 2.)+Cos(lat1)*Cos(lat2)*Pow(Sin((lon1-lon2)/2.), 2.)))
	if Sin(lon2-lon1) > 0 {
		crs12 = Acos((Sin(lat2) - Sin(lat1)*Cos(dst12)) / (Sin(dst12) * Cos(lat1)))
		crs21 = 2.*Pi - Acos((Sin(lat1)-Sin(lat2)*Cos(dst12))/(Sin(dst12)*Cos(lat2)))
	} else {
		crs12 = 2.*Pi - Acos((Sin(lat2)-Sin(lat1)*Cos(dst12))/(Sin(dst12)*Cos(lat1)))
		crs21 = Acos((Sin(lat1) - Sin(lat2)*Cos(dst12)) / (Sin(dst12) * Cos(lat2)))
	}
	//fmt.Println("crs12:", ToDegrees(crs12))
	//fmt.Println("crs21:", ToDegrees(crs21))

	//ang1 := Mod(crs13-crs12+Pi, 2.*Pi) - Pi
	//ang2 := Mod(crs21-crs23+Pi, 2.*Pi) - Pi
	ang1 := crs13 - crs12
	ang2 := crs21 - crs23

	if Sin(ang1) == 0 && Sin(ang2) == 0 {
		return nil, errors.New("infinity of intersections")
	} else if Sin(ang1)*Sin(ang2) < 0 {
		return nil, errors.New("intersection ambiguous")
	} else {
		//ang1 = Abs(ang1)
		//ang2 = Abs(ang2)
		ang3 := Acos(-Cos(ang1)*Cos(ang2) + Sin(ang1)*Sin(ang2)*Cos(dst12))
		dst13 := Atan2(Sin(dst12)*Sin(ang1)*Sin(ang2), Cos(ang2)+Cos(ang1)*Cos(ang3))
		lat3 := Asin(Sin(lat1)*Cos(dst13) + Cos(lat1)*Sin(dst13)*Cos(crs13))
		dlon := Atan2(Sin(crs13)*Sin(dst13)*Cos(lat1), Cos(dst13)-Sin(lat1)*Sin(lat3))
		lon3 := lon1 + dlon
		//fmt.Printf("POINT(%v %v)\n", ToDegrees(lon3), ToDegrees(lat3))
		return NewPoint(ToDegrees(lon3), ToDegrees(lat3), GeoCtx), nil
	}
}
