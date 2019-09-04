package geo

import (
	"errors"
	"fmt"
	"github.com/Tony-Kwan/go-geo/geo/internal/ds"
	"github.com/deckarep/golang-set"
	"math"
)

type earClippingPoint struct {
	Point
	isReflex bool
	isEar    bool

	tmp int
}

func (p Polygon) triangulateWithHole() (tris []Triangle, err error) {
	shell := p.Shell.MakeCCW()[:p.Shell.GetNumPoints()-1]
	hole := p.Holes[0].MakeCW()[:p.Holes[0].GetNumPoints()-1]

	for j, holePoint := range hole {
		var i int
		minDis := math.MaxFloat64
		calc := p.GetContext().GetCalculator()
		for k := 0; k < shell.GetNumPoints(); k++ {
			dis := calc.Distance(holePoint, shell[k])
			if dis < minDis {
				minDis = dis
				i = k
			}
		}

		if i == shell.GetNumPoints() || j == hole.GetNumPoints() {
			return nil, errors.New("can not triangulate")
		}
		ring := make(LinearRing, 0)
		ring = append(ring, shell[i])
		for ii := 1; ii <= shell.GetNumPoints(); ii++ {
			ring = append(ring, shell[(ii+i)%shell.GetNumPoints()])
		}
		ring = append(ring, hole[j])
		for jj := 1; jj <= hole.GetNumPoints(); jj++ {
			ring = append(ring, hole[(jj+j)%hole.GetNumPoints()])
		}
		ring = append(ring, ring[0])
		plg := NewPolygon(ring)
		//fmt.Println("triangulateWithHole:\n", plg)
		tris, err = plg.triangulate()
		if err == nil {
			break
		}
	}
	return
}

func (p Polygon) Triangulate() ([]Triangle, error) {
	//TODO: validate polygon
	if len(p.Holes) > 1 {
		return nil, errors.New("unsupported multiple holes")
	}
	if len(p.Holes) == 1 {
		return p.triangulateWithHole()
	}
	return p.triangulate()
}

// O(n²)
func (p Polygon) triangulate() ([]Triangle, error) {
	var shell = p.Shell.MakeCCW()[:p.Shell.GetNumPoints()-1]
	n := shell.GetNumPoints()
	points := ds.NewCircularLinkedList()
	for i, p := range shell {
		points.Add(&earClippingPoint{Point: p, tmp: i})
	}
	// check convex/reflex
	node := points.Head
	for i := 0; i < points.Size(); i, node = i+1, node.Next {
		curr := node.Elem.(*earClippingPoint)
		curr.isReflex = p.checkIsReflex(node)
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

	//clip ear
	ret := make([]Triangle, 0, n-2)
	var cnt int
	for node := points.Head; ; {
		curr := node.Elem.(*earClippingPoint)
		prev := node.Prev.Elem.(*earClippingPoint)
		next := node.Next.Elem.(*earClippingPoint)
		if points.Size() <= 3 {
			ret = append(ret, Triangle{A: prev.Point, B: curr.Point, C: next.Point})
			break
		}
		if !curr.isEar {
			node = node.Next
			cnt++
			if cnt >= points.Size() {
				it := points.Iterator()
				shell := make(LinearRing, 0)
				for it.Next() {
					pt := it.Value().(*earClippingPoint)
					shell = append(shell, pt.Point)
					//fmt.Println(pt.isReflex, pt.isEar)
				}
				shell = append(shell, shell[0])
				//fmt.Println(NewPolygon(shell).String())
				return nil, errors.New("infinite loop: program should not run here")
			}
			continue
		}
		cnt = 0
		ret = append(ret, Triangle{A: prev.Point, B: curr.Point, C: next.Point})
		node = points.RemoveNode(node)
		next.isReflex = p.checkIsReflex(node)
		prev.isReflex = p.checkIsReflex(node.Prev)
		next.isEar = checkIsEar(node, points)
		prev.isEar = checkIsEar(node.Prev, points)
		//printVertexType(points)
	}
	//for i := 2; i < n; i++ {
	//	ear, err := findBestEar(points)
	//	if err != nil {
	//		return nil, err
	//	}
	//	curr := ear.Elem.(*earClippingPoint)
	//	prev := ear.Prev.Elem.(*earClippingPoint)
	//	next := ear.Next.Elem.(*earClippingPoint)
	//	ret = append(ret, Triangle{A: prev.Point, B: curr.Point, C: next.Point})
	//	node = points.RemoveNode(ear)
	//	next.isReflex = p.checkIsReflex(node)
	//	prev.isReflex = p.checkIsReflex(node.Prev)
	//	next.isEar = checkIsEar(node, points)
	//	prev.isEar = checkIsEar(node.Prev, points)
	//}
	return ret, nil
}

func findBestEar(list *ds.CircularLinkedList) (*ds.Node, error) {
	var chosen *ds.Node
	var best, tmp float64
	var curr, prev, next *ds.Node
	it := list.Iterator()
	for it.Next() {
		if !it.Value().(*earClippingPoint).isEar {
			continue
		}
		curr = it.Node()
		prev = curr.Prev
		next = curr.Next
		v1 := newNEWithPoint(prev.Elem.(*earClippingPoint).Point)
		v2 := newNEWithPoint(curr.Elem.(*earClippingPoint).Point)
		v3 := newNEWithPoint(next.Elem.(*earClippingPoint).Point)
		tmp = math.Abs(v1.cross(v3).unit().dot(v2) + 0.707)
		if chosen == nil || tmp < best {
			best = tmp
			chosen = curr
		}
	}
	if chosen == nil {
		return nil, errors.New("can not find ear vertex")
	}
	return chosen, nil
}

func (p *Polygon) checkIsReflex(node *ds.Node) bool {
	curr := node.Elem.(*earClippingPoint)
	prev := node.Prev.Elem.(*earClippingPoint)
	next := node.Next.Elem.(*earClippingPoint)
	switch p.GetContext().(type) {
	case *SpatialContext:
		v1, v2, v3 := newNEWithPoint(prev.Point), newNEWithPoint(curr.Point), newNEWithPoint(next.Point)
		return v1.cross(v3).dot(v2) > 0
	default:
		return curr.cross(&prev.Point, &next.Point) > 0.0
	}
}

func checkIsEar(node *ds.Node, list *ds.CircularLinkedList) bool {
	if list.Size() == 3 {
		return true
	}

	curr := node.Elem.(*earClippingPoint)
	if curr.isReflex {
		return false
	}
	prev := node.Prev.Elem.(*earClippingPoint)
	next := node.Next.Elem.(*earClippingPoint)
	tri := Triangle{
		A: prev.Point,
		B: curr.Point,
		C: next.Point,
	}
	it := list.Iterator()
	for it.Next() {
		e := it.Value().(*earClippingPoint)
		if e.isReflex {
			if e != prev && e != next && tri.contains(e.Point) {
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
