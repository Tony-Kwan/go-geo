package ds

type OrderObj struct {
	Obj   interface{}
	Order int
}

type OrderObjs []OrderObj

func (this OrderObjs) Len() int           { return len(this) }
func (this OrderObjs) Less(i, j int) bool { return this[i].Order < this[j].Order }
func (this OrderObjs) Swap(i, j int)      { this[i], this[j] = this[j], this[i] }
