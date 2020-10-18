package sort

import (
    "fmt"
    // "time"
)

func Partition(a Sortable, lo, hi int) int {
	v := a.Get(lo)
	i := lo
	j := hi + 1
	for {
        //time.Sleep(1 * time.Second)
		for i++; a.Compare(i, v) < 0;i++ {
			if i == hi {
				break
			}
		}
		for j--; a.Compare(j, v) > 0;j-- {
			if j == lo {
				break
			}
		}
		if i >= j {
			break
		}
		a.Swap(i, j)
        a.Out()
	}
	a.Swap(lo, j)
    a.Out()
	return j
}

func QuickSort(a Sortable, lo, hi int) {
	if hi <= lo {
		return
	}
	j := Partition(a, lo, hi)
    fmt.Printf("partition j=%d LO: %d %d \n", j, lo, j - 1)
	QuickSort(a, lo, j-1)
    fmt.Printf("partition j=%d HI: %d %d \n", j, j + 1, hi)
	QuickSort(a, j+1, hi)
}

func ThreeWayQuickSort(a Sortable, lo, hi int) {
    if hi <= lo {
        return
    }
    v := a.Get(lo)
    lt := lo
    i := lo + 1
    gt := hi
    for i <= gt {
        cmp := a.Compare(i, v)
        if cmp < 0 {
            a.Swap(lt, i)
            lt++
            i++
        } else if cmp > 0 {
            a.Swap(gt, i)
            gt--
            // 注意这里i位置上换来的新元素比一定比v大，所以i不需要移动
        } else {
            i++
        }
    }
    ThreeWayQuickSort(a, lo, lt - 1)
    ThreeWayQuickSort(a, gt + 1, hi)
}
