// Copyright (c) 2019, Benjamin Wang (benjamin_wang@aliyun.com). All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

package list

import (
	"testing"
)

func TestArrayListSize(t *testing.T) {
	al := NewArrayList()

	al.Add(5, 6, 7)

	if al.Size() != 3 {
		t.Errorf("The length isn't expected, expect: 3, actual: %d\n", al.Size())
	}

	// remove the element at the position 1
	v, err := al.Remove(1)
	if err != nil {
		t.Errorf("Failed to remove element with index 1, error: %v\n", err)
	}

	if v != 6 {
		t.Errorf("The removed element isn't expected, expect: 6, actual: %v\n", v)
	}

	if al.Size() != 2 {
		t.Errorf("The length isn't expected, expect: 2, actual: %d\n", al.Size())
	}

	if al.IsEmpty() {
		t.Errorf("The arraylist shouldn't be empty, actual length: %d\n", al.Size())
	}

	// clear al the elements
	al.Clear()
	if al.Size() != 0 {
		t.Errorf("The length isn't expected, expect: 0, actual: %d\n", al.Size())
	}
	if !al.IsEmpty() {
		t.Errorf("The arraylist should be empty, actual length: %d\n", al.Size())
	}
}

func TestArrayListValue(t *testing.T) {
	al := NewArrayList()

	al.Add(5, 6, 7)

	err := al.AddTo(2, 8)
	if err != nil {
		t.Errorf("Failed to add element at specific index, error: %v\n", err)
	}

	v, err := al.Get(2)
	if err != nil {
		t.Errorf("Failed to get element at specific index, error: %v\n", err)
	}
	if v != 8 {
		t.Errorf("The element isn't expected, expect: 8, actual: %v\n", v)
	}

	v, err = al.Get(3)
	if err != nil {
		t.Errorf("Failed to get element at specific index, error: %v\n", err)
	}
	if v != 7 {
		t.Errorf("The element isn't expected, expect: 7, actual: %v\n", v)
	}

	// check an element which doesn't exist
	if al.Contains(9) {
		t.Error("The array list shouldn't contain 9")
	}

	// check element 8
	if !al.Contains(8) {
		t.Error("The array list should contain 8")
	}
	if !al.RemoveByValue(8) {
		t.Error("Failed to remove element 8")
	}
	if al.Contains(8) {
		t.Error("The array list shouldn't contain 8")
	}

	// check length at last
	if al.Size() != 3 {
		t.Errorf("The length isn't expected, expect: 3, actual: %d\n", al.Size())
	}
}

func TestArrayListIterator(t *testing.T) {
	al := NewArrayList()

	al.Add(5, 6, 7)

	/* In production use cases, it should like this:
	it, hasNext := al.Iterator()
	var v interface{}
	for hasNext {
		v, hasNext = it()
		fmt.Printf("value: %v", v)
	} */

	it, hasNext := al.Iterator()
	if !hasNext {
		t.Error("The iterator should has elements")
	}

	// first element: 5
	v, hasNext := it()
	if !hasNext {
		t.Error("The iterator should has next element")
	}
	if v != 5 {
		t.Errorf("The element isn't expected, expect: 5, actual: %v\n", v)
	}

	// second element: 6
	v, hasNext = it()
	if !hasNext {
		t.Error("The iterator should has next element")
	}
	if v != 6 {
		t.Errorf("The element isn't expected, expect: 6, actual: %v\n", v)
	}

	// third element: 7
	v, hasNext = it()
	if hasNext {
		t.Error("The iterator shouldn't has next element")
	}
	if v != 7 {
		t.Errorf("The element isn't expected, expect: 7, actual: %v\n", v)
	}
}

func TestArrayListReverseIterator(t *testing.T) {
	al := NewArrayList()

	al.Add(5, 6, 7)

	/* In production use cases, it should like this:
	it, hasPrev := al.ReverseIterator()
	var v interface{}
	for hasPrev {
		v, hasPrev = it()
		fmt.Printf("value: %v", v)
	} */

	it, hasPrev := al.ReverseIterator()
	if !hasPrev {
		t.Error("The iterator should has elements")
	}

	// first element: 7
	v, hasPrev := it()
	if !hasPrev {
		t.Error("The iterator should has previous element")
	}
	if v != 7 {
		t.Errorf("The element isn't expected, expect: 7, actual: %v\n", v)
	}

	// second element: 6
	v, hasPrev = it()
	if !hasPrev {
		t.Error("The iterator should has previous element")
	}
	if v != 6 {
		t.Errorf("The element isn't expected, expect: 6, actual: %v\n", v)
	}

	// third element: 5
	v, hasPrev = it()
	if hasPrev {
		t.Error("The iterator shouldn't has previous element")
	}
	if v != 5 {
		t.Errorf("The element isn't expected, expect: 5, actual: %v\n", v)
	}
}

