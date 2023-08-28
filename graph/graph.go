package graph

import (
	"fmt"

	"github.com/op-y/algs/golang/bag"
)

type Graph struct {
	V   int
	E   int
	Adj []*bag.Bag
}

func NewGraph(v int) *Graph {
	graph := new(Graph)
	graph.V = v
	graph.Adj = make([]*bag.Bag, v)
	for i := 0; i < v; i++ {
		graph.Adj[i] = bag.NewBag()
	}
	return graph
}

func (g *Graph) Vertex() int {
	return g.V
}

func (g *Graph) Edge() int {
	return g.E
}

// 如果每个节点的边都添加则无向图中一条边会被计算两次
func (g *Graph) AddEdge(v, w int) {
	g.Adj[v].Add(w)
	g.Adj[w].Add(v)
	g.E++
}

func (g *Graph) Adjacency(v int) *bag.Bag {
	return g.Adj[v]
}

func (g *Graph) Out() {
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
