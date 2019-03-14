package geo

import (
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

func NewPolygon(shell LinearRing) *Polygon {
	return &Polygon{shell: shell}
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

func (p *Polygon) String() string {
	return "POLYGON(" + p.shell.String()[10:] + ")"
}
