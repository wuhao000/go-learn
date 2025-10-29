package task1

func plusOne(digits []int) []int {
  length := len(digits)
  b := false
  for i := length - 1; i >= 0; i-- {
    n := digits[i]
    if i == length-1 || b {
      if n == 9 {
        digits[i] = 0
        b = true
      } else {
        digits[i] = n + 1
        return digits
      }
    }
  }
  return append([]int{1}, digits...)
}
