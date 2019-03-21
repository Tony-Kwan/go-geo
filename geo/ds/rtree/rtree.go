package rtree

import (
	"bytes"
	"fmt"
	"github.com/Tony-Kwan/go-geo/geo"
	"math"
	"strings"
)

type Spatial interface {
	Bounds() *geo.Rectangle
}

type Rtree struct {
	root       *rnode
	minEntries int
	maxEntries int
	numEntries int
}

func NewRtree(m, M int) (*Rtree, error) {
	//TODO: verify m and M
	return &Rtree{
		root:       &rnode{bounds: *geo.NewRectangle(0, 0, 0, 0, geo.CartesianCtx), isLeaf: true},
		minEntries: m,
		maxEntries: M,
	}, nil
}

func (r *Rtree) Insert(entry Spatial) *Rtree {
	// I1
	l := r.chooseLeaf(r.root, entry).(*rnode)

	// I2
	l.entries = append(l.entries, entry)
	var ll *rnode
	if l.NumEntries() > r.maxEntries {
		l, ll = r.splitNode(l)
	}

	// I3
	root, splitRoot := r.adjustTree(l, ll)

	// I4
	if splitRoot != nil {
		newRoot := &rnode{
			bounds:  *root.bounds.Union(&splitRoot.bounds),
			entries: []Spatial{root, splitRoot},
			isLeaf:  false,
		}
		root.parent, splitRoot.parent = newRoot, newRoot
		r.root = newRoot
	}
	r.numEntries++
	return r
}

func (r *Rtree) chooseLeaf(node Spatial, entry Spatial) Spatial {
	n, ok := node.(*rnode)
	if !ok {
		panic("program should not run here")
	}
	// CL2
	if n.isLeaf {
		return node
	}

	// CL3
	var chosenIdx int
	var area, enlargement float64
	var minArea = math.MaxFloat64
	var minEnlargement = math.MaxFloat64
	entryBounds := entry.Bounds()
	for i, f := range n.entries {
		area = f.Bounds().GetArea()
		unionArea := f.Bounds().Union(entryBounds).GetArea()
		enlargement = unionArea - area
		if enlargement < minEnlargement || (enlargement == minEnlargement && area < minArea) {
			minEnlargement = enlargement
			minArea = area
			chosenIdx = i
		}
	}

	// CL4
	return r.chooseLeaf(n.entries[chosenIdx], entry)
}

func (r *Rtree) splitNode(node *rnode) (*rnode, *rnode) {
	l := node
	ll := &rnode{parent: node.parent, entries: node.entries[r.maxEntries:node.NumEntries()], isLeaf: node.isLeaf}
	l.entries = node.entries[:r.maxEntries]
	return l, ll
}

func (r *Rtree) adjustTree(n, nn *rnode) (*rnode, *rnode) {
	n.bounds = *n.calcBounds()
	if nn != nil {
		nn.bounds = *nn.calcBounds()
	}

	// AT2
	if n == r.root {
		return n, nn
	}

	// AT3
	p := n.parent

	// AT4
	if nn != nil {
		nn.bounds = *nn.calcBounds()
		p.entries = append(p.entries, nn)
		if p.NumEntries() > r.maxEntries {
			var pp *rnode
			p, pp = r.adjustTree(r.splitNode(p))
			for i := range pp.entries {
				pp.entries[i].(*rnode).parent = pp
			}
			return p, pp
		}
	}

	// AT5
	return r.adjustTree(p, nil)
}

func (r *Rtree) NumEntries() int { return r.numEntries }

func (r *Rtree) travel(node Spatial, deep int, f func(node Spatial, deep int) error) error {
	f(node, deep)

	rnodePtr, ok := node.(*rnode)
	if !ok {
		return nil
	}

	for _, entry := range rnodePtr.entries {
		if err := r.travel(entry, deep+1, f); err != nil {
			return err
		}
	}
	return nil
}

func (r *Rtree) String() string {
	var buf bytes.Buffer
	f := func(node Spatial, deep int) error {
		rnodePtr, ok := node.(*rnode)
		if !ok {
			return nil
		}
		indent := strings.Repeat("\t", deep)
		if _, err := buf.Write([]byte(indent + rnodePtr.String() + "\n")); err != nil {
			return err
		}
		return nil
	}
	if err := r.travel(r.root, 1, f); err != nil {
		return fmt.Sprintf("error in travel tree: %v", err)
	}
	return fmt.Sprintf("Rtree{\n%s}", buf.String())
}
