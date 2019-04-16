package rtree

import (
	"fmt"
	"github.com/Tony-Kwan/go-geo/geo"
	"github.com/fogleman/gg"
	"golang.org/x/image/colornames"
	"image/color"
	"math"
	"math/rand"
	"testing"
	"time"
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
	//return fmt.Sprintf("%s %d", e.rect.String(), e.value)
}

func TestMain(m *testing.M) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	rt, _ = NewRtree(2, 4, &UglySplitter{})
	m.Run()
}

func TestRtree_Insert(t *testing.T) {
	rand.Seed(time.Now().Unix())
	n := 10
	for i := 0; i < 10; i++ {
		r := rand.Int() % n
		minX, minY := float64(r), float64(r)
		e := Entry{rect: geo.NewRectangle(minX, minY, minX+1, minY+1, geo.CartesianCtx), value: r}
		rt.Insert(&e)
	}
	t.Log(rt.String())
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

	colors := []color.Color{colornames.Black, colornames.Blue, colornames.Green, colornames.Red, colornames.Yellow, colornames.Cyan}
	f := func(node Spatial, deep int) error {
		b := node.Bounds()
		dc.SetColor(colors[deep%len(colors)])
		dc.DrawRectangle((b.GetMinX()*100)+float64(deep), (b.GetMinY()*100)+float64(deep), (b.GetWidth()*100)-2*float64(deep), (b.GetHeight()*100)-2*float64(deep))
		dc.Stroke()
		return nil
	}
	rt.travel(rt.root, 0, f)

	if err := dc.SavePNG("rtree.png"); err != nil {
		fmt.Println(err)
	}
}
