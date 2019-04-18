package geo

import (
	"bytes"
	"fmt"
	"strconv"
)

type LineString []Point

func (r LineString) GetNumPoints() int { return len(r) }

func (r LineString) String() string {
	var buf bytes.Buffer
	buf.WriteString("LINESTRING (")
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
