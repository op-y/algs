package st

import (
    "fmt"
)


type LinearProbeHashST struct {
    N int
    M int
    keys []StringKey
    values []interface{}
}

func NewLinearProbeHashST(m int) *LinearProbeHashST {
    hst := new(LinearProbeHashST)
    hst.N = 0
    hst.M = m
    hst.keys = make([]StringKey, m)
    hst.values = make([]interface{}, m)
    return hst
}

// StringKey 底层数据类型是string，这里用每个byte除留余数法来做hash
func (hst *LinearProbeHashST) Hash(key StringKey) int {
    R := 31
    hash := 0
    bs := []byte(string(key))
    for _, v := range bs {
        hash = (R * hash + int(v)) % hst.M
    }
    return hash
}

func (hst *LinearProbeHashST) Resize(m int) *LinearProbeHashST {
    newhst := NewLinearProbeHashST(m)
    for i := 0; i < hst.M; i++ {
        if string(hst.keys[i]) != "" {
            newhst.Put(hst.keys[i], hst.values[i])
        }
    }
    return newhst
}

func (hst *LinearProbeHashST) Put(key StringKey, value interface{}) {
    if hst.N >= hst.M / 2 {
        hst.Resize(2 * hst.M)
    }
    pos := hst.Hash(key)
    for string(hst.keys[pos]) != "" {
        if key.CompareTo(hst.keys[pos]) == 0 {
            hst.values[pos] = value
            return
        }
        pos = (pos + 1) % hst.M
    }
    hst.keys[pos] = key
    hst.values[pos] = value
    hst.N++
}

func (hst *LinearProbeHashST) Get(key StringKey) interface{} {
    for pos := hst.Hash(key); string(hst.keys[pos]) != ""; pos = (pos + 1) % hst.M {
        if key.CompareTo(hst.keys[pos]) == 0 {
            return hst.values[pos]
        }
    }
    return nil
}

func (hst *LinearProbeHashST) Delete(key StringKey) {
    pos := hst.Hash(key)
    for key.CompareTo(hst.keys[pos]) != 0 {
         pos = (pos + 1) % hst.M
    }
    hst.keys[pos] = ""
    hst.values[pos] = nil
    pos = (pos + 1) % hst.M
    for string(hst.keys[pos]) != "" {
         rekey := hst.keys[pos]
         revalue := hst.values[pos]
         hst.keys[pos] = ""
         hst.values[pos] = nil
         hst.N--
         hst.Put(rekey, revalue)
         pos = (pos + 1) % hst.M
    }
    hst.N--
    if (hst.N > 0 && hst.N == hst.M / 8) {
         hst.Resize(hst.M / 2)
    }
}

func (hst *LinearProbeHashST) Out() {
    fmt.Printf("Hash: M[%d] N[%d] and ST \n", hst.M, hst.N)
    for i := 0; i < hst.M; i+=1 {
        fmt.Printf("%v ==> %v\n", hst.keys[i], hst.values[i])
    }
    fmt.Printf("\n")
}
