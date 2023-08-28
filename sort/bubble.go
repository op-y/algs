package sort

func BubbleSort(a Sortable) {
	N := a.Len()
	for i := 1; i < N; i++ {
		for j := 0; j < N - i; j++ {
			if a.Less(j + 1, j) {
		        a.Swap(j, j + 1)
			}
		}
	}
}
