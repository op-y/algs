package st

import (
    "fmt"
)

type SeparateChainingHashST struct {
    N int
    M int
    st []*SequentialSearchST
}

func NewSeparateChainingHashST(m int) *SeparateChainingHashST {
    hst := new(SeparateChainingHashST)
    hst.N = 0
    hst.M = m
    hst.st = make([]*SequentialSearchST, m)
    for i := 0; i < m; i+=1 {
        hst.st[i] = NewSequentialSearchST()
    }
    return hst
}

// IntKey 底层类型是int，所以这里的hash函数直接用了取模
func (hst *SeparateChainingHashST) Hash(key IntKey) int {
    hash := int(key) % hst.M
    return hash
}

func (hst *SeparateChainingHashST) Get(key IntKey) interface{} {
    hash := hst.Hash(key)
    return hst.st[hash].Get(key)
}

func (hst *SeparateChainingHashST) Put(key IntKey, value interface{}) {
    hash := hst.Hash(key)
    hst.st[hash].Put(key, value)
    // N 变化处理
}

func (hst *SeparateChainingHashST) Delete(key IntKey) {
    hash := hst.Hash(key)
    hst.st[hash].Delete(key)
    // N 变化处理
}


func (hst *SeparateChainingHashST) Out() {
    fmt.Printf("Hash: M[%d] N[%d] and ST \n", hst.M, hst.N)
    for i := 0; i < hst.M; i+=1 {
        fmt.Printf("%v\n", hst.st[i])
    }
    fmt.Printf("\n")
}
