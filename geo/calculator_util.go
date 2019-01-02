package geo

import "math"

const (
	deg2Rad = math.Pi / 180
	rad2Deg = 1 / deg2Rad
)

func ToRadians(degrees float64) float64 {
	return degrees * deg2Rad
}

func ToDegrees(radians float64) float64 {
	return radians * rad2Deg
}

func Radians2Dist(radians, radius float64) float64 {
	return radians * radius
}

func Dist2Radians(dist, radius float64) float64 {
	return dist / radius
}

func Degrees2Dist(degrees, radius float64) float64 {
	return Radians2Dist(ToRadians(degrees), radius)
}

func Dist2Degrees(dist, radius float64) float64 {
	return ToDegrees(Dist2Radians(dist, radius))
}
