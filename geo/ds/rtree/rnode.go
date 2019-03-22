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
	minX, minY, maxX, maxY := math.MaxFloat64, math.MaxFloat64, -math.MaxFloat64, -math.MaxFloat64
	var bounds *geo.Rectangle
	for _, entry := range node.entries {
		bounds = entry.Bounds()
		minX = math.Min(minX, bounds.GetMinX())
		minY = math.Min(minY, bounds.GetMinY())
		maxX = math.Max(maxX, bounds.GetMaxX())
		maxY = math.Max(maxY, bounds.GetMaxY())
	}
	return geo.NewRectangle(minX, minY, maxX, maxY, node.entries[0].Bounds().GetContext())
}

func (node *rnode) String() string {
	if node.isLeaf {
		esl := make([]string, len(node.entries))
		for i := range node.entries {
			esl[i] = fmt.Sprintf("%v", node.entries[i])
		}
		return fmt.Sprintf("rnode{bounds=%s, isLeaf=%t, entries=[%s]", node.bounds.String(), node.isLeaf, strings.Join(esl, ", "))
	}
	return fmt.Sprintf("rnode{bounds=%s, isLeaf=%t}", node.bounds.String(), node.isLeaf)
}
