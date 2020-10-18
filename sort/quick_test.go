package sort

import (
    "testing"
)

func TestThreeWayQuickSort(t *testing.T) {
    N := 10
    data := NewVoidSlice(N)
    data.Set(0, 49)
    data.Set(1, 56)
    data.Set(2, 9)
    data.Set(3, 56)
    data.Set(4, 77)
    data.Set(5, 9)
    data.Set(6, 56)
    data.Set(7, 12)
    data.Set(8, 22)
    data.Set(9, 43)
    data.Out()
    lo := 0
    hi := 9
    ThreeWayQuickSort(data, lo, hi)
    data.Out()
    if data.IsSorted() {
        t.Logf("ok")
    } else {
        t.Errorf("fail")
    }
}

func TestQuickSort(t *testing.T) {
    N := 10
    data := NewVoidSlice(N)
    data.Set(0, 49)
    data.Set(1, 56)
    data.Set(2, 9)
    data.Set(3, 56)
    data.Set(4, 77)
    data.Set(5, 9)
    data.Set(6, 56)
    data.Set(7, 12)
    data.Set(8, 22)
    data.Set(9, 43)
    data.Out()
    lo := 0
    hi := 9
    QuickSort(data, lo, hi)
    data.Out()
    if data.IsSorted() {
        t.Logf("ok")
    } else {
        t.Errorf("fail")
    }
}
