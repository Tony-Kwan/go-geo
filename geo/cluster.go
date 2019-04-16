package geo

type Clustering interface {
	Partition(dataset Observations, k int) (*ClusterResult, error)
}

type Observation struct {
	Position *Point
	Value    interface{}
}

type Observations []Observation

type ClusterResult struct {
	Centers  []*Point
	Clusters []Observations
	Labels   []int
}

func (r *ClusterResult) reset(k int) {
	r.Clusters = make([]Observations, k)
}
