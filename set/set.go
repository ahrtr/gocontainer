// Copyright (c) 2019, Benjamin Wang (benjamin_wang@aliyun.com). All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

// Package set implements a set, which contains no duplicate elements.
//
// To iterate over a set (where s is a *Set):
//   s.Iterate(func(v interface{}) bool {
//       // do something with v
//	     return true
//   })
// Returning false in the callback function will break the iterating.
//
package set

import (
	"github.com/ahrtr/gocontainer/collection"
)

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

// IterateCallback is the signature of the callback function called by Iterate.
// If the callback function returns false, then the Iterate breaks.
type IterateCallback func(interface{}) bool

// Set is the definition of a set data structure, which contains no duplicate elements.
type Set struct {
	items map[interface{}]struct{}
}

// New creates a set.
func New() Interface {
	return &Set{
		items: map[interface{}]struct{}{},
	}
}

func (s *Set) Len() int {
	return len(s.items)
}

// IsEmpty returns true if this set contains no elements.
func (s *Set) IsEmpty() bool {
	return s.Len() == 0
}

func (s *Set) Add(val interface{}) bool {
	if _, ok := s.items[val]; !ok {
		s.items[val] = struct{}{}
		return true
	}
	return false
}

func (s *Set) Contains(val interface{}) bool {
	if _, ok := s.items[val]; ok {
		return true
	}
	return false
}

func (s *Set) Remove(val interface{}) bool {
	if _, ok := s.items[val]; ok {
		delete(s.items, val)
		return true
	}
	return false
}

// Clear removes all the elements from this set.
func (s *Set) Clear() {
	s.items = map[interface{}]struct{}{}
}

func (s *Set) Iterate(cb IterateCallback) {
	for k, _ := range s.items {
		if !cb(k) {
			break
		}
	}
}
