package wkt

import (
	"testing"
)

func TestWktReader_Read(t *testing.T) {
	cases := []string{
		"POLYGON EMPTY",
		"POLYGON ()",
		"POLYGON ((30 10, 40 40, 20 40, 10 20, 30 10     ))",
		"POLYGON ((35 10, 45 45, 15 40, 10 20, 35 10), (20 30, 35 35, 30 20, 20 30))",
		"POLYGON ((30.0 10.1, 40.2 40.3, 20.4 40.5, 10.6 20.7, 30.8 10.9))",
		"POLYGON ((.0 -10.1, -.2 .3, 1e10 -1E8, 10.6 20.7, 30.8 10.9))",
		"POLYGON ((30 10, 40 40, 20 40, 10 20, 30 10)))",

		"POINT EMPTY",
		"POINT (1e6 1E7)",
		"POINT (30 -10)",
		"POINT (-30.012300 10.0000000)",
		"POINT (-30.012300 10.0000000))",

		"LINESTRING EMPTY",
		"LINESTRING (30 10, 10 30, 40 40)",
		"LINESTRING (.1 -0.1, -.200000 .23450001, 40.1 -40)",
		"LINESTRING (30 10, 10 30, 40 40))",

		"UNKNOWN (0 0)",
	}
	for _, c := range cases {
		t.Log("Input: ", c)
		s, err := WktReader{}.Read(c)
		if err != nil {
			t.Error("[ERROR]", err)
		} else {
			t.Logf("output: %v", s)
		}
	}

}
