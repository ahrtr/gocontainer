// Copyright (c) 2019, Benjamin Wang (benjamin_wang@aliyun.com). All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

// Package linkedmap implements a linked hashmap, based on a map and a doubly linked list. The iteration ordering is normally
// the order in which keys were inserted into the map, or the order in which the keys were accessed if the accessOrder flag is set.
// If a linkedMap is configured as access-order, then the first element in the list is the eldest element, which means it's the least recently inserted
// or accessed element; while the last element is the newest element, which means it's the most recently inserted or accessed element.
//
// To iterate over an linkedMap (where lm is an instance of linkedMap.Interface):
//	it, hasNext := lm.Iterator()
//  var k, v interface{}
//	for hasNext {
//		k, v, hasNext = it()
//		// do something with k & v
//	}
//
// To iterate over an linkedMap in reverse order (where lm is an instance of linkedMap.Interface):
//	it, hasPrev := lm.ReverseIterator()
//  var k, v interface{}
//	for hasPrev {
//		k, v, hasPrev = it()
//		// do something with k & v
//	}
//
package linkedmap

import (
	"github.com/ahrtr/gocontainer/collection"
)

// Interface is a type of linked map, and linkedMap implements this interface.
type Interface interface {
	collection.Interface

	// Len returns the number of elements in the linkedMap.
	Len() int
	// Put associates the specified value with the specified key in this map. If the map previously contained a mapping for the key,
	// the old value is replaced by the specified value.
	// It returns the previous value associated with the specified key, or nil if there was no mapping for the key.
	// A nil return can also indicate that the map previously associated nil with the specified key.
	Put(k, v interface{}) interface{}
	// WithAccessOrder configures the iteration ordering for this linked map,
	// true for access-order, and false for insertion-order.
	WithAccessOrder(accessOrder bool) Interface
	// Get returns the value to which the specified key is mapped, or nil if this map contains no mapping for the key.
	Get(k interface{}) interface{}
	// GetOrDefault returns the value to which the specified key is mapped, or the defaultValue if this map contains no mapping for the key.
	GetOrDefault(k, defaultValue interface{}) interface{}
	// ContainsKey returns true if this map contains a mapping for the specified key.
	ContainsKey(k interface{}) bool
	// ContainsValue returns true if this map maps one or more keys to the specified value.
	ContainsValue(v interface{}) bool
	// Remove removes the mapping for a key from this map if it is present.
	// It returns the value to which this map previously associated the key, and true,
	// or nil and false if the map contained no mapping for the key.
	Remove(k interface{}) (interface{}, bool)
	// RemoveFirstElement removes the first element from this map, which is the head of the list.
	// It returns the (key, value, true) if the map isn't empty, or (nil, nil, false) if the map is empty.
	RemoveFirstElement() (interface{}, interface{}, bool)
	// RemoveLastElement removes the last element from this map, which is the tail of the list.
	// It returns the (key, value, true) if the map isn't empty, or (nil, nil, false) if the map is empty.
	RemoveLastElement() (interface{}, interface{}, bool)

	// Iterator returns an iterator over the elements in this map in proper sequence.
	Iterator() (func() (interface{}, interface{}, bool), bool)
	// ReverseIterator returns an iterator over the elements in this map in reverse sequence as Iterator.
	ReverseIterator() (func() (interface{}, interface{}, bool), bool)
}

type element struct {
	key   interface{}
	value interface{}
	prev  *element
	next  *element
}

// linkedMap implements the Interface.
type linkedMap struct {
	data        map[interface{}]*element
	accessOrder bool
	head        *element
	tail        *element
	length      int
}

// New creates a linkedMap.
func New() Interface {
	return &linkedMap{
		data:        map[interface{}]*element{},
		accessOrder: false,
		head:        nil,
		tail:        nil,
		length:      0,
	}
}

func (lm *linkedMap) WithAccessOrder(accessOrder bool) Interface {
	lm.accessOrder = accessOrder
	return lm
}

func (lm *linkedMap) Len() int {
	return lm.length
}

func (lm *linkedMap) IsEmpty() bool {
	return lm.Len() == 0
}

