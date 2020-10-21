package st

import (
    "fmt"
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
    size int // 直接记录符号表大小，空间换时间。
}

func NewSequentialSearchST() *SequentialSearchST {
    st := new(SequentialSearchST)
    st.first = nil
    st.size = 0
    return st
}

func (st *SequentialSearchST) Put(key IntKey, value interface{}) {
    node := st.first
    // 符号表为空
    if node == nil {
        n := NewNode(key, value, nil)
        st.first = n
        st.size = 1
        return
    }
    // 符号表不为空
    for node != nil {
		if 0 == node.key.CompareTo(key) {
            node.value = value
            return
        }
        if node.next == nil {
            n := NewNode(key, value, nil)
            node.next = n
            st.size += 1
            return
        }
        node = node.next
    }
    return
}

func (st *SequentialSearchST) Get(key IntKey) interface{} {
    for node := st.first; node != nil; node = node.next {
		if 0 == node.key.CompareTo(key) {
            return node.value
        }
    }
    return nil
}

func (st *SequentialSearchST) Delete(key IntKey) {
    node := st.first
    // 符号表为空
    if node == nil {
        return
    }
    fast := node.next
    // 删除首节点
	if 0 == node.key.CompareTo(key) {
        st.first = node.next
        st.size -= 1
        return
    }
    // 符号表不为空
    for fast != nil {
	    if 0 == fast.key.CompareTo(key) {
            node.next = fast.next
            st.size -= 1
            return
        }
        node = fast
        fast = fast.next
    }
    return
}

func (st *SequentialSearchST) Contains(key IntKey) bool {
    for node := st.first; node != nil; node = node.next {
		if 0 == node.key.CompareTo(key) {
            return true
        }
    }
    return false
}

func (st *SequentialSearchST) IsEmpty() bool {
    return st.size == 0
}

func (st *SequentialSearchST) Size() int {
    return st.size
}

// Go 语言中没有迭代器，这里返回一个可以用于 for...range 的切片
func (st *SequentialSearchST) Keys() []IntKey {
    if st.size == 0 {
        return nil
    }
    keys := make([]IntKey, st.size)
    i := 0
    for node := st.first; node != nil; node = node.next {
        keys[i] = node.key
        i++
    }
    return keys
}

func (st *SequentialSearchST) Out() {
    fmt.Printf("st size [%d] and item(s) is(are): ", st.size)
    for node := st.first; node != nil; node = node.next {
        fmt.Printf(" %v=>%v", node.key, node.value)
    }
    fmt.Printf("\n")
}
