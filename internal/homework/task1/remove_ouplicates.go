package task1

func removeDuplicates(nums []int) int {
  i := 0
  for k := 1; k < len(nums); k++ {
    if nums[i] != nums[k] {
      nums[i+1] = nums[k]
      i++
    }
  }
  return i + 1
}
