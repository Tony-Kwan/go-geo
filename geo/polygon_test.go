package geo

import (
	"fmt"
	"github.com/atotto/clipboard"
	"strings"
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

		//*NewPoint(-126.04545593261719, 43.99404946578386, nil),
		//*NewPoint(-126.13334655761719, 31.03969759558136, nil),
		//*NewPoint(-110.92826843261719, 27.279958882565637, nil),
		//*NewPoint(-96.95365905761717, 31.715025580423827, nil),
		//*NewPoint(-90.88920593261717, 42.34712656235513, nil),
		//*NewPoint(-89.8956298828125, 47.01654600163033, nil),
		//*NewPoint(-100.32783508300781, 48.53411250954363, nil),
		//*NewPoint(-109.2047882080078, 49.055195215685586, nil),
		//*NewPoint(-119.61021423339844, 47.77186813430575, nil),
		//*NewPoint(-126.04545593261719, 43.99404946578386, nil),

		//*NewPoint(-113.70197296142578, 44.86438632934659, nil),
		//*NewPoint(-130.13751983642578, 29.455143906645745, nil),
		//*NewPoint(-75.5574417114258, 34.59365010878254, nil),
		//*NewPoint(-114.75666046142578, 36.1700311594995, nil),
		//*NewPoint(-91.90509796142578, 40.777421721005965, nil),
		//*NewPoint(-112.47528076171875, 40.242846824049764, nil),
		//*NewPoint(-99.29168701171876, 45.086126831109596, nil),
		//*NewPoint(-113.35418701171875, 42.48526384858917, nil),
		//*NewPoint(-107.81330108642578, 47.21676985912018, nil),
		//*NewPoint(-113.70197296142578, 44.86438632934659, nil),

		//*NewPoint(-92.38059997558592, 45.38157243512828, nil),
		//*NewPoint(-90.47378540039062, 40.582670638095294, nil),
		//*NewPoint(-81.50894165039061, 39.77397788285171, nil),
		//*NewPoint(-87.19505310058592, 37.1594957106433, nil),
		//*NewPoint(-86.93138122558594, 32.01972036197235, nil),
		//*NewPoint(-91.06224060058594, 35.909908145897035, nil),
		//*NewPoint(-95.95733642578125, 31.974007590177635, nil),
		//*NewPoint(-94.81475830078125, 37.60580020781012, nil),
		//*NewPoint(-102.54913330078125, 38.29882852868994, nil),
		//*NewPoint(-93.93585205078125, 40.27140563877154, nil),
		//*NewPoint(-92.38059997558592, 45.38157243512828, nil),
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
