package graph

import (
	"fmt"

	"github.com/op-y/algs/golang/bag"
)

type Digraph struct {
	V   int
	E   int
	Adj []*bag.Bag
}

func NewDigraph(v int) *Digraph {
	digraph := new(Digraph)
	digraph.V = v
	digraph.Adj = make([]*bag.Bag, v)
	for i := 0; i < v; i++ {
		digraph.Adj[i] = bag.NewBag()
	}
	return digraph
}

func (g *Digraph) Vertex() int {
	return g.V
}

func (g *Digraph) Edge() int {
	return g.E
}

// 如果每个节点的边都添加则无向图中一条边会被计算两次
func (g *Digraph) AddEdge(v, w int) {
	g.Adj[v].Add(w)
	g.E++
}

func (g *Digraph) Adjacency(v int) *bag.Bag {
	return g.Adj[v]
}

func (g *Digraph) Reverse() *Digraph {
	r := NewDigraph(g.V)
	for v := 0; v < g.V; v++ {
		adj := g.Adjacency(v)
		for i := adj.First; i != nil; i = i.Next {
			w := i.Value.(int)
			r.AddEdge(w, v)
		}
	}
	return r
}

func (g *Digraph) Out() {
	fmt.Println("graph info: ")
	fmt.Printf("V: %d\n", g.V)
	fmt.Printf("E: %d\n", g.E)
	for v, Adj := range g.Adj {
		fmt.Printf("%d: ", v)
		for i := Adj.First; i != nil; i = i.Next {
			fmt.Printf(" %v", i.Value)
		}
		fmt.Printf("\n")
	}
}
