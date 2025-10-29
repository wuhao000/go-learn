package task1

import (
  "fmt"
  "testing"
)

func TestMerge(t *testing.T) {

  res := merge(
    [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
  )
  fmt.Println(res)

}
