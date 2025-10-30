package task2

import (
  "fmt"
  "testing"
  "time"
)

func TestSendAndReceive(t *testing.T) {

  var ch chan int
  ch = make(chan int)

  go send(ch, 10)
  go receive(ch)

  timer := time.After(2 * time.Second)
  for {
    select {
    case <-timer:
      fmt.Println("end")
      return
    }
  }

}

func TestSendAndReceiveWithBuffer(t *testing.T) {

  var ch chan int
  ch = make(chan int, 5)

  go send(ch, 100)
  go receive(ch)

  timer := time.After(2 * time.Second)
  for {
    select {
    case <-timer:
      fmt.Println("end")
      return
    }
  }

}
