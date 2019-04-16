package geo

import (
	"errors"
	"math"
	"math/rand"
	"time"
)

type Kmeans struct {
	calc Calculator
}

func NewKmeans(calc Calculator) *Kmeans {
	return &Kmeans{calc: calc}
}

func (km *Kmeans) Partition(dataset Observations, k int) (*ClusterResult, error) {
	if len(dataset) < k {
		return nil, errors.New("num of dataset must gte k")
	}
	return km.partition(dataset, k)
}

func (km *Kmeans) partition(dataset Observations, k int) (*ClusterResult, error) {
	n := len(dataset)
	result := &ClusterResult{
		Centers: km.initCenteres(dataset, k),
		Labels:  make([]int, n),
	}

	var change = 1
	for t := 0; change != 0 && t < 100; t++ {
		change = 0
		result.reset(k)

		labels := make([]int, n)
		for i, d := range dataset {
			labels[i] = km.nearestCenter(result.Centers, d.Position)
			result.Clusters[labels[i]] = append(result.Clusters[labels[i]], d)
			if labels[i] != result.Labels[i] {
				change++
			}
		}

		//{ //log
		//	fmt.Println(t, change, result.Centers)
		//	//count distinct label
		//	count := make([]int, k)
		//	for _, label := range labels {
		//		count[label]++
		//	}
		//	fmt.Println(count)
		//}

		result.Centers = km.recenter(result.Clusters)
		result.Labels = labels
	}

	return result, nil
}

func (km *Kmeans) initCenteres(dataset Observations, k int) []*Point {
	rand.Seed(time.Now().UnixNano())

	n := len(dataset)
	firstIdx := rand.Intn(n)
	centers := make([]*Point, k)
	centers[0] = dataset[firstIdx].Position
	for i := 1; i < k; i++ {
		maxDis, maxIdx := -math.MaxFloat64, -1
		for m, o := range dataset {
			//TODO: optimize, may all the same
			flag := false
			for j := 0; j < i; j++ {
				if centers[j].ApproxEqualWithEps(o.Position, E15) {
					flag = true
					break
				}
			}
			if flag {
				continue
			}

			dis := 0.
			for j := 0; j < i; j++ {
				dis += km.calc.Distance(o.Position, centers[j])
			}
			if dis > maxDis {
				maxDis = dis
				maxIdx = m
			}
		}
		centers[i] = dataset[maxIdx].Position
	}
	return centers
}

func (km *Kmeans) nearestCenter(centers []*Point, point *Point) int {
	minDis, minIdx := math.MaxFloat64, -1
	var dis float64
	for i := range centers {
		dis = km.calc.Distance(centers[i], point)
		if dis < minDis {
			minDis = dis
			minIdx = i
		}
	}
	return minIdx
}

func (km *Kmeans) recenter(obsl []Observations) []*Point {
	centers := make([]*Point, len(obsl))
	for i, obs := range obsl {
		points := make([]*Point, len(obs))
		for j, ob := range obs {
			points[j] = ob.Position
		}
		centers[i] = km.calc.MeanPosition(points...)
	}
	return centers
}
