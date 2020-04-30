package main

import (
	"fmt"

	"github.com/ahrtr/gocontainer/stack"
)

func stackExample1() {
	printFuncName()
	s := stack.New()

	values := []int{5, 6, 7}
	for _, v := range values {
		s.Push(v)
	}

	for s.Len() > 0 {
		fmt.Printf("s.Pop() = %v\n", s.Pop())
	}
}
