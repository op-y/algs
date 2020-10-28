package st

import (
    "testing"
)

func TestSequentialSearchST(t *testing.T) {
    t.Logf("=====>>>测试开始<<<=====")
    st := NewSequentialSearchST()
    if st == nil {
        t.Errorf("符号表创建失败！")
    } else {
        t.Logf("符号表创建成功")
    }
    st.Out()

    t.Logf("空符号表情况下删除节点 key=666")
    st.Delete(IntKey(666))
    st.Out()
    t.Logf("空符号表情况下获取节点 key=666 的value")
    item := st.Get(IntKey(666))
    t.Logf("value: %v", item)
    st.Out()

    t.Logf("空符号表情况下添加节点 key=8 value=\"Hello\"")
    st.Put(IntKey(8), "Hello")
    st.Out()
    t.Logf("获取这个节点 key=8 的value")
    item = st.Get(IntKey(8))
    t.Logf("value: %v", item)
    st.Out()
    t.Logf("然后删除节点 key=8")
    st.Delete(IntKey(8))
    st.Out()

    t.Logf("空符号表情况下添加节点 key=21 value=\"Hello\"")
    st.Put(IntKey(21), "Hello")
    st.Out()
    t.Logf("添加节点 key=22 value=\"my\"")
    st.Put(IntKey(22), "my")
    st.Out()
    t.Logf("添加节点 key=23 value=\"friends\"")
    st.Put(IntKey(23), "friend")
    st.Out()

    t.Logf("然后删除首节点 key=21")
    st.Delete(IntKey(21))
    st.Out()
    t.Logf("然后删除尾节点 key=23")
    st.Delete(IntKey(23))
    st.Out()

    t.Logf("=====>>>测试结束<<<=====")
}
