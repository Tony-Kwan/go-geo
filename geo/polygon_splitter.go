package geo

import (
	"github.com/Tony-Kwan/go-geo/geo/internal/ds"
	"github.com/pkg/errors"
	"sort"
)

func (p *Polygon) Split(vertexLimit int) ([]Polygon, error) {
	order := make(map[uint64]int)
	vertexCnt := len(p.Shell) - 1
	for i := 0; i < vertexCnt; i++ {
		order[p.Shell[i].pointHash()] = i
	}

	tris, err := p.Triangulate()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	n := len(tris)
	graph := ds.New2DBoolSlice(n, n)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			graph[i][j] = tris[i].IsConnected(&tris[j])
			graph[j][i] = graph[i][j]
		}
	}
	sg := splitterGroup{graph: graph, tris: tris, n: n, gid: make([]int, n), now: 1, cnt: make(map[int]int), father: make([]int, n), vertexLimit: vertexLimit}
	sg.search(0)

	m := make(map[int]map[uint64]Point)
	for i, tri := range tris {
		if _, exists := m[sg.gid[i]]; !exists {
			m[sg.gid[i]] = make(map[uint64]Point)
		}
		m[sg.gid[i]][tri.A.pointHash()] = tri.A
		m[sg.gid[i]][tri.B.pointHash()] = tri.B
		m[sg.gid[i]][tri.C.pointHash()] = tri.C
	}
	ps := make([]Polygon, 0)
	for _, v := range m {
		ops, i := make(ds.OrderObjs, len(v)), 0
		for _, p := range v {
			ops[i], i = ds.OrderObj{Obj: p, Order: order[p.pointHash()]}, i+1
		}
		sort.Sort(ops)
		shell := make(LinearRing, len(v)+1)
		for i, op := range ops {
			shell[i] = op.Obj.(Point)
		}
		shell[len(v)] = shell[0]
		ps = append(ps, NewPolygon(shell))
	}
	return ps, nil
}

type splitterGroup struct {
	graph       ds.TwoDimBoolSlice
	tris        []Triangle
	n           int
	gid         []int
	now         int
	cnt         map[int]int
	father      []int
	vertexLimit int
}

func (sg *splitterGroup) search(k int) {
	sg.dfs(k)
	//sg.bfs(k)
}

func (sg *splitterGroup) dfs(k int) {
	for i := 0; i < sg.n; i++ {
		if sg.gid[i] == 0 && sg.graph[k][i] {
			if sg.cnt[sg.gid[k]] >= sg.vertexLimit-3 {
				sg.now++
				sg.gid[i] = sg.now
			} else {
				sg.gid[i] = sg.gid[k]
			}
			sg.cnt[sg.gid[i]]++
			sg.search(i)
		}
	}
}

func (sg *splitterGroup) bfs(k int) {
	sg.father[0] = -1
	q, pos := []int{0}, 0
	for pos < len(q) {
		k, pos = q[pos], pos+1
		if k == 0 {
			sg.gid[k] = sg.now
		} else {
			if sg.cnt[sg.gid[sg.father[k]]] >= sg.vertexLimit-3 {
				sg.now++
				sg.gid[k] = sg.now
			} else {
				sg.gid[k] = sg.gid[sg.father[k]]
			}
		}
		sg.cnt[sg.gid[k]]++
		for i := 0; i < sg.n; i++ {
			if sg.gid[i] == 0 && sg.graph[k][i] {
				q = append(q, i)
				sg.father[i] = k
				sg.gid[i] = -1
			}
		}
	}
}
