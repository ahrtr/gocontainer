// Copyright (c) 2019, Benjamin Wang (benjamin_wang@aliyun.com). All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

package priorityqueue

import (
	"reflect"
	"testing"
)

func TestPQLen(t *testing.T) {
	pq := New()

	// add 3 elements
	pq.Add(5)
	pq.Add(6)
	pq.Add(7)
	if pq.Len() != 3 {
		t.Errorf("The length isn't expected, expect: 3, actual: %d\n", pq.Len())
	}
	if pq.IsEmpty() {
		t.Error("The queue shouldn't be empty")
	}

	// remove one element
	if !pq.Remove(6) {
		t.Error("Failed to remove element 6")
	}
	if pq.Len() != 2 {
		t.Errorf("The length isn't expected, expect: 2, actual: %d\n", pq.Len())
	}

	// Clear all elements
	pq.Clear()
	if pq.Len() != 0 {
		t.Errorf("The length isn't expected, expect: 0, actual: %d\n", pq.Len())
	}
	if !pq.IsEmpty() {
		t.Error("The queue should be empty")
	}
}

func TestPQValue(t *testing.T) {
	// create priority queue
	pq := New()
	pq.Add(15)
	pq.Add(19)
	pq.Add(12)
	pq.Add(8)
	pq.Add(13)
	if pq.Len() != 5 {
		t.Errorf("The length isn't expected, expect: 5, actual: %d\n", pq.Len())
	}

	// Peek
	v1 := pq.Peek()
	if v1 != 8 {
		t.Errorf("The head element isn't expected, expect: 8, actual: %v\n", v1)
	}
	if pq.Len() != 5 {
		t.Errorf("The length isn't expected, expect: 5, actual: %d\n", pq.Len())
	}

	// Contains
	if !pq.Contains(12) {
		t.Error("The queue should contain 12")
	}

	// Poll
	v1 = pq.Poll()
	if v1 != 8 {
		t.Errorf("The head element isn't expected, expect: 8, actual: %v\n", v1)
	}
	if pq.Len() != 4 {
		t.Errorf("The length isn't expected, expect: 4, actual: %d\n", pq.Len())
	}

	v1 = pq.Poll()
	if v1 != 12 {
		t.Errorf("The head element isn't expected, expect: 12, actual: %v\n", v1)
	}
	if pq.Len() != 3 {
		t.Errorf("The length isn't expected, expect: 3, actual: %d\n", pq.Len())
	}

	// Contains (again)
	if pq.Contains(12) {
		t.Error("The queue shouldn't contain 12")
	}

	// Remove
	if !pq.Contains(15) {
		t.Error("The queue should contain 15")
	}
	if !pq.Remove(15) {
		t.Error("Failed to remove element 15")
	}
	if pq.Contains(15) {
		t.Error("The queue shouldn't contain 15")
	}
}

func TestPQReverseSort(t *testing.T) {
	// create priority queue
	pq := Reverse(New())
	reverseSort(t, pq)
}

func TestPQReverseSortWithComparator(t *testing.T) {
	// create priority queue
	pq := NewWithComparator(&myInt{})
	reverseSort(t, pq)
}

func reverseSort(t *testing.T, pq Interface) {
	pq.Add(15)
	pq.Add(19)
	pq.Add(12)
	pq.Add(8)
	pq.Add(13)
	if pq.Len() != 5 {
		t.Errorf("The length isn't expected, expect: 5, actual: %d\n", pq.Len())
	}

	// Peek
	v1 := pq.Peek()
	if v1 != 19 {
		t.Errorf("The head element isn't expected, expect: 19, actual: %v\n", v1)
	}
	if pq.Len() != 5 {
		t.Errorf("The length isn't expected, expect: 5, actual: %d\n", pq.Len())
	}

	// Contains
	if !pq.Contains(12) {
		t.Error("The queue should contain 12")
	}

	// Poll
	v1 = pq.Poll()
	if v1 != 19 {
		t.Errorf("The head element isn't expected, expect: 19, actual: %v\n", v1)
	}
	if pq.Len() != 4 {
		t.Errorf("The length isn't expected, expect: 4, actual: %d\n", pq.Len())
	}

	v1 = pq.Poll()
	if v1 != 15 {
		t.Errorf("The head element isn't expected, expect: 15, actual: %v\n", v1)
	}
	if pq.Len() != 3 {
		t.Errorf("The length isn't expected, expect: 3, actual: %d\n", pq.Len())
	}

	// Contains (again)
	if pq.Contains(15) {
		t.Error("The queue shouldn't contain 15")
	}

	// Remove
	if !pq.Contains(12) {
		t.Error("The queue should contain 12")
	}
	if !pq.Remove(12) {
		t.Error("Failed to remove element 12")
	}
	if pq.Contains(12) {
		t.Error("The queue shouldn't contain 12")
	}
}

