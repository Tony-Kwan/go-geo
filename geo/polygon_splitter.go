package geo

import (
	"fmt"
	"github.com/Tony-Kwan/go-geo/geo/internal/ds"
	"github.com/atotto/clipboard"
	"github.com/pkg/errors"
	"strings"
)

func (p *Polygon) Split(vertexLimit int) ([]Polygon, error) {
	tris, err := p.Triangulate()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	n := len(tris)
	graph := ds.New2DBoolSlice(n, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			graph[i][j] = tris[i].IsConnected(&tris[j])
		}
	}
	sg := splitterGroup{graph: graph, tris: tris, n: n, gid: make([]int, n), now: 1, cnt: make([]int, n)}
	sg.search(0, vertexLimit)

	{
		wkts := make([]string, 0)
		for i, tri := range tris {
			if sg.gid[i] == 1 {
				wkts = append(wkts, tri.String())
			}
		}
		wkt := fmt.Sprintf("GEOMETRYCOLLECTION(%s, %s)", strings.Join(wkts, ","), p.String())
		//wkt := fmt.Sprintf("GEOMETRYCOLLECTION(%s)", strings.Join(wkts, ","))
		clipboard.WriteAll(wkt)
	}

	return nil, nil
}

type splitterGroup struct {
	graph ds.TwoDimBoolSlice
	tris  []Triangle
	n     int
	gid   []int
	now   int
	cnt   []int
}

func (sg *splitterGroup) search(k, vertexLimit int) {
	for i := 0; i < sg.n; i++ {
		if sg.gid[i] == 0 && sg.graph[k][i] {
			sg.gid[i] = sg.now
			sg.cnt[sg.now]++
			if sg.cnt[sg.now] >= vertexLimit-1 {
				sg.now++
			}
			sg.search(i, vertexLimit)
		}
	}
}
