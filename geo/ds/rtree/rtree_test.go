package rtree

import (
	"fmt"
	"github.com/Tony-Kwan/go-geo/geo"
	"github.com/fogleman/gg"
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

func (e *Entry) String() string {
	return fmt.Sprintf("%d", e.value)
	//return e.rect.String()
}

func TestMain(m *testing.M) {
	rt, _ = NewRtree(1, 2)
	m.Run()
}

func TestRtree_Insert(t *testing.T) {
	for i := 0; i < 5; i++ {
		minX, minY := float64(i), float64(i)
		e := Entry{rect: geo.NewRectangle(minX*100, (minX+1)*100, minY*100, (minY+1)*100, geo.CartesianCtx), value: i}
		rt.Insert(&e)
	}

	t.Log(rt.String())
	drawRtree()
}

func drawRtree() {
	bounds := rt.root.bounds

	dc := gg.NewContext(int(math.Ceil(bounds.GetWidth())), int(math.Ceil(bounds.GetHeight())))
	dc.SetLineWidth(1)
	dc.SetLineCap(gg.LineCapSquare)
	dc.DrawRectangle(0, 0, bounds.GetWidth(), bounds.GetHeight())
	dc.SetRGB(0, 0, 0)
	dc.Stroke()
	dc.SavePNG("/Users/zhongxian.guan/tmp/rtree.png")
}
