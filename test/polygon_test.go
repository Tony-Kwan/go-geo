package test

import (
	"fmt"
	"github.com/Tony-Kwan/go-geo/geo"
	wkt2 "github.com/Tony-Kwan/go-geo/geo/io/wkt"
	"github.com/atotto/clipboard"
	"strings"
	"testing"
)

var wktReader wkt2.WktReader
var wktStr = "POLYGON((-119.14159855412011 41.92499096272812,-102.87886464247828 26.767453854346442,-90.9950327117614 32.905865438530256,-93.78787570746942 36.051610613180415,-98.38245074916146 33.171901863337226,-102.05698491988139 36.07009780383788,-96.33267709169169 38.664513680653954,-88.4319360597416 37.10206976053486,-87.28074151678913 30.775914227484662,-87.59809404666174 26.844389706437113,-79.64598553204937 29.904465341243238,-80.44042235294994 37.59795124939845,-84.81650967729065 43.819424657392034,-96.2837724335406 40.82461783417634,-98.73181999767975 46.03621512184026,-103.30774793950121 42.50917847188873,-108.37624078356396 46.00372124466807,-119.14159855412011 41.92499096272812))"
var polygon = wkt2.MustPolygon(wktReader.Read(wktStr))

func TestPolygon_Triangulate(t *testing.T) {
	var polygo *geo.Polygon
	fmt.Println(polygo)
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
	t.Log(wkt)
	clipboard.WriteAll(wkt)
}

func TestPolygon_ConvexHull(t *testing.T) {
	hull, _ := polygon.ConvexHull()
	t.Log(hull.String())

	wkt := fmt.Sprintf("GEOMETRYCOLLECTION(%s,%s)", polygon.String(), hull.String())
	t.Log(wkt)
	clipboard.WriteAll(wkt)
}
