package main

import (
	"fmt"

	stack "github.com/truongnhatanh7/go-adt/types/stack"
)

func main() {
	s := stack.Stack{}
  
	s.Push(1)
	s.Push(2)
  
  for !s.IsEmpty() {
    fmt.Println(s.Pop())
  }

}
