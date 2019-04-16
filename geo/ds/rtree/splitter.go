package rtree

type Splitter interface {
	Split(*Rtree, *rnode) (*rnode, *rnode)
}

type UglySplitter struct{}

func (splitter *UglySplitter) Split(r *Rtree, node *rnode) (*rnode, *rnode) {
	l := node
	ll := &rnode{parent: node.parent, entries: node.entries[r.maxEntries:node.NumEntries():node.NumEntries()], isLeaf: node.isLeaf}
	l.entries = node.entries[0:r.maxEntries:r.maxEntries]
	if !ll.isLeaf {
		for i := range ll.entries {
			ll.entries[i].(*rnode).parent = ll
		}
	}
	return l, ll
}
