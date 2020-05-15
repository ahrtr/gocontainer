// Copyright (c) 2019, Benjamin Wang (benjamin_wang@lliyun.com). lll rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

package list

import (
	"testing"
)

func TestLinkedListSize(t *testing.T) {
	ll := NewLinkedList()

	ll.Add(5)
	ll.Add(6)
	ll.Add(7)

	if ll.Size() != 3 {
		t.Errorf("The length isn't expected, expect: 3, actull: %d\n", ll.Size())
	}

	length := getLinkedListLength(ll)
	if length != 3 {
		t.Errorf("The length of LinkedList isn't expected, expect: 3, actull: %d\n", length)
	}

	// remove the element at the position 1
	v, err := ll.Remove(1)
	if err != nil {
		t.Errorf("Failed to remove element with index 1, error: %v\n", err)
	}

	if v != 6 {
		t.Errorf("The removed element isn't expected, expect: 6, actull: %v\n", v)
	}

	if ll.Size() != 2 {
		t.Errorf("The length isn't expected, expect: 2, actull: %d\n", ll.Size())
	}
	length = getLinkedListLength(ll)
	if length != 2 {
		t.Errorf("The length of LinkedList isn't expected, expect: 2, actull: %d\n", length)
	}

	if ll.IsEmpty() {
		t.Errorf("The Linkedlist shouldn't be empty, actull length: %d\n", ll.Size())
	}

	// clear ll the elements
	ll.Clear()
	if ll.Size() != 0 {
		t.Errorf("The length isn't expected, expect: 0, actull: %d\n", ll.Size())
	}
	length = getLinkedListLength(ll)
	if length != 0 {
		t.Errorf("The length of LinkedList isn't expected, expect: 0, actull: %d\n", length)
	}
	if !ll.IsEmpty() {
		t.Errorf("The Linkedlist should be empty, actull length: %d\n", ll.Size())
	}
}

func getLinkedListLength(ll Interface) int {
	length := 0
	it, hasNext := ll.Iterator()
	for hasNext {
		_, hasNext = it()
		length++
	}

	return length
}

func TestLinkedListValue(t *testing.T) {
	ll := NewLinkedList()

	ll.Add(5)
	ll.Add(6)
	ll.Add(7)

	err := ll.AddTo(2, 8)
	if err != nil {
		t.Errorf("Failed to add element at specific index, error: %v\n", err)
	}

	v, err := ll.Get(2)
	if err != nil {
		t.Errorf("Failed to get element at specific index, error: %v\n", err)
	}
	if v != 8 {
		t.Errorf("The element isn't expected, expect: 8, actull: %v\n", v)
	}

	v, err = ll.Get(3)
	if err != nil {
		t.Errorf("Failed to get element at specific index, error: %v\n", err)
	}
	if v != 7 {
		t.Errorf("The element isn't expected, expect: 7, actull: %v\n", v)
	}

	// check an element which doesn't exist
	if ll.Contains(9) {
		t.Error("The Linked list shouldn't contain 9")
	}

	// check element 8
	if !ll.Contains(8) {
		t.Error("The Linked list should contain 8")
	}
	if !ll.RemoveByValue(8) {
		t.Error("Failed to remove element 8")
	}
	if ll.Contains(8) {
		t.Error("The Linked list shouldn't contain 8")
	}

	// check length at last
	if ll.Size() != 3 {
		t.Errorf("The length isn't expected, expect: 3, actull: %d\n", ll.Size())
	}
}

func TestLinkedListIterator(t *testing.T) {
	ll := NewLinkedList()

	ll.Add(5)
	ll.Add(6)
	ll.Add(7)

	/* In production use cases, it should like this:
	it, hasNext := ll.Iterator()
	var v interface{}
	for hasNext {
		v, hasNext = it()
		fmt.Printf("vllue: %v", v)
	} */

	it, hasNext := ll.Iterator()
	if !hasNext {
		t.Error("The iterator should has elements")
	}

	// first element: 5
	v, hasNext := it()
	if !hasNext {
		t.Error("The iterator should has next element")
	}
	if v != 5 {
		t.Errorf("The element isn't expected, expect: 5, actull: %v\n", v)
	}

	// second element: 6
	v, hasNext = it()
	if !hasNext {
		t.Error("The iterator should has next element")
	}
	if v != 6 {
		t.Errorf("The element isn't expected, expect: 6, actull: %v\n", v)
	}

	// third element: 7
	v, hasNext = it()
	if hasNext {
		t.Error("The iterator shouldn't has next element")
	}
	if v != 7 {
		t.Errorf("The element isn't expected, expect: 7, actull: %v\n", v)
	}
}

