// Copyright (c) 2019, Benjamin Wang (benjamin_wang@aliyun.com). All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

// Package priorityqueue implements an unbounded priority queue based on a priority heap.
// The elements of the priority queue are ordered according to their natural ordering, or by a Comparator provided at PriorityQueue construction time.
package priorityqueue

import (
	"container/heap"

	"github.com/ahrtr/gocontainer/queue"

	gsort "github.com/ahrtr/gocontainer/sort"
)

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

// priorityQueue represents an unbounded priority queue based on a priority heap.
// It implements heap.Interface.
type priorityQueue struct {
	items []interface{}
	cmp   gsort.Comparator
}

// New initializes and returns an priorityQueue.
func New() Interface {
	return &priorityQueue{
		items: []interface{}{},
		cmp:   nil,
	}
}

func (pq *priorityQueue) WithComparator(c gsort.Comparator) Interface {
	pq.cmp = c
	return pq
}

// Size returns the length of this priority queue.
func (pq *priorityQueue) Size() int { return len(pq.items) }

// Len returns the length of this priority queue.
// Len is supposed to be called only by the functions in package container/heap.
// It isn't recommended for applications  call this method directly.
func (pq *priorityQueue) Len() int { return pq.Size() }

// Less reports whether the element with index i should sort before the element with index j.
// Less is supposed to be called only by the functions in package container/heap.
// Applications shouldn't call this method directly.
func (pq *priorityQueue) Less(i, j int) bool {
	var cmpRet int
	var err error
	if nil != pq.cmp {
		cmpRet, err = pq.cmp.Compare(pq.items[i], pq.items[j])
	} else {
		cmpRet, err = gsort.Compare(pq.items[i], pq.items[j])
	}
	if err != nil {
		panic(err)
	}
	return cmpRet < 0
}

// Swap swaps the elements with indexes i and j.
// Swap is supposed to be called only by the functions in package container/heap.
// Applications shouldn't call this method directly.
func (pq *priorityQueue) Swap(i, j int) {
	pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
}

// IsEmpty returns true if this list contains no elements.
func (pq *priorityQueue) IsEmpty() bool {
	return pq.Len() == 0
}

// Push is supposed to be called only by heap.Push.
// Applications shouldn't call this method directly.
func (pq *priorityQueue) Push(val interface{}) {
	pq.items = append(pq.items, val)
}

// Pop is supposed to be called only by heap.Pop.
// Applications shouldn't call this method directly.
func (pq *priorityQueue) Pop() interface{} {
	size := pq.Len()

	if size > 0 {
		val := pq.items[size-1]
		pq.items = pq.items[:(size - 1)]
		return val
	}
	return nil
}

// Clear removes all of the elements from this priority queue.
func (pq *priorityQueue) Clear() {
	for i := 0; i < len(pq.items); i++ {
		pq.items[i] = nil
	}
	pq.items = []interface{}{}
}

// Add inserts the specified element into this priority queue.
func (pq *priorityQueue) Add(val interface{}) {
	heap.Push(pq, val)
}

// Peek retrieves, but does not remove, the head of this queue, or return nil if this queue is empty.
func (pq *priorityQueue) Peek() interface{} {
	if pq.Len() > 0 {
		return pq.items[0]
	}
	return nil
}

// Poll retrieves and removes the head of the this queue, or return nil if this queue is empty.
func (pq *priorityQueue) Poll() interface{} {
	if pq.Len() > 0 {
		return heap.Pop(pq)
	}
	return nil
}

func (pq *priorityQueue) Contains(val interface{}) bool {
	return pq.indexOf(val) >= 0
}

func (pq *priorityQueue) Remove(val interface{}) bool {
	if pq.Len() == 0 {
		return false
	}

	i := pq.indexOf(val)
	if i < 0 {
		return false
	}

	heap.Remove(pq, i)
	return true
}

func (pq *priorityQueue) indexOf(val interface{}) int {
	if nil != val {
		for i := 0; i < len(pq.items); i++ {
			if val == pq.items[i] {
				return i
			}
		}
	}
	return -1
}

/*-----------------------------------------------------------------------------
The function Reverse() returns the reverse order for the data.
Previously the head of the queue is the the highest priority element, now it's the lowest priority element.

Re-implements the following methods in Interface:
	Less(i, j int) bool   // sort.Interface
	Swap(i, j int)        // sort.Interface
	Push(val interface{}) // heap.Interface
	Pop() interface{}     // heap.Interface

	// Add inserts the specified element into this priority queue.
	Add(val interface{})
	// Poll retrieves and removes the head of the this queue, or return nil if this queue is empty.
	Poll() interface{}
	// Remove a single instance of the specified element from this queue, if it is present.
	// It returns false if the target value isn't present, otherwise returns true.
	Remove(val interface{}) bool
-----------------------------------------------------------------------------*/
type reverse struct {
	// This embedded Interface permits Reverse to use the methods of
	// another Interface implementation.
	Interface
}

// Reverse returns the reverse order for data.
func Reverse(data Interface) Interface {
	if _, ok := data.(*priorityQueue); !ok {
		panic("The parameter must be a pointer to priorityQueue")
	}
	return &reverse{data}
}

func (r *reverse) Len() int { return r.Interface.Size() }

// Less returns the opposite of the embedded implementation's Less method.
func (r *reverse) Less(i, j int) bool {
	return r.Interface.(*priorityQueue).Less(j, i)
}

func (r *reverse) Swap(i, j int) {
	r.Interface.(*priorityQueue).Swap(i, j)
}

// Push is supposed to be called only by heap.Push.
// Developers shouldn't call this method directly.
func (r *reverse) Push(val interface{}) {
	r.Interface.(*priorityQueue).Push(val)
}

// Pop is supposed to be called only by heap.Pop.
// Developers shouldn't call this method directly.
func (r *reverse) Pop() interface{} {
	return r.Interface.(*priorityQueue).Pop()
}

func (r *reverse) Add(val interface{}) {
	heap.Push(r, val)
}

func (r *reverse) Poll() interface{} {
	if r.Interface.Size() > 0 {
		return heap.Pop(r)
	}
	return nil
}

func (r *reverse) Remove(val interface{}) bool {
	if r.Interface.Size() == 0 {
		return false
	}

	i := r.Interface.(*priorityQueue).indexOf(val)
	if i < 0 {
		return false
	}

	heap.Remove(r, i)
	return true
}
