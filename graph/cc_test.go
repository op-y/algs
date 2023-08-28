package graph

import (
	"testing"

	"github.com/op-y/algs/golang/bag"
)

func TestCC(t *testing.T) {
	g := NewGraph(13)
	g.AddEdge(0, 6)
	g.AddEdge(0, 2)
	g.AddEdge(0, 1)
	g.AddEdge(0, 5)
	g.AddEdge(1, 0)
	g.AddEdge(2, 0)
	g.AddEdge(3, 5)
	g.AddEdge(3, 4)
	g.AddEdge(4, 5)
	g.AddEdge(4, 6)
	g.AddEdge(4, 3)
	g.AddEdge(5, 3)
	g.AddEdge(5, 4)
	g.AddEdge(5, 0)
	g.AddEdge(6, 0)
	g.AddEdge(6, 4)
	g.AddEdge(7, 8)
	g.AddEdge(8, 7)
	g.AddEdge(9, 11)
	g.AddEdge(9, 10)
	g.AddEdge(9, 12)
	g.AddEdge(10, 9)
	g.AddEdge(11, 9)
	g.AddEdge(11, 12)
	g.AddEdge(12, 11)
	g.AddEdge(12, 9)
	if g.Vertex() != 13 {
		t.Errorf("V should be 13!")
	}
	g.Out()

	cc := NewCC(g)
	cc.FindCC()
	bags := make([]*bag.Bag, cc.Count())
	for i := 0; i < cc.Count(); i++ {
		bags[i] = bag.NewBag()
	}
	for i := 0; i < cc.GD.G.Vertex(); i++ {
		bags[cc.Id(i)].Add(i)
	}
	for i := 0; i < cc.Count(); i++ {
		bags[i].Out()
	}
}
