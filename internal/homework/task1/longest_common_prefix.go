package task1

func longestCommonPrefix(strs []string) string {
  if len(strs) == 0 {
    return ""
  }
  prefix := strs[0]
  for _, s := range strs {
    if len(s) < len(prefix) {
      prefix = prefix[:len(s)]
    }
    for i := range min(len(s), len(prefix)) {
      if s[i] != prefix[i] {
        prefix = prefix[:i]
        break
      }
    }
  }
  return prefix
}
