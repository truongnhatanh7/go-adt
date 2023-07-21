package main

import (
	"fmt"

	stack "github.com/truongnhatanh7/go-adt/stack"
)

func main() {
	s := stack.Stack{}

	s.Push(1)
	s.Push(2)

  fmt.Println(s.Seek())
  
  for !s.IsEmpty() {
    fmt.Println(s.Pop())
  }

}
