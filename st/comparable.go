package st

import (
    "errors"
    "fmt"
)

var TypeErr = errors.New("不能比较不同类型数据！")
var KeyNotFoundErr = errors.New("没有查找到这个key！")

var NOT_FOUND_KEY = IntKey(404)
var NOT_FOUND_STRING_KEY = StringKey("404")

type Comparable interface {
	CompareTo(other interface{}) int
}

// int类型自身没有实现CompareTo函数，这里为了演示封装一下。
type IntKey int

// 实现 Comparable 接口
func (key IntKey) CompareTo(other interface{}) int {
    if o, ok := other.(IntKey); ok {
        if key > o {
            return 1
        } else if key < o {
            return -1
        } else {
            return 0
        }
    }
    fmt.Printf("failed to compare: %s\n", TypeErr.Error())
    return 255
}

type StringKey string

func (key StringKey) CompareTo(other interface{}) int {
    if o, ok := other.(StringKey); ok {
        if key > o {
            return 1
        } else if key < o {
            return -1
        } else {
            return 0
        }
    }
    fmt.Printf("failed to compare: %s\n", TypeErr.Error())
    return 255
}
