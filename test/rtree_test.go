package test

import (
	"encoding/csv"
	"fmt"
	"github.com/Tony-Kwan/go-geo/geo"
	"github.com/Tony-Kwan/go-geo/geo/ds/rtree"
	"github.com/Tony-Kwan/go-geo/geo/io/wkt"
	"io"
	"os"
	"strconv"
	"testing"
)

type Entry struct {
	rect  *geo.Rectangle
	value int
}

func (e *Entry) Bounds() *geo.Rectangle {
	return e.rect
}

func (e *Entry) String() string {
	return fmt.Sprintf("%d", e.value)
}

func TestRtree(t *testing.T) {
	f, err := os.Open("rtree_test.csv")
	if err != nil {
		t.Error(err)
		return
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	wktReader := wkt.NewReader()
	rt, _ := rtree.NewRtree(20, 40)
	for {
		row, err := csvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			t.Error(err)
			return
		}
		e := &Entry{}
		e.value, err = strconv.Atoi(row[0])
		if err != nil {
			t.Error(err)
			return
		}
		shape, err := wktReader.Read(row[1])
		if err != nil {
			t.Error(err)
			return
		}
		polygon := shape.(*geo.Polygon)
		e.rect = polygon.Bounds()
		e.rect.SetContext(geo.CartesianCtx)
		rt.Insert(e)
	}
	t.Log("Bounds:", rt.Bounds())
	t.Log("NumEntries:", rt.NumEntries())

	maxDeep := 0
	_ = rt.Travel(func(node rtree.Spatial, deep int) error {
		if deep > maxDeep {
			maxDeep = deep
		}
		return nil
	})
	t.Log("Deep:", maxDeep)
}
