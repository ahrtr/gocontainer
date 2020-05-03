// Copyright (c) 2019, Benjamin Wang (benjamin_wang@aliyun.com). All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

package main

import (
	"fmt"

	"github.com/ahrtr/gocontainer/set"
)

func setExample1() {
	printFuncName()
	s := set.New()

	values := []int{5, 3, 9, 7, 6}
	for _, v := range values {
		s.Add(v)
	}

	for _, v := range values {
		fmt.Printf("s.Contains(%v) = %t\n", v, s.Contains(v))
	}

	// iterate all the elements, the callback function's signature:
	//   type IterateCallback func(interface{}) bool
	s.Iterate(func(v interface{}) bool {
		fmt.Printf("Iterate callback: %v\n", v)
		return true
	})

	s.Remove(6)
}
