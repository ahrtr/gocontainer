// Copyright (c) 2019, Benjamin Wang (benjamin_wang@aliyun.com). All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

package main

import (
	"fmt"

	"github.com/ahrtr/gocontainer/map/linkedmap"
)

func linkedmapExample1() {
	lm := linkedmap.New()

	keys := []interface{}{24, 43, 18, 23, 35}
	values := []interface{}{"benjamin", "alice", "john", "tom", "bill"}
	for i := 0; i < len(keys); i++ {
		lm.Put(keys[i], values[i])
	}

	for _, k := range keys {
		fmt.Printf("Get(%v) = %v\n", k, lm.Get(k))
	}

	v, _ := lm.Remove(18)
	fmt.Printf("The value associated with 18 is %v\n", v)

	k, v, _ := lm.RemoveFirstElement()
	fmt.Printf("The first element removed is (%v, %v)\n", k, v)

	k, v, _ = lm.RemoveLastElement()
	fmt.Printf("The last element removed is (%v, %v)\n", k, v)
}

// example for accessOrder
func linkedmapExample2() {
	lm := linkedmap.New().WithAccessOrder(true)

	keys := []interface{}{24, 43, 18, 23, 35}
	values := []interface{}{"benjamin", "alice", "john", "tom", "bill"}
	for i := 0; i < len(keys); i++ {
		lm.Put(keys[i], values[i])
	}

	lm.Get(23)
	lm.Get(24)
	lm.Get(35)
	lm.Get(18)
	lm.Get(43)

	// The first element should be (23, "tom")
	k, v, _ := lm.RemoveFirstElement()
	fmt.Printf("The first element removed is (%v, %v)\n", k, v)

	// The last element should be (43, "alice")
	k, v, _ = lm.RemoveLastElement()
	fmt.Printf("The last element removed is (%v, %v)\n", k, v)
}

// example for iterator & reverse iterator
func linkedmapExample3() {
	lm := linkedmap.New()

	keys := []interface{}{24, 43, 18, 23, 35}
	values := []interface{}{"benjamin", "alice", "john", "tom", "bill"}
	for i := 0; i < len(keys); i++ {
		lm.Put(keys[i], values[i])
	}

	fmt.Println("Iterating elements......")
	it, hasNext := lm.Iterator()
	var k, v interface{}
	for hasNext {
		k, v, hasNext = it()
		fmt.Printf("    Iterate element, key: %v, value: %v\n", k, v)
	}

	fmt.Println("Reverse iterating elements......")
	it, hasPrev := lm.ReverseIterator()
	for hasPrev {
		k, v, hasPrev = it()
		fmt.Printf("    Iterate element, key: %v, value: %v\n", k, v)
	}
}
