package geo

import "math"

func sign(a float64) int {
	switch {
	case a < 0:
		return -1
	case a > 0:
		return +1
	}
	return 0
}

func AbsInt(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func ApproxEqual(f1, f2, eps float64) bool {
	return math.Abs(f1-f2) < eps
}
