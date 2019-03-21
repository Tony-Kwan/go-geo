package rtree

import (
	"fmt"
	"github.com/Tony-Kwan/go-geo/geo"
	"github.com/fogleman/gg"
	"golang.org/x/image/colornames"
	"image/color"
	"math"
	"testing"
)

var rt *Rtree

type Entry struct {
	rect  *geo.Rectangle
	value int
}

func (e *Entry) Bounds() *geo.Rectangle {
	return e.rect
}

//func (e *Entry) String() string {
//	return fmt.Sprintf("%d", e.value)
//	//return e.rect.String()
//}

func TestMain(m *testing.M) {
	rt, _ = NewRtree(1, 2)
	m.Run()
}

func TestRtree_Insert(t *testing.T) {
	ts := []int{2, 1, 0, 3}
	for _, i := range ts {
		minX, minY := float64(i), float64(i)
		e := Entry{rect: geo.NewRectangle(minX, minX+1, minY, minY+1, geo.CartesianCtx), value: i}
		rt.Insert(&e)
		t.Log(rt.String())
	}

	drawRtree()
}

func drawRtree() {
	bounds := rt.root.bounds

	dc := gg.NewContext(int(math.Ceil(bounds.GetWidth()))*100, int(math.Ceil(bounds.GetHeight()))*100)
	dc.SetLineWidth(0.5)
	dc.SetLineCap(gg.LineCapRound)
	dc.SetColor(colornames.Lightgray)
	dc.DrawRectangle(0, 0, bounds.GetWidth()*100, bounds.GetHeight()*100)
	dc.Fill()

	colors := []color.Color{colornames.Black, colornames.Blue, colornames.Green, colornames.Red, colornames.Yellow}
	f := func(node Spatial, deep int) error {
		b := node.Bounds()
		dc.SetColor(colors[deep%len(colors)])
		dc.DrawRectangle((b.MinX*100)+float64(deep), (b.MinY*100)+float64(deep), (b.GetWidth()*100)-2*float64(deep), (b.GetHeight()*100)-2*float64(deep))
		dc.Stroke()
		return nil
	}
	rt.travel(rt.root, 0, f)

	if err := dc.SavePNG("rtree.png"); err != nil {
		fmt.Println(err)
	}
}
