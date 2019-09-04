package geo

import (
	"fmt"
	"github.com/Tony-Kwan/go-geo/geo/internal/ds"
	"github.com/pkg/errors"
)

func (p *Polygon) Split(vertexLimit int) ([]Polygon, error) {
	if vertexLimit < 4 {
		return nil, fmt.Errorf("wrong vertexLimit: expect>3, found=%d", vertexLimit)
	}
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
			graph[i][j] = tris[i].IsConnected(tris[j])
			graph[j][i] = graph[i][j]
		}
	}
	sg := newSplitterGroup(graph, tris, n, vertexLimit)
	sg.search(0)
	if len(sg.errs) != 0 {
		return nil, sg.errs[len(sg.errs)-1]
	}

	ps := make([]Polygon, 0, sg.now)
	for _, list := range sg.polygonMap {
		shell := make(LinearRing, 0, list.Size()+1)
		it := list.Iterator()
		for it.Next() {
			shell = append(shell, it.Value().(Point))
		}
		shell = append(shell, shell[0])
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

	polygonMap map[int]*ds.CircularLinkedList
	errs       []error
}

func newSplitterGroup(graph ds.TwoDimBoolSlice, tris []Triangle, n, vertexLimit int) *splitterGroup {
	return &splitterGroup{
		graph:       graph,
		tris:        tris,
		n:           n,
		gid:         make([]int, n),
		now:         1,
		cnt:         make(map[int]int),
		father:      make([]int, n),
		vertexLimit: vertexLimit,
		polygonMap:  make(map[int]*ds.CircularLinkedList),
	}
}

func (sg *splitterGroup) search(k int) {
	sg.gid[k] = 1
	sg.cnt[1] = 1

	list := ds.NewCircularLinkedList()
	list.Add(sg.tris[k].A)
	list.Add(sg.tris[k].B)
	list.Add(sg.tris[k].C)
	sg.polygonMap[1] = list

	sg.dfs(k)
}

func (sg *splitterGroup) dfs(k int) {
	for i := 0; i < sg.n; i++ {
		if sg.gid[i] != 0 || !sg.graph[k][i] {
			continue
		}

		var notConnectPoint *Point
		if !sg.tris[k].isVertex(sg.tris[i].A) {
			notConnectPoint = &sg.tris[i].A
		} else if !sg.tris[k].isVertex(sg.tris[i].B) {
			notConnectPoint = &sg.tris[i].B
		} else if !sg.tris[k].isVertex(sg.tris[i].C) {
			notConnectPoint = &sg.tris[i].C
		}
		flag := false
		if notConnectPoint != nil {
			for l := 0; l < sg.n; l++ {
				if sg.gid[l] == sg.now && sg.tris[l].isVertex(*notConnectPoint) {
					flag = true
					break
				}
			}
		}
		if sg.cnt[sg.gid[k]] >= sg.vertexLimit-3 || flag {
			sg.now++
			sg.gid[i] = sg.now

			list := ds.NewCircularLinkedList()
			list.Add(sg.tris[i].A)
			list.Add(sg.tris[i].B)
			list.Add(sg.tris[i].C)
			sg.polygonMap[sg.now] = list
		} else {
			sg.gid[i] = sg.gid[k]

			tri := sg.tris[i]
			list := sg.polygonMap[sg.gid[i]]
			it := list.Iterator()
			flag = false
			for it.Next() {
				node := it.Node()
				curr := node.Elem.(Point)
				next := node.Next.Elem.(Point)
				if (curr.Equals(tri.A) && next.Equals(tri.B)) || (curr.Equals(tri.B) && next.Equals(tri.A)) {
					list.Insert(it.Node(), tri.C)
					flag = true
					break
				} else if (curr.Equals(tri.B) && next.Equals(tri.C)) || (curr.Equals(tri.C) && next.Equals(tri.B)) {
					list.Insert(it.Node(), tri.A)
					flag = true
					break
				} else if (curr.Equals(tri.A) && next.Equals(tri.C)) || (curr.Equals(tri.C) && next.Equals(tri.A)) {
					list.Insert(it.Node(), tri.B)
					flag = true
					break
				}
			}
			if !flag {
				sg.errs = append(sg.errs, errors.New("can not insert point to list"))
			}
		}

		sg.cnt[sg.gid[i]]++
		sg.dfs(i)
	}
}
