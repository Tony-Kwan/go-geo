package geo

import "testing"

func TestLinearRing_IsCCW(t *testing.T) {
	cases := []struct {
		ring  LinearRing
		isCCW bool
	}{
		{
			LinearRing{
				*NewPoint(-66.91566467285156, 39.93632920085673, nil),
				*NewPoint(-57.973823547363295, 33.81309907738368, nil),
				*NewPoint(-72.24437713623047, 34.240188956777786, nil),
				*NewPoint(-66.91566467285156, 39.93632920085673, nil),
			},
			false,
		},
		{
			LinearRing{
				*NewPoint(-66.91566467285156, 39.93632920085673, nil),
				*NewPoint(-72.24437713623047, 34.240188956777786, nil),
				*NewPoint(-57.973823547363295, 33.81309907738368, nil),
				*NewPoint(-66.91566467285156, 39.93632920085673, nil),
			},
			true,
		},
		{
			LinearRing{
				*NewPoint(-67.43545532226564, 40.40460782849249, nil),
				*NewPoint(-56.04469299316408, 40.518497187767764, nil),
				*NewPoint(-60.80829620361329, 38.05728288377773, nil),
				*NewPoint(-55.934829711914084, 35.04967312791719, nil),
				*NewPoint(-67.47871398925783, 35.729513159796724, nil),
				*NewPoint(-67.43545532226564, 40.40460782849249, nil),
			},
			false,
		},
		{
			LinearRing{
				*NewPoint(-67.43545532226564, 40.40460782849249, nil),
				*NewPoint(-67.47871398925783, 35.729513159796724, nil),
				*NewPoint(-55.934829711914084, 35.04967312791719, nil),
				*NewPoint(-60.80829620361329, 38.05728288377773, nil),
				*NewPoint(-56.04469299316408, 40.518497187767764, nil),
				*NewPoint(-67.43545532226564, 40.40460782849249, nil),
			},
			true,
		},
	}
	for _, c := range cases {
		if c.ring.IsCCW() != c.isCCW {
			t.Errorf("Test IsCCW(): expect=%v, found=%v, ring=%v", c.isCCW, c.ring.IsCCW(), c.ring)
		}
	}
}

func TestLinearRing_MakeCCW(t *testing.T) {
	ring := LinearRing{
		*NewPoint(-67.43545532226564, 40.40460782849249, nil),
		*NewPoint(-56.04469299316408, 40.518497187767764, nil),
		*NewPoint(-60.80829620361329, 38.05728288377773, nil),
		*NewPoint(-55.934829711914084, 35.04967312791719, nil),
		*NewPoint(-67.47871398925783, 35.729513159796724, nil),
		*NewPoint(-67.43545532226564, 40.40460782849249, nil),
	}
	ring.MakeCCW()
	if !ring.IsCCW() {
		t.Errorf("Test MakeCCW(): fail: ring=%v", ring)
	}
}
