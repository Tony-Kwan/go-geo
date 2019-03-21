package rtree

import (
	"fmt"
	"github.com/Tony-Kwan/go-geo/geo"
	"math"
	"strings"
)

type rnode struct {
	bounds  geo.Rectangle
	parent  *rnode
	entries []Spatial
	isLeaf  bool
}

func (node *rnode) Bounds() *geo.Rectangle {
	return &node.bounds
}

func (node *rnode) NumEntries() int {
	return len(node.entries)
}

func (node *rnode) calcBounds() *geo.Rectangle {
	minX, maxX, minY, maxY := math.MaxFloat64, -math.MaxFloat64, math.MaxFloat64, -math.MaxFloat64
	var bounds *geo.Rectangle
	for _, entry := range node.entries {
		bounds = entry.Bounds()
		minX = math.Min(minX, bounds.MinX)
		maxX = math.Max(maxX, bounds.MaxX)
		minY = math.Min(minY, bounds.MinY)
		maxY = math.Max(maxY, bounds.MaxY)
	}
	return geo.NewRectangle(minX, maxX, minY, maxY, node.entries[0].Bounds().GetContext())
}

func (node *rnode) String() string {
	if node.isLeaf {
		esl := make([]string, len(node.entries))
		for i, entry := range node.entries {
			esl[i] = fmt.Sprintf("%v", entry)
		}
		return fmt.Sprintf("rnode{bounds=%s, isLeaf=%t, entries=[%s]", node.bounds.String(), node.isLeaf, strings.Join(esl, ", "))
	}
	return fmt.Sprintf("rnode{bounds=%s, isLeaf=%t}", node.bounds.String(), node.isLeaf)
}
