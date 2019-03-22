package geo

import (
	"fmt"
	"math"
	"strconv"
)

type Rectangle struct {
	AbstractShape

	min *Point
	max *Point
}

func NewRectangle(minX, minY, maxX, maxY float64, ctx GeoContext) *Rectangle {
	return &Rectangle{
		AbstractShape: AbstractShape{ctx: ctx},
		min:           NewPoint(minX, minY, ctx),
		max:           NewPoint(maxX, maxY, ctx),
	}
}

func (rect *Rectangle) GetMinX() float64 { return rect.min.x }
func (rect *Rectangle) GetMinY() float64 { return rect.min.y }
func (rect *Rectangle) GetMaxX() float64 { return rect.max.x }
func (rect *Rectangle) GetMaxY() float64 { return rect.max.y }

func (rect *Rectangle) GetWidth() float64 {
	return rect.max.x - rect.min.x
}

func (rect *Rectangle) GetHeight() float64 {
	return rect.max.y - rect.min.y
}

func (rect *Rectangle) Union(rects ...*Rectangle) *Rectangle {
	ret := rect.clone().(*Rectangle)
	for _, r := range rects {
		ret.min.x = math.Min(ret.min.x, r.min.x)
		ret.max.x = math.Max(ret.max.x, r.max.x)
		ret.min.y = math.Min(ret.min.y, r.min.y)
		ret.max.y = math.Max(ret.max.y, r.max.y)
	}
	return ret
}

func (rect *Rectangle) clone() Shape {
	return NewRectangle(rect.min.x, rect.min.y, rect.max.x, rect.max.y, rect.ctx)
}

func (rect *Rectangle) GetArea() float64 {
	return rect.ctx.GetCalculator().Area(rect)
}

func (rect *Rectangle) String() string {
	//return fmt.Sprintf(
	//	"POLYGON ((%s %s, %s %s, %s %s, %s %s, %s %s))",
	//	strconv.FormatFloat(rect.MinX, 'G', -1, 64), strconv.FormatFloat(rect.MinY, 'G', -1, 64),
	//	strconv.FormatFloat(rect.MinX, 'G', -1, 64), strconv.FormatFloat(rect.MaxY, 'G', -1, 64),
	//	strconv.FormatFloat(rect.MaxX, 'G', -1, 64), strconv.FormatFloat(rect.MaxY, 'G', -1, 64),
	//	strconv.FormatFloat(rect.MaxX, 'G', -1, 64), strconv.FormatFloat(rect.MinY, 'G', -1, 64),
	//	strconv.FormatFloat(rect.MinX, 'G', -1, 64), strconv.FormatFloat(rect.MinY, 'G', -1, 64),
	//)
	return fmt.Sprintf(
		"RECTANGLE ((%s %s), (%s %s))",
		strconv.FormatFloat(rect.min.x, 'G', -1, 64), strconv.FormatFloat(rect.min.y, 'G', -1, 64),
		strconv.FormatFloat(rect.max.x, 'G', -1, 64), strconv.FormatFloat(rect.max.y, 'G', -1, 64),
	)
}
