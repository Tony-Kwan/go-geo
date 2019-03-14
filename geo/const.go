package geo

const (
	E6  = 1e-6
	E7  = 1e-7
	E8  = 1e-8
	E9  = 1e-9
	E10 = 1e-10
	E11 = 1e-11
	E12 = 1e-12
	E13 = 1e-13
	E14 = 1e-14
	E15 = 1e-15

	EarthRadius  float64 = 6371000
	EarthRadius2 float64 = EarthRadius * EarthRadius
)

var (
	GeoCtx = NewSpatialContext()

	cartesianCalc = &CartesianCalculator{}
	sphereCalc    = &SphereCalculator{}
	vectorCalc    = &VectorCalculator{}
)
