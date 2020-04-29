// Copyright (c) 2019, Benjamin Wang (benjamin_wang@aliyun.com). All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

// Package stack implements a stack, which orders elements in a LIFO (last-in-first-out) manner.
package stack

import (
	"github.com/ahrtr/gocontainer/collection"
)

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

// Stack represents a stack.
type Stack struct {
	items []interface{}
}

// New creates a stack.
func New() *Stack {
	return &Stack{
		items: []interface{}{},
	}
}

func (s *Stack) Len() int {
	return len(s.items)
}

// IsEmpty returns true if this stack contains no elements.
func (s *Stack) IsEmpty() bool {
	return s.Len() == 0
}

func (s *Stack) Push(val interface{}) {
	s.items = append(s.items, val)
}

func (s *Stack) Pop() interface{} {
	if len(s.items) > 0 {
		val := s.items[len(s.items)-1]
		s.items = s.items[:len(s.items)-1]
		return val
	} else {
		return nil
	}
}

// Clear removes all the elements from this stack.
func (s *Stack) Clear() {
	for i := 0; i < len(s.items); i++ {
		s.items[i] = nil
	}
	s.items = []interface{}{}
}
