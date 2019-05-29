package ds

type TwoDimIntSlice [][]int

func New2DIntSlice(n, m int) TwoDimIntSlice {
	slc := make([][]int, n)
	for i := 0; i < n; i++ {
		slc[i] = make([]int, m)
	}
	return slc
}

func (this TwoDimIntSlice) Fill(v int) TwoDimIntSlice {
	for i := 0; i < len(this); i++ {
		for j := 0; j < len(this[i]); j++ {
			this[i][j] = v
		}
	}
	return this
}

type TwoDimBoolSlice [][]bool

func New2DBoolSlice(n, m int) TwoDimBoolSlice {
	slc := make([][]bool, n)
	for i := 0; i < n; i++ {
		slc[i] = make([]bool, m)
	}
	return slc
}

func (this TwoDimBoolSlice) Fill(b bool) TwoDimBoolSlice {
	for i := 0; i < len(this); i++ {
		for j := 0; j < len(this[i]); j++ {
			this[i][j] = b
		}
	}
	return this
}
