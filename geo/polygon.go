package geo

import (
	"github.com/emirpasic/gods/lists/arraylist"
)

type Polygon struct {
	shell LinearRing
	//Holes []LinearRing  //TODO: Support holes
}

func (p *Polygon) GetNumPoints() int {
	return p.shell.GetNumPoints()
}

func (p *Polygon) IsSimple() bool {
	return p.shell.IsSimple()
}

// O(n²)
func (p *Polygon) Triangulate() ([]Triangle, error) {
	//TODO: validate polygon
	convexVertices := arraylist.New()
	reflexVertices := arraylist.New()
	earVertices := arraylist.New()
	var p1, p3 Point
	var v1, v2 vector2
	var shell = p.shell.MakeCCW()[:p.shell.GetNumPoints()-1]
	n := shell.GetNumPoints()
	for i, p2 := range shell {
		p1 = shell[(n+i-1)%n]
		p3 = shell[(i+1)%n]
		v1.x, v1.y = p1.X()-p2.X(), p1.Y()-p2.Y()
		v2.x, v2.y = p3.X()-p2.X(), p3.Y()-p2.Y()
		if v1.cross(&v2) > 0.0 {
			reflexVertices.Add(i)
		} else {
			convexVertices.Add(i)
		}
	}
	// find ear
	cIt := convexVertices.Iterator()
	for cIt.Next() {
		convexVertex := cIt.Value().(int)
		tri := Triangle{
			a: &shell[(n+convexVertex-1)%n],
			b: &shell[convexVertex],
			c: &shell[(convexVertex+1)%n],
		}
		isEar := true
		rIt := reflexVertices.Iterator()
		for rIt.Next() {
			reflexVertex := rIt.Value().(int)
			abs := AbsInt(convexVertex - reflexVertex)
			if abs != 1 && abs != n-1 && tri.contains(&shell[reflexVertex]) {
				isEar = false
				break
			}
		}
		if isEar {
			earVertices.Add(convexVertex)
		}
	}
	//remove ear
	eIt := earVertices.Iterator()
	for eIt.Next() {
		//TODO: impl
	}
	return []Triangle{}, nil
}
func (p *Polygon) String() string {
	return "POLYGON(" + p.shell.String()[10:] + ")"
}
