package task2

import "testing"

func TestCounter(t *testing.T) {
  count := count()
  if count != 1000 {
    t.Error("结果应该是 1000， 实际是 ", count)
  }
}

func TestCounterUseAtomic(t *testing.T) {
  count := countUseAtomic()
  if count != 1000 {
    t.Error("结果应该是 1000， 实际是 ", count)
  }
}
