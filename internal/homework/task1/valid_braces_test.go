package task1

import (
  "testing"
)

func TestIsValid(t *testing.T) {
  str := "()"
  res := isValid(str)
  if !res {
    t.Errorf("%s是有效的括号", str)
  }
}

func TestIsValid2(t *testing.T) {
  str := "(){}"
  res := isValid(str)
  if !res {
    t.Errorf("%s是有效的括号", str)
  }
}

func TestIsValid3(t *testing.T) {
  str := "({})"
  res := isValid(str)
  if !res {
    t.Errorf("%s是有效的括号", str)
  }
}

func TestIsValid4(t *testing.T) {
  str := "({})["
  res := isValid(str)
  if res {
    t.Errorf("%s不是有效的括号", str)
  }
}
