package task1

// singleNumber 返回切片中唯一不重复的数字
func singleNumber(nums []int) int {
  var m = make(map[int]int)
  for _, e := range nums {
    m[e] = m[e] + 1
  }
  for k, v := range m {
    if v == 1 {
      return k
    }
  }
  return -1
}
