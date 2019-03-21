package rtree

import (
	"fmt"
	"github.com/Tony-Kwan/go-geo/geo"
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
		e := Entry{rect: geo.NewRectangle(minX, minX+1, minY, minY+1, geo.CartesianCtx), value: i}
		rt.Insert(&e)
	}

	t.Log(rt.String())
}
