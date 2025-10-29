package task1

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
  expect := "ab"
  prefix := longestCommonPrefix([]string{"abc", "abd"})
  if prefix != expect {
    t.Errorf("公共前缀应该为 %s", expect)
  }
}

func TestLongestCommonPrefix2(t *testing.T) {
  expect := "fl"
  prefix := longestCommonPrefix([]string{"flower", "flow", "flight"})
  if prefix != expect {
    t.Errorf("公共前缀应该为 %s", expect)
  }
}

func TestLongestCommonPrefix3(t *testing.T) {
  expect := "a"
  prefix := longestCommonPrefix([]string{"ab", "a"})
  if prefix != expect {
    t.Errorf("公共前缀应该为 %s, 运行结果为 %s", expect, prefix)
  }
}

func TestLongestCommonPrefix4(t *testing.T) {
  expect := ""
  prefix := longestCommonPrefix([]string{"reflower", "flow", "flight"})
  if prefix != expect {
    t.Errorf("公共前缀应该为 %s, 运行结果为 %s", expect, prefix)
  }
}
