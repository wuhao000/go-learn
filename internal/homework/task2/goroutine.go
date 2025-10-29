package task2

import "fmt"

func printEven() {

  for i := 2; i <= 10; i++ {
    if i&1 == 0 {
      fmt.Println(i)
    }
  }

}

func printOdd() {

  for i := 1; i <= 10; i++ {
    if i&1 == 1 { // 奇数
      fmt.Println(i)
    }
  }

}

func goroutine() {

  go printEven()

  go printOdd()

}
