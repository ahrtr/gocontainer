// Copyright (c) 2019, Benjamin Wang (benjamin_wang@aliyun.com). All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

// Package stack implements a stack, which orders elements in a LIFO (last-in-first-out) manner.
package stack

import (
	"github.com/ahrtr/gocontainer/collection"
	"github.com/ahrtr/gocontainer/list"
)

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

// stack is LIFO.
type stack struct {
	l list.Interface
}

// New creates a stack.
func New() Interface {
	return &stack{list.NewArrayList()}
}

func (s *stack) Size() int {
	return s.l.Size()
}

// IsEmpty returns true if this stack contains no elements.
func (s *stack) IsEmpty() bool {
	return s.l.Size() == 0
}

func (s *stack) Push(val interface{}) {
	s.l.Add(val)
}

func (s *stack) Pop() interface{} {
	size := s.l.Size()
	if size > 0 {
		val, _ := s.l.Get(size - 1)
		if _, err := s.l.Remove(size - 1); err != nil {
			//todo: what should we do if failing to remove the element?
			return nil
		}

		return val
	}
	return nil
}

func (s *stack) Peek() interface{} {
	size := s.l.Size()
	if size > 0 {
		val, _ := s.l.Get(size - 1)
		return val
	}
	return nil
}

// Clear removes all the elements from this stack.
func (s *stack) Clear() {
	s.l.Clear()
}
