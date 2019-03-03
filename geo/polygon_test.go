package geo

import (
	"testing"
)

func TestPolygon_Triangulate(t *testing.T) {
	polygon := Polygon{shell: LinearRing{
		*NewPoint(-114.48303222656251, 38.39468429873875, nil),
		*NewPoint(-97.49473571777345, 28.884362064951418, nil),
		*NewPoint(-86.58977508544923, 36.15700384567333, nil),
		*NewPoint(-76.65985107421875, 31.486650038733302, nil),
		*NewPoint(-66.9386672973633, 41.65239288426815, nil),
		*NewPoint(-78.3706283569336, 40.22109212030119, nil),
		*NewPoint(-90.11192321777345, 48.20911695037711, nil),
		*NewPoint(-90.74810028076169, 34.93013417230951, nil),
		*NewPoint(-103.67694854736328, 39.471185234712465, nil),
		*NewPoint(-105.26824951171875, 47.96693928840199, nil),
		*NewPoint(-114.48303222656251, 38.39468429873875, nil),
	}}
	//fmt.Println(&polygon)
	tris, err := polygon.Triangulate()
	if err != nil {
		t.Error(err)
		return
	}
	for _, tri := range tris {
		t.Log(tri)
	}
}
