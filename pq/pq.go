package pq

import (
    "errors"
    "fmt"
    "math"
)

var TypeErr = errors.New("不同类型数据不能比较！")

type MaxPQ struct {
	items    []interface{}
	capacity int
	N        int
}

func NewMaxPQ(capacity int) *MaxPQ {
	pq := new(MaxPQ)
	// 为了方便操作0索引位置空置
	q := make([]interface{}, capacity+1)
	pq.items = q
	pq.capacity = capacity
	pq.N = 0
	return pq
}

func (pq *MaxPQ) less(i, j int) bool {
	// 这里可以加上越界判断
	if vi, oki := pq.items[i].(int); oki {
		if vj, okj := pq.items[j].(int); okj {
			if vi < vj {
				return true
			} else {
				return false
			}
		}
	}
	// 类型不同的情况下返回值没有意义
	fmt.Printf("%s \n", TypeErr)
	return false
}

func (pq *MaxPQ) swap(i, j int) {
	tmp := pq.items[i]
	pq.items[i] = pq.items[j]
	pq.items[j] = tmp
}

func (pq *MaxPQ) swim(k int) {
	for k > 1 && pq.less(k/2, k) {
		pq.swap(k/2, k)
		k = k / 2
	}
}

func (pq *MaxPQ) sink(k int) {
	for 2*k <= pq.N {
		j := 2 * k
		if j < pq.N && pq.less(j, j+1) {
			j++
		}
		if !pq.less(k, j) {
			break
		}
		pq.swap(k, j)
		k = j
	}
}

func (pq *MaxPQ) IsEmpty() bool {
	return pq.N == 0
}

func (pq *MaxPQ) IsFull() bool {
	return pq.N == pq.capacity
}

func (pq *MaxPQ) Size() int {
	return pq.N
}

func (pq *MaxPQ) Insert(v interface{}) {
	if pq.IsFull() {
		fmt.Println("the queue is full!")
		return
	}
	pq.items[pq.N+1] = v
	pq.N++
	pq.swim(pq.N)
}

func (pq *MaxPQ) DelMax() interface{} {
	if pq.IsEmpty() {
		fmt.Println("the queue is empty!")
		return nil
	}
	max := pq.items[1]
	pq.swap(1, pq.N)
	pq.items[pq.N] = nil
	pq.N--
	pq.sink(1)
	return max
}

func (pq *MaxPQ) Out() {
    fmt.Printf("current queue size is [%d], and item(s) is(are):", pq.N)
    for i := 1; i <= pq.N ; i++ {
        fmt.Printf(" %v", pq.items[i])
    }
    fmt.Printf("\n")
}

// Pretty() 是为了按照树状结构输出，便于观察，还需要更多调试。
func (pq *MaxPQ) Pretty() {
    layerN := int(math.Ceil(math.Log2(float64(pq.N) + 1)))
    width := int(math.Pow(float64(2), float64(layerN))) - 1
    fmt.Printf("current queue size is [%d], layerN is [%d], width is [%d]: \n", pq.N, layerN, width)
    fmt.Printf("\n")
    gap := 0
    k := 1
    for i := 1; i <= layerN ; i++ {
        margin := int(math.Pow(float64(2), float64(layerN - i))) - 1
        itemNLayer := int(math.Pow(float64(2), float64(i - 1)))
        for m := 0; m < margin; m++ {
            fmt.Printf("    ")
        }
        for t := 1; t <= itemNLayer; t++ {
            fmt.Printf("%4v", pq.items[k])
            k++
            if t != itemNLayer {
                for g := 0; g < gap; g++ {
                    fmt.Printf("    ")
                }
            }
        }
        for m := 0; m < margin; m++ {
            fmt.Printf("    ")
        }
        fmt.Printf("\n")
        gap = margin
    }
    fmt.Printf("\n")
}
