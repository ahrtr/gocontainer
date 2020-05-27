// Copyright (c) 2019, Benjamin Wang (benjamin_wang@aliyun.com). All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

package priorityqueue_test

import (
	"reflect"
	"testing"

	"github.com/ahrtr/gocontainer/queue/priorityqueue"
)

func TestPQSize(t *testing.T) {
	pq := priorityqueue.New()

	// add 3 elements
	pq.Add(5, 6, 7)
	if pq.Size() != 3 {
		t.Errorf("The length isn't expected, expect: 3, actual: %d\n", pq.Size())
	}
	if pq.IsEmpty() {
		t.Error("The queue shouldn't be empty")
	}

	// remove one element
	if !pq.Remove(6) {
		t.Error("Failed to remove element 6")
	}
	if pq.Size() != 2 {
		t.Errorf("The length isn't expected, expect: 2, actual: %d\n", pq.Size())
	}

	// Clear all elements
	pq.Clear()
	if pq.Size() != 0 {
		t.Errorf("The length isn't expected, expect: 0, actual: %d\n", pq.Size())
	}
	if !pq.IsEmpty() {
		t.Error("The queue should be empty")
	}
}

func TestPQValue(t *testing.T) {
	// create priority queue
	pq := priorityqueue.New()
	pq.Add(15, 19, 12, 8, 13)
	if pq.Size() != 5 {
		t.Errorf("The length isn't expected, expect: 5, actual: %d\n", pq.Size())
	}

	// Peek
	v1 := pq.Peek()
	if v1 != 8 {
		t.Errorf("The head element isn't expected, expect: 8, actual: %v\n", v1)
	}
	if pq.Size() != 5 {
		t.Errorf("The length isn't expected, expect: 5, actual: %d\n", pq.Size())
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
	if pq.Size() != 4 {
		t.Errorf("The length isn't expected, expect: 4, actual: %d\n", pq.Size())
	}

	v1 = pq.Poll()
	if v1 != 12 {
		t.Errorf("The head element isn't expected, expect: 12, actual: %v\n", v1)
	}
	if pq.Size() != 3 {
		t.Errorf("The length isn't expected, expect: 3, actual: %d\n", pq.Size())
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

/*-----------------------------------------------------------------------------
// Test: Add, Poll, Contains
-----------------------------------------------------------------------------*/
func TestPQMinHeap(t *testing.T) {
	pq := priorityqueue.New()
	pqTestPQSortImpl(t, pq, []interface{}{15, 19, 12, 8, 13}, []interface{}{8, 12, 13, 15, 19})
}

func TestPQMinHeapWithComparator(t *testing.T) {
	pq := priorityqueue.New().WithComparator(&myInt{})
	pqTestPQSortImpl(t, pq, []interface{}{15, 19, 12, 8, 13}, []interface{}{19, 15, 13, 12, 8})
}

func TestPQMaxHeap(t *testing.T) {
	pq := priorityqueue.New().WithMinHeap(false)
	pqTestPQSortImpl(t, pq, []interface{}{15, 19, 12, 8, 13}, []interface{}{19, 15, 13, 12, 8})
}

func TestPQMaxHeapWithComparator(t *testing.T) {
	pq := priorityqueue.New().WithComparator(&myInt{}).WithMinHeap(false)
	pqTestPQSortImpl(t, pq, []interface{}{15, 19, 12, 8, 13}, []interface{}{8, 12, 13, 15, 19})
}

func pqTestPQSortImpl(t *testing.T, pq priorityqueue.Interface, input, expected []interface{}) {
	pq.Add(input...)

	if pq.Size() != len(input) {
		t.Errorf("The length isn't expected, expect: %d, actual: %d\n", len(input), pq.Size())
	}

	for i := 0; i < len(expected); i++ {
		v := pq.Poll()
		if v != expected[i] {
			t.Errorf("pq.Poll() returned an unexpected value, expected: %v, actual: %v\n", expected[i], v)
		}
	}

	if pq.Size() != 0 {
		t.Errorf("The length isn't expected, expect: 0, actual: %d\n", pq.Size())
	}
}

/*-----------------------------------------------------------------------------
// Test: Add, Remove, Contains, Poll
-----------------------------------------------------------------------------*/
func TestPQDeleteMinHeap(t *testing.T) {
	pq := priorityqueue.New()
	pqTestPQDeleteImpl(t, pq, []interface{}{15, 19, 12, 8, 13}, []interface{}{8, 12, 13, 15}, 19)
}

func TestPQDeleteMinHeapWithComparator(t *testing.T) {
	pq := priorityqueue.New().WithComparator(&myInt{})
	pqTestPQDeleteImpl(t, pq, []interface{}{15, 19, 12, 8, 13}, []interface{}{19, 13, 12, 8}, 15)
}

func TestPQDeleteMaxHeap(t *testing.T) {
	pq := priorityqueue.New().WithMinHeap(false)
	pqTestPQDeleteImpl(t, pq, []interface{}{15, 19, 12, 8, 13}, []interface{}{19, 15, 13, 8}, 12)
}

func TestPQDeleteMaxHeapWithComparator(t *testing.T) {
	pq := priorityqueue.New().WithComparator(&myInt{}).WithMinHeap(false)
	pqTestPQDeleteImpl(t, pq, []interface{}{15, 19, 12, 8, 13}, []interface{}{12, 13, 15, 19}, 8)
}

func pqTestPQDeleteImpl(t *testing.T, pq priorityqueue.Interface, input, expected []interface{}, val interface{}) {
	pq.Add(input...)

	if !pq.Remove(val) {
		t.Errorf("Failed to remove value: %v\n", val)
	}

	if pq.Size() != len(input)-1 {
		t.Errorf("The length isn't expected, expect: %d, actual: %d\n", len(input)-1, pq.Size())
	}

	if pq.Contains(val) {
		t.Errorf("The PQ shouldn't contain value: %v\n", val)
	}

	for i := 0; i < len(expected); i++ {
		v := pq.Poll()
		if v != expected[i] {
			t.Errorf("pq.Poll() returned an unexpected value, expected[%d]: %v, actual: %v\n", i, expected[i], v)
		}
	}

	if pq.Size() != 0 {
		t.Errorf("The length isn't expected, expect: 0, actual: %d\n", pq.Size())
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
	pq := priorityqueue.New().WithComparator(&student{})

	pq.Add(&student{name: "benjamin", age: 34})
	pq.Add(&student{name: "alice", age: 21})
	pq.Add(&student{name: "john", age: 42})
	pq.Add(&student{name: "roy", age: 28})
	pq.Add(&student{name: "moss", age: 25})

	if pq.Size() != 5 {
		t.Errorf("The length isn't expected, expect: 5, actual: %d\n", pq.Size())
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
	if pq.Size() != 0 {
		t.Errorf("The length isn't expected, expect: 0, actual: %d\n", pq.Size())
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
