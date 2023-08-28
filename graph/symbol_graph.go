package graph

import (
	"fmt"

	"github.com/op-y/algs/golang/bag"
)

type SymbolGraph struct {
	V    int
	E    int
	Adj  []*bag.Bag
	ST   map[string]int // 由于类型的原因这里没有使用之前章节里的st
	Keys []string
}

func NewSymbolGraph() *SymbolGraph {
	g := new(SymbolGraph)
	g.V = 0
	g.E = 0
	g.Adj = make([]*bag.Bag, 0)
	g.ST = make(map[string]int, 0)
	g.Keys = make([]string, 0)
	return g
}

func (g *SymbolGraph) AddEdge(v, w string) {
	if _, ok := g.ST[v]; !ok {
		g.ST[v] = g.V
		g.Keys = append(g.Keys, v)
		bag := bag.NewBag()
		g.Adj = append(g.Adj, bag)
		g.V++
	}

	if _, ok := g.ST[w]; !ok {
		g.ST[w] = g.V
		g.Keys = append(g.Keys, w)
		bag := bag.NewBag()
		g.Adj = append(g.Adj, bag)
		g.V++
	}

	g.Adj[g.ST[v]].Add(w)
	g.Adj[g.ST[w]].Add(v)
	g.E++
}

func (g *SymbolGraph) Vertex() int {
	return g.V
}

func (g *SymbolGraph) Edge() int {
	return g.E
}

func (g *SymbolGraph) Adjacency(v int) *bag.Bag {
	return g.Adj[v]
}

func (g *SymbolGraph) Contains(key string) bool {
	if _, ok := g.ST[key]; ok {
		return true
	} else {
		return false
	}
}

func (g *SymbolGraph) Index(key string) int {
	if idx, ok := g.ST[key]; ok {
		return idx
	} else {
		return -1
	}
}

func (g *SymbolGraph) Name(idx int) string {
	return g.Keys[idx]
}

func (g *SymbolGraph) Out() {
	fmt.Println("symbol graph info: ")
	fmt.Printf("V   : %d\n", g.V)
	fmt.Printf("E   : %d\n", g.E)
	fmt.Printf("Keys: %q\n", g.Keys)
	for key, idx := range g.ST {
		fmt.Printf("vertex key %s: ", key)
		adj := g.Adj[idx]
		for i := adj.First; i != nil; i = i.Next {
			fmt.Printf(" %v", i.Value)
		}
		fmt.Printf("\n")
	}
}
