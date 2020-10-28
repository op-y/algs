package st

import (
    "testing"
)

func TestIntKey(t *testing.T) {
    t.Logf("=====>>>测试开始<<<=====")
    key1 := IntKey(98)
    key2 := IntKey(198)
    key3 := IntKey(298)
    if key2.CompareTo(key1) <= 0 {
        t.Errorf("192 比 98 大！")
    }
    if key2.CompareTo(key3) >= 0 {
        t.Errorf("192 比 298 小！")
    }
    if key2.CompareTo(key2) != 0 {
        t.Errorf("自己和自己比要相等！")
    }
    t.Logf("=====>>>测试结束<<<=====")
}

func TestStringKey(t *testing.T) {
    t.Logf("=====>>>测试开始<<<=====")
    key1 := StringKey("alpha")
    key2 := StringKey("beta")
    key3 := StringKey("gamma")
    if key2.CompareTo(key1) <= 0 {
        t.Errorf("beta 比 alpha 大！")
    }
    if key2.CompareTo(key3) >= 0 {
        t.Errorf("beta 比 gamma 小！")
    }
    if key2.CompareTo(key2) != 0 {
        t.Errorf("自己和自己比要相等！")
    }
    t.Logf("=====>>>测试结束<<<=====")
}
