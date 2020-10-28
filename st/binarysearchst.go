package st

import (
	"errors"
	"fmt"
)

var DEFAULT_LEN = 1

type BinarySearchST struct {
	keys   []IntKey
	values []interface{}
	size   int
}

func NewBinarySearchST() *BinarySearchST {
	st := new(BinarySearchST)
	keys := make([]IntKey, DEFAULT_LEN)
	values := make([]interface{}, DEFAULT_LEN)
	st.keys = keys
	st.values = values
	st.size = 0
	return st
}

func (st *BinarySearchST) resize(capacity int) {
	keys := make([]IntKey, capacity)
	values := make([]interface{}, capacity)
	for i := 0; i < st.size; i++ {
		keys[i] = st.keys[i]
		values[i] = st.values[i]
	}
	st.keys = keys
	st.values = values
}

func (st *BinarySearchST) IsEmpty() bool {
	return st.size == 0
}

// Rank是核心的方法
func (st *BinarySearchST) Rank(key IntKey) int {
	lo := 0
	hi := st.size - 1
	for lo <= hi {
		mid := lo + (lo+hi)/2
		if st.keys[mid].CompareTo(key) > 0 {
			hi = mid - 1
		} else if st.keys[mid].CompareTo(key) < 0 {
			lo = mid + 1
		} else {
			return mid
		}
	}
	return lo
}

func (st *BinarySearchST) Select(k int) (IntKey, error) {
	if st.size < k+1 {
		return NOT_FOUND_KEY, KeyNotFoundErr
	}
	return st.keys[k], nil
}

func (st *BinarySearchST) Min() (IntKey, error) {
	if st.IsEmpty() {
		return NOT_FOUND_KEY, KeyNotFoundErr
	}
	return st.keys[0], nil
}

func (st *BinarySearchST) Max() (IntKey, error) {
	if st.IsEmpty() {
		return NOT_FOUND_KEY, KeyNotFoundErr
	}
	return st.keys[st.size-1], nil
}

// 下边这些方法就需要仔细分析了

func (st *BinarySearchST) Contains(key IntKey) bool {
	// 这个空的判断是防御性的，Rank返回的最小值也应该是0，下同。
	if st.IsEmpty() {
		return false
	}
	i := st.Rank(key)
	if i < st.size && st.keys[i].CompareTo(key) == 0 {
		return true
	} else {
		return false
	}
}

func (st *BinarySearchST) Get(key IntKey) interface{} {
	if st.IsEmpty() {
		return nil
	}
	i := st.Rank(key)
	if i < st.size && st.keys[i].CompareTo(key) == 0 {
		return st.values[i]
	} else {
		return nil
	}
}

func (st *BinarySearchST) Ceiling(key IntKey) (IntKey, error) {
	if st.IsEmpty() {
		return NOT_FOUND_KEY, KeyNotFoundErr
	}
	i := st.Rank(key)
	if i < st.size {
		// Rank返回最小值也是0，查找key小于符号表中所有key的也包含在里。
		return st.keys[i], nil
	} else {
		// 查找的key比所有在符号表中key都大的情况
		return NOT_FOUND_KEY, KeyNotFoundErr
	}
}

func (st *BinarySearchST) Floor(key IntKey) (IntKey, error) {
	if st.IsEmpty() {
		return NOT_FOUND_KEY, KeyNotFoundErr
	}
	i := st.Rank(key)
	if i > st.size {
		// 查找的key比所有在符号表中key都大的情况
		return st.keys[st.size-1], nil
	}
	// Rank返回结果不会比0小，所以要分情况考虑。
	if i < st.size && i > 0 {
		return st.keys[i], nil
	}
	// 查找的key正好是第0号位置
	if i == 0 && st.keys[i].CompareTo(key) == 0 {
		return st.keys[0], nil
	}
	// 查找的key比所有在符号表中key都小的情况
	return NOT_FOUND_KEY, KeyNotFoundErr
}

func (st *BinarySearchST) Put(key IntKey, value interface{}) {
	// 提前扩容
	N := len(st.keys)
	if st.size == len(st.keys) {
		st.resize(2 * N)
	}
	i := st.Rank(key)
	// key存在直接覆盖value
	if i < st.size && st.keys[i].CompareTo(key) == 0 {
		st.values[i] = value
		return
	}
	// key不存在插入合适的位置，size+1
	for j := st.size; j > i; j-- {
		st.keys[j] = st.keys[j-1]
		st.values[j] = st.keys[j-1]
	}
	st.keys[i] = key
	st.values[i] = value
	st.size += 1
	return
}

func (st *BinarySearchST) Delete(key IntKey) {
	if st.IsEmpty() {
		return
	}
	// 提前缩容
	N := len(st.keys)
	if st.size == len(st.keys)/4 {
		st.resize(N / 2)
	}
	i := st.Rank(key)
	// 要删除的key在符号表中
	if i < st.size && st.keys[i].CompareTo(key) == 0 {
		// 后边的元素往前移动
		for j := i + 1; j < st.size; j++ {
			st.keys[j-1] = st.keys[j]
			st.values[j-1] = st.keys[j]
		}
		// 最后一位置为nil，size-1
		// st.keys 用的整型封装不能直接置nil
		st.values[st.size-1] = nil
		st.size -= 1
		return
	}
	return
}

func (st *BinarySearchST) DeleteMin() {
	min, err := st.Min()
	if err == nil {
		st.Delete(min)
	}
}

func (st *BinarySearchST) DeleteMax() {
	max, err := st.Max()
	if err == nil {
		st.Delete(max)
	}
}

func (st *BinarySearchST) Size() int {
	return st.size
}

func (st *BinarySearchST) SizeRange(lo, hi IntKey) int {
	if hi.CompareTo(lo) < 0 {
		return 0
	} else if st.Contains(hi) {
		return st.Rank(hi) - st.Rank(lo) + 1
	} else {
		return st.Rank(hi) - st.Rank(lo)
	}
}

func (st *BinarySearchST) Keys() []IntKey {
	return st.keys
}

func (st *BinarySearchST) KeysRange(lo, hi IntKey) []IntKey {
	if hi.CompareTo(lo) < 0 {
		return nil
	} else if st.Contains(hi) {
		rl := st.Rank(lo)
		rh := st.Rank(hi)
		keys := make([]IntKey, rh-rl+1)
		for i := rl; i <= rh; i++ {
			keys[i-rl] = st.keys[i]
		}
		return keys
	} else if st.Contains(lo) {
		rl := st.Rank(lo)
		rh := st.Rank(hi)
		keys := make([]IntKey, rh-rl)
		for i := rl; i < st.size; i++ {
			keys[i-rl] = st.keys[i]
		}
		return keys
	} else {
		return nil
	}
}

func (st *BinarySearchST) Out() {
	fmt.Printf("binary search st size [%d], and item(s) is(are): ", st.size)
	for i := 0; i < st.size; i++ {
		fmt.Printf(" %v=>%v", st.keys[i], st.values[i])
	}
	fmt.Printf("\n")
}
