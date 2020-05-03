// Copyright (c) 2019, Benjamin Wang (benjamin_wang@aliyun.com). All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

package main

import (
	"fmt"

	"github.com/ahrtr/gocontainer/priorityqueue"
)

func priorityqueueExample1() {
	printFuncName()

	pq := priorityqueue.New()

	values := []string{"benjamin", "alice", "john", "tom", "bill"}

	for _, v := range values {
		pq.Add(v)
	}

	for _, v := range values {
		fmt.Printf("pq.Contains(%v) = %t\n", v, pq.Contains(v))
	}

	fmt.Printf("pq.Remove(john) = %t\n", pq.Remove("john"))

	for {
		if pq.Peek() == nil {
			break
		}

		fmt.Printf("pq.Peek() = %v\n", pq.Peek())
		fmt.Printf("pq.Poll() = %v\n", pq.Poll())
	}
}

// priorityqueueExample2 demos how to reverse order for the build-in data types.
func priorityqueueExample2() {
	printFuncName()

	pq := priorityqueue.Reverse(priorityqueue.New())

	values := []string{"benjamin", "alice", "john", "tom", "bill"}

	for _, v := range values {
		pq.Add(v)
	}

	for _, v := range values {
		fmt.Printf("pq.Contains(%v) = %t\n", v, pq.Contains(v))
	}

	fmt.Printf("pq.Remove(john) = %t\n", pq.Remove("john"))

	for {
		if pq.Peek() == nil {
			break
		}

		fmt.Printf("pq.Peek() = %v\n", pq.Peek())
		fmt.Printf("pq.Poll() = %v\n", pq.Poll())
	}
}

// priorityqueueExample3 demos how to reverse order for the build-in data types using a comparator.
func priorityqueueExample3() {
	printFuncName()

	pq := priorityqueue.NewWithComparator(&reverseString{})

	values := []string{"benjamin", "alice", "john", "tom", "bill"}

	for _, v := range values {
		pq.Add(v)
	}

	for _, v := range values {
		fmt.Printf("pq.Contains(%v) = %t\n", v, pq.Contains(v))
	}

	fmt.Printf("pq.Remove(john) = %t\n", pq.Remove("john"))

	for {
		if pq.Peek() == nil {
			break
		}

		fmt.Printf("pq.Peek() = %v\n", pq.Peek())
		fmt.Printf("pq.Poll() = %v\n", pq.Poll())
	}
}

// priorityqueueExample3 demos how to order the customized data types (struct).
func priorityqueueExample4() {
	printFuncName()

	pq := priorityqueue.NewWithComparator(&student{})

	values := []student{
		{name: "benjamin", age: 28},
		{name: "alice", age: 42},
		{name: "john", age: 35},
		{name: "tom", age: 18},
		{name: "bill", age: 25},
	}

	for _, v := range values {
		pq.Add(v)
	}

	for _, v := range values {
		fmt.Printf("pq.Contains(%v) = %t\n", v, pq.Contains(v))
	}

	for {
		if pq.Peek() == nil {
			break
		}

		fmt.Printf("pq.Peek() = %v\n", pq.Peek())
		fmt.Printf("pq.Poll() = %v\n", pq.Poll())
	}
}
