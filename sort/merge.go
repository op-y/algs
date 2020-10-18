package sort


func Merge(a Sortable, lo, mid, hi int) {
    i := lo
    j := mid + 1
    aux := a.Copy()
    for k := lo; k <= hi; k++ {
        if i > mid {
            a.Set(k, aux.Get(j))
            j += 1
        } else if j > hi {
            a.Set(k, aux.Get(i))
            i += 1
        } else if aux.Less(j, i) {
            a.Set(k, aux.Get(j))
            j += 1
        } else {
            a.Set(k, aux.Get(i))
            i += 1
        }
    }
}

func TopDownMergeSort(a Sortable, lo, hi int) {
   // 参数判断: 上界不能小于下界,等于情况下只有一个元素排序直接结束
   if hi <= lo {
       return
   }
   mid := lo + (hi - lo) / 2
   TopDownMergeSort(a, lo, mid)
   TopDownMergeSort(a, mid + 1, hi)
   Merge(a, lo, mid, hi)
}

func BottomUpMergeSort(a Sortable) {
    N := a.Len()
    for sz := 1; sz < N; sz += sz {
        for lo := 0; lo < N - sz; lo += sz + sz {
            mid := lo + sz - 1
            hi := lo + sz + sz - 1
            if hi > N - 1 {
                hi = N - 1
            }
            Merge(a, lo, mid, hi)
        }
    }
}
