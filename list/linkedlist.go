// Copyright (c) 2019, Benjamin Wang (benjamin_wang@aliyun.com). All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

// Package list implements both an ArrayList and a LinkedList.
//
// To iterate over an LinkedList (where ll is a *LinkedList):
//	it, hasNext := ll.Iterator()
//  var v interface{}
//	for hasNext {
//		v, hasNext = it()
//		// do something with v
//	}
//
// To iterate over an LinkedList in reverse order (where ll is a *LinkedList):
//	it, hasPrev := ll.ReverseIterator()
//  var v interface{}
//	for hasPrev {
//		v, hasPrev = it()
//		// do something with v
//	}
//
package list

import (
	"fmt"

	"github.com/ahrtr/gocontainer/sort"
)

type element struct {
	next, prev *element
	// The value stored with this element.
	value interface{}
}

// LinkedList represents a doubly linked list.
// It implements the interface list.Interface.
type LinkedList struct {
	head   *element
	tail   *element
	length int
	cmp    sort.Comparator
}

// NewLinkedList initializes and returns an LinkedList.
func NewLinkedList() Interface {
	return &LinkedList{
		head:   nil,
		tail:   nil,
		length: 0,
		cmp:    nil,
	}
}

// NewLinkedListWithComparator initializes and returns an LinkedList with a comparator.
func NewLinkedListWithComparator(c sort.Comparator) Interface {
	return &LinkedList{
		head:   nil,
		tail:   nil,
		length: 0,
		cmp:    c,
	}
}

func (ll *LinkedList) Len() int {
	return ll.length
}

func (ll *LinkedList) Less(i, j int) bool {
	v1, v2 := ll.getElement(i).value, ll.getElement(j).value

	var cmpRet int
	var err error
	if nil != ll.cmp {
		cmpRet, err = ll.cmp.Compare(v1, v2)
	} else {
		cmpRet, err = sort.Compare(v1, v2)
	}
	if err != nil {
		panic(err)
	}

	return cmpRet < 0
}

func (ll *LinkedList) Swap(i, j int) {
	if i == j {
		return
	}

	// For simplicity, just swap the values directly.
	e1, e2 := ll.getElement(i), ll.getElement(j)
	e1.value, e2.value = e2.value, e1.value
}

func (ll *LinkedList) IsEmpty() bool {
	return ll.Len() == 0
}

func (ll *LinkedList) Add(val interface{}) {
	ll.linkLast(val)
}

// linkLast links val as last element.
func (ll *LinkedList) linkLast(val interface{}) {
	e := element{
		prev:  ll.tail,
		next:  nil,
		value: val,
	}

	if nil == ll.tail {
		ll.head, ll.tail = &e, &e
	} else {
		ll.tail.next = &e
		ll.tail = &e
	}
	ll.length++
}

func (ll *LinkedList) AddTo(index int, val interface{}) error {
	size := ll.Len()
	if index < 0 || index > size {
		return fmt.Errorf("Index out of range, index:%d, len:%d", index, size)
	}

	if index == size {
		ll.linkLast(val)
	} else {
		ll.linkBefore(val, ll.getElement(index))
	}

	return nil
}

// linkBefore inserts val before non-null element e.
func (ll *LinkedList) linkBefore(val interface{}, e *element) {
	newElement := element{
		prev:  nil,
		next:  e,
		value: val,
	}

	if e != nil {
		newElement.prev = e.prev
		if e.prev != nil {
			e.prev.next = &newElement
		} else {
			ll.head = &newElement
		}
		e.prev = &newElement
	} else {
		ll.head, ll.tail = &newElement, &newElement
	}

	ll.length++
}

// getElement returns the element at the specified positon.
func (ll *LinkedList) getElement(index int) *element {
	size := ll.Len()
	var e *element
	if index < (size >> 1) {
		e = ll.head
		for i := 0; i < index; i++ {
			e = e.next
		}
	} else {
		e = ll.tail
		for i := (size - 1); i > index; i-- {
			e = e.prev
		}
	}
	return e
}

func (ll *LinkedList) Contains(val interface{}) bool {
	return ll.indexOf(val) >= 0
}

// indexOf returns the index of the first occurence of the specified element
// in this list, or -1 if this list does not contain the element.
func (ll *LinkedList) indexOf(val interface{}) int {
	index := 0

	for e := ll.head; e != nil; e = e.next {
		if val == e.value {
			return index
		}
		index++
	}

	return -1
}

func (ll *LinkedList) Get(index int) (interface{}, error) {
	size := ll.Len()
	if index < 0 || index >= size {
		return nil, fmt.Errorf("Index out of range, index:%d, len:%d", index, size)
	}

	return ll.getElement(index).value, nil
}

func (ll *LinkedList) Remove(index int) (interface{}, error) {
	size := ll.Len()
	if index < 0 || index >= size {
		return nil, fmt.Errorf("Index out of range, index:%d, len:%d", index, size)
	}

	return ll.unlink(ll.getElement(index)), nil
}

// unlink removes the specified element e in this list.
func (ll *LinkedList) unlink(e *element) interface{} {
	if nil == e {
		return nil
	}

	retValue := e.value

	if nil == e.prev {
		ll.head = e.next
	} else {
		e.prev.next = e.next
	}

	if nil == e.next {
		ll.tail = e.prev
	} else {
		e.next.prev = e.prev
	}

	e.prev, e.next, e.value = nil, nil, nil
	ll.length--

	return retValue
}

func (ll *LinkedList) RemoveByValue(val interface{}) bool {
	if ll.Len() == 0 {
		return false
	}

	for e := ll.head; e != nil; e = e.next {
		if val == e.value {
			ll.unlink(e)
			return true
		}
	}

	return false
}

func (ll *LinkedList) Clear() {
	for e := ll.head; e != nil; {
		next := e.next
		e.prev, e.next, e.value = nil, nil, nil
		e = next
	}

	ll.head, ll.tail, ll.length = nil, nil, 0
}

func (ll *LinkedList) Iterator() (func() (interface{}, bool), bool) {
	e := ll.head

	return func() (interface{}, bool) {
		var element interface{}
		if e != nil {
			element = e.value
			e = e.next
		} else {
			element = nil
		}
		return element, e != nil
	}, e != nil
}

func (ll *LinkedList) ReverseIterator() (func() (interface{}, bool), bool) {
	e := ll.tail

	return func() (interface{}, bool) {
		var element interface{}
		if e != nil {
			element = e.value
			e = e.prev
		} else {
			element = nil
		}
		return element, e != nil
	}, e != nil
}
