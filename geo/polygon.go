package geo

import (
	"github.com/emirpasic/gods/lists/arraylist"
	"math"
)

//TODO: optimize to O(nlogn)
var polarAngleSort = func(ps []Point) {
	n := len(ps)
	for i := 1; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			ai := math.Atan2(ps[i].y-ps[0].y, ps[i].x-ps[0].x)
			aj := math.Atan2(ps[j].y-ps[0].y, ps[j].x-ps[0].x)
			if ai != aj {
				if aj < ai {
					ps[i], ps[j] = ps[j], ps[i]
				}
			} else {
				if ps[j].x < ps[i].x {
					ps[i], ps[j] = ps[j], ps[i]
				}
			}
		}
	}
}

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

func (p *Polygon) ConvexHull() (*Polygon, error) {
	//TODO: validation
	n := p.GetNumPoints() - 1
	points := make([]Point, n)
	for i := 0; i < n; i++ {
		points[i] = p.shell[i]
	}
	bottomLeft := Point{x: math.MaxFloat64, y: math.MaxFloat64}
	k := 0
	for i, point := range points {
		if point.y < bottomLeft.y || (bottomLeft.y == point.y && point.x < bottomLeft.x) {
			bottomLeft = points[i]
			k = i
		}
	}
	points[k], points[0] = points[0], points[k]
	polarAngleSort(points)
	s := make([]Point, n)
	s[0], s[1] = points[0], points[1]
	top := 1
	for i := 2; i < n; {
		if top > 0 && s[top-1].cross(&points[i], &s[top]) >= 0.0 {
			top--
		} else {
			top++
			s[top] = points[i]
			i++
		}
	}
	result := make([]Point, top+2)
	for i := 0; i <= top; i++ {
		result[i] = s[i]
	}
	result[top+1] = Point{x: result[0].x, y: result[0].y}
	return &Polygon{shell: result}, nil
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
