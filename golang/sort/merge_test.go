package sort

import (
    "testing"
)

func TestTopDownMergeSort(t *testing.T) {
    N := 100
    data := NewVoidSlice(N)
    for i:=0; i < N; i++ {
        data.Set(i, N - i)
    }
    data.Out()
    lo := 0
    hi := 99
    TopDownMergeSort(data, lo, hi)
    data.Out()
    if data.IsSorted() {
        t.Logf("ok")
    } else {
        t.Errorf("fail")
    }
}

func TestBottomUpMergeSort(t *testing.T) {
    N := 100
    data := NewVoidSlice(N)
    for i:=0; i < N; i++ {
        data.Set(i, N - i)
    }
    data.Out()
    BottomUpMergeSort(data)
    data.Out()
    if data.IsSorted() {
        t.Logf("ok")
    } else {
        t.Errorf("fail")
    }
}
