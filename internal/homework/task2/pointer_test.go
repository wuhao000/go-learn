package task2

import (
  "testing"
)

func TestPointer(t *testing.T) {
  n := 1
  modifyNum(&n)
  if n != 11 {
    t.Error("n should be 11, but is ", n)
  }
}

func TestPointer2(t *testing.T) {
  a := &[]int{1, 2, 3}
  scale(a)
  if (*a)[0] != 2 || (*a)[1] != 4 || (*a)[2] != 6 {
    t.Error("a should be [2,4,6], but is", *a)
  }
}
