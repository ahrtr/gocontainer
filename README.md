gocontainer ([中文版](README_cn.md))
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
  - [LinkedMap](#linkedMap)
  - [Others](#others)
- **[Sort](#sort)**
- **[Contribute to this repo](#contribute-to-this-repo)**
- **[Support](#support)**

# How to use this repo
It's very straightforward, just imports the containers you need and then use them directly. The following is an example for ArrayList, 
```go
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
```go
// Interface is a type of collection, all containers should implement this interface.
type Interface interface {
	// Size returns the number of elements in the collection.
	Size() int
	// IsEmpty returns true if this container contains no elements.
	IsEmpty() bool
	// Clear removes all of the elements from this container.
	Clear()
}
```

# Containers
Currently this library implements the following containers:
- Stack
- Queue
- Set
- List (ArrayList, LinkedList)
- PriorityQueue
- LinkedMap

## Stack
Stack is a LIFO(last-in-first-out) container. It implements the following interface. Click **[here](examples/stack_example.go)** to find examples on how to use a stack. 
```go
// Interface is a stack, which is LIFO (last-in-first-out).
type Interface interface {
	collection.Interface

	// Push pushes an element into this stack.
	Push(val interface{})
	// Pop pops the element on the top of this stack.
	Pop() interface{}
}
```

Please import the following package in order to use stack,
```go
import (
	"github.com/ahrtr/gocontainer/stack"
)
```

Call stack.New() to create a stack,
```go
New() Interface
```

The following is a simple example for stack,
```go
package main

import (
	"fmt"

	"github.com/ahrtr/gocontainer/stack"
)

func main() {
	s := stack.New()

	values := []int{5, 6, 7}
	for _, v := range values {
		s.Push(v)
	}

	for s.Size() > 0 {
		fmt.Printf("s.Pop() = %v\n", s.Pop())
	}
}
```

## Queue
Queue is a FIFO(first-in-first-out) container. It implements the following interface. Click **[here](examples/queue_example.go)** to find examples on how to use a queue.
```go
// Interface is a type of queue, which is FIFO(first-in-first-out).
type Interface interface {
	collection.Interface

	// Add inserts an element into the tail of this queue.
	Add(val interface{})
	// Peek retrieves, but does not remove, the head of this queue, or return nil if this queue is empty.
	Peek() interface{}
	// Poll retrieves and removes the head of the this queue, or return nil if this queue is empty.
	Poll() interface{}
}
```

Please import the following package in order to use queue,
```go
import (
	"github.com/ahrtr/gocontainer/queue"
)
```

Call queue.New() to create a queue,
```go
New() Interface
```

The following is a simple example for queue,
```go
package main

import (
	"fmt"

	"github.com/ahrtr/gocontainer/queue"
)

func main() {
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
```

## Set
A set contains no duplicate elements. The values contained in a set may be any type that is comparable, please refer to the golang [language spec](https://golang.org/ref/spec#Comparison_operators) to get more detailed info on comparison operators. 

Set implements the following interface. Click **[here](examples/set_example.go)** to find examples on how to use a set. 
```go
// Interface is a type of set, which contains no duplicate elements.
type Interface interface {
	collection.Interface

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

Please import the following package in order to use set,
```go
import (
	"github.com/ahrtr/gocontainer/set"
)
```

Call set.New() to create a set,
```go
New() Interface
```

The following is a simple example for set,
```go
package main

import (
	"fmt"

	"github.com/ahrtr/gocontainer/set"
)

func main() {
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
```

Applications are supposed to define a callback function (see below) when iterating a set. 
```go
// IterateCallback is the signature of the callback function called by Iterate.
// If the callback function returns false, then the iteration breaks.
type IterateCallback func(interface{}) bool
```

The following snip shows how to iterate a set. Please see the **[example](examples/set_example.go)** to get more detailed info.
```go
// To iterate over a set (where s is an instance of set.Interface):
s.Iterate(func(v interface{}) bool {
	// Do something with v

	// If you want to break the iteration, then return a false
	return true
})
```

## List
This library implements two kinds of list, which are **ArrayList** and **LinkedList**, both of which implement the following interface. Click **[here](examples/list_example.go)** to find examples on how to use a list.
```go
// Interface is a type of list, both ArrayList and LinkedList implement this interface.
type Interface interface {
	collection.Interface
	sort.Interface

	// Add appends the specified element to the end of this list.
	Add(val interface{})
	// AddTo inserts the specified element at the specified position in this list.
	AddTo(index int, val interface{}) error

	// WithComparator sets a gsort.Comparator instance for the list.
	// It's used to imposes a total ordering on the elements in the list.
	WithComparator(c gsort.Comparator) Interface

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

Please import the following package in order to use list (arrayList or linkedList),
```go
import (
	"github.com/ahrtr/gocontainer/list"
)
```

Call list.NewArrayList() and list.NewLinkedList() to create a ArrayList and a LinkedList respectively, 
```go
NewArrayList() Interface
NewLinkedList() Interface
```

The following is a simple example for arrayList,
```go
package main

import (
	"fmt"

	"github.com/ahrtr/gocontainer/list"
)

func main() {
	al := list.NewArrayList()
	values := []int{5, 7, 12, 9}
	for _, v := range values {
		al.Add(v)
	}

	al.AddTo(2, 18)
	v3, _ := al.Remove(3)
	fmt.Printf("al.Remove(3) = %v\n", v3)

	// Iterate all the elements 
	fmt.Println("Iterate: ")
	for i := 0; i < al.Size(); i++ {
		v, _ := al.Get(i)
		fmt.Printf("    Index: %d, value: %v\n", i, v)
	}
}
```

The following is a simple example for linkedList,
```go
package main

import (
	"fmt"

	"github.com/ahrtr/gocontainer/list"
)
func main() {
	ll := list.NewLinkedList()
	values := []int{5, 7, 12, 9}
	for _, v := range values {
		ll.Add(v)
	}

	ll.AddTo(2, 18)
	v3, _ := ll.Remove(3)
	fmt.Printf("ll.Remove(3) = %v\n", v3)


	// Iterate all the elements
	fmt.Println("Iterate: ")
	it, hasNext := ll.Iterator()
	var v interface{}
	for hasNext {
		v, hasNext = it()
		fmt.Printf("    Value: %v\n", v)
	}
}
```

A sort.Comparator instance can be provided for a list (ArrayList or LinkedList) by method WithComparator, please get more detailed info in **[Sort](#sort)**.
```go
WithComparator(c gsort.Comparator) Interface 
```

The list.Interface has a nested sort.Interface, so a list can be sorted using sort.Sort(data) into ascending order, according to the natural ordering of its elements for some golang build-in data types, or sorted into a customized order, according to the comparator provided by applications. 

There are multiple ways to iterate a list. The following snips show how to iterate a list (arrayList or linkedList),
```go
// To iterate over a list (where l is an instance of list.Interface):
it, hasNext := l.Iterator()
var v interface{}
for hasNext {
	v, hasNext = it()
	// do something with v
}
```

```go
// To iterate over a list (where l is an instance of list.Interface):
for i:=0; i<l.Len(); i++ {
	v, _ := l.Get(i)
	// Do something with v
}
```

```go
// To iterate over a list in reverse order (where l is an instance of list.Interface):
it, hasPrev := l.ReverseIterator()
var v interface{}
for hasPrev {
	v, hasPrev = it()
	// do something with v
}
```

```go
// To iterate over a list in reverse order (where l is an instance of list.Interface):
for i:=l.Len()-1; i>=0; i-- {
	v, _ := l.Get(i)
	// Do something with v
}
```

## PriorityQueue
PriorityQueue is an unbounded priority queue based on a priority heap. It implements the following interface. Click **[here](examples/priorityqueue_example.go)** to find examples on how to use a priority queue.
```go
// Interface is a type of priority queue, and priorityQueue implement this interface.
type Interface interface {
	queue.Interface

	// WithComparator sets a gsort.Comparator instance for the queue.
	// It's used to imposes a total ordering on the elements in the queue.
	WithComparator(c gsort.Comparator) Interface

	// Contains returns true if this queue contains the specified element.
	Contains(val interface{}) bool
	// Remove a single instance of the specified element from this queue, if it is present.
	// It returns false if the target value isn't present, otherwise returns true.
	Remove(val interface{}) bool
}
```

Please import the following package in order to use priorityQueue,
```go
import (
	"github.com/ahrtr/gocontainer/queue/priorityqueue"
)
```

Call priorityqueue.New() to create a PriorityQueue,
```go
New() Interface 
```

The following is a simple example for priorityQueue,
```go
package main

import (
	"fmt"

	"github.com/ahrtr/gocontainer/queue/priorityqueue"
)

func main() {
	pq := priorityqueue.New()

	values := []string{"benjamin", "alice", "john", "tom", "bill"}

	for _, v := range values {
		pq.Add(v)
	}

	for _, v := range values {
		fmt.Printf("pq.Contains(%v) = %t\n", v, pq.Contains(v))
	}

	fmt.Printf("pq.Remove(john) = %t\n", pq.Remove("john"))

	for pq.Peek() != nil {
		fmt.Printf("pq.Peek() = %v\n", pq.Peek())
		fmt.Printf("pq.Poll() = %v\n", pq.Poll())
	}
}
```

A sort.Comparator instance can be provided for a PriorityQueue by method WithComparator, please get more detailed info in **[Sort](#sort)**.
```go
WithComparator(c gsort.Comparator) Interface
```

The elements of a PriorityQueue are ordered according to their natural ordering, or by a sort.Comparator instance. 

If the reverse order for the elements is expected, then makes use of the priorityqueue.Reverse function, 
```go
pq := priorityqueue.Reverse(priorityqueue.New())
```

## LinkedMap
LinkedMap is based on a map and a doubly linked list. The iteration ordering is normally the order in which keys were inserted into the map, or the order in which the keys were accessed if the accessOrder flag is set. It implements the following interface. Click **[here](examples/linkedmap_example.go)** to find examples on how to use a linked map.
```go
// Interface is a type of linked map, and linkedMap implements this interface.
type Interface interface {
	collection.Interface

	// Put associates the specified value with the specified key in this map. If the map previously contained a mapping for the key,
	// the old value is replaced by the specified value.
	// It returns the previous value associated with the specified key, or nil if there was no mapping for the key.
	// A nil return can also indicate that the map previously associated nil with the specified key.
	Put(k, v interface{}) interface{}

	// WithAccessOrder configures the iteration ordering for this linked map,
	// true for access-order, and false for insertion-order.
	WithAccessOrder(accessOrder bool) Interface

	// Get returns the value to which the specified key is mapped, or nil if this map contains no mapping for the key.
	Get(k interface{}) interface{}
	// GetOrDefault returns the value to which the specified key is mapped, or the defaultValue if this map contains no mapping for the key.
	GetOrDefault(k, defaultValue interface{}) interface{}

	// ContainsKey returns true if this map contains a mapping for the specified key.
	ContainsKey(k interface{}) bool
	// ContainsValue returns true if this map maps one or more keys to the specified value.
	ContainsValue(v interface{}) bool

	// Remove removes the mapping for a key from this map if it is present.
	// It returns the value to which this map previously associated the key, and true,
	// or nil and false if the map contained no mapping for the key.
	Remove(k interface{}) (interface{}, bool)
	// RemoveFirstElement removes the first element from this map, which is the head of the list.
	// It returns the (key, value, true) if the map isn't empty, or (nil, nil, false) if the map is empty.
	RemoveFirstElement() (interface{}, interface{}, bool)
	// RemoveLastElement removes the last element from this map, which is the tail of the list.
	// It returns the (key, value, true) if the map isn't empty, or (nil, nil, false) if the map is empty.
	RemoveLastElement() (interface{}, interface{}, bool)

	// Iterator returns an iterator over the elements in this map in proper sequence.
	Iterator() (func() (interface{}, interface{}, bool), bool)
	// ReverseIterator returns an iterator over the elements in this map in reverse sequence as Iterator.
	ReverseIterator() (func() (interface{}, interface{}, bool), bool)
}
```

Please import the following package in order to use linkedMap,
```go
import (
	"github.com/ahrtr/gocontainer/map/linkedmap"
)
```

Call linkedmap.New() to create a linked map,
```go
New() Interface
```

The following is a simple example for linkedMap,
```go
package main

import (
	"fmt"

	"github.com/ahrtr/gocontainer/map/linkedmap"
)

func main() {
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
```

If the order in which the keys were accessed is expected for the iteration ordering, then the accessOrder flag should be set, 
```go
// WithAccessOrder configures the iteration ordering for this linked map,
// true for access-order, and false for insertion-order.
WithAccessOrder(accessOrder bool) Interface
```

The following snips show how to interate a linkedMap,
```go
// To iterate over an linkedMap (where lm is an instance of linkedmap.Interface):
it, hasNext := lm.Iterator()
var k, v interface{}
for hasNext {
	k, v, hasNext = it()
	// do something with k & v
}
```

```go
// To iterate over an linkedMap in reverse order (where lm is an instance of linkedmap.Interface):
it, hasPrev := lm.ReverseIterator()
var k, v interface{}
for hasPrev {
	k, v, hasPrev = it()
	// do something with k & v
}
```

## Others
More containers will be added soon. Please also kindly let me know if you need any other kinds of containers. Feel free to raise issues. 

# Sort
Some containers implement interface **sort.Interface**, such as ArrayList and LinkedList, which means that they can be sorted directly by sort.Sort(data). For the following golang build-in data types, the elements can be ordered into ascending order according to their natural ordering. Note that for **bool**, a false is regarded as less than a true. 
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
- byte
- rune
- time.Time

Applications can also provide a sort.Comparator instance using method WithComparator for a container which implements sort.Interface.
```go
// Comparator imposes a total ordering on some collection of objects.
// Comparators can be passed to the construction function of a container(such as ArrayList, LinkedList or PriorityQueue) to allow precise control over the sort order.
type Comparator interface {
	// Compare compares its two arguments for order.
	// It returns a negative integer, zero, or a positive integer as the first argument is less than, equal to, or greater than the second.
	Compare(v1 interface{}, v2 interface{}) (int, error)
}
```

The rough logic should be something like below. Please find more examples in **[List](examples/list_example.go)** and **[PriorityQueue](examples/priorityqueue_example.go)**.
```go
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

