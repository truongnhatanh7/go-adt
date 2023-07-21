package main

import (
	"fmt"

	"github.com/truongnhatanh7/go-adt/queue"
)

func main() {
  q := queue.Queue{}

  q.Push(1)
  q.Push(2)
  q.Push(3)

  fmt.Println(q.Seek())

  for !q.IsEmpty() {
    fmt.Println(q.Pop())
  }

}