package task1

func twoSum(nums []int, target int) []int {
  var res [2]int
  var mp = make(map[int]int)
  for _, n := range nums {
    m := target - n
    value, exists := mp[n]
    if exists {
      res[0] = value
      res[1] = n
    } else {
      mp[m] = n
    }
  }
  return res[:]
}
