package graph

import (
    "github.com/op-y/algs/golang/queue"
)

type DFSPath struct {
    GD *DFS
    S int
    EdgeTo []int
}

type BFSPath struct {
    GB *BFS
    S int
    EdgeTo []int
}

func NewDFSPath(graph *Graph, vertex int) *DFSPath {
    dfsp := new(DFSPath)
    dfsp.GD = NewDFS(graph)
    dfsp.S = vertex
    dfsp.EdgeTo = make([]int, graph.Vertex())
    return dfsp
}

func NewBFSPath(graph *Graph, vertex int) *BFSPath {
    bfsp := new(BFSPath)
    bfsp.GB = NewBFS(graph)
    bfsp.S = vertex
    bfsp.EdgeTo = make([]int, graph.Vertex())
    return bfsp
}

// v 是路径起点
func (dfsp *DFSPath) Search(v int) {
    dfsp.GD.Marked[v] = true
    dfsp.GD.N++
    //fmt.Printf("%d ==> ", v)
    adj := dfsp.GD.G.Adjacency(v)
    for i := adj.First; i != nil; i = i.Next {
        w := i.Value.(int)
        if !dfsp.GD.IsMarked(w) {
            dfsp.EdgeTo[w] = v
            dfsp.Search(w)
        }
    }
    //fmt.Printf("Done\n")
}

func (dfsp *DFSPath) HasPathTo(v int) bool {
    return dfsp.GD.Marked[v]
}

func (dfsp *DFSPath) PathTo(v int) []int {
    if !dfsp.HasPathTo(v) {
        return nil
    }
    path := make([]int, 0)
    for i := v; i != dfsp.S; i = dfsp.EdgeTo[i] {
        path = append([]int{i}, path...)
    }
    path = append([]int{dfsp.S}, path...)
    return path
}

// v 是路径起点
func (bfsp *BFSPath) Search(v int) {
    q := queue.NewQueue()
    bfsp.GB.Marked[v] = true
    bfsp.GB.N++
    q.Enqueue(v)
    for !q.IsEmpty() {
        item := q.Dequeue()
        x := item.(int)
        adj := bfsp.GB.G.Adjacency(x)
        for i := adj.First; i != nil; i = i.Next {
            w := i.Value.(int)
            if !bfsp.GB.IsMarked(w) {
                bfsp.GB.Marked[w] = true
                bfsp.EdgeTo[w] = x
                q.Enqueue(w)
            }
        }
    }
}

func (bfsp *BFSPath) HasPathTo(v int) bool {
    return bfsp.GB.Marked[v]
}

func (bfsp *BFSPath) PathTo(v int) []int {
    if !bfsp.HasPathTo(v) {
        return nil
    }
    path := make([]int, 0)
    for i := v; i != bfsp.S; i = bfsp.EdgeTo[i] {
        path = append([]int{i}, path...)
    }
    path = append([]int{bfsp.S}, path...)
    return path
}

