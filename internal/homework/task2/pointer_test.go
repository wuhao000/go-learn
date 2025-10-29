package task2

import (
  "fmt"
  "testing"
)

func TestPointer(t *testing.T) {
  n := 1
  modifyNum(&n)
  fmt.Println(n)

  a := &[]int{1, 2, 3}
  scale(a)

  fmt.Println(a)
}
