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
	AbstractShape

	Shell LinearRing
	//Holes []LinearRing  //TODO: Support holes
}

func NewPolygon(shell LinearRing) Polygon {
	return Polygon{Shell: shell}
}

func (p Polygon) GetNumPoints() int {
	return p.Shell.GetNumPoints()
}

func (p Polygon) IsSimple() bool {
	return p.Shell.IsSimple()
}

func (p Polygon) GetArea() float64 {
	return p.GetContext().GetCalculator().Area(p)
}

func (p Polygon) Bounds() Rectangle {
	rect := NewRectangle(math.MaxFloat64, -math.MaxFloat64, math.MaxFloat64, -math.MaxFloat64, p.GetContext())
	for _, pt := range p.Shell {
		rect.minX = math.Min(rect.minX, pt.x)
		rect.maxX = math.Max(rect.maxX, pt.x)
		rect.minY = math.Min(rect.minY, pt.y)
		rect.maxY = math.Max(rect.maxY, pt.y)
	}
	return rect
}

func (p Polygon) ConvexHull() (Polygon, error) {
	//TODO: validation
	n := p.GetNumPoints() - 1
	points := make([]Point, n)
	for i := 0; i < n; i++ {
		points[i] = p.Shell[i]
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
	return Polygon{Shell: result}, nil
}

func (p Polygon) Contain(pt Point) bool {
	n := p.GetNumPoints()
	var p1, p2 Point
	var c int
	for i := 0; i < n-1; i++ {
		p1 = p.Shell[i]
		p2 = p.Shell[i+1]
		if ((pt.y >= p1.y) && (pt.y <= p2.y)) || ((pt.y >= p2.y) && (pt.y <= p1.y)) {
			duT := (pt.y - p1.y) / (p2.y - p1.y)
			duXT := p1.x + duT*(p2.x-p1.x)
			if pt.x == duXT {
				return false
			}
			if pt.x > duXT {
				c++
			}
		}
	}
	return c%2 == 1
}

func (p Polygon) CoverByCircles(k int) ([]Circle, error) {
	var density = 50
	ps := make([]Point, 0)
	n := p.GetNumPoints()
	calc := p.GetContext().GetCalculator()
	for i := 0; i < n; i++ {
		p1, p2 := p.Shell[i], p.Shell[(i+1)%n]
		dDis := calc.Distance(p1, p2) / float64(density)
		bearing := calc.Bearing(p1, p2)
		for j := 0; j < density; j++ {
			ps = append(ps, calc.PointOnBearing(p1, dDis*float64(j), bearing, p.GetContext()))
		}
	}

	bound := p.Bounds()
	dx, dy := (bound.maxX-bound.minX)/float64(density), (bound.maxY-bound.minY)/float64(density)
	for i := 1; i < density; i++ {
		for j := 1; j < density; j++ {
			pt := NewPoint(bound.minX+dx*float64(i), bound.minY+dy*float64(j), p.GetContext())
			if p.Contain(pt) {
				ps = append(ps, pt)
			}
		}
	}

	dataset := Observations{}
	for _, p := range ps {
		dataset = append(dataset, Observation{
			Position: NewPoint(p.X(), p.Y(), nil),
			Value:    nil,
		})
	}
	km := NewKmeans(VectorCalculator{})
	result, err := km.Partition(dataset, k)
	if err != nil {
		return nil, err
	}

	circles := make([]Circle, 0)
	for _, cluster := range result.Clusters {
		pts := make([]Point, len(cluster))
		for i, o := range cluster {
			pts[i] = o.Position
		}
		circle, err := vectorCalc.MinCoverCircle(pts...)
		if err != nil {
			return nil, err
		}
		circles = append(circles, circle)
	}
	return circles, nil
}

func (p Polygon) String() string {
	return "POLYGON(" + p.Shell.String()[10:] + ")"
}
