package sort

func InsertionSort(a Sortable) {
    N := a.Len()
    for i:=0; i < N; i++ {
        for j := i; j > 0 && a.Less(j, j - 1); j-- {
            a.Swap(j, j - 1)
        }
    }
}

func ShellSort(a Sortable) {
    N := a.Len()
    h := 1
    for h < N / 3 {
        h = h * 3 + 1
    }
    for h >= 1 {
        for i := h; i < N; i++ {
            for j := i; j >= h && a.Less(j, j - h); j -= h {
                a.Swap(j, j - h)
            }
        }
        h = h / 3
    }
}
