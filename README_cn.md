gocontainer
======
gocontainer实现了一些Java中存在，而Golang中没有的容器。这个开源容器库不依赖于任何其它第三方软件包，可以说是**零依赖**。目前该项目中实现的容器不是线程安全的。

# 目录

- **[如何使用这个项目中的容器](#如何使用这个项目中的容器)**
- **[公共接口](#公共接口)**
- **[容器](#容器)**
  - [Stack](#stack)
  - [Queue](#queue)
  - [Set](#set)
  - [List](#list)
  - [PriorityQueue](#priorityqueue)
  - [LinkedMap](#linkedMap)
  - [BTree](#bTree)
  - [其它容器](#其它容器)
- **[工具箱](#工具箱)**
  - [Comparator](#Comparator)
  - [Sort](#sort)
  - [Heap](#heap)
- **[为该项目提供帮助](#为该项目提供帮助)**
- **[技术支持](#技术支持)**

# 如何使用这个项目中的容器
使用方法非常简单直接，只需要imports您所需要的容器所在的package，然后直接使用即可。下面是一个完整的使用ArrayList的例子，
```go
package main

import (
	"fmt"

	"github.com/ahrtr/gocontainer/list"
)

func main() {
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

您可以在 **[这里](examples)** 找到更多的示例。

# 公共接口
这个项目中的所有容器都实现了接口**collection.Interface**,
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

# 容器
目前这个项目实现了下面这些容器：
- Stack
- Queue
- Set
- List (ArrayList, LinkedList)
- PriorityQueue
- LinkedMap
- BTree

## Stack
Stack(栈)是一种后进先出(LIFO:last-in-first-out)的容器. 它实现了下面的接口。点击 **[这里](examples/stack_example.go)** 查看关于stack的示例。
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

为了使用stack，必须import下面这个package，
```go
import (
	"github.com/ahrtr/gocontainer/stack"
)
```

调用stack.New()可以创建一个stack，
```go
New() Interface
```

下面是一个简单的使用stack的例子,
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
Queue（队列）是一种先进先出(FIFO: first-in-first-out)的容器。它实现了下面的接口。点击 **[这里](examples/queue_example.go)** 查看关于queue的示例。
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

为了使用queue，必须import下面这个package，
```go
import (
	"github.com/ahrtr/gocontainer/queue"
)
```

调用queue.New()可以创建一个queue，
```go
New() Interface
```

下面是一个简单的使用queue的例子,
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
一个set（集合）内不允许包含重复的元素。保存在set中的值必须是可比较(comparable)的，请参考Golang的[语言规范](https://golang.org/ref/spec#Comparison_operators)获取关于比较操作符的详细信息。

set（集合）实现了下面的接口。点击 **[这里](examples/set_example.go)** 查看关于set的示例。
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

为了使用set，必须import下面这个package，
```go
import (
	"github.com/ahrtr/gocontainer/set"
)
```

调用set.New()可以创建一个set,
```go
New() Interface
```

下面是一个简单的使用set的例子，
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

上层应用程序在遍历一个set时，需要定义一个回调函数（如下），
```go
// IterateCallback is the signature of the callback function called by Iterate.
// If the callback function returns false, then the iteration breaks.
type IterateCallback func(interface{}) bool
```

下面的代码片段演示了如何遍历一个set. 请查看 **[example](examples/set_example.go)** 获取更详细的信息。

```go
// To iterate over a set (where s is an instance of set.Interface):
s.Iterate(func(v interface{}) bool {
	// Do something with v

	// If you want to break the iteration, then return a false
	return true
})
```

## List
这个容器库中实现了两种类型的list(列表)，分别是 **ArrayList** 和 **LinkedList**, 他们都实现了下面的接口。点击 **[这里](examples/list_example.go)** 查看关于list的示例。
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

为了使用list(arrayList或linkedList)，必须import下面这个package，
```go
import (
	"github.com/ahrtr/gocontainer/list"
)
```

调用list.NewArrayList()和list.NewLinkedList()可以分别创建一个ArrayList和LinkedList,
```go
NewArrayList() Interface
NewLinkedList() Interface
```

下面是一个简单的使用arrayList的例子，
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

下面是一个简单的使用linkedList的例子，
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

可以通过下面两个方法对一个list进行排序。第一个方法Sort()默认是根据list中元素的自然顺序按升序排序；它实际上是用默认参数直接调用第二个方法SortWithOptions(false, nil)。第二个方法  SortWithOptions根据传入的参数值对list进行排序。具体请参考 **[Comparator](#comparator)**
```go
Sort()
SortWithOptions(reverse bool, c utils.Comparator)
```

有多种方法可以遍历一个list，下面的代码片段演示了如何遍历一个list(arrayList或linkedList),
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
PriorityQueue (优先级队列)是一种基于优先级堆实现的队列。 它实现了下面的接口。点击 **[这里](examples/priorityqueue_example.go)** 查看关于PriorityQueue的示例。
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

为了使用priorityQueue，必须import下面这个package，
```go
import (
	"github.com/ahrtr/gocontainer/queue/priorityqueue"
)
```

调用priorityqueue.New()可以创建一个PriorityQueue,
```go
New() Interface 
```

下面是一个简单的使用priorityQueue的例子，
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

可以通过方法WithComparator为一个priorityQueue设置一个utils.Comparator实例，具体请参考 **[Comparator](#comparator)**.
```go
WithComparator(c utils.Comparator) Interface
```

可以通过方法WithMinHeap将一个priorityQueue设置成小顶堆或大顶堆。如果传给这个方法的参数为true，那就是小顶堆，这也是默认选项。如果为false，则是大顶堆。
```go
WithMinHeap(isMinHeap bool) Interface
```

## LinkedMap
LinkedMap是基于一个map和一个双向链表实现的。元素的遍历顺序默认是根据插入的顺序；但是如果设置了标志accessOrder，那就是根据访问的顺序。LinkedMap实现了下面的接口。点击 **[这里](examples/linkedmap_example.go)** 查看linkedMap的示例。
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

为了使用linkedMap，必须import下面这个package，
```go
import (
	"github.com/ahrtr/gocontainer/map/linkedmap"
)
```

调用linkedmap.New()可以创建一个linkedMap，
```go
New() Interface
```

下面是一个简单的使用linkedMap的例子，
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

如果设置了标志accessOrder，那么元素的访问顺序就是其遍历顺序，
```go
// WithAccessOrder configures the iteration ordering for this linked map,
// true for access-order, and false for insertion-order.
WithAccessOrder(accessOrder bool) Interface
```

下面的代码片段演示了如何遍历一个linkedMap,
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
BTree是一个B-Tree的实现。最初的实现是从github.com/google/btree拷贝来的，但是做了一定的重构以适应该项目一贯的接口约定。另外，也对原来的一些设计与实现做了一些改进，因此使用起来更加用户友好！
具体来说，它实现了下面的接口。点击**[这里](examples/btree_example.go)**查看关于btree的示例。
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

为了使用btree，必须import下面这个package，
```go
import (
	"github.com/ahrtr/gocontainer/btree"
)
```

调用btree.New()可以创建一个btree,
```go
New() Interface 
```

下面是一个简单的使用btree的例子，
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

可以通过方法WithComparator为一个btree设置一个utils.Comparator实例，具体请参考**[Comparator](#comparator)**.
```go
WithComparator(c utils.Comparator) Interface
```

## 其它容器
更多的容器将来可能会加入进来。如果您需要任何其它类型的容器，或者有任何建议，欢迎通过issues反馈给我。


# 工具箱
## Comparator
Comparator包含一个函数"Compare"和一个接口"Comparator"，
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

函数"Compare"用提供的Comparator(第三个参数)比较两个值。如果Comparator为nil，则默认对golang内置的数据类型进行比较，支持的内置类型见下面的列表。这个函数的前两个参数的数据类型必须相同，否则Compare就会返回一个错误（第二个返回参数）。在第一个参数小于、等于、大于第二个参数的情况下，第一个返回值分别是一个负整数、零、一个正整数。对于 **bool** 类型, false被认为是比true要小。
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

上层应用程序也可以提供一个utils.Comparator实例来定制排序。下面的例子演示了如何通过定制的utils.Comparator来根据age对两个student对象进行排序。
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
Sort提供了下面两个函数，对传入的slice中的元素进行排序。
```go
// Sort sorts values into ascending sequence according to their natural ordering, or according to the provided comparator.
func Sort(values []interface{}, c Comparator)

// ReverseSort sorts the values into opposite sequence to Sort
func ReverseSort(values []interface{}, c Comparator)
```

上面两个函数都是原地操作，所以对slice元素的操作会反映到调用者的原始slice中。第一个函数“Sort”根据元素的自然顺序或者根据传入的comparator来排序。第二个函数"ReverseSort"的排序顺序与第一个函数正好相反。

## Heap
Heap(堆)提供了下面这些方法。Heap对于像priorityQueue这样的容器非常有用。每一个函数都有比较详细的注释，请参考这些注释。
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

# 为该项目提供帮助
欢迎任何人为该开源项目提供任何帮助以及提出任何建议，哪怕只是指出一个拼写错误，谢谢。如果您觉得该项目对您有些用处，麻烦给个星(star)。请通过issue反馈任何问题。

# 技术支持
如果您遇到任何问题，或需要任何支持，请直接创建一个issue，在issue里面把你遇到的问题或需要的支持说清楚。


