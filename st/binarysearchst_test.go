package st

import (
    "testing"
)

func TestBinarySearchST(t *testing.T) {
    t.Logf("=====>>>测试开始<<<=====")
    st := NewBinarySearchST()
    if st == nil {
        t.Errorf("顺序查找表创建失败！")
    } else {
        t.Logf("顺序查找表创建成功")
    }
    st.Out()
    t.Logf("=====>>>测试结束<<<=====")
}
