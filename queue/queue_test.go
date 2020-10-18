package queue

import (
    "testing"
)

func TestQueue(t *testing.T) {
    t.Logf("=====>>>测试开始<<<=====")
    q := NewQueue() // 不同于使用顺序结构的栈，这里是链式队列。
    if q == nil {
        t.Errorf("队列创建失败！")
    } else {
        t.Logf("队列创建成功")
    }
    t.Logf("Hello 入队..")
    q.Enqueue("Hello")
    t.Logf("2020 入队..")
    q.Enqueue(2020)
    t.Logf("空int切片入队..")
    q.Enqueue(make([]int, 1))
    q.Out()
    if q.Count() != 3 {
        t.Errorf("队列长度应该为3！")
    }
    value := q.Dequeue()
    if v, ok := value.(string); ok {
        if v == "Hello" {
            t.Logf("%s 出队，符合预期", v)
        } else {
            t.Errorf("%s 出队，不符合预期！", v)
        }
    }
    value = q.Dequeue()
    if v, ok := value.(int); ok {
        if v == 2020 {
            t.Logf("%d 出队，符合预期", v)
        } else {
            t.Errorf("%d 出队，不符合预期！", v)
        }
    }
    value = q.Dequeue()
    if v, ok := value.([]int); ok {
        t.Logf("%v 空切片出栈，符合预期", v)
    }
    if !q.IsEmpty() {
        t.Errorf("队列应该空了！")
    }
    q.Out()
    t.Logf("=====>>>测试结束<<<=====")
}
