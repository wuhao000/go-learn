package task2

import (
  "sync"
  "sync/atomic"
  "time"
)

//题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
//考察点 ： sync.Mutex 的使用、并发数据安全。
//题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
//考察点 ：原子操作、并发数据安全。

type Counter struct {
  lock    sync.Mutex
  counter int
}

func count() int {

  counter := Counter{}

  for i := 0; i < 10; i++ {
    go func() {

      for j := 0; j < 100; j++ {
        counter.lock.Lock()
        counter.counter++
        counter.lock.Unlock()
      }

    }()
  }
  time.Sleep(100 * time.Millisecond)
  return counter.counter
}

func countUseAtomic() int64 {

  var counter int64 = 0

  for i := 0; i < 10; i++ {
    go func() {

      for j := 0; j < 100; j++ {
        atomic.AddInt64(&counter, 1)
      }

    }()
  }
  time.Sleep(100 * time.Millisecond)
  return counter
}
