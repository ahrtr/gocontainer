gocontainer
======
gocontainer implements some containers which exist in Java, but are missing in golang. This library is **zero dependency**, which means it does NOT depend on any 3rd party packages. Currently the containers are not thread-safe. 

# Table of Contents

- **[How to use this repo](#how-to-use-this-repo)**
- **[Common Interface](#Common-Interface)**
- **[Containers](#Containers)**
  - [Stack](#stack)
  - [Queue](#queue)
  - [Set](#set)
  - [List](#list)
  - [PriorityQueue](#priorityqueue)
  - [Others](#others)
- **[Sort](#sort)**
- **[Contribute to this repo](#contribute-to-this-repo)**
- **[Support](#support)**

# How to use this repo
It's very straightforward, just imports the containers you need and then use them directly. The following is an example for ArrayList, 
```
package main

import (
	"fmt"

	"github.com/ahrtr/gocontainer/list"
)

func main() {
	al := list.NewArrayList()

	al.Add(5)
	al.Add(6)
	al.Add(7)

	// Iterate all the elements 
	fmt.Println("Iterate (method 1): ")
	for i := 0; i < al.Len(); i++ {
		v, _ := al.Get(i)
		fmt.Printf("    Index: %d, value: %v\n", i, v)
	}
}
```

Please find more examples **[here](examples)**. 

# Common Interface
All containers in this repository implement interface **collection.Interface**,
```
// Interface is a type of collection, all containers should implement this interface.
type Interface interface {
	// IsEmpty returns true if this container contains no elements.
	IsEmpty() bool
	// Clear removes all of the elements from this container.
	Clear()

	// Len is the number of elements in the container.
	// Len() is also included in sort.Interface. Only golang 1.14 supports embedding of Interfaces with overlapping method sets,
	// so let's add it in this interface in the future.
	//Len() int
}
```

# Containers
Currently this library implements the following containers:
- Stack
- Queue
- Set
- List (ArrayList, LinkedList)
- PriorityQueue

## Stack
Stack is a LIFO(last-in-first-out) container. It implements the following interface. Click **[here](examples/stack_example.go)** to find examples on how to use a stack. 
```
// Interface is a stack, which is LIFO (last-in-first-out).
type Interface interface {
	collection.Interface

	// Len returns the length of this stack.
	Len() int
	// Push pushes an element into this stack.
	Push(val interface{})
	// Pop pops the element on the top of this stack.
	Pop() interface{}
}
```

## Queue
Queue is a FIFO(first-in-first-out) container. It implements the following interface. Click **[here](examples/queue_example.go)** to find examples on how to use a queue.
```
// Interface is a type of queue, which is FIFO(first-in-first-out).
type Interface interface {
	collection.Interface

	// Len returns the length of this queue.
	Len() int
	// Add inserts an element into the tail of this queue.
	Add(val interface{})
	// Peek retrieves but does not remove the head of this queue.
	Peek() interface{}
	// Poll retrieves and removes the head of the this queue.
	Poll() interface{}
}
```

## Set
A set contains no duplicate elements. It implements the following interface. Click **[here](examples/set_example.go)** to find examples on how to use a set. 
```
// Interface is a type of set, which contains no duplicate elements.
type Interface interface {
	collection.Interface

	// Len returns the length of this set.
	Len() int
	// Add adds the specified element to this set if it is not already present.
	// It returns false if the value is already present.
	Add(val interface{}) bool
	// Contains returns true if this set contains the specified element.
	Contains(val interface{}) bool
	// Remove removes the specified element from this set if it is present.
	// It returns false if the target value isn't present, otherwise returns true.
	Remove(val interface{}) bool
	// Iterate iterates all the elements in this set.
	Iterate(cb IterateCallback)
}
```

Applications are supposed to define a callback function (see below) when iterating a set. Please see the example on how to iterate a set. 
```
// IterateCallback is the signature of the callback function called by Iterate.
// If the callback function returns false, then the Iterate breaks.
type IterateCallback func(interface{}) bool
```

## List
This library implements two kinds of list, which are **ArrayList** and **LinkedList**, both of which implement the following interface. Click **[here](examples/list_example.go)** to find examples on how to use a list.
```
// Interface is a type of list, both ArrayList and LinkedList implement this interface.
type Interface interface {
	collection.Interface
	sort.Interface

	// Add appends the specified element to the end of this list.
	Add(val interface{})
	// AddTo inserts the specified element at the specified position in this list.
	AddTo(index int, val interface{}) error

	// Contains returns true if this list contains the specified element.
	Contains(val interface{}) bool
	// Get returns the element at the specified positon in this list. The index must be in the range of [0, size).
	Get(index int) (interface{}, error)

	// Remove removes the element at the specified position in this list.
	// It returns an error if the index is out of range.
	Remove(index int) (interface{}, error)
	// RemoveByValue removes the first occurence of the specified element from this list, if it is present.
	// It returns false if the target value isn't present, otherwise returns true.
	RemoveByValue(val interface{}) bool

	// Iterator returns an iterator over the elements in this list in proper sequence.
	Iterator() (func() (interface{}, bool), bool)
	// ReverseIterator returns an iterator over the elements in this list in reverse sequence as Iterator.
	ReverseIterator() (func() (interface{}, bool), bool)
}
```

The list.Interface has a nested sort.Interface, so a list can be sorted into ascending order, according to the natural ordering of its elements for some golang build-in data types, or sorted into a customized order, according to the comparator provided by applications. Please see [Sort](#sort) to get more detailed info.

## PriorityQueue
PriorityQueue is an unbounded priority queue based on a priority heap. It implements the following interface. Click **[here](examples/priorityqueue_example.go)** to find examples on how to use a priority queue.
```
// Interface is a type of priority queue, and PriorityQueue implement this interface.
type Interface interface {
	queue.Interface

	// Contains returns true if this queue contains the specified element.
	Contains(val interface{}) bool
	// Remove a single instance of the specified element from this queue, if it is present.
	// It returns false if the target value isn't present, otherwise returns true.
	Remove(val interface{}) bool
}
```

The elements of a PriorityQueue are ordered according to their natural ordering, or by a Comparator provided at PriorityQueue construction time. Please see [Sort](#sort) to get more detailed info.

If the reverse order for the elements is expected, then makes use of the priorityqueue.Reverse function, 
```
pq := priorityqueue.Reverse(priorityqueue.New())
```

## Others
More containers will be added soon. Please also kindly let me know if you need any other kinds of containers. Feel free to raise issues. 

# Sort
Some containers implement interface **sort.Interface**, such as ArrayList, LinkedList and PriorityQueue. For the following golang build-in data types, the elements can be ordered into ascending order according to their natural ordering. Note that for **bool**, a false is regarded as less than a true. 
- bool
- int
- int8
- int16
- int32
- int64
- uint
- uint8
- uint16
- uint32
- uint64
- float32
- float64
- string

Applications can also provide a sort.Comparator instance when constructing a container which implements sort.Interface. The rough logic should be something like below. Please find more examples in **[List](examples/list_example.go)** and **[PriorityQueue](examples/priorityqueue_example.go)**.
```
type MyComparator struct{}

func (c *MyComparator) Compare(v1, v2 interface{}) (int, error) {
    //......
}
```

# Contribute to this repo
Anyone is welcome to contribute to this repo. Please raise an issue firstly, then fork this repo and submit a pull request.

Currently this repo is under heavily development, any helps are appreciated! 

# Support
If you need any support, please raise issues. 

If you have any suggestions or proposals, please also raise issues. Thanks!

