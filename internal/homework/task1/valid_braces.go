package task1

func isValid(s string) bool {
  var arr []int32
  for _, c := range s {
    if len(arr) == 0 {
      arr = append(arr, c)
      continue
    }
    last := arr[len(arr)-1]
    if (c == ')' && last == '(') || (c == '}' && last == '{') || (c == ']' && last == '[') {
      arr = arr[:len(arr)-1]
    } else {
      arr = append(arr, c)
    }
  }
  return len(arr) == 0
}