func TestArrayListSort(t *testing.T) {
	al := NewArrayList()
	al.Add(15, 6, 7, 4)

	al.Sort()

	// check length after sorting
	if al.Size() != 4 {
		t.Errorf("The length isn't expected, expect: 4, actual: %d\n", al.Size())
	}

	// check values after sorting
	if v, _ := al.Get(0); v != 4 {
		t.Errorf("The first element isn't correct, excepted: 4, actual: %d\n", v)
	}
	if v, _ := al.Get(1); v != 6 {
		t.Errorf("The second element isn't correct, excepted: 6, actual: %d\n", v)
	}
	if v, _ := al.Get(2); v != 7 {
		t.Errorf("The thrid element isn't correct, excepted: 7, actual: %d\n", v)
	}
	if v, _ := al.Get(3); v != 15 {
		t.Errorf("The fourth element isn't correct, excepted: 15, actual: %d\n", v)
	}

	// reverse sorting
	al.SortWithOptions(true, nil)
	if al.Size() != 4 {
		t.Errorf("The length isn't expected, expect: 4, actual: %d\n", al.Size())
	}

	// check values after reverse sorting
	if v, _ := al.Get(0); v != 15 {
		t.Errorf("The first element isn't correct, excepted: 15, actual: %d\n", v)
	}
	if v, _ := al.Get(1); v != 7 {
		t.Errorf("The second element isn't correct, excepted: 7, actual: %d\n", v)
	}
	if v, _ := al.Get(2); v != 6 {
		t.Errorf("The thrid element isn't correct, excepted: 6, actual: %d\n", v)
	}
	if v, _ := al.Get(3); v != 4 {
		t.Errorf("The fourth element isn't correct, excepted: 4, actual: %d\n", v)
	}
}

func TestArrayListComparatorSort(t *testing.T) {
	al := NewArrayList()
	al.Add(&arrayListNode{age: 32})
	al.Add(&arrayListNode{age: 20})
	al.Add(&arrayListNode{age: 27})
	al.Add(&arrayListNode{age: 25})

	al.SortWithOptions(false, &arrayListNode{})
	// check length after sorting
	if al.Size() != 4 {
		t.Errorf("The length isn't expected, expect: 4, actual: %d\n", al.Size())
	}

	// check values after sorting
	if v, _ := al.Get(0); v.(*arrayListNode).age != 20 {
		t.Errorf("The first element isn't correct, excepted: 20, actual: %d\n", v.(*arrayListNode).age)
	}
	if v, _ := al.Get(1); v.(*arrayListNode).age != 25 {
		t.Errorf("The first element isn't correct, excepted: 25, actual: %d\n", v.(*arrayListNode).age)
	}
	if v, _ := al.Get(2); v.(*arrayListNode).age != 27 {
		t.Errorf("The first element isn't correct, excepted: 27, actual: %d\n", v.(*arrayListNode).age)
	}
	if v, _ := al.Get(3); v.(*arrayListNode).age != 32 {
		t.Errorf("The first element isn't correct, excepted: 32, actual: %d\n", v.(*arrayListNode).age)
	}

	// reverse sorting
	al.SortWithOptions(true, &arrayListNode{})
	if al.Size() != 4 {
		t.Errorf("The length isn't expected, expect: 4, actual: %d\n", al.Size())
	}

	// check values after reverse sorting
	if v, _ := al.Get(0); v.(*arrayListNode).age != 32 {
		t.Errorf("The first element isn't correct, excepted: 32, actual: %d\n", v.(*arrayListNode).age)
	}
	if v, _ := al.Get(1); v.(*arrayListNode).age != 27 {
		t.Errorf("The first element isn't correct, excepted: 27, actual: %d\n", v.(*arrayListNode).age)
	}
	if v, _ := al.Get(2); v.(*arrayListNode).age != 25 {
		t.Errorf("The first element isn't correct, excepted: 25, actual: %d\n", v.(*arrayListNode).age)
	}
	if v, _ := al.Get(3); v.(*arrayListNode).age != 20 {
		t.Errorf("The first element isn't correct, excepted: 20, actual: %d\n", v.(*arrayListNode).age)
	}
}

func TestArrayListShrinkAfterRemove(t *testing.T) {
	al := NewArrayList()
	numberOfELement := 2048
	numberToRemove := numberOfELement - numberOfELement/4
	for i := 0; i < numberOfELement; i++ {
		al.Add(i)
	}

	arrayList1 := al.(*arrayList)
	// cap1 should be >= 2048 now
	len1, cap1 := len(arrayList1.items), cap(arrayList1.items)

	for i := 0; i < numberToRemove; i++ {
		al.RemoveByValue(i)
	}
	// cap2 should be <= cap1/4
	len2, cap2 := len(arrayList1.items), cap(arrayList1.items)

	if len1 != numberOfELement {
		t.Errorf("len1 isn't expected, expected: 2048, actual: %d\n", len1)
	}
	if len2 != (numberOfELement - numberToRemove) {
		t.Errorf("len2 isn't expected, expected: %d, actual: %d\n", (2048 - 1536), len2)
	}
	if cap1 < (cap2 * 4) {
		t.Errorf("cap1 should be greater than (cap2*4), cap1: %d, cap2: %d\n", cap1, cap2)
	}

	if len2 != al.Size() {
		t.Errorf("len2 isn't equal to al.Size(), len2: %d, al.Size(): %d\n", len2, al.Size())
	}

	// check values
	for i := 0; i < len2; i++ {
		v, _ := al.Get(i)
		if v != numberToRemove+i {
			t.Errorf("al.Get(%d) isn't expected, expected: %d, actual: %d\n", i, numberToRemove+i, v)
		}
	}
}

type arrayListNode struct {
	age int
}

func (aln *arrayListNode) Compare(v1, v2 interface{}) (int, error) {
	n1, n2 := v1.(*arrayListNode), v2.(*arrayListNode)

	if n1.age < n2.age {
		return -1, nil
	}

	if n1.age == n2.age {
		return 0, nil
	}

	return 1, nil
}
