// Copyright (c) 2019, Benjamin Wang (benjamin_wang@aliyun.com). All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

package main

import (
	"fmt"

	"github.com/ahrtr/gocontainer/queue"
)

func queueExample1() {
	q := queue.New()

	values := []string{"benjamin", "alice", "john", "tom", "bill"}

	for _, v := range values {
		q.Add(v)
	}

	for q.Peek() != nil {
		fmt.Printf("q.Peek() = %v\n", q.Peek())
		fmt.Printf("q.Poll() = %v\n", q.Poll())
	}
}
