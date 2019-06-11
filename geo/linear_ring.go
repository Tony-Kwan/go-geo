package geo

import (
	"bytes"
	"fmt"
	"strconv"
)

type LinearRing LineString

func (r LinearRing) GetNumPoints() int { return len(r) }

func (r LinearRing) IsValid() bool {
	n := r.GetNumPoints()
	if n < 4 {
		return false
	}
	return r[0].Equals(r[n-1])
}

func (r LinearRing) IsSimple() bool {
	return true //TODO: impl
}

func (r LinearRing) IsCCW() bool {
	var sum float64
	var nextPoint Point
	n := r.GetNumPoints()
	for i, point := range r {
		nextPoint = r[(i+1)%n]
		sum += (nextPoint.X() - point.X()) * (nextPoint.Y() + point.Y())
	}
	return sum <= 0.0
}

func (r LinearRing) MakeCCW() LinearRing {
	n := r.GetNumPoints()
	ret := make(LinearRing, n)
	if !r.IsCCW() {
		for i := range r {
			ret[n-i-1] = r[i]
		}
	} else {
		for i := range r {
			ret[i] = r[i]
		}
	}
	return ret
}

func (r LinearRing) String() string {
	var buf bytes.Buffer
	buf.WriteString("LINEARRING (")
	n := r.GetNumPoints()
	for i, point := range r {
		buf.WriteString(fmt.Sprintf("%s %s", strconv.FormatFloat(point.X(), 'f', -1, 64), strconv.FormatFloat(point.Y(), 'f', -1, 64)))
		if i != n-1 {
			buf.WriteString(", ")
		}
	}
	buf.WriteRune(')')
	return buf.String()
}
