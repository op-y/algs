package bag

import (
    "fmt"
)

type Node struct {
   value interface{}
   next *Node
}

type Bag struct {
    first *Node
    size int
}

func NewBag() *Bag {
    return &Bag{first: nil, size: 0}
}

func (b *Bag) IsEmpty() bool {
    return b.first == nil
}

func (b *Bag) Count() int {
    return b.size
}

func (b *Bag) Add(item interface{}) {
    pre_first := b.first
    b.first = new(Node)
    b.first.value = item
    b.first.next = pre_first
    b.size += 1
}

func (b *Bag) Out() {
    fmt.Printf("bag current size is [%d], and item(s) is(are):", b.size)
   for i := b.first; i != nil ; i = i.next {
       fmt.Printf(" %v", i.value)
   }
   fmt.Printf("\n")
}
