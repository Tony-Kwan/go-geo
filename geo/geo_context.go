package geo

var (
	GeoCtx = NewSpatialContext()
)

type GeoContext interface {
	GetCalculator() Calculator
}
