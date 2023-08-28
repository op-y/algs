package graph

import (
	"testing"
)

func TestDFS(t *testing.T) {
	g := NewGraph(6)
	g.AddEdge(0, 5)
	g.AddEdge(2, 4)
	g.AddEdge(2, 3)
	g.AddEdge(1, 2)
	g.AddEdge(0, 1)
	g.AddEdge(3, 4)
	g.AddEdge(3, 5)
	g.AddEdge(0, 2)
	if g.Vertex() != 6 {
		t.Errorf("V should be 6!")
	}
	if g.Edge() != 8 {
		t.Errorf("E should be 8!")
	}
	g.Out()

	dfs := NewDFS(g)
	dfs.Search(0)
}
