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
  - [BTree](#bTree)
  - [Others](#others)
- **[Utilities](#Utilities)**
  - [Comparator](#Comparator)
  - [Sort](#sort)
  - [Heap](#heap)
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
	// Create an array list
	al := list.NewArrayList()
	al.Add(5, 6, 7)

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
- BTree 

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
	// Peek retrieves, but does not remove, the element on the top of this stack, or return nil if this stack is empty.
	Peek() interface{}
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
	Add(vals ...interface{})
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

	// Add adds the specified values to this set if they are not already present.
	// It returns false if any value is already present.
	Add(vals ...interface{}) bool
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

	// Add appends the specified elements to the end of this list.
	Add(vals ...interface{})
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

	// Sort sorts the element using default options below. It sorts the elements into ascending sequence according to their natural ordering.
	//     reverse: false
	//     comparator: nil
	Sort()
	// SortWithOptions sorts the elements in the list.
	// Parameters:
	//     reverse: whether sort the data in reverse ordering
	//     c:       sort the data according to the provided comparator
	// If reverse is true, and a comparator is also provided, then the result will be the reverse sequence as the comparator generates.
	SortWithOptions(reverse bool, c utils.Comparator)

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

A list can be sorted using one of the following two methods. The first method Sort() sorts the list into ascending sequence according to the natural ordering of its elements by default; actually it just calls the second method SortWithOptions(false, nil) using the default parameters. SortWithOptions sorts the list according to the provided parameters. Please get more detailed info in **[Comparator](#comparator)**
```go
Sort()
SortWithOptions(reverse bool, c utils.Comparator)
```

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
// This approach isn't efficient for linkedList.
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
// This approach isn't efficient for linkedList.
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

	// WithComparator sets a utils.Comparator instance for the queue.
	// It's used to imposes a total ordering on the elements in the queue.
	WithComparator(c utils.Comparator) Interface
	// WithMinHeap configures whether or not using min-heap.
	// If not configured, then it's min-heap by default.
	WithMinHeap(isMinHeap bool) Interface

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

A utils.Comparator instance can be provided for a priorityQueue by method WithComparator, please get more detailed info in **[Comparator](#comparator)**.
```go
WithComparator(c utils.Comparator) Interface
```

A priorityQueue can be configured to use min-heap or max-heap using method WithMinHeap. If the parameter is true, then it's a min-heap, which is the default option as well; otherwise, it's a max-heap.
```go
WithMinHeap(isMinHeap bool) Interface
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

## **BTree**
BTree is a B-Tree implementation. It was originally copied from github.com/google/btree, but it is refactored to adapt to the interface convention in this repository. Some improvements are also applied on top of the original design & implementation, so that it's more user-friendly.
It implements the following interface. Click **[here](examples/btree_example.go)** to find examples on how to use a BTree.
```go
// Interface is a type of btree, and bTree implements this interface
type Interface interface {
	collection.Interface

	// WithComparator sets an utils.Comparator instance for the btree.
	// It's used to impose a total ordering on the elements in the btree.
	WithComparator(c utils.Comparator) Interface

	// Clone clones the btree, lazily. The internal tree structure is marked read-only and
	// shared between the old and new btree. Writes to both the old and the new btree use copy-on-write logic.
	Clone() Interface
	// ReplaceOrInsert adds the given item to the tree.  If an item in the tree
	// already equals the given one, it is removed from the tree and returned.
	// Otherwise, nil is returned.
	ReplaceOrInsert(item interface{}) interface{}
	// Delete removes an item equal to the passed in item from the tree, returning
	// it.  If no such item exists, returns nil.
	Delete(item interface{}) interface{}
	// DeleteMin removes the smallest item in the tree and returns it.
	// If no such item exists, returns nil.
	DeleteMin() interface{}
	// DeleteMax removes the largest item in the tree and returns it.
	// If no such item exists, returns nil.
	DeleteMax() interface{}

	// AscendRange calls the iterator for every value in the tree within the range
	// [greaterOrEqual, lessThan), until iterator returns false.
	AscendRange(greaterOrEqual, lessThan interface{}, iterator ItemIterator)
	// AscendLessThan calls the iterator for every value in the tree within the range
	// [first, pivot), until iterator returns false.
	AscendLessThan(pivot interface{}, iterator ItemIterator)
	// AscendGreaterOrEqual calls the iterator for every value in the tree within
	// the range [pivot, last], until iterator returns false.
	AscendGreaterOrEqual(pivot interface{}, iterator ItemIterator)
	// Ascend calls the iterator for every value in the tree within the range
	// [first, last], until iterator returns false.
	Ascend(iterator ItemIterator)

	// DescendRange calls the iterator for every value in the tree within the range
	// [lessOrEqual, greaterThan), until iterator returns false.
	DescendRange(lessOrEqual, greaterThan interface{}, iterator ItemIterator)
	// DescendLessOrEqual calls the iterator for every value in the tree within the range
	// [pivot, first], until iterator returns false.
	DescendLessOrEqual(pivot interface{}, iterator ItemIterator)
	// DescendGreaterThan calls the iterator for every value in the tree within
	// the range [last, pivot), until iterator returns false.
	DescendGreaterThan(pivot interface{}, iterator ItemIterator)
	// Descend calls the iterator for every value in the tree within the range
	// [last, first], until iterator returns false.
	Descend(iterator ItemIterator)

	// Get looks for the key item in the tree, returning it.  It returns nil if
	// unable to find that item.
	Get(key interface{}) interface{}
	// Min returns the smallest item in the tree, or nil if the tree is empty.
	Min() interface{}
	// Max returns the largest item in the tree, or nil if the tree is empty.
	Max() interface{}
	// Has returns true if the given key is in the tree.
	Has(key interface{}) bool
}
```

Please import the following package in order to use btree,
```go
import (
	"github.com/ahrtr/gocontainer/btree"
)
```

Call btree.New() to create a BTree,
```go
New() Interface 
```

The following is a simple example for btree,
```go
package main

import (
	"fmt"

	"github.com/ahrtr/gocontainer/btree"
)

func main() {
    items := []int {5, 9, 2, 4, 11, 6}
    tr := btree.New(2)

    fmt.Printf("tr.Size(): %d\n", tr.Size()) // should be 0 in the beginning

    // Insert values
    fmt.Printf("Inserting %d items: %v\n", len(items), items)
    for _, item := range items {
    	tr.ReplaceOrInsert(item)
    }

    // Search values
    fmt.Printf("    tr.Size(): %d\n", tr.Size()) // should be len(items): 6 now
    fmt.Printf("    tr.Min(): %v\n", tr.Min()) // should be 2
    fmt.Printf("    tr.Max(): %v\n", tr.Max()) // should be 11
    fmt.Printf("    tr.Has(6): %t\n", tr.Has(6))  // true
    fmt.Printf("    tr.Get(6): %v\n", tr.Get(6))  // 6
    fmt.Printf("    tr.Has(7): %t\n", tr.Has(7))  // false
    fmt.Printf("    tr.Get(7): %v\n", tr.Get(7))  // nil

    // Delete values
    fmt.Println("Deleting items:")
    fmt.Printf("    tr.DeleteMin(): %v\n", tr.DeleteMin()) // 2 is deleted and returned
    fmt.Printf("    tr.Min(): %v\n", tr.Min()) // should be 4 now
    fmt.Printf("    tr.DeleteMax(): %v\n", tr.DeleteMax()) // 11 is deleted and returned
    fmt.Printf("    tr.Max(): %v\n", tr.Max()) // should be 9 now
    fmt.Printf("    tr.Delete(6): %v\n", tr.Delete(6)) // 6 is deleted and returned
    fmt.Printf("    tr.Delete(7): %v\n", tr.Delete(7)) // 7 doesn't exist, so nil is returned

    fmt.Printf("tr.Size(): %d\n", tr.Size()) // should be 3 now because 3 items have already been removed
}
```

A utils.Comparator instance can be provided for a btree by method WithComparator, please get more detailed info in **[Comparator](#comparator)**.
```go
WithComparator(c utils.Comparator) Interface
```

## Others
More containers will be added soon. Please also kindly let me know if you need any other kinds of containers. Feel free to raise issues. 

# Utilities
## Comparator
The comparator utility contains a function "Compare" and an interface "Comparator", 
```go
// Compare compares two arguments using the given Comparator. If the Comparator isn't provided, then the two values are compared according to their natural ordering.
// They must be the same type, otherwise returns an error in the second return value.
// It returns a negative integer, zero, or a positive integer as the first argument is less than, equal to, or greater than the second.
func Compare(v1 interface{}, v2 interface{}, cmp Comparator) (int, error)

// Comparator imposes a total ordering on some collection of objects, and it allows precise control over the sort order.
type Comparator interface {
	// Compare compares its two arguments for order.
	// It returns a negative integer, zero, or a positive integer as the first argument is less than, equal to, or greater than the second.
	Compare(v1 interface{}, v2 interface{}) (int, error)
}
```

The function "Compare" is used to compare two values using the given Comparator of the third parameter. If the Comparator is nil, then they are compared according to their natural ordering of golang build-in data types listed below. The two arguments must be the same data type, otherwise an error in the second return value will be returned. It returns a negative integer, zero, or a positive integer as the first argument is less than, equal to, or greater than the second one. Note that for **bool**, a false is regarded as less than a true.
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

Applications can also provide a utils.Comparators instance to customize the comparing. The following example demonstrates how to compare two students by age.
```go
type student struct {
	name string
	age int
}

type MyComparator struct{}

func (c *MyComparator) Compare(v1, v2 interface{}) (int, error) {
	e1, e2 := v1.(*student), v2.(*student)
	if e1.age < e2.age {
		return -1, nil
	}
	if e1.age > e2.age {
		return 1, nil
	}
	return 0, nil
}
```

## Sort
The sort utility provides the following two functions to sort the values in the provided slice.
```go
// Sort sorts values into ascending sequence according to their natural ordering, or according to the provided comparator.
func Sort(values []interface{}, c Comparator)

// ReverseSort sorts the values into opposite sequence to Sort
func ReverseSort(values []interface{}, c Comparator)
```

Both of the above functions sort values in-place. The first function "Sort" sorts the values into ascending sequence according to their natural ordering, or according to the provided comparator. The second function "ReverseSort" sorts the values into opposite sequence to the first function "Sort".

## Heap
The heap utility provides the following functions. It's useful for containers like priorityQueue. Please read the comment for each function to get more detailed info.
```go
// HeapInit establishes the heap from scratch. The operation is in-place.
// Parameters:
//     values:    the data source of the heap
//     isMinHeap: true for min-hap, false for max-heap
//     c:         an utils.Comparator instance
func HeapInit(values []interface{}, isMinHeap bool, c Comparator) 

// HeapPostPush moves the new element up until it gets to the right place. The operation is in-place.
// Push workflow (this functions takes care of the second step):
//     1.  add a new element to the end of the slice;
//     2*. call this method to move the new element up until it gets to the right place.
// Parameters:
//     values:    the data source of the heap
//     isMinHeap: true for min-hap, false for max-heap
//     c:         an utils.Comparator instance
func HeapPostPush(values []interface{}, isMinHeap bool, c Comparator) 

// HeapPrePop move the top element down until it gets to the right place. The operation is in-place.
// Pop workflow (this function takes care of step 1 and 2):
//    1*. swap the first and the last element;
//    2*. move the first/top element down until it gets to the right place;
//    3.  remove the last element, and return the removed element to users.
// Parameters:
//     values:    the data source of the heap
//     isMinHeap: true for min-hap, false for max-heap
//     c:         an utils.Comparator instance
func HeapPrePop(values []interface{}, isMinHeap bool, c Comparator)

// HeapPreRemove move the element with the specified index down or up until it gets to the right place. The operation is in-place.
// Remove workflow(this function takes care of step 1 and 2):
//    1*. swap the element with the specifed index and the last element;
//    2*. move the element with the specified index down or up until it gets to the right place;
//    3.  remove the last element, and return the removed element to users.
// Parameters:
//     values:    the data source of the heap
//     index:     the element at the specified index will be removed after calling this function
//     isMinHeap: true for min-hap, false for max-heap
//     c:         an utils.Comparator instance
func HeapPreRemove(values []interface{}, index int, isMinHeap bool, c Comparator) 

// HeapPostUpdate re-establishes the heap ordering after the element at the specified index has changed its value. The operation is in-place.
// Update workflow (this function takes care of the second step):
//    1.  update the element's value at the specified index;
//    2*. call this function to move the updated element down or up until it gets to the right place.
// Parameters:
//     values:    the data source of the heap
//     index:     the element at the specified index should have already been updated before calling this function
//     isMinHeap: true for min-hap, false for max-heap
//     c:         an utils.Comparator instance
func HeapPostUpdate(values []interface{}, index int, isMinHeap bool, c Comparator)
```

# Contribute to this repo
Anyone is welcome to contribute to this repo. Please raise an issue firstly, then fork this repo and submit a pull request.

Currently this repo is under heavily development, any helps are appreciated! 

# Support
If you need any support, please raise issues. 

If you have any suggestions or proposals, please also raise issues. Thanks!

