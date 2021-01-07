package sort

import (
    "testing"
)

func TestBubbleSort(t *testing.T) {
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
    BubbleSort(data)
    data.Out()
    if data.IsSorted() {
        t.Logf("ok")
    } else {
        t.Errorf("fail")
    }
}
