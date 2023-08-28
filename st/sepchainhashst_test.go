package st

import (
    "testing"
)

func TestSeparateChainingHashST(t *testing.T) {
    t.Logf("=====>>>测试开始<<<=====")
    hst := NewSeparateChainingHashST(997)
    if hst == nil {
        t.Errorf("哈希表创建失败！")
    } else {
        t.Logf("哈希表创建成功")
    }
    hst.Out()

    t.Logf("key=1024 求hash值")
    hash := hst.Hash(IntKey(1024))
    t.Logf("hash: %d", hash)
    if hash == 27 {
        t.Logf("哈希值正确")
    } else {
        t.Errorf("哈希值错误")
    }

    t.Logf("添加 key=1024 value=\"Hello\"")
    hst.Put(IntKey(1024), "Hello")
    hst.Out()

    t.Logf("获取 key=1024 的value")
    item := hst.Get(IntKey(1024))
    t.Logf("value: %v", item)

    t.Logf("=====>>>测试结束<<<=====")
}
