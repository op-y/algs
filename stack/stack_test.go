package stack

import (
    "testing"
)

func TestStack(t *testing.T) {
    t.Logf("=====>>>测试开始<<<=====")
    N := 9 // 指定栈容量
    s := NewFixedCapacityStack(N)
    if s == nil {
        t.Errorf("创建定容栈失败!")
    } else {
        t.Logf("定容栈创建成功")
    }
    if s.Count() != 0 {
        t.Errorf("栈长度应该为0!")
    }
    t.Logf("A B C D E F G H I 依次入栈..")
    s.Push("A")
    s.Push("B")
    s.Push("C")
    s.Push("D")
    s.Push(1)
    s.Push(2)
    s.Push(3)
    s.Push("H")
    s.Push("I")
    if s.Count() != N {
        t.Errorf("栈长度应该为9！")
    }
    s.Out()
    t.Logf("J 入栈..")
    t.Logf("栈满，扩容2倍，注意观察..")
    s.Push("J")
    s.Out()
    item := s.Pop()
    if v, ok := item.(string); ok {
        if v == "J" {
            t.Logf("%s 出栈，符合预期", v)
        } else {
            t.Errorf("%s 出栈，不符合预期！", v)
        }
    }
    item = s.Pop()
    if v, ok := item.(string); ok {
        if v == "I" {
            t.Logf("%s 出栈，符合预期", v)
        } else {
            t.Errorf("%s 出栈，不符合预期！", v)
        }
    }
    item = s.Pop()
    if v, ok := item.(string); ok {
        if v == "H" {
            t.Logf("%s 出栈，符合预期", v)
        } else {
            t.Errorf("%s 出栈，不符合预期！", v)
        }
    }
    item = s.Pop()
    if v, ok := item.(int); ok {
        if v == 3 {
            t.Logf("%d 出栈，符合预期", v)
        } else {
            t.Errorf("%d 出栈，不符合预期！", v)
        }
    }
    item = s.Pop()
    if v, ok := item.(int); ok {
        if v == 2 {
            t.Logf("%d 出栈，符合预期", v)
        } else {
            t.Errorf("%d 出栈，不符合预期！", v)
        }
    }
    item = s.Pop()
    if v, ok := item.(int); ok {
        if v == 1 {
            t.Logf("%d 出栈，符合预期", v)
        } else {
            t.Errorf("%d 出栈，不符合预期！", v)
        }
    }
    t.Logf("栈内容不够1/4，扩容1/2，注意观察..")
    s.Out()
    t.Logf("=====>>>测试结束<<<=====")
}
