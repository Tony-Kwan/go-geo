package geo

type SpatialContext struct {
	calc Calculator
}

func NewSpatialContext() *SpatialContext {
	return &SpatialContext{calc: &SphereCalculator{}}
}

func (ctx *SpatialContext) GetCalculator() Calculator {
	return ctx.calc
}
