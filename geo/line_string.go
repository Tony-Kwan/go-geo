package geo

type LineString []Point

func (r LineString) GetNumPoints() int { return len(r) }
