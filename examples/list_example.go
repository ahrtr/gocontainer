package main

import (
	"fmt"
	"sort"

	"github.com/ahrtr/gocontainer/list"
)

/*
	examples for ArrayList
*/
func listArrayListExample1() {
	printFuncName()
	al := list.NewArrayList()
	listBasicOperation(al)
}

func listArrayListExample2() {
	printFuncName()
	al := list.NewArrayList()
	listSortData(al)
}

func listArrayListExample3() {
	printFuncName()
	al := list.NewArrayList()
	listReverseSortData(al)
}

func listArrayListExample4() {
	printFuncName()
	al := list.NewArrayListWithComparator(&student{})
	listCustomizedData(al)
}

/*
	examples for LinkedList
*/
func listLinkedListExample1() {
	printFuncName()
	al := list.NewLinkedList()
	listBasicOperation(al)
}

func listLinkedListExample2() {
	printFuncName()
	al := list.NewLinkedList()
	listSortData(al)
}

func listLinkedListExample3() {
	printFuncName()
	al := list.NewLinkedList()
	listReverseSortData(al)
}

func listLinkedListExample4() {
	printFuncName()
	al := list.NewLinkedListWithComparator(&student{})
	listCustomizedData(al)
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

	// Iterator all the elements (method 1)
	fmt.Println("Iterate (method 1): ")
	for i := 0; i < h.Len(); i++ {
		v, _ := h.Get(i)
		fmt.Printf("    Index: %d, value: %v\n", i, v)
	}

	// Iterator all the elements (method 2)
	fmt.Println("Iterate (method 2): ")
	it, hasNext := h.Iterator()
	var v interface{}
	for hasNext {
		v, hasNext = it()
		fmt.Printf("    Value: %v\n", v)
	}

	// Iterator all the elements (method 3: reverse iterator)
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
	sort.Sort(h)

	// Iterator all the elements (method 1)
	fmt.Println("Iterate (method 1): ")
	for i := 0; i < h.Len(); i++ {
		v, _ := h.Get(i)
		fmt.Printf("    Index: %d, value: %v\n", i, v)
	}

	// Iterator all the elements (method 2)
	fmt.Println("Iterate (method 2): ")
	it, hasNext := h.Iterator()
	var v interface{}
	for hasNext {
		v, hasNext = it()
		fmt.Printf("    Value: %v\n", v)
	}

	// Iterator all the elements (method 3: reverse iterator)
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
	sort.Sort(sort.Reverse(h))

	// Iterator all the elements (method 1)
	fmt.Println("Iterate (method 1): ")
	for i := 0; i < h.Len(); i++ {
		v, _ := h.Get(i)
		fmt.Printf("    Index: %d, value: %v\n", i, v)
	}

	// Iterator all the elements (method 2)
	fmt.Println("Iterate (method 2): ")
	it, hasNext := h.Iterator()
	var v interface{}
	for hasNext {
		v, hasNext = it()
		fmt.Printf("    Value: %v\n", v)
	}

	// Iterator all the elements (method 3: reverse iterator)
	fmt.Println("Reverse iterate (method 2): ")
	it, hasPrev := h.ReverseIterator()
	for hasPrev {
		v, hasPrev = it()
		fmt.Printf("    Value: %v\n", v)
	}
}

func listCustomizedData(h list.Interface) {
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
	sort.Sort(h)

	// Iterator all the elements (method 1)
	fmt.Println("Iterate (method 1): ")
	for i := 0; i < h.Len(); i++ {
		v, _ := h.Get(i)
		fmt.Printf("    Index: %d, value: %v\n", i, v)
	}

	// Iterator all the elements (method 2)
	fmt.Println("Iterate (method 2): ")
	it, hasNext := h.Iterator()
	var v interface{}
	for hasNext {
		v, hasNext = it()
		fmt.Printf("    Value: %v\n", v)
	}

	// Iterator all the elements (method 3: reverse iterator)
	fmt.Println("Reverse iterate (method 2): ")
	it, hasPrev := h.ReverseIterator()
	for hasPrev {
		v, hasPrev = it()
		fmt.Printf("    Value: %v\n", v)
	}
}
