// Copyright (c) 2019, Benjamin Wang (benjamin_wang@aliyun.com). All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

package list

import (
	"github.com/ahrtr/gocontainer/collection"
	"github.com/ahrtr/gocontainer/utils"
)

// Interface is a type of list, both ArrayList and LinkedList implement this interface.
type Interface interface {
	collection.Interface

	// Add appends the specified elements to the end of this list.
	Add(vals ...interface{})
	// AddTo inserts the specified element at the specified position in this list.
	AddTo(index int, val interface{}) error

	// Contains returns true if this list contains the specified element.
	Contains(val interface{}) bool
	// Get returns the element at the specified position in this list. The index must be in the range of [0, size).
	Get(index int) (interface{}, error)

	// Remove removes the element at the specified position in this list.
	// It returns an error if the index is out of range.
	Remove(index int) (interface{}, error)
	// RemoveByValue removes the first occurrence of the specified element from this list, if it is present.
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
