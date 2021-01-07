package st

import (
    "testing"
)

func TestBRT(t *testing.T) {
    t.Logf("=====>>>测试开始<<<=====")
    tree := NewBRT()
    if tree == nil {
        t.Errorf("BRT创建失败！")
    } else {
        t.Logf("BRT创建成功")
    }
    tree.Out()
    t.Logf("=====>>>测试结束<<<=====")
}
