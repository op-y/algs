package stack

import (
	"fmt"
)

type FixedCapacityStack struct {
	items []interface{}
	size int
}

func NewFixedCapacityStack(capacity int) *FixedCapacityStack {
    e := make([]interface{}, capacity)
    s := new(FixedCapacityStack)
    s.size = 0
    s.items = e
    return s
}

func (s *FixedCapacityStack) resize(capacity int) {
    e := make([]interface{}, capacity)
    for i := 0; i < s.size; i++ {
        e[i] = s.items[i]
    }
    s.items = e
}

func (s *FixedCapacityStack) Push(item interface{}) {
    if s.size == len(s.items) {
        s.resize(2 * len(s.items))
    }
    s.items[s.size] = item
    s.size += 1
}

func (s *FixedCapacityStack) Pop() interface{} {
    s.size -= 1
    e := s.items[s.size]
    if s.size > 0 && s.size == len(s.items) / 4 {
        s.resize(len(s.items) / 2)
    }
    return e
}

func (s *FixedCapacityStack) IsEmpty() bool {
	return s.size == 0
}

func (s *FixedCapacityStack) Count() int {
	return s.size
}

// Go 语言中没有迭代器，这里设计Iterator方法返回一个可以 for...range 的切片
func (s *FixedCapacityStack) Iterator() []interface{} {
	return s.items
}

// Out() 方法打印栈信息用于调试
func (s *FixedCapacityStack) Out() {
    fmt.Printf("stack capacity is [%d], size is [%d], and item(s) is(are):", len(s.items), s.size)
    for i := 0; i < s.size; i++ {
        fmt.Printf(" %v", s.items[i])
    }
    fmt.Printf("\n")
}
