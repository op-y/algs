package graph

type CC struct {
	GD    *DFS
	id    []int
	count int
}

func NewCC(graph *Graph) *CC {
	cc := new(CC)
	cc.GD = NewDFS(graph)
	cc.id = make([]int, graph.Vertex())
	cc.count = 0
	return cc
}

func (cc *CC) FindCC() {
	for i := 0; i < cc.GD.G.Vertex(); i++ {
		if !cc.GD.IsMarked(i) {
			cc.Search(i)
			cc.count++
		}
	}
}

func (cc *CC) Search(v int) {
	cc.GD.Marked[v] = true
	cc.GD.N++
	cc.id[v] = cc.count
	adj := cc.GD.G.Adjacency(v)
	for i := adj.First; i != nil; i = i.Next {
		w := i.Value.(int)
		if !cc.GD.IsMarked(w) {
			cc.Search(w)
		}
	}
}

func (cc *CC) Connected(v, w int) bool {
	return cc.id[v] == cc.id[w]
}

func (cc *CC) Id(v int) int {
	return cc.id[v]
}

func (cc *CC) Count() int {
	return cc.count
}
