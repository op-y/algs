package bag

import (
    "fmt"
)

type Node struct {
   Value interface{}
   Next *Node
}

type Bag struct {
    First *Node
    Size int
}

func NewBag() *Bag {
    return &Bag{First: nil, Size: 0}
}

func (b *Bag) IsEmpty() bool {
    return b.First == nil
}

func (b *Bag) Count() int {
    return b.Size
}

func (b *Bag) Add(item interface{}) {
    pre_first := b.First
    b.First = new(Node)
    b.First.Value = item
    b.First.Next = pre_first
    b.Size += 1
}

func (b *Bag) Out() {
    fmt.Printf("bag current size is [%d], and item(s) is(are):", b.Size)
   for i := b.First; i != nil ; i = i.Next {
       fmt.Printf(" %v", i.Value)
   }
   fmt.Printf("\n")
}
