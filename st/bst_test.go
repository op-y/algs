package st

import (
    "testing"
)

func TestBST(t *testing.T) {
    t.Logf("=====>>>测试开始<<<=====")
    tree := NewBST()
    if tree == nil {
        t.Errorf("BST创建失败！")
    } else {
        t.Logf("BST创建成功")
    }
    tree.Out()
    t.Logf("=====>>>测试结束<<<=====")
}
