package graph

import (
	"fmt"
)

type DirectedDFS struct {
	DG     *Digraph
	Marked []bool
}

func NewDirectedDFS(digraph *Digraph) *DirectedDFS {
	ddfs := new(DirectedDFS)
	ddfs.DG = digraph
	ddfs.Marked = make([]bool, digraph.Vertex())
	return ddfs
}

func (ddfs *DirectedDFS) IsMarked(v int) bool {
	return ddfs.Marked[v]
}

func (ddfs *DirectedDFS) SearchSources(sources []int) {
	for _, source := range sources {
		if !ddfs.IsMarked(source) {
			ddfs.Search(source)
		}
	}
}

func (ddfs *DirectedDFS) Search(v int) {
	ddfs.Marked[v] = true
	adj := ddfs.DG.Adjacency(v)
	for i := adj.First; i != nil; i = i.Next {
		w := i.Value.(int)
		if !ddfs.IsMarked(w) {
			ddfs.Search(w)
		}
	}
}

func (ddfs *DirectedDFS) Out() {
	for idx, v := range ddfs.Marked {
		fmt.Printf("%d => %t\n", idx, v)
	}
}
