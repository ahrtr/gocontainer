// Copyright (c) 2019, Benjamin Wang (benjamin_wang@aliyun.com). All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

// Package list implements both an arrayList and a linkedList.
//
// To iterate over an arrayList (where al is a *arrayList):
//	it, hasNext := al.Iterator()
//  var v interface{}
//	for hasNext {
//		v, hasNext = it()
//		// do something with v
//	}
//
// To iterate over an arrayList in reverse order (where al is a *arrayList):
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

	"github.com/ahrtr/gocontainer/utils"
)

// arrayList represents an array list.
// It implements the interface list.Interface.
type arrayList struct {
	items []interface{}
}

// NewArrayList initializes and returns an ArrayList.
func NewArrayList() Interface {
	return &arrayList{
		items: []interface{}{},
	}
}

func (al *arrayList) Size() int {
	return len(al.items)
}

func (al *arrayList) IsEmpty() bool {
	return al.Size() == 0
}

func (al *arrayList) Add(vals ...interface{}) {
	al.items = append(al.items, vals...)
}

func (al *arrayList) AddTo(index int, val interface{}) error {
	if index < 0 || index > len(al.items) {
		return fmt.Errorf("index out of range, index:%d, len:%d", index, al.Size())
	}

	if index == al.Size() {
		al.Add(val)
	} else {
		curLen := al.Size()
		al.items = append(al.items, val)
		copy(al.items[(index+1):(curLen+1)], al.items[index:curLen])
		al.items[index] = val
	}

	return nil
}

func (al *arrayList) Contains(val interface{}) bool {
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

func (al *arrayList) Get(index int) (interface{}, error) {
	if index < 0 || index >= len(al.items) {
		return nil, fmt.Errorf("index out of range, index:%d, len:%d", index, al.Size())
	}

	return al.items[index], nil
}

func (al *arrayList) Remove(index int) (interface{}, error) {
	if index < 0 || index >= len(al.items) {
		return nil, fmt.Errorf("index out of range, index:%d, len:%d", index, al.Size())
	}

	val := al.items[index]

	al.items = append(al.items[:index], al.items[(index+1):]...)

	al.shrinkList()
	return val, nil
}

func (al *arrayList) RemoveByValue(val interface{}) bool {
	if al.Size() == 0 {
		return false
	}

	for i, v := range al.items {
		if v == val {
			al.items = append(al.items[:i], al.items[(i+1):]...)
			al.shrinkList()
			return true
		}
	}

	return false
}

func (al *arrayList) Clear() {
	for i := 0; i < len(al.items); i++ {
		al.items[i] = nil
	}
	al.items = []interface{}{}
}

func (al *arrayList) Sort() {
	al.SortWithOptions(false, nil)
}

func (al *arrayList) SortWithOptions(reverse bool, c utils.Comparator) {
	if reverse {
		utils.ReverseSort(al.items, c)
	} else {
		utils.Sort(al.items, c)
	}
}

func (al *arrayList) Iterator() (func() (interface{}, bool), bool) {
	index := 0

	return func() (interface{}, bool) {
		var element interface{}
		if index < al.Size() {
			element = al.items[index]
			index++
		} else {
			element = nil
		}
		return element, index < al.Size()
	}, index < al.Size()
}

func (al *arrayList) ReverseIterator() (func() (interface{}, bool), bool) {
	index := al.Size() - 1

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

func (al *arrayList) shrinkList() {
	oldcap := cap(al.items)
	if oldcap <= 1024 {
		// no need to shrink when the capacity is less than 1024
		return
	}
	oldlen := len(al.items)
	if oldlen <= oldcap/4 { // shrink when len(list) <= cap(list)/4
		newItems := make([]interface{}, oldlen)
		copy(newItems, al.items)
		al.Clear()
		al.items = newItems
	}
}
