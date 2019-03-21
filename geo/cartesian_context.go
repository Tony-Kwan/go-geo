package geo

type CartesianContext struct {
	calc Calculator
}

func NewCartesianContext() *CartesianContext {
	return &CartesianContext{calc: &CartesianCalculator{}}
}

func (ctx *CartesianContext) GetCalculator() Calculator {
	return ctx.calc
}
