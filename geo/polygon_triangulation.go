package geo

import (
	"fmt"
	"github.com/Tony-Kwan/go-geo/geo/internal/ds"
	"github.com/deckarep/golang-set"
)

type earClippingPoint struct {
	Point
	isReflex bool
	isEar    bool

	tmp int
}

// O(n²)
func (p *Polygon) Triangulate() ([]Triangle, error) {
	//TODO: validate polygon
	var shell = p.shell.MakeCCW()[:p.shell.GetNumPoints()-1]
	n := shell.GetNumPoints()
	points := ds.NewCircularLinkedList()
	for i, p := range shell {
		points.Add(&earClippingPoint{Point: p, tmp: i})
	}
	// check convex/reflex
	node := points.Head
	for i := 0; i < points.Size(); i, node = i+1, node.Next {
		curr := node.Elem.(*earClippingPoint)
		curr.isReflex = checkIsReflex(node)
	}
	// find ear
	node = points.Head
	for i := 0; i < points.Size(); i, node = i+1, node.Next {
		curr := node.Elem.(*earClippingPoint)
		if curr.isReflex {
			continue
		}
		curr.isEar = checkIsEar(node, points)
	}
	//printVertexType(points)

	//clip ear
	ret := make([]Triangle, 0, n-2)
	for node := points.Head; ; {
		curr := node.Elem.(*earClippingPoint)
		prev := node.Prev.Elem.(*earClippingPoint)
		next := node.Next.Elem.(*earClippingPoint)
		if points.Size() < 3 {
			//ret = append(ret, Triangle{A: &prev.Point, B: &curr.Point, C: &next.Point})
			break
		}
		if !curr.isEar {
			node = node.Next
			continue
		}
		ret = append(ret, Triangle{A: &prev.Point, B: &curr.Point, C: &next.Point})
		node = points.RemoveNode(node)
		next.isReflex = checkIsReflex(node)
		prev.isReflex = checkIsReflex(node.Prev)
		next.isEar = checkIsEar(node, points)
		prev.isEar = checkIsEar(node.Prev, points)
		//printVertexType(points)
	}
	return ret, nil
}

func checkIsReflex(node *ds.Node) bool {
	curr := node.Elem.(*earClippingPoint)
	prev := node.Prev.Elem.(*earClippingPoint)
	next := node.Next.Elem.(*earClippingPoint)
	return curr.cross(&prev.Point, &next.Point) > 0.0
}

func checkIsEar(node *ds.Node, list *ds.CircularLinkedList) bool {
	curr := node.Elem.(*earClippingPoint)
	if curr.isReflex {
		return false
	}
	prev := node.Prev.Elem.(*earClippingPoint)
	next := node.Next.Elem.(*earClippingPoint)
	tri := Triangle{
		A: &prev.Point,
		B: &curr.Point,
		C: &next.Point,
	}
	it := list.Iterator()
	for it.Next() {
		e := it.Value().(*earClippingPoint)
		if e.isReflex {
			if e != prev && e != next && tri.contains(&e.Point) {
				return false
			}
		}
	}
	return true
}

func printVertexType(points *ds.CircularLinkedList) {
	it := points.Iterator()
	convexSet := mapset.NewSet()
	reflexSet := mapset.NewSet()
	earSet := mapset.NewSet()
	for it.Next() {
		e := it.Value().(*earClippingPoint)
		if e.isReflex {
			reflexSet.Add(e.tmp)
		} else {
			convexSet.Add(e.tmp)
		}
		if e.isEar {
			earSet.Add(e.tmp)
		}
	}
	fmt.Println("convex:", convexSet)
	fmt.Println("reflex:", reflexSet)
	fmt.Println("ear:", earSet)
	fmt.Println("-----------------------")
}
