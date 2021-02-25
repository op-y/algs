package graph

import (
    "fmt"
)

type DFS struct {
    G *Graph
    Marked []bool
    N int
}

func NewDFS(graph *Graph) *DFS {
    dfs := new(DFS)
    dfs.G = graph
    dfs.Marked = make([]bool, graph.Vertex())
    dfs.N = 0
    return dfs
}

func (dfs *DFS) IsMarked(w int) bool {
    return dfs.Marked[w]
}

func (dfs *DFS) Count() int {
    return dfs.N
}

func (dfs *DFS) Search(v int) {
    dfs.Marked[v] = true
    dfs.N++
    fmt.Printf("%d ==> ", v)
    adj := dfs.G.Adjacency(v)
    for i := adj.First; i != nil; i = i.Next {
        w := i.Value.(int)
        if !dfs.IsMarked(w) {
            dfs.Search(w)
        }
    }
    fmt.Printf("Done\n")
}
