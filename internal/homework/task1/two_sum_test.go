package task1

import (
  "testing"
)

func TestTwoSum(t *testing.T) {
  res := twoSum([]int{2, 3, 5, 6, 9}, 12)
  if len(res) != 2 {
    t.Errorf("返回结果中应该有2个数字")
  }
  if res[0] != 3 {
    t.Errorf("第1个数字应该是3")
  }
  if res[1] != 9 {
    t.Errorf("第2个数字应该是9")
  }
}
