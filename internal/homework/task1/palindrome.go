package task1

import (
  "fmt"
  "strconv"
)

// isPalindrome 判断一个数字是不是回文数
func isPalindrome(n int) bool {
  str := strconv.Itoa(n)
  for i := 0; i < len(str)/2; i++ {
    if str[i] != str[len(str)-i-1] {
      return false
    }
  }
  return true
}

func main() {
  fmt.Println(isPalindrome(1234))
  fmt.Println(isPalindrome(1234321))
}
