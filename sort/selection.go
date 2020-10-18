package sort

func SelectionSort(a Sortable) {
	N := a.Len()
	for i := 0; i < N; i++ {
        min := i
		for j := i + 1; j < N; j++ {
			if a.Less(j, min) {
                min = j
			}
		}
		a.Swap(i, min)
	}
}
