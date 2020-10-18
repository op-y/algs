package queue

import (
    "fmt"
)

type Node struct {
   value interface{}
   next *Node
}

type Queue struct {
    first *Node
    last *Node
    size int
}

func NewQueue() *Queue {
    return &Queue{first: nil, last: nil, size: 0}
}

func (q *Queue) IsEmpty() bool {
    return q.first == nil
}

func (q *Queue) Count() int {
    return q.size
}

func (q *Queue) Enqueue(item interface{}) {
    pre_last := q.last
    q.last = new(Node)
    q.last.value = item
    q.last.next = nil
    if q.first != nil {
       pre_last.next = q.last
    } else {
       q.first = q.last
    }
    q.size += 1
}

func (q *Queue) Dequeue() interface{} {
    if 0 == q.size {
        return nil
    }
    item := q.first
    value := item.value
    q.first = q.first.next
    if q.first == nil {
        q.last = nil
    }
    q.size -= 1
    item = nil
    return value
}

func (q *Queue) Out() {
    fmt.Printf("queue current size is [%d], and item(s) is(are):", q.size)
    for i := q.first; i != nil ; i = i.next {
        fmt.Printf(" %v", i.value)
    }
    fmt.Printf("\n")
}
