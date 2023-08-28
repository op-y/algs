package graph

import (
	"fmt"

	"github.com/op-y/algs/golang/queue"
)

type BFS struct {
	G      *Graph
	Marked []bool
	N      int
}

func NewBFS(graph *Graph) *BFS {
	bfs := new(BFS)
	bfs.G = graph
	bfs.Marked = make([]bool, graph.Vertex())
	bfs.N = 0
	return bfs
}

func (bfs *BFS) IsMarked(w int) bool {
	return bfs.Marked[w]
}

func (bfs *BFS) Count() int {
	return bfs.N
}

func (bfs *BFS) Search(v int) {
	q := queue.NewQueue()
	bfs.Marked[v] = true
	bfs.N++
	q.Enqueue(v)
	for !q.IsEmpty() {
		item := q.Dequeue()
		x := item.(int)
		adj := bfs.G.Adjacency(x)
		fmt.Printf("Layer ")
		for i := adj.First; i != nil; i = i.Next {
			w := i.Value.(int)
			if !bfs.IsMarked(w) {
				fmt.Printf("%d ==> ", w)
				bfs.Marked[w] = true
				q.Enqueue(w)
			}
		}
		fmt.Printf("Done\n")
	}
}
