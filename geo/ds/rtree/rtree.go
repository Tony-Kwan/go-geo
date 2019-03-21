package rtree

import (
	"bytes"
	"fmt"
	"github.com/Tony-Kwan/go-geo/geo"
	"io"
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
	l := r.chooseLeaf(r.root, entry).(*rnode)
	l.entries = append(l.entries, entry)
	var ll *rnode
	if l.NumEntries() > r.maxEntries {
		l, ll = r.splitNode(l)
	} else {
		l.bounds = *l.calcBounds()
	}
	r.adjustTree(l, ll)
	r.numEntries++
	return r
}

func (r *Rtree) chooseLeaf(node Spatial, entry Spatial) Spatial {
	rnodePtr, ok := node.(*rnode)
	if !ok {
		panic("program should not run here")
	}
	if rnodePtr.isLeaf {
		return node
	}

	var chosenIdx int
	var area, enlargement float64
	var minArea = math.MaxFloat64
	var minEnlargement = math.MaxFloat64
	entryBounds := entry.Bounds()
	for i, childNode := range rnodePtr.entries {
		area = childNode.Bounds().GetArea()
		unionArea := childNode.Bounds().Union(entryBounds).GetArea()
		enlargement = unionArea - area
		if enlargement < minEnlargement || (enlargement == minEnlargement && area < minArea) {
			minEnlargement = enlargement
			minArea = area
			chosenIdx = i
		}
	}
	return r.chooseLeaf(rnodePtr.entries[chosenIdx], entry)
}

func (r *Rtree) splitNode(node *rnode) (l, ll *rnode) {
	ll = &rnode{parent: node.parent, entries: node.entries[r.maxEntries:], isLeaf: node.isLeaf}
	l = node
	l.entries = node.entries[:r.maxEntries]
	l.bounds, ll.bounds = *l.calcBounds(), *ll.calcBounds()
	if !ll.isLeaf {
		var rnodePtr *rnode
		for i := range ll.entries {
			rnodePtr = ll.entries[i].(*rnode)
			rnodePtr.parent = ll
		}
	}
	return
}

func (r *Rtree) adjustTree(n, nn *rnode) {
	p := n.parent
	if p == nil {
		if nn != nil {
			newRoot := &rnode{
				bounds:  *n.bounds.Union(&nn.bounds),
				entries: []Spatial{n, nn},
				isLeaf:  false,
			}
			r.root = newRoot
			n.parent, nn.parent = r.root, r.root
		}
		return
	}

	var pp *rnode
	if nn != nil {
		p.entries = append(p.entries, nn)
		if p.NumEntries() > r.maxEntries {
			p, pp = r.splitNode(p)
		} else {
			p.bounds = *p.calcBounds()
		}
	}
	r.adjustTree(p, pp)
}

func (r *Rtree) NumEntries() int { return r.numEntries }

func (r *Rtree) travel(node Spatial, deep int, w io.Writer) error {
	rnodePtr, ok := node.(*rnode)
	if !ok {
		return nil
	}

	indent := strings.Repeat("\t", deep)
	if _, err := w.Write([]byte(indent + rnodePtr.String() + "\n")); err != nil {
		return err
	}

	for _, entry := range rnodePtr.entries {
		if err := r.travel(entry, deep+1, w); err != nil {
			return err
		}
	}
	return nil
}

func (r *Rtree) String() string {
	var buf bytes.Buffer
	if err := r.travel(r.root, 1, &buf); err != nil {
		return fmt.Sprintf("error in travel tree: %v", err)
	}
	return fmt.Sprintf("Rtree{\n%s}", buf.String())
}
