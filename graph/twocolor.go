package graph

type TwoColor struct {
	GD             *DFS
	color          []bool
	isTwoColorable bool
}

func NewTwoColor(graph *Graph) *TwoColor {
	tc := new(TwoColor)
	tc.GD = NewDFS(graph)
	tc.color = make([]int, graph.Vertex())
	tc.isTwoColorable = false
	return tc
}

func (tc *TwoColor) Detect() {
	for i := 0; i < tc.GD.G.Vertex(); i++ {
		if !tc.GD.IsMarked(i) {
			tc.Search(i)
		}
	}
}

func (tc *TwoColor) Search(v int) {
	tc.GD.Marked[v] = true
	tc.GD.N++
	adj := tc.GD.G.Adjacency(v)
	for i := adj.First; i != nil; i = i.Next {
		w := i.Value.(int)
		if !tc.GD.IsMarked(w) {
			tc.color[w] = !tc.color[v]
			tc.Search(w)
		} else if tc.color[w] == tc.color[v] {
			tc.isTwoColorable = false
		}
	}
}

func (tc *TwoColor) IsTwoColorable() bool {
	return tc.isTwoColorable
}
