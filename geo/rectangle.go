package geo

type Rectangle struct {
	AbstractShape

	minX float64
	maxX float64
	minY float64
	maxY float64
}

func NewRectangle(minX, maxX, minY, maxY float64, ctx GeoContext) Rectangle {
	return Rectangle{
		AbstractShape: AbstractShape{ctx: ctx},
		minX:          minX,
		maxX:          maxX,
		minY:          minY,
		maxY:          maxY,
	}
}

func (rect Rectangle) GetArea() float64 {
	return rect.ctx.GetCalculator().Area(rect)
}
