package geo

import (
	"fmt"
	"github.com/atotto/clipboard"
	"strings"
	"testing"
)

func TestPolygon_Triangulate(t *testing.T) {
	polygon := Polygon{shell: LinearRing{
		//*NewPoint(-114.48303222656251, 38.39468429873875, nil),
		//*NewPoint(-97.49473571777345, 28.884362064951418, nil),
		//*NewPoint(-86.58977508544923, 36.15700384567333, nil),
		//*NewPoint(-76.65985107421875, 31.486650038733302, nil),
		//*NewPoint(-66.9386672973633, 41.65239288426815, nil),
		//*NewPoint(-78.3706283569336, 40.22109212030119, nil),
		//*NewPoint(-90.11192321777345, 48.20911695037711, nil),
		//*NewPoint(-90.74810028076169, 34.93013417230951, nil),
		//*NewPoint(-103.67694854736328, 39.471185234712465, nil),
		//*NewPoint(-105.26824951171875, 47.96693928840199, nil),
		//*NewPoint(-114.48303222656251, 38.39468429873875, nil),

		*NewPoint(-126.04545593261719, 43.99404946578386, nil),
		*NewPoint(-126.13334655761719, 31.03969759558136, nil),
		*NewPoint(-110.92826843261719, 27.279958882565637, nil),
		*NewPoint(-96.95365905761717, 31.715025580423827, nil),
		*NewPoint(-90.88920593261717, 42.34712656235513, nil),
		*NewPoint(-89.8956298828125, 47.01654600163033, nil),
		*NewPoint(-100.32783508300781, 48.53411250954363, nil),
		*NewPoint(-109.2047882080078, 49.055195215685586, nil),
		*NewPoint(-119.61021423339844, 47.77186813430575, nil),
		*NewPoint(-126.04545593261719, 43.99404946578386, nil),
	}}
	fmt.Println(&polygon)
	tris, err := polygon.Triangulate()
	if err != nil {
		t.Error(err)
		return
	}
	wkts := make([]string, len(tris))
	for i, tri := range tris {
		wkts[i] = tri.String()
	}
	wkt := fmt.Sprintf("GEOMETRYCOLLECTION(%s)", strings.Join(wkts, ", "))
	//t.Log(wkt)
	clipboard.WriteAll(wkt)
}

func TestPolygon_ConvexHull(t *testing.T) {
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
	t.Log(polygon.String())
	hull, _ := polygon.ConvexHull()
	t.Log(hull.String())

	wkt := fmt.Sprintf("GEOMETRYCOLLECTION(%s,%s)", polygon.String(), hull.String())
	t.Log(wkt)
}
