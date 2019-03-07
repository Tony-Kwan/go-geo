package geo

type Triangle struct {
	AbstractShape

	A *Point
	B *Point
	C *Point
}

func NewTriangle(a, b, c *Point, ctx GeoContext) *Triangle {
	return &Triangle{
		AbstractShape: AbstractShape{ctx: ctx},
		A:             a,
		B:             b,
		C:             c,
	}
}

func (tri *Triangle) GetArea() float64 {
	return tri.ctx.GetCalculator().Area(tri)
}