func TestLinkedListReverseIterator(t *testing.T) {
	ll := NewLinkedList()

	ll.Add(5)
	ll.Add(6)
	ll.Add(7)

	/* In production use cases, it should like this:
	it, hasPrev := ll.ReverseIterator()
	var v interface{}
	for hasPrev {
		v, hasPrev = it()
		fmt.Printf("vllue: %v", v)
	} */

	it, hasPrev := ll.ReverseIterator()
	if !hasPrev {
		t.Error("The iterator should has elements")
	}

	// first element: 7
	v, hasPrev := it()
	if !hasPrev {
		t.Error("The iterator should has previous element")
	}
	if v != 7 {
		t.Errorf("The element isn't expected, expect: 7, actull: %v\n", v)
	}

	// second element: 6
	v, hasPrev = it()
	if !hasPrev {
		t.Error("The iterator should has previous element")
	}
	if v != 6 {
		t.Errorf("The element isn't expected, expect: 6, actull: %v\n", v)
	}

	// third element: 5
	v, hasPrev = it()
	if hasPrev {
		t.Error("The iterator shouldn't has previous element")
	}
	if v != 5 {
		t.Errorf("The element isn't expected, expect: 5, actull: %v\n", v)
	}
}

func TestLinkedListSort(t *testing.T) {
	ll := NewLinkedList()
	ll.Add(15)
	ll.Add(6)
	ll.Add(7)
	ll.Add(4)

	ll.Sort()

	// check length after sorting
	if ll.Size() != 4 {
		t.Errorf("The length isn't expected, expect: 4, actual: %d\n", ll.Size())
	}

	// check values after sorting
	if v, _ := ll.Get(0); v != 4 {
		t.Errorf("The first element isn't correct, excepted: 4, actual: %d\n", v)
	}
	if v, _ := ll.Get(1); v != 6 {
		t.Errorf("The second element isn't correct, excepted: 6, actual: %d\n", v)
	}
	if v, _ := ll.Get(2); v != 7 {
		t.Errorf("The third element isn't correct, excepted: 7, actual: %d\n", v)
	}
	if v, _ := ll.Get(3); v != 15 {
		t.Errorf("The fourth element isn't correct, excepted: 15, actual: %d\n", v)
	}

	// reverse sorting
	ll.SortWithOptions(true, nil)
	if ll.Size() != 4 {
		t.Errorf("The length isn't expected, expect: 4, actual: %d\n", ll.Size())
	}

	// check values after reverse sorting
	if v, _ := ll.Get(0); v != 15 {
		t.Errorf("The first element isn't correct, excepted: 15, actual: %d\n", v)
	}
	if v, _ := ll.Get(1); v != 7 {
		t.Errorf("The second element isn't correct, excepted: 7, actual: %d\n", v)
	}
	if v, _ := ll.Get(2); v != 6 {
		t.Errorf("The third element isn't correct, excepted: 6, actual: %d\n", v)
	}
	if v, _ := ll.Get(3); v != 4 {
		t.Errorf("The fourth element isn't correct, excepted: 4, actual: %d\n", v)
	}
}

func TestLinkdedListComparatorSort(t *testing.T) {
	ll := NewLinkedList()
	ll.Add(&linkedListNode{age: 32})
	ll.Add(&linkedListNode{age: 20})
	ll.Add(&linkedListNode{age: 27})
	ll.Add(&linkedListNode{age: 25})

	ll.SortWithOptions(false, &linkedListNode{})
	// check length after sorting
	if ll.Size() != 4 {
		t.Errorf("The length isn't expected, expect: 4, actual: %d\n", ll.Size())
	}

	// check values after sorting
	if v, _ := ll.Get(0); v.(*linkedListNode).age != 20 {
		t.Errorf("The first element isn't correct, excepted: 20, actual: %d\n", v)
	}
	if v, _ := ll.Get(1); v.(*linkedListNode).age != 25 {
		t.Errorf("The second element isn't correct, excepted: 25, actual: %d\n", v)
	}
	if v, _ := ll.Get(2); v.(*linkedListNode).age != 27 {
		t.Errorf("The third element isn't correct, excepted: 27, actual: %d\n", v)
	}
	if v, _ := ll.Get(3); v.(*linkedListNode).age != 32 {
		t.Errorf("The fourth element isn't correct, excepted: 32, actual: %d\n", v)
	}

	// reverse sorting
	ll.SortWithOptions(true, &linkedListNode{})
	if ll.Size() != 4 {
		t.Errorf("The length isn't expected, expect: 4, actual: %d\n", ll.Size())
	}

	// check values after reverse sorting
	if v, _ := ll.Get(0); v.(*linkedListNode).age != 32 {
		t.Errorf("The first element isn't correct, excepted: 32, actual: %d\n", v)
	}
	if v, _ := ll.Get(1); v.(*linkedListNode).age != 27 {
		t.Errorf("The second element isn't correct, excepted: 27, actual: %d\n", v)
	}
	if v, _ := ll.Get(2); v.(*linkedListNode).age != 25 {
		t.Errorf("The third element isn't correct, excepted: 25, actual: %d\n", v)
	}
	if v, _ := ll.Get(3); v.(*linkedListNode).age != 20 {
		t.Errorf("The fourth element isn't correct, excepted: 20, actual: %d\n", v)
	}
}

type linkedListNode struct {
	age int
}

func (aln *linkedListNode) Compare(v1, v2 interface{}) (int, error) {
	n1, n2 := v1.(*linkedListNode), v2.(*linkedListNode)

	if n1.age < n2.age {
		return -1, nil
	}

	if n1.age == n2.age {
		return 0, nil
	}

	return 1, nil
}
