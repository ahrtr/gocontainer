// Copyright (c) 2019, Benjamin Wang (benjamin_wang@aliyun.com). All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

// Package list implements both an ArrayList and a LinkedList.
//
// To iterate over an ArrayList (where al is a *ArrayList):
//	it, hasNext := al.Iterator()
//  var v interface{}
//	for hasNext {
//		v, hasNext = it()
//		// do something with v
//	}
//
// To iterate over an ArrayList in reverse order (where al is a *ArrayList):
//	it, hasPrev := al.ReverseIterator()
//  var v interface{}
//	for hasPrev {
//		v, hasPrev = it()
//		// do something with v
//	}
//
package list

import (
	"fmt"

	gsort "github.com/ahrtr/gocontainer/sort"
)

// ArrayList represents an array list.
// It implements the interface list.Interface.
type ArrayList struct {
	items []interface{}
	cmp   gsort.Comparator
}

// NewArrayList initializes and returns an ArrayList.
func NewArrayList() Interface {
	return &ArrayList{
		items: []interface{}{},
		cmp:   nil,
	}
}

// NewArrayListWithComparator initializes and returns an ArrayList with a comparator.
func NewArrayListWithComparator(c gsort.Comparator) Interface {
	return &ArrayList{
		items: []interface{}{},
		cmp:   c,
	}
}

func (al *ArrayList) Len() int {
	return len(al.items)
}

func (al *ArrayList) Less(i, j int) bool {
	var cmpRet int
	var err error
	if nil != al.cmp {
		cmpRet, err = al.cmp.Compare(al.items[i], al.items[j])
	} else {
		cmpRet, err = gsort.Compare(al.items[i], al.items[j])
	}
	if err != nil {
		panic(err)
	}
	return cmpRet < 0
}

func (al *ArrayList) Swap(i, j int) {
	if i != j {
		al.items[i], al.items[j] = al.items[j], al.items[i]
	}
}

func (al *ArrayList) IsEmpty() bool {
	return al.Len() == 0
}

func (al *ArrayList) Add(val interface{}) {
	al.items = append(al.items, val)
}

func (al *ArrayList) AddTo(index int, val interface{}) error {
	if index < 0 || index > len(al.items) {
		return fmt.Errorf("Index out of range, index:%d, len:%d", index, al.Len())
	}

	if index == al.Len() {
		al.Add(val)
	} else {
		curLen := al.Len()
		al.items = append(al.items, val)
		copy(al.items[(index+1):(curLen+1)], al.items[index:curLen])
		al.items[index] = val
	}

	return nil
}

func (al *ArrayList) Contains(val interface{}) bool {
	if al.IsEmpty() || nil == val {
		return false
	}

	for _, v := range al.items {
		if v == val {
			return true
		}
	}

	return false
}

func (al *ArrayList) Get(index int) (interface{}, error) {
	if index < 0 || index >= len(al.items) {
		return nil, fmt.Errorf("Index out of range, index:%d, len:%d", index, al.Len())
	}

	return al.items[index], nil
}

func (al *ArrayList) Remove(index int) (interface{}, error) {
	if index < 0 || index >= len(al.items) {
		return nil, fmt.Errorf("Index out of range, index:%d, len:%d", index, al.Len())
	}

	val := al.items[index]

	al.items = append(al.items[:index], al.items[(index+1):]...)

	return val, nil
}

func (al *ArrayList) RemoveByValue(val interface{}) bool {
	if al.Len() == 0 {
		return false
	}

	for i, v := range al.items {
		if v == val {
			al.items = append(al.items[:i], al.items[(i+1):]...)
			return true
		}
	}

	return false
}

func (al *ArrayList) Clear() {
	for i := 0; i < len(al.items); i++ {
		al.items[i] = nil
	}
	al.items = []interface{}{}
}

func (al *ArrayList) Iterator() (func() (interface{}, bool), bool) {
	index := 0

	return func() (interface{}, bool) {
		var element interface{}
		if index < al.Len() {
			element = al.items[index]
			index++
		} else {
			element = nil
		}
		return element, index < al.Len()
	}, index < al.Len()
}

func (al *ArrayList) ReverseIterator() (func() (interface{}, bool), bool) {
	index := al.Len() - 1

	return func() (interface{}, bool) {
		var element interface{}
		if index >= 0 {
			element = al.items[index]
			index--
		} else {
			element = nil
		}
		return element, index >= 0
	}, index >= 0
}
