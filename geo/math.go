package geo

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
