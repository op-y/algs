package bag

import (
    "testing"
)

func TestBag(t *testing.T) {
    b := NewBag()
    if !b.IsEmpty() {
        t.Errorf("new bag must be empty!")
    }
    b.Add(1)
    b.Add(2)
    b.Add(3)
    if b.IsEmpty() {
        t.Errorf("now the bag has 3 items!")
    }
    t.Logf("bag size: %d", b.Count())
    b.Out()
}