func (lm *linkedMap) Put(k, v interface{}) interface{} {
	var retVal interface{} = nil
	if oldElement, ok := lm.data[k]; ok {
		retVal = oldElement.value
		oldElement.value = v
		// move the element to the end of the list
		if lm.accessOrder {
			lm.unlink(oldElement)
			lm.linkLast(oldElement)
		}
	} else {
		e := &element{
			key:   k,
			value: v,
		}
		lm.data[k] = e
		lm.linkLast(e)
	}

	return retVal
}

func (lm *linkedMap) Get(k interface{}) interface{} {
	if oldElement, ok := lm.data[k]; ok {
		// move the element to the end of the list
		if lm.accessOrder {
			lm.unlink(oldElement)
			lm.linkLast(oldElement)
		}
		return oldElement.value
	}

	return nil
}

func (lm *linkedMap) GetOrDefault(k, defaultValue interface{}) interface{} {
	if oldElement, ok := lm.data[k]; ok {
		// move the element to the end of the list
		if lm.accessOrder {
			lm.unlink(oldElement)
			lm.linkLast(oldElement)
		}
		return oldElement.value
	}

	return defaultValue
}

func (lm *linkedMap) ContainsKey(k interface{}) bool {
	_, ok := lm.data[k]
	return ok
}

func (lm *linkedMap) ContainsValue(v interface{}) bool {
	e := lm.head
	for e != nil {
		if e.value == v {
			return true
		}
		e = e.next
	}

	return false
}

func (lm *linkedMap) Remove(k interface{}) (interface{}, bool) {
	if oldElement, ok := lm.data[k]; ok {
		retVal := oldElement.value
		delete(lm.data, k)
		lm.unlink(oldElement)
		oldElement.key, oldElement.value = nil, nil
		return retVal, true
	}

	return nil, false
}

func (lm *linkedMap) RemoveFirstElement() (interface{}, interface{}, bool) {
	if lm.head != nil {
		e := lm.head
		k, v := e.key, e.value

		lm.unlink(e)
		e.key, e.value = nil, nil

		return k, v, true
	}

	return nil, nil, false
}

func (lm *linkedMap) RemoveLastElement() (interface{}, interface{}, bool) {
	if lm.tail != nil {
		e := lm.tail
		k, v := e.key, e.value

		lm.unlink(e)
		e.key, e.value = nil, nil

		return k, v, true
	}

	return nil, nil, false
}

func (lm *linkedMap) Clear() {
	lm.data = map[interface{}]*element{}

	for e := lm.head; e != nil; {
		next := e.next
		e.prev, e.next, e.key, e.value = nil, nil, nil, nil
		e = next
	}

	lm.head, lm.tail, lm.length = nil, nil, 0
}

func (lm *linkedMap) Iterator() (func() (interface{}, interface{}, bool), bool) {
	e := lm.head

	return func() (interface{}, interface{}, bool) {
		var k, v interface{}
		if e != nil {
			k = e.key
			v = e.value
			e = e.next
		} else {
			k, v = nil, nil
		}
		return k, v, e != nil
	}, e != nil
}

func (lm *linkedMap) ReverseIterator() (func() (interface{}, interface{}, bool), bool) {
	e := lm.tail

	return func() (interface{}, interface{}, bool) {
		var k, v interface{}
		if e != nil {
			k = e.key
			v = e.value
			e = e.prev
		} else {
			k, v = nil, nil
		}
		return k, v, e != nil
	}, e != nil
}

// linkLast links val as last element.
func (lm *linkedMap) linkLast(e *element) {
	e.prev, e.next = lm.tail, nil

	if nil == lm.tail {
		lm.head, lm.tail = e, e
	} else {
		lm.tail.next = e
		lm.tail = e
	}

	lm.length++
}

// unlink removes the specified element e in this list.
func (lm *linkedMap) unlink(e *element) {
	ePrev, eNext := e.prev, e.next
	e.prev, e.next = nil, nil

	if nil == ePrev {
		lm.head = eNext
	} else {
		ePrev.next = eNext
	}

	if nil == eNext {
		lm.tail = ePrev
	} else {
		eNext.prev = ePrev
	}

	lm.length--
}
