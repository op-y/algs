package st

import (
    "fmt"
)

// Binary Search Tree Node 数据结构
type BSTNode struct {
    key StringKey
    value interface{}
    left *BSTNode
    right *BSTNode
    N int
}

func NewBSTNode(key StringKey, value interface{}, n int) *BSTNode {
    node := new(BSTNode)
    node.key = key
    node.value = value
    node.left = nil
    node.right = nil
    node.N = n
    return node
}

func (node *BSTNode) IsEmpty() bool {
    if node == nil {
        return true
    }
    // 根节点存在就不会为0
    return node.N == 0
}

func (node *BSTNode) Size() int {
    if node == nil {
        return 0
    }
    return node.N
}

func (node *BSTNode) Keys() []StringKey {
    if node == nil {
        return nil
    }
    result := make([]StringKey, 16) //这里省事取了一个默认值
    if node.left != nil {
        result =  append(result, node.left.Keys()...)
    }
    result =  append(result, node.key)
    if node.right != nil {
        result =  append(result, node.right.Keys()...)
    }
    return result
}

func (node *BSTNode) Select(k int) (StringKey, error) {
    if node == nil {
        return NOT_FOUND_STRING_KEY, KeyNotFoundErr
    }
    if node.left.Size() == k {
        return node.key, nil
    } else if node.left.Size() > k {
        return node.left.Select(k)
    } else {
        return node.right.Select(k - node.left.Size() -1)
    }
}

func (node *BSTNode) Rank(key StringKey) int {
    if node == nil {
        return 0
    }
    cmp := node.key.CompareTo(key)
    if cmp > 0 {
        return node.left.Rank(key)
    } else if cmp < 0 {
        return 1 + node.left.Size() + node.right.Rank(key)
    } else {
        return node.left.Size()
    }
}

func (node *BSTNode) Min() (*BSTNode, error) {
    if node != nil && node.left != nil {
        return node.left.Min()
    } else if node != nil {
        return node, nil
    } else {
        return nil, KeyNotFoundErr
    }
}

func (node *BSTNode) Max() (*BSTNode, error) {
    if node != nil && node.right != nil {
        return node.right.Max()
    } else if node != nil {
        return node, nil
    } else {
        return nil, KeyNotFoundErr
    }
}

func (node *BSTNode) Contains(key StringKey) bool {
    if node == nil  {
        return false
    }
    cmp := node.key.CompareTo(key)
    if cmp > 0 {
        return node.left.Contains(key)
    } else if cmp < 0 {
        return node.right.Contains(key)
    } else {
        return true
    }
}

func (node *BSTNode) Get(key  StringKey) interface{} {
    if node == nil {
        return nil
    }
    cmp := node.key.CompareTo(key)
    if cmp > 0 {
        return node.left.Get(key)
    } else if cmp < 0 {
        return node.right.Get(key)
    } else {
        return node.value
    }
}

func (node *BSTNode) Ceiling(key StringKey) (StringKey, error) {
    if node == nil {
        return NOT_FOUND_STRING_KEY, KeyNotFoundErr
    }
    cmp := node.key.CompareTo(key)
    if cmp == 0 {
        return node.key, nil
    }
    if cmp < 0 {
        return node.right.Ceiling(key)
    }
    k, err := node.left.Ceiling(key)
    if err != nil {
        return k, nil
    } else {
        return node.key, nil
    }
}

func (node *BSTNode) Floor(key StringKey) (StringKey, error) {
    if node == nil {
        return NOT_FOUND_STRING_KEY, KeyNotFoundErr
    }
    cmp := node.key.CompareTo(key)
    if cmp == 0 {
        return node.key, nil
    }
    if cmp > 0 {
        return node.left.Floor(key)
    }
    k, err := node.right.Floor(key)
    if err != nil {
        return k, nil
    } else {
        return node.key, nil
    }
}

func (node *BSTNode) Put(key StringKey, value interface{}) {
    if node == nil {
        node = NewBSTNode(key, value, 1)
        return
    }
    cmp := node.key.CompareTo(key)
    if cmp > 0 {
        node.left.Put(key, value)
    } else if cmp < 0 {
        node.right.Put(key, value)
    } else {
        node.value = value
    }
    node.N = node.left.Size() + node.right.Size() + 1
}

