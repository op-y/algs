package graph

import (
	"testing"
)

func TestGraph(t *testing.T) {
	g := NewGraph(13)
	g.AddEdge(0, 5)
	g.AddEdge(4, 3)
	g.AddEdge(0, 1)
	g.AddEdge(9, 12)
	g.AddEdge(6, 4)
	g.AddEdge(5, 4)
	g.AddEdge(0, 2)
	g.AddEdge(11, 12)
	g.AddEdge(9, 10)
	g.AddEdge(0, 6)
	g.AddEdge(7, 8)
	g.AddEdge(9, 11)
	g.AddEdge(5, 3)
	if g.Vertex() != 13 {
		t.Errorf("V should be 13!")
	}
	if g.Edge() != 13 {
		t.Errorf("E should be 13!")
	}
	g.Out()
}
