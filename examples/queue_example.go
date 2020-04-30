package main

import (
	"fmt"

	"github.com/ahrtr/gocontainer/queue"
)

func queueExample1() {
	printFuncName()
	q := queue.New()

	values := []string{"benjamin", "alice", "john", "tom", "bill"}

	for _, v := range values {
		q.Add(v)
	}

	for {
		if q.Peek() == nil {
			break
		}
		fmt.Printf("q.Peek() = %v\n", q.Peek())
		fmt.Printf("q.Poll() = %v\n", q.Poll())
	}
}
