// Copyright (c) 2019, Benjamin Wang (benjamin_wang@aliyun.com). All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

// Package priorityqueue implements an unbounded priority queue based on a priority heap.
// The elements of the priority queue are ordered according to their natural ordering, or by a Comparator provided at PriorityQueue construction time.
package priorityqueue

import (
	"container/heap"
	"sort"

	"github.com/ahrtr/gocontainer/collection"

	gsort "github.com/ahrtr/gocontainer/sort"
)

// Interface is a type of priority queue, and PriorityQueue implement this interface.
type Interface interface {
	collection.Interface
	sort.Interface

	// Add inserts the specified element into this priority queue.
	Add(val interface{})
	// Contains returns true if this queue contains the specified element.
	Contains(val interface{}) bool
	// Peek retrieves, but does not remove, the head of this queue, or return nil if this queue is empty.
	Peek() interface{}
	// Poll retrieves and removes the head of the this queue, or return nil if this queue is empty.
	Poll() interface{}
	// Remove a single instance of the specified element from this queue, if it is present.
	// It returns false if the target value isn't present, otherwise returns true.
	Remove(val interface{}) bool
}

// PriorityQueue represents an unbounded priority queue based on a priority heap.
// It implements heap.Interface.
type PriorityQueue struct {
	items []interface{}
	cmp   gsort.Comparator
}

// New initializes and returns an PriorityQueue.
func New() Interface {
	return &PriorityQueue{
		items: []interface{}{},
		cmp:   nil,
	}
}

// NewWithComparator initializes and returns an PriorityQueue with a comparator.
func NewWithComparator(c gsort.Comparator) Interface {
	return &PriorityQueue{
		items: []interface{}{},
		cmp:   c,
	}
}

func (pq *PriorityQueue) Len() int { return len(pq.items) }

func (pq *PriorityQueue) Less(i, j int) bool {
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

func (pq *PriorityQueue) Swap(i, j int) {
	pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
}

// IsEmpty returns true if this list contains no elements.
func (pq *PriorityQueue) IsEmpty() bool {
	return pq.Len() == 0
}

// Push is supposed to be called only by heap.Push.
// Developers shouldn't call this method directly.
func (pq *PriorityQueue) Push(val interface{}) {
	pq.items = append(pq.items, val)
}

// Pop is supposed to be called only by heap.Pop.
// Developers shouldn't call this method directly.
func (pq *PriorityQueue) Pop() interface{} {
	size := pq.Len()

	if size > 0 {
		val := pq.items[size-1]
		pq.items = pq.items[:(size - 1)]
		return val
	}
	return nil
}

// Clear removes all of the elements from this priority queue.
func (pq *PriorityQueue) Clear() {
	for i := 0; i < len(pq.items); i++ {
		pq.items[i] = nil
	}
	pq.items = []interface{}{}
}

func (pq *PriorityQueue) Add(val interface{}) {
	heap.Push(pq, val)
}

func (pq *PriorityQueue) Peek() interface{} {
	if pq.Len() > 0 {
		return pq.items[0]
	}
	return nil
}

func (pq *PriorityQueue) Poll() interface{} {
	if pq.Len() > 0 {
		return heap.Pop(pq)
	}
	return nil
}

func (pq *PriorityQueue) Contains(val interface{}) bool {
	return pq.indexOf(val) >= 0
}

func (pq *PriorityQueue) Remove(val interface{}) bool {
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

func (pq *PriorityQueue) indexOf(val interface{}) int {
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
	return &reverse{data}
}

// Less returns the opposite of the embedded implementation's Less method.
func (r *reverse) Less(i, j int) bool {
	return r.Interface.Less(j, i)
}

// Push is supposed to be called only by heap.Push.
// Developers shouldn't call this method directly.
func (r *reverse) Push(val interface{}) {
	r.Interface.(*PriorityQueue).Push(val)
}

// Pop is supposed to be called only by heap.Pop.
// Developers shouldn't call this method directly.
func (r *reverse) Pop() interface{} {
	return r.Interface.(*PriorityQueue).Pop()
}

func (r *reverse) Add(val interface{}) {
	heap.Push(r, val)
}

func (r *reverse) Poll() interface{} {
	if r.Interface.Len() > 0 {
		return heap.Pop(r)
	}
	return nil
}

func (r *reverse) Remove(val interface{}) bool {
	if r.Interface.Len() == 0 {
		return false
	}

	i := r.Interface.(*PriorityQueue).indexOf(val)
	if i < 0 {
		return false
	}

	heap.Remove(r, i)
	return true
}
