package st

import (
    "fmt"
)

var (
    RED = true
    BLACK = false
)

type BRTNode struct {
    key StringKey
    value interface{}
    left *BRTNode
    right *BRTNode
    N int
    color bool
}

func NewBRTNode(key StringKey, value interface{}, n int, color bool) *BRTNode {
    node := new(BRTNode)
    node.key = key
    node.value = value
    node.left = nil
    node.right = nil
    node.N = 0
    node.color = color
    return node
}

func (node *BRTNode) IsEmpty() bool {
    if node == nil {
        return true
    }
    return node.N == 0
}

func (node *BRTNode) Size() int {
    if node == nil {
        return 0
    }
    return node.N
}

func (node *BRTNode) IsRed() bool {
    if node == nil {
        return false
    }
    return node.color == RED
}

func (node *BRTNode) RotateLeft() *BRTNode {
    if node == nil {
        return nil
    }
    x := node.left
    x.left = node
    x.color = node.color
    node.color = RED
    x.N = node.N
    node.N = node.left.Size() + node.right.Size() + 1
    return x
}

func (node *BRTNode) RotateRight() *BRTNode {
    if node == nil {
        return nil
    }
    x := node.right
    x.right = node
    x.color = node.color
    node.color = RED
    x.N = node.N
    node.N = node.left.Size() + node.right.Size() + 1
    return x
}

type BRT struct {
    root *BRTNode
}

func NewBRT() *BRT {
    brt := new(BRT)
    brt.root = nil
    return brt
}


func (brt *BRT) Out() {
    fmt.Println("Black Red Tree")
}
