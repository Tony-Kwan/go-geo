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
var calc = geo.VectorCalculator{}

func TestPolygon_Triangulate(t *testing.T) {
	var wktStr = "POLYGON((-119.14159855412011 41.92499096272812,-102.87886464247828 26.767453854346442,-90.9950327117614 32.905865438530256,-93.78787570746942 36.051610613180415,-98.38245074916146 33.171901863337226,-102.05698491988139 36.07009780383788,-96.33267709169169 38.664513680653954,-88.4319360597416 37.10206976053486,-87.28074151678913 30.775914227484662,-87.59809404666174 26.844389706437113,-79.64598553204937 29.904465341243238,-80.44042235294994 37.59795124939845,-84.81650967729065 43.819424657392034,-96.2837724335406 40.82461783417634,-98.73181999767975 46.03621512184026,-103.30774793950121 42.50917847188873,-108.37624078356396 46.00372124466807,-119.14159855412011 41.92499096272812))"
	var polygon = wkt2.MustPolygon(wktReader.Read(wktStr))
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
	var wktStr = "POLYGON((-119.14159855412011 41.92499096272812,-102.87886464247828 26.767453854346442,-90.9950327117614 32.905865438530256,-93.78787570746942 36.051610613180415,-98.38245074916146 33.171901863337226,-102.05698491988139 36.07009780383788,-96.33267709169169 38.664513680653954,-88.4319360597416 37.10206976053486,-87.28074151678913 30.775914227484662,-87.59809404666174 26.844389706437113,-79.64598553204937 29.904465341243238,-80.44042235294994 37.59795124939845,-84.81650967729065 43.819424657392034,-96.2837724335406 40.82461783417634,-98.73181999767975 46.03621512184026,-103.30774793950121 42.50917847188873,-108.37624078356396 46.00372124466807,-119.14159855412011 41.92499096272812))"
	var polygon = wkt2.MustPolygon(wktReader.Read(wktStr))
	hull, _ := polygon.ConvexHull()
	t.Log(hull.String())

	wkt := fmt.Sprintf("GEOMETRYCOLLECTION(%s,%s)", polygon.String(), hull.String())
	t.Log(wkt)
	clipboard.WriteAll(wkt)
}

func TestPolygon_Split(t *testing.T) {
	//var wktStr = "POLYGON((-125.5413513456855 48.918235435184954,-125.53242428182067 48.91078058459627,-125.51220907541939 48.92262611322883,-125.5036752527155 48.91671817587559,-125.52484990228118 48.903322476385796,-125.51609743096496 48.89653048983544,-125.5084496246527 48.90121911868874,-125.49573210760602 48.894251985666045,-125.48920038824009 48.8987413119778,-125.50176289321091 48.90518442646572,-125.49520506654801 48.90940037154863,-125.45722384454442 48.891888625154735,-125.44534828782567 48.90030631327758,-125.49946055595753 48.92932021612094,-125.46315020096135 48.952403911273876,-125.47667704421477 48.95696515535485,-125.48256423968623 48.951052708636325,-125.51450325414248 48.96976269837617,-125.52205968491528 48.96395246809763,-125.49369084332254 48.94848092923701,-125.50890650235004 48.93822249705744,-125.53998071259348 48.958768322866774,-125.54472081870341 48.95476333836433,-125.51928741631608 48.92973508374098,-125.5413513456855 48.918235435184954))"
	var wktStr = "POLYGON((-125.54215148091313 48.930035701162296,-125.53744688630101 48.926151753870045,-125.53423628211017 48.92848762135765,-125.53126171231268 48.92569619851514,-125.52080780267713 48.931483292976594,-125.48642724752423 48.90817046433605,-125.51268473267552 48.891247001395584,-125.50677850842473 48.887622822954256,-125.50316959619519 48.88983968738856,-125.49983024597165 48.887863561199964,-125.43472766876218 48.92681437242206,-125.44480741024014 48.93277930578961,-125.44831708073615 48.93043306621078,-125.45301765203473 48.93224539839986,-125.47871053218837 48.91527104004945,-125.51022246479985 48.93480034904374,-125.46667814254756 48.95849877085138,-125.47594919800757 48.96363402932619,-125.47962382435796 48.96207466431636,-125.48432573676106 48.96457702087133,-125.50467833876607 48.95422691075413,-125.52337467670436 48.96769994771978,-125.53015127778052 48.96232384901117,-125.5145555734634 48.94981288750151,-125.5323453247547 48.93684157253844,-125.54731205105779 48.9492685934228,-125.5515298247337 48.945825663708376,-125.536997616291 48.9339669208932,-125.54215148091313 48.930035701162296))"
	var polygon = wkt2.MustPolygon(wktReader.Read(wktStr))
	ps, err := polygon.Split()
	if err != nil {
		t.Error(err)
		return
	}
	wkts := make([]string, len(ps))
	for i, p := range ps {
		wkts[i] = p.String()
	}
	//wkt := fmt.Sprintf("GEOMETRYCOLLECTION(%s)", strings.Join(wkts, ", "))
	//t.Log(wkt)
	//clipboard.WriteAll(wkt)
}
