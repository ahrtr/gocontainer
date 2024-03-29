// Copyright (c) 2019, Benjamin Wang (benjamin_wang@aliyun.com). All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

// Package queue implements a queue, which orders elements in a FIFO (first-in-first-out) manner.
package queue

import (
	"github.com/ahrtr/gocontainer/collection"
)

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

// element is an element of the queue.
type element struct {
	next  *element
	value interface{}
}

// queue represents a singly linked list.
type queue struct {
	head   *element
	tail   *element
	length int
}

// New creates a queue.
func New() Interface {
	return &queue{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

func (q *queue) Size() int {
	return q.length
}

// IsEmpty returns true if this queue contains no elements.
func (q *queue) IsEmpty() bool {
	return q.Size() == 0
}

// Add (todo): add a capacity for the queue, and return an error when this queue is full.
func (q *queue) Add(vals ...interface{}) {
	for _, v := range vals {
		e := element{
			next:  nil,
			value: v,
		}

		if nil == q.tail {
			q.head, q.tail = &e, &e
		} else {
			q.tail.next = &e
			q.tail = &e
		}

		q.length++
	}
}

func (q *queue) Peek() interface{} {
	if q.head != nil {
		return q.head.value
	}
	return nil
}

func (q *queue) Poll() interface{} {
	if q.head != nil {
		val := q.head.value

		q.head = q.head.next
		if nil == q.head {
			q.tail = nil
		}
		q.length--

		return val
	}

	return nil
}

// Clear removes all the elements from this queue.
func (q *queue) Clear() {
	for e := q.head; e != nil; {
		next := e.next
		e.next, e.value = nil, nil
		e = next
	}
	q.head, q.tail, q.length = nil, nil, 0
}
