package sort

import (
	"errors"
	"fmt"
)

var UnsupportedTypeErr = errors.New("目前示例只支持int类型数据！")

type Sortable interface {
	Len() int
	Less(int, int) bool
	Compare(int, interface{}) int
	Swap(int, int)
	IsSorted() bool
	Out() // 调试输出用的
	Get(int) interface{}
	Set(int, interface{})
	Copy() Sortable
	Reverse()
}

// Go 语言中切片本身就实现了排序接口
// 这里自定义类型由于后续的验证
type VoidSlice []interface{}

func NewVoidSlice(capacity int) VoidSlice {
	s := make([]interface{}, capacity)
	return VoidSlice(s)
}

func (vs VoidSlice) Len() int {
	return len(vs)
}

func (vs VoidSlice) Less(i, j int) bool {
	if vi, oki := vs[i].(int); oki {
		if vj, okj := vs[j].(int); okj {
			if vi < vj {
				return true
			} else {
				return false
			}
		}
	}
	// 类型不同的情况下返回值没有意义
	fmt.Printf("%s \n", UnsupportedTypeErr)
	return false
}

func (vs VoidSlice) Compare(i int, v interface{}) int {
	if vi, oki := vs[i].(int); oki {
		if vv, okv := v.(int); okv {
			if vi < vv {
				return -1
			} else if vi > vv {
				return 1
			} else {
				return 0
			}
		}
	}
	// 类型不同的情况下返回值没有意义
	fmt.Printf("%s \n", UnsupportedTypeErr)
	return 0
}

func (vs VoidSlice) Swap(i, j int) {
	tmp := vs[i]
	vs[i] = vs[j]
	vs[j] = tmp
}

func (vs VoidSlice) IsSorted() bool {
	c := vs.Len() - 1
	for i := 0; i < c; i++ {
		if vs.Less(i+1, i) {
			return false
		}
	}
	return true
}

func (vs VoidSlice) Out() {
	fmt.Println("current sequence: ")
	for _, v := range vs {
		fmt.Printf(" %v", v)
	}
	fmt.Println()
}

func (vs VoidSlice) Get(idx int) interface{} {
	return vs[idx]
}

func (vs VoidSlice) Set(idx int, value interface{}) {
	vs[idx] = value
}

func (vs VoidSlice) Copy() Sortable {
	capacity := vs.Len()
	ns := VoidSlice(make([]interface{}, capacity))
	for idx, v := range vs {
		ns[idx] = v
	}
	return ns
}

func (vs VoidSlice) Reverse() {
	i := 0
	j := vs.Len() - 1
	for i < j {
		vs.Swap(i, j)
		i++
		j--
	}
}