func (node *BSTNode) DeleteMin() *BSTNode {
    if node == nil {
        return nil
    }
    if node.left != nil {
        node.left = node.left.DeleteMin()
        node.N = node.left.Size() + node.right.Size() + 1
        return node
    } else {
        return node.right
    }
}

func (node *BSTNode) DeleteMax() *BSTNode {
    if node == nil {
        return nil
    }
    if node.right != nil {
        node.right = node.right.DeleteMax()
        node.N = node.left.Size() + node.right.Size() + 1
        return node
    } else {
        return node.left
    }
}

func (node *BSTNode) Delete(key StringKey) *BSTNode {
    if node == nil {
        return nil
    }
    cmp := node.key.CompareTo(key)
    if cmp > 0 {
        node.left.Delete(key)
    } else if cmp < 0 {
       node.right.Delete(key)
    } else {
        if node.right == nil {
            return node.left
        }
        if node.left == nil {
            return node.right
        }
        // 暂存当前node
        t := node
        // 右子树为nil的情况已经排除
        pn, _ := t.right.Min()
        node = pn
        node.right = t.right.DeleteMin()
        node.left = t.left
    }
    node.N = node.left.Size() + node.right.Size() + 1
    return node
}

func (node *BSTNode) SizeRange(lo, hi StringKey) int {
    if node == nil {
        return 0
    }
    s := 0
    cmplo := node.key.CompareTo(lo)
    cmphi := node.key.CompareTo(hi)
    if cmplo > 0 {
        s += node.left.SizeRange(lo, hi)
    }
    if cmplo >= 0 && cmphi <= 0 {
        s += 1
    }
    if cmphi < 0 {
        s += node.right.SizeRange(lo, hi)
    }
    return s
}

func (node *BSTNode) KeysRange(lo, hi StringKey) []StringKey {
    if node == nil {
        return nil
    }
    result := make([]StringKey, 16) //这里省事儿取了一个默认值
    cmplo := node.key.CompareTo(lo)
    cmphi := node.key.CompareTo(hi)
    if cmplo > 0 {
        result = append(result, node.left.KeysRange(lo, hi)...)
    }
    if cmplo >= 0 && cmphi <= 0 {
        result = append(result, node.key)
    }
    if cmphi < 0 {
        result = append(result, node.right.KeysRange(lo, hi)...)
    }
    return result
}

// Binary Search Tree 数据结构就是简单封装了BSTNode
type BST struct {
    root *BSTNode
}

func NewBST() *BST {
    bst := new(BST)
    bst.root = nil
    return bst
}

// 调用 BSTNode 的方法
func (bst *BST) IsEmpty() bool {
    return bst.root.IsEmpty()
}

func (bst *BST) Size() int {
    return bst.root.Size()
}

func (bst *BST) Keys() []StringKey {
    return bst.root.Keys()
}

func (bst *BST) Select(k int) (StringKey, error) {
    return bst.root.Select(k)
}

func (bst *BST) Rank(key StringKey) int {
    return bst.root.Rank(key)
}

func (bst *BST) Min() (*BSTNode, error) {
    return bst.root.Min()
}

func (bst *BST) Max() (*BSTNode, error) {
    return bst.root.Max()
}

func (bst *BST) Contains(key StringKey) bool {
    return bst.root.Contains(key)
}

func (bst *BST) Get(key StringKey) interface{} {
    return bst.root.Get(key)
}

func (bst *BST) Ceiling(key StringKey) (StringKey, error) {
    return bst.root.Ceiling(key)
}

func (bst *BST) Floor(key StringKey) (StringKey, error) {
    return bst.root.Floor(key)
}

func (bst *BST) Put(key StringKey, value interface{}) {
    bst.root.Put(key, value)
}

func (bst *BST) Delete(key StringKey) {
    bst.root.Delete(key)
}

func (bst *BST) DeleteMin() {
    bst.root.DeleteMin()
}

func (bst *BST) DeleteMax() {
    bst.root.DeleteMax()
}

func (bst *BST) SizeRange(lo, hi StringKey) int {
    return bst.root.SizeRange(lo, hi)
}

func (bst *BST) KeysRange(lo, hi StringKey) []StringKey {
    return bst.root.KeysRange(lo, hi)
}


func (bst *BST) Out() {
    fmt.Println()
}
