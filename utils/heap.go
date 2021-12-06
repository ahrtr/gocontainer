// Copyright (c) 2019, Benjamin Wang (benjamin_wang@aliyun.com). All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

package utils

import "sort"

// The GO's heap.Init()
/*
// Init establishes the heap invariants required by the other routines in this package.
// Init is idempotent with respect to the heap invariants
// and may be called whenever the heap invariants may have been invalidated.
// The complexity is O(n) where n = h.Len().
func Init(h Interface) {
	// heapify
	n := h.Len()
	for i := n/2 - 1; i >= 0; i-- {
		down(h, i, n)
	}
}
*/

// HeapInit establishes the heap from scratch. The operation is in-place.
// Parameters:
//     values:    the data source of the heap
//     isMinHeap: true for min-hap, false for max-heap
//     c:         an utils.Comparator instance
func HeapInit(values []interface{}, isMinHeap bool, c Comparator) {
	sc := constructHeapContainer(values, isMinHeap, c)
	n := sc.Len()
	for i := n/2 - 1; i >= 0; i-- {
		down(sc, i, n)
	}
}

// The Go's heap.Push()
/*
// Push pushes the element x onto the heap.
// The complexity is O(log n) where n = h.Len().
func Push(h Interface, x interface{}) {
	h.Push(x)
	up(h, h.Len()-1)
}
*/

// HeapPostPush moves the new element up until it gets to the right place. The operation is in-place.
// Push workflow (this functions takes care of the second step):
//     1.  add a new element to the end of the slice;
//     2*. call this method to move the new element up until it gets to the right place.
// Parameters:
//     values:    the data source of the heap
//     isMinHeap: true for min-hap, false for max-heap
//     c:         an utils.Comparator instance
func HeapPostPush(values []interface{}, isMinHeap bool, c Comparator) {
	sc := constructHeapContainer(values, isMinHeap, c)
	up(sc, sc.Len()-1)
}

// The GO's heap.Pop()
/*
// Pop removes and returns the minimum element (according to Less) from the heap.
// The complexity is O(log n) where n = h.Len().
// Pop is equivalent to Remove(h, 0).
func Pop(h Interface) interface{} {
	n := h.Len() - 1
	h.Swap(0, n)
	down(h, 0, n)
	return h.Pop()
}
*/

// HeapPrePop move the top element down until it gets to the right place. The operation is in-place.
// Pop workflow (this function takes care of step 1 and 2):
//    1*. swap the first and the last element;
//    2*. move the first/top element down until it gets to the right place;
//    3.  remove the last element, and return the removed element to users.
// Parameters:
//     values:    the data source of the heap
//     isMinHeap: true for min-hap, false for max-heap
//     c:         an utils.Comparator instance
func HeapPrePop(values []interface{}, isMinHeap bool, c Comparator) {
	// swap the first element (values[0]) and the last element (values[n])
	n := len(values) - 1
	values[0], values[n] = values[n], values[0]

	sc := constructHeapContainer(values, isMinHeap, c)
	down(sc, 0, n)
}

// The GO's heap.Remove()
/*
// Remove removes and returns the element at index i from the heap.
// The complexity is O(log n) where n = h.Len().
func Remove(h Interface, i int) interface{} {
	n := h.Len() - 1
	if n != i {
		h.Swap(i, n)
		if !down(h, i, n) {
			up(h, i)
		}
	}
	return h.Pop()
}
*/

// HeapPreRemove move the element with the specified index down or up until it gets to the right place. The operation is in-place.
// Remove workflow(this function takes care of step 1 and 2):
//    1*. swap the element with the specified index and the last element;
//    2*. move the element with the specified index down or up until it gets to the right place;
//    3.  remove the last element, and return the removed element to users.
// Parameters:
//     values:    the data source of the heap
//     index:     the element at the specified index will be removed after calling this function
//     isMinHeap: true for min-hap, false for max-heap
//     c:         an utils.Comparator instance
func HeapPreRemove(values []interface{}, index int, isMinHeap bool, c Comparator) {
	n := len(values) - 1
	if n != index {
		values[index], values[n] = values[n], values[index]

		sc := constructHeapContainer(values, isMinHeap, c)
		if !down(sc, index, n) {
			up(sc, index)
		}
	}
}

// The GO's heap.Fix()
/*
// Fix re-establishes the heap ordering after the element at index i has changed its value.
// Changing the value of the element at index i and then calling Fix is equivalent to,
// but less expensive than, calling Remove(h, i) followed by a Push of the new value.
// The complexity is O(log n) where n = h.Len().
func Fix(h Interface, i int) {
	if !down(h, i, h.Len()) {
		up(h, i)
	}
}
*/

// HeapPostUpdate re-establishes the heap ordering after the element at the specified index has changed its value. The operation is in-place.
// Update workflow (this function takes care of the second step):
//    1.  update the element's value at the specified index;
//    2*. call this function to move the updated element down or up until it gets to the right place.
// Parameters:
//     values:    the data source of the heap
//     index:     the element at the specified index should have already been updated before calling this function
//     isMinHeap: true for min-hap, false for max-heap
//     c:         an utils.Comparator instance
func HeapPostUpdate(values []interface{}, index int, isMinHeap bool, c Comparator) {
	sc := constructHeapContainer(values, isMinHeap, c)
	if !down(sc, index, sc.Len()) {
		up(sc, index)
	}
}

func constructHeapContainer(values []interface{}, isMinHeap bool, c Comparator) sort.Interface {
	if isMinHeap {
		return &sortableContainer{values, c}
	}
	return &reverseSortableContainer{&sortableContainer{values, c}}
}

// copied from Go's package container/heap, but changed the first parameter from heap.Interface to sort.Interface.
func up(h sort.Interface, j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		j = i
	}
}

// copied from Go's package container/heap, but changed the first parameter from heap.Interface to sort.Interface.
func down(h sort.Interface, i0, n int) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && h.Less(j2, j1) {
			j = j2 // = 2*i + 2  // right child
		}
		if !h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		i = j
	}
	return i > i0
}
