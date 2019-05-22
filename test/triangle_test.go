package test

import (
	"fmt"
	"github.com/Tony-Kwan/go-geo/geo/io/wkt"
	"github.com/atotto/clipboard"
	"math/rand"
	"strings"
	"testing"
	"time"
)

func TestTriangle_IsDisjoint(t *testing.T) {
	var wktStr = "POLYGON((-153.402099609375 43.192662019750884,-157.005615234375 26.760632389116736,-141.44519805908203 25.760010550981164,-142.14832305908203 35.78189218363366,-132.41924285888675 36.137597450383765,-131.89189910888672 26.54553708313331,-116.01047515869142 27.35316775063474,-117.15305328369142 39.011447820620134,-127.38269805908205 38.671840771106446,-127.73426055908203 44.0118301164365,-121.58191680908203 40.900279369669875,-115.86902618408205 44.82787379200292,-125.44910430908205 47.74232479513432,-136.76021575927734 43.248453753571255,-143.70357513427734 38.898247462673254,-135.61763763427734 38.14184742464704,-141.24263763427734 36.81670599350721,-147.71976470947268 36.67585384333481,-146.22562408447266 42.734152628905065,-145.61038970947266 45.3082172149164,-153.402099609375 43.192662019750884))"
	var polygon = wkt.MustPolygon(wktReader.Read(wktStr))
	tris, err := polygon.Triangulate()
	if err != nil {
		t.Error(err)
		return
	}
	n := len(tris)
	mat := make([][]int, n)
	for i := 0; i < n; i++ {
		mat[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if tris[i].IsDisjoint(&tris[j]) {
				mat[i][j] = 0
			} else {
				mat[i][j] = 1
			}
		}
	}
	for _, m := range mat {
		fmt.Println(m)
	}

	rand.Seed(time.Now().UnixNano())
	rIdx := rand.Intn(n)
	wkts := []string{wktStr, tris[rIdx].String(), tris[rIdx].String(), tris[rIdx].String(), tris[rIdx].String()}
	cnt := 0
	for i := range mat[rIdx] {
		if mat[rIdx][i] == 1 {
			cnt++
			wkts = append(wkts, tris[i].String())
		}
	}
	fmt.Println("joint cnt:", cnt)
	s := fmt.Sprintf("GEOMETRYCOLLECTION(%s)", strings.Join(wkts, ", "))
	clipboard.WriteAll(s)
}
