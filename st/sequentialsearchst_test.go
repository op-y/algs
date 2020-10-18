package st

import (
    "testing"
)

func TestComparable(t *testing.T) {
    key := IntKey(100)
    n := NewNode(key, 666, nil)
    t.Logf("new node: %v", n)
}
