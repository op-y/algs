package sort

import (
	"github.com/op-y/algs/golang/pq"
)

func PQSort(a Sortable) {
	n := a.Len()
	q := pq.NewMaxPQ(n)
	for i := 0; i < n; i++ {
		q.Insert(a.Get(i))
	}
	for i := 0; i < n; i++ {
		a.Set(i, q.DelMax())
	}
}
