package st

import (
    "testing"
)

func TestLinearProbeHashST(t *testing.T) {
    t.Logf("=====>>>测试开始<<<=====")
    hst := NewLinearProbeHashST(13)
    if hst == nil {
        t.Errorf("哈希表创建失败！")
    } else {
        t.Logf("哈希表创建成功")
    }
    hst.Out()

    t.Logf("key=hello 求hash值")
    hash := hst.Hash(StringKey("hello"))
    t.Logf("hash: %d", hash)

    t.Logf("添加 key=hello value=111")
    hst.Put(StringKey("hello"), 111)
    hst.Out()

    t.Logf("获取 key=hello 的value")
    item := hst.Get(StringKey("hello"))
    t.Logf("value: %v", item)

    t.Logf("=====>>>测试结束<<<=====")
}
