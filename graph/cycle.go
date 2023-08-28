package graph

type CycleDetector struct {
	GD       *DFS
	hasCycle bool
}

func NewCycleDetector(graph *Graph) *CycleDetector {
	cd := new(CycleDetector)
	cd.GD = NewDFS(graph)
	cd.hasCycle = false
	return cd
}

func (cd *CycleDetector) Detect() {
	for i := 0; i < cd.GD.G.Vertex(); i++ {
		if !cd.GD.IsMarked(i) {
			cd.Search(i, i)
		}
	}
}

func (cd *CycleDetector) Search(v, u int) {
	cd.GD.Marked[v] = true
	cd.GD.N++
	adj := cd.GD.G.Adjacency(v)
	for i := adj.First; i != nil; i = i.Next {
		w := i.Value.(int)
		if !cd.GD.IsMarked(w) {
			cd.Search(w, v)
		} else if w != u { // 排除当前两个节点之间的边
			cd.hasCycle = true
		}
	}
}

func (cd *CycleDetector) HasCycle() bool {
	return cd.hasCycle
}
