// Copyright (c) 2019, Benjamin Wang (benjamin_wang@aliyun.com). All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

package main

import (
	"fmt"

	"github.com/ahrtr/gocontainer/list"
	"github.com/ahrtr/gocontainer/utils"
)

/*
	examples for ArrayList
*/
func listArrayListExample1() {
	al := list.NewArrayList()
	listBasicOperation(al)
}

func listArrayListExample2() {
	al := list.NewArrayList()
	listSortData(al)
}

func listArrayListExample3() {
	al := list.NewArrayList()
	listReverseSortData(al)
}

func listArrayListExample4() {
	al := list.NewArrayList()
	listCustomizedData(al, &student{})
}

/*
	examples for LinkedList
*/
func listLinkedListExample1() {
	ll := list.NewLinkedList()
	listBasicOperation(ll)
}

func listLinkedListExample2() {
	ll := list.NewLinkedList()
	listSortData(ll)
}

func listLinkedListExample3() {
	ll := list.NewLinkedList()
	listReverseSortData(ll)
}

func listLinkedListExample4() {
	ll := list.NewLinkedList()
	listCustomizedData(ll, &student{})
}

/*
	common functions for both ArrayList and LinkedList
*/
func listBasicOperation(h list.Interface) {
	values := []int{5, 7, 12, 9}
	for _, v := range values {
		h.Add(v)
	}

	h.AddTo(2, 18)
	v3, _ := h.Remove(3)
	fmt.Printf("h.Remove(3) = %v\n", v3)

	// Iterate all the elements (method 1)
	fmt.Println("Iterate (method 1): ")
	for i := 0; i < h.Size(); i++ {
		v, _ := h.Get(i)
		fmt.Printf("    Index: %d, value: %v\n", i, v)
	}

	// Iterate all the elements (method 2)
	fmt.Println("Iterate (method 2): ")
	it, hasNext := h.Iterator()
	var v interface{}
	for hasNext {
		v, hasNext = it()
		fmt.Printf("    Value: %v\n", v)
	}

	// Iterate all the elements (method 3: reverse iterator)
	fmt.Println("Reverse iterate (method 2): ")
	it, hasPrev := h.ReverseIterator()
	for hasPrev {
		v, hasPrev = it()
		fmt.Printf("    Value: %v\n", v)
	}
}

func listSortData(h list.Interface) {
	values := []int{5, 7, 12, 9}
	for _, v := range values {
		h.Add(v)
	}

	fmt.Println("Sorting the data...")
	h.Sort()

	// Iterate all the elements (method 1)
	fmt.Println("Iterate (method 1): ")
	for i := 0; i < h.Size(); i++ {
		v, _ := h.Get(i)
		fmt.Printf("    Index: %d, value: %v\n", i, v)
	}

	// Iterate all the elements (method 2)
	fmt.Println("Iterate (method 2): ")
	it, hasNext := h.Iterator()
	var v interface{}
	for hasNext {
		v, hasNext = it()
		fmt.Printf("    Value: %v\n", v)
	}

	// Iterate all the elements (method 3: reverse iterator)
	fmt.Println("Reverse iterate (method 2): ")
	it, hasPrev := h.ReverseIterator()
	for hasPrev {
		v, hasPrev = it()
		fmt.Printf("    Value: %v\n", v)
	}
}

func listReverseSortData(h list.Interface) {
	values := []int{5, 7, 12, 9}
	for _, v := range values {
		h.Add(v)
	}

	fmt.Println("Reverse sorting the data...")
	h.SortWithOptions(true, nil)

	// Iterate all the elements (method 1)
	fmt.Println("Iterate (method 1): ")
	for i := 0; i < h.Size(); i++ {
		v, _ := h.Get(i)
		fmt.Printf("    Index: %d, value: %v\n", i, v)
	}

	// Iterate all the elements (method 2)
	fmt.Println("Iterate (method 2): ")
	it, hasNext := h.Iterator()
	var v interface{}
	for hasNext {
		v, hasNext = it()
		fmt.Printf("    Value: %v\n", v)
	}

	// Iterate all the elements (method 3: reverse iterator)
	fmt.Println("Reverse iterate (method 2): ")
	it, hasPrev := h.ReverseIterator()
	for hasPrev {
		v, hasPrev = it()
		fmt.Printf("    Value: %v\n", v)
	}
}

func listCustomizedData(h list.Interface, cmp utils.Comparator) {
	values := []student{
		{name: "benjamin", age: 28},
		{name: "alice", age: 42},
		{name: "john", age: 35},
		{name: "tom", age: 18},
		{name: "bill", age: 25},
	}

	for _, v := range values {
		h.Add(v)
	}

	fmt.Println("Sorting the data...")
	h.SortWithOptions(false, cmp)

	// Iterate all the elements (method 1)
	fmt.Println("Iterate (method 1): ")
	for i := 0; i < h.Size(); i++ {
		v, _ := h.Get(i)
		fmt.Printf("    Index: %d, value: %v\n", i, v)
	}

	// Iterate all the elements (method 2)
	fmt.Println("Iterate (method 2): ")
	it, hasNext := h.Iterator()
	var v interface{}
	for hasNext {
		v, hasNext = it()
		fmt.Printf("    Value: %v\n", v)
	}

	// Iterate all the elements (method 3: reverse iterator)
	fmt.Println("Reverse iterate (method 2): ")
	it, hasPrev := h.ReverseIterator()
	for hasPrev {
		v, hasPrev = it()
		fmt.Printf("    Value: %v\n", v)
	}
}