type myInt struct{}

// Compare returns reverse order
func (i myInt) Compare(v1, v2 interface{}) (int, error) {
	i1, i2 := v1.(int), v2.(int)
	if i1 < i2 {
		return 1, nil
	}

	if i1 > i2 {
		return -1, nil
	}

	return 0, nil
}

func TestPQComparator(t *testing.T) {
	pq := NewWithComparator(&student{})

	pq.Add(&student{name: "benjamin", age: 34})
	pq.Add(&student{name: "alice", age: 21})
	pq.Add(&student{name: "john", age: 42})
	pq.Add(&student{name: "roy", age: 28})
	pq.Add(&student{name: "moss", age: 25})

	if pq.Len() != 5 {
		t.Errorf("The length isn't expected, expect: 5, actual: %d\n", pq.Len())
	}

	// Peek
	v, ok := pq.Peek().(*student)
	if !ok {
		t.Errorf("The value type peeked from queue isn't expected, expect: &student, actual: %s\n", reflect.TypeOf(v))
	}
	if v.name != "john" || v.age != 42 {
		t.Errorf("The head element isn't expected, expected: {john, 42}, actual: %v\n", v)
	}

	// Poll: 1 {"john", 42}
	v, ok = pq.Poll().(*student)
	if !ok {
		t.Errorf("The value type polled from queue isn't expected, expect: &student, actual: %s\n", reflect.TypeOf(v))
	}
	if v.name != "john" || v.age != 42 {
		t.Errorf("The head element isn't expected, expected: {john, 42}, actual: %v\n", v)
	}
	// Poll: 2 {"benjamin", 34}
	v, ok = pq.Poll().(*student)
	if !ok {
		t.Errorf("The value type polled from queue isn't expected, expect: &student, actual: %s\n", reflect.TypeOf(v))
	}
	if v.name != "benjamin" || v.age != 34 {
		t.Errorf("The head element isn't expected, expected: {john, 42}, actual: %v\n", v)
	}
	// Poll: 3 {"roy", 28}
	v, ok = pq.Poll().(*student)
	if !ok {
		t.Errorf("The value type polled from queue isn't expected, expect: &student, actual: %s\n", reflect.TypeOf(v))
	}
	if v.name != "roy" || v.age != 28 {
		t.Errorf("The head element isn't expected, expected: {john, 42}, actual: %v\n", v)
	}
	// Poll: 4 {"moss", 25}
	v, ok = pq.Poll().(*student)
	if !ok {
		t.Errorf("The value type polled from queue isn't expected, expect: &student, actual: %s\n", reflect.TypeOf(v))
	}
	if v.name != "moss" || v.age != 25 {
		t.Errorf("The head element isn't expected, expected: {john, 42}, actual: %v\n", v)
	}
	// Poll: 5 {"alice", 21}
	v, ok = pq.Poll().(*student)
	if !ok {
		t.Errorf("The value type polled from queue isn't expected, expect: &student, actual: %s\n", reflect.TypeOf(v))
	}
	if v.name != "alice" || v.age != 21 {
		t.Errorf("The head element isn't expected, expected: {john, 42}, actual: %v\n", v)
	}

	// The queue should be empty now
	if pq.Len() != 0 {
		t.Errorf("The length isn't expected, expect: 0, actual: %d\n", pq.Len())
	}
	last := pq.Poll()
	if last != nil {
		t.Errorf("The queue should be empty, but got: %v\n", last)
	}
}

type student struct {
	name string
	age  int
}

// Compare returns -1, 0 or 1 when the first student's age is greater, equal to, or less than the second student's age.
func (s *student) Compare(v1, v2 interface{}) (int, error) {
	s1, s2 := v1.(*student), v2.(*student)
	if s1.age < s2.age {
		return 1, nil
	}
	if s1.age > s2.age {
		return -1, nil
	}
	return 0, nil
}
