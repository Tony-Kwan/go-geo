package geo

type SpatialContext struct {
	calc Calculator
}

func NewSpatialContext() *SpatialContext {
	//return &SpatialContext{calc: &SphereCalculator{}}
	return &SpatialContext{calc: &VectorCalculator{}}
}

func (ctx *SpatialContext) GetCalculator() Calculator {
	return ctx.calc
}
