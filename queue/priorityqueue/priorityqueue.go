// Copyright (c) 2019, Benjamin Wang (benjamin_wang@aliyun.com). All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

// Package priorityqueue implements an unbounded priority queue based on a priority heap.
// The elements of the priority queue are ordered according to their natural ordering, or by a Comparator provided at PriorityQueue construction time.
package priorityqueue

import (
	"github.com/ahrtr/gocontainer/queue"
	"github.com/ahrtr/gocontainer/utils"
)

// Interface is a type of priority queue, and priorityQueue implement this interface.
type Interface interface {
	queue.Interface

	// WithComparator sets an utils.Comparator instance for the queue.
	// It's used to impose a total ordering on the elements in the queue.
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

// priorityQueue represents an unbounded priority queue based on a priority heap.
// It implements heap.Interface.
type priorityQueue struct {
	items     []interface{}
	cmp       utils.Comparator
	isMinHeap bool
}

// New initializes and returns an priorityQueue.
func New() Interface {
	return &priorityQueue{
		items:     []interface{}{},
		cmp:       nil,
		isMinHeap: true,
	}
}

func (pq *priorityQueue) WithComparator(c utils.Comparator) Interface {
	pq.cmp = c
	return pq
}

func (pq *priorityQueue) WithMinHeap(isMinHeap bool) Interface {
	pq.isMinHeap = isMinHeap
	return pq
}

// Size returns the length of this priority queue.
func (pq *priorityQueue) Size() int { return len(pq.items) }

// IsEmpty returns true if this list contains no elements.
func (pq *priorityQueue) IsEmpty() bool {
	return pq.Size() == 0
}

// Clear removes all of the elements from this priority queue.
func (pq *priorityQueue) Clear() {
	size := pq.Size()
	for i := 0; i < size; i++ {
		pq.items[i] = nil
	}
	pq.items = []interface{}{}
}

// Add inserts the specified element into this priority queue.
func (pq *priorityQueue) Add(vals ...interface{}) {
	for _, v := range vals {
		pq.push(v)
		utils.HeapPostPush(pq.items, pq.isMinHeap, pq.cmp)
	}
}

// Peek retrieves, but does not remove, the head of this queue, or return nil if this queue is empty.
func (pq *priorityQueue) Peek() interface{} {
	if pq.Size() > 0 {
		return pq.items[0]
	}
	return nil
}

// Poll retrieves and removes the head of the this queue, or return nil if this queue is empty.
func (pq *priorityQueue) Poll() interface{} {
	if pq.Size() > 0 {
		utils.HeapPrePop(pq.items, pq.isMinHeap, pq.cmp)
		return pq.pop()
	}
	return nil
}

func (pq *priorityQueue) Contains(val interface{}) bool {
	return pq.indexOf(val) >= 0
}

func (pq *priorityQueue) Remove(val interface{}) bool {
	if pq.Size() == 0 {
		return false
	}

	i := pq.indexOf(val)
	if i < 0 {
		return false
	}

	utils.HeapPreRemove(pq.items, i, pq.isMinHeap, pq.cmp)
	pq.pop()

	return true
}

// push appends the provided value to the end.
func (pq *priorityQueue) push(val interface{}) {
	pq.items = append(pq.items, val)
}

// pop removes and returns the last element.
func (pq *priorityQueue) pop() interface{} {
	size := pq.Size()

	if size > 0 {
		val := pq.items[size-1]
		pq.items = pq.items[:(size - 1)]
		return val
	}
	return nil
}

func (pq *priorityQueue) indexOf(val interface{}) int {
	size := pq.Size()
	if nil != val {
		for i := 0; i < size; i++ {
			if val == pq.items[i] {
				return i
			}
		}
	}
	return -1
}
