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
  - [其它容器](#其它容器)
- **[关于排序](#关于排序)**
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

您可以在 **[这里](examples)** 找到更多的示例。

# 公共接口
这个项目中的所有容器都实现了接口**collection.Interface**,
```go
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

# 容器
目前这个项目实现了下面这些容器：
- Stack
- Queue
- Set
- List (ArrayList, LinkedList)
- PriorityQueue
- LinkedMap

## Stack
Stack(栈)是一种后进先出(LIFO:last-in-first-out)的容器. 它实现了下面的接口。点击 **[这里](examples/stack_example.go)** 查看关于stack的示例。
```go
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

为了使用stack，必须import下面这个package，
```go
import (
	"github.com/ahrtr/gocontainer/stack"
)
```

调用stack.New()可以创建要一个stack，
```go
New() Interface
```

## Queue
Queue（队列）是一种先进先出(FIFO: first-in-first-out)的容器。它实现了下面的接口。点击 **[这里](examples/queue_example.go)** 查看关于queue的示例。
```go
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

## Set
一个set（集合）内不允许包含重复的元素。保存在set中的值必须是可比较(comparable)的，请参考Golang的[语言规范](https://golang.org/ref/spec#Comparison_operators)获取关于比较操作符的详细信息。

set（集合）实现了下面的接口。点击 **[这里](examples/set_example.go)** 查看关于set的示例。
```go
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
可以通过方法WithComparator为一个list设置一个sort.Comparator实例，具体请参考 **[关于排序](#关于排序)**.
```go
WithComparator(c gsort.Comparator) Interface 
```

因为接口list.Interface"继承"（内嵌）了接口sort.Interface，因此可以用sort.Sort(data)直接对一个list进行排序；要么根据元素的自然顺序按升序排序，要么根据上层应用提供的sort.Comparator实例进行排序。

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
PriorityQueue (优先级队列)是一种基于优先级堆实现的队列。 它实现了下面的接口。点击 **[这里](examples/priorityqueue_example.go)** 查看关于PriorityQueue的示例。
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
可以通过方法WithComparator为一个priorityQueue设置一个sort.Comparator实例，具体请参考 **[关于排序](#关于排序)**.
```go
WithComparator(c gsort.Comparator) Interface
```

PriorityQueue中的元素根据它们的自然顺序排序，或者根据一个sort.Comparator实例来排序。

如何期望是倒序，那么使用函数priorityqueue.Reverse,
```go
pq := priorityqueue.Reverse(priorityqueue.New())
```

## LinkedMap
LinkedMap是基于一个map和一个双向链表实现的。元素的遍历顺序默认是根据插入的顺序；但是如果设置了标志accessOrder，那就是根据访问的顺序。LinkedMap实现了下面的接口。点击 **[这里](examples/linkedmap_example.go)** 查看linkedMap的示例。
```go
// Interface is a type of linked map, and linkedMap implements this interface.
type Interface interface {
	collection.Interface

	// Len returns the number of elements in the linkedMap.
	Len() int
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
	"github.com/ahrtr/gocontainer/linkedmap"
)
```

调用linkedmap.New()可以创建一个linkedMap，
```go
New() Interface
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

## 其它容器
更多的容器将来可能会加入进来。如果您需要任何其它类型的容器，或者有任何建议，欢迎通过issues反馈给我。

# 关于排序
一些容器实现了接口 **sort.Interface**, 比如ArrayList和LinkedList。这就意味着这些容器可以直接通过sort.Sort(data)排序。对于下面这些Golang内置的类型，默认是根据它们的自然顺序按升序排序。对于 **bool** 类型, false被认为是比true要小。 
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

上层应用程序也可以通过方法WithComparator为实现了接口sort.Interface的容器提供一个sort.Comparator实例，
```go
// Comparator imposes a total ordering on some collection of objects.
// Comparators can be passed to the construction function of a container(such as ArrayList, LinkedList or PriorityQueue) to allow precise control over the sort order.
type Comparator interface {
	// Compare compares its two arguments for order.
	// It returns a negative integer, zero, or a positive integer as the first argument is less than, equal to, or greater than the second.
	Compare(v1 interface{}, v2 interface{}) (int, error)
}
```

实现sort.Comparator的大致逻辑如下。请在 **[List](examples/list_example.go)** 和 **[PriorityQueue](examples/priorityqueue_example.go)** 中分别查看更多的示例。
```go
type MyComparator struct{}

func (c *MyComparator) Compare(v1, v2 interface{}) (int, error) {
    //......
}
```

# 为该项目提供帮助
欢迎任何人为该开源项目提供任何帮助以及提出任何建议，哪怕只是指出一个拼写错误，谢谢。如果您觉得该项目对您有些用处，麻烦给个星(star)。请通过issue反馈任何问题。

# 技术支持
如果您遇到任何问题，或需要任何支持，请直接创建一个issue，在issue里面把你遇到的问题或需要的支持说清楚。


