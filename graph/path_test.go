package graph

import (
	"fmt"
	"testing"
)

func TestDFSPath(t *testing.T) {
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

	dfsp := NewDFSPath(g, 0)
	dfsp.Search(0)
	for i := 0; i < dfsp.GD.G.Vertex(); i++ {
		if dfsp.HasPathTo(i) {
			path := dfsp.PathTo(i)
			fmt.Printf("Path from %d to %d: ", 0, i)
			for _, x := range path {
				fmt.Printf("%d -> ", x)
			}
			fmt.Printf("\n")
		}
	}
	//path := dfsp.PathTo(4)
	//fmt.Printf("Path from %d to %d: ", 0, 4)
	//for _, x := range path {
	//    fmt.Printf("%d -> ", x)
	//}
	//fmt.Printf("\n")
}

func TestBFSPath(t *testing.T) {
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

	bfsp := NewBFSPath(g, 0)
	bfsp.Search(0)
	for i := 0; i < bfsp.GB.G.Vertex(); i++ {
		if bfsp.HasPathTo(i) {
			path := bfsp.PathTo(i)
			fmt.Printf("Path from %d to %d: ", 0, i)
			for _, x := range path {
				fmt.Printf("%d -> ", x)
			}
			fmt.Printf("\n")
		}
	}
}
