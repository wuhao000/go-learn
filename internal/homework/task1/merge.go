package task1

import (
  "sort"
)

func merge(intervals [][]int) [][]int {
  sort.Slice(intervals, func(i, j int) bool {
    return intervals[i][0] < intervals[j][0]
  })
  var i = 1
  for j := 1; j < len(intervals); j++ {
    pair := intervals[j]
    if intervals[i-1][1] >= pair[1] {
      continue
    } else if intervals[i-1][1] >= pair[0] {
      intervals[i-1][1] = pair[1]
    } else {
      intervals[i] = pair
      i++
    }
  }
  return intervals[:i]
}
