package geo

import (
	"fmt"
	"math"
	"strconv"
)

type Rectangle struct {
	AbstractShape

	MinX float64
	MaxX float64
	MinY float64
	MaxY float64
}

func NewRectangle(minX, maxX, minY, maxY float64, ctx GeoContext) *Rectangle {
	return &Rectangle{
		AbstractShape: AbstractShape{ctx: ctx},
		MinX:          minX,
		MaxX:          maxX,
		MinY:          minY,
		MaxY:          maxY,
	}
}

func (rect *Rectangle) GetWidth() float64 {
	return rect.MaxX - rect.MinY
}

func (rect *Rectangle) GetHeight() float64 {
	return rect.MaxY - rect.MinY
}

func (rect *Rectangle) Union(rects ...*Rectangle) *Rectangle {
	ret := rect.clone().(*Rectangle)
	for _, r := range rects {
		ret.MinX = math.Min(ret.MinX, r.MinX)
		ret.MaxX = math.Max(ret.MaxX, r.MaxX)
		ret.MinY = math.Min(ret.MinY, r.MinY)
		ret.MaxY = math.Max(ret.MaxY, r.MaxY)
	}
	return ret
}

func (rect *Rectangle) clone() Shape {
	return NewRectangle(rect.MinX, rect.MaxX, rect.MinY, rect.MaxY, rect.ctx)
}

func (rect *Rectangle) GetArea() float64 {
	return rect.ctx.GetCalculator().Area(rect)
}

func (rect *Rectangle) String() string {
	return fmt.Sprintf(
		"POLYGON ((%s %s, %s %s, %s %s, %s %s, %s %s))",
		strconv.FormatFloat(rect.MinX, 'G', -1, 64), strconv.FormatFloat(rect.MinY, 'G', -1, 64),
		strconv.FormatFloat(rect.MinX, 'G', -1, 64), strconv.FormatFloat(rect.MaxY, 'G', -1, 64),
		strconv.FormatFloat(rect.MaxX, 'G', -1, 64), strconv.FormatFloat(rect.MaxY, 'G', -1, 64),
		strconv.FormatFloat(rect.MaxX, 'G', -1, 64), strconv.FormatFloat(rect.MinY, 'G', -1, 64),
		strconv.FormatFloat(rect.MinX, 'G', -1, 64), strconv.FormatFloat(rect.MinY, 'G', -1, 64),
	)
}
