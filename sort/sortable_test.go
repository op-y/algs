package sort

import (
    "testing"
)

func TestNewVoidSlice(t *testing.T) {
    N := 9
    data := NewVoidSlice(N)
    data.Set(0, 1)
    data.Set(1, 2)
    data.Set(2, 3)
    data.Set(3, 4)
    data.Set(4, 5)
    data.Set(5, 6)
    data.Set(6, 7)
    data.Set(7, 8)
    data.Set(8, 9)
    t.Logf("0：%v", data.Get(0))
    t.Logf("1：%v", data.Get(1))
    t.Logf("2：%v", data.Get(2))
    t.Logf("3：%v", data.Get(3))
    t.Logf("4：%v", data.Get(4))
    t.Logf("5：%v", data.Get(5))
    t.Logf("6：%v", data.Get(6))
    t.Logf("7：%v", data.Get(7))
    t.Logf("8：%v", data.Get(8))
    t.Logf("长度：%v", data.Len())
    data.Swap(0, 8)
    data.Out()
    isLess := data.Less(0, 1)
    t.Logf("Is data[0] less data[1]: %t", isLess)
    isSorted := data.IsSorted()
    t.Logf("Is sorted: %t", isSorted)
    data.Swap(0, 8)
    data.Out()
    isSorted = data.IsSorted()
    t.Logf("Is sorted: %t", isSorted)
    replica := data.Copy()
    replica.Out()
    t.Logf("data %p and replica %p: ", data, replica)
    t.Logf("ok")
}
