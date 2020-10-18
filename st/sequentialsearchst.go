package st

import (
    //"fmt"
)

type Node struct {
   key IntKey
   value interface{}
   next *Node
}

func NewNode(key IntKey, value interface{}, next *Node) *Node {
    node := new(Node)
    node.key = key
    node.value = value
    node.next = next
    return node
}

type SequentialSearchST struct {
    first *Node
}
