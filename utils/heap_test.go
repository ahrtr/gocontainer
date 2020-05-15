// Copyright (c) 2019, Benjamin Wang (benjamin_wang@aliyun.com). All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

package utils

import "testing"

/*-----------------------------------------------------------------------------
// Test: HeapPostPush and HeapPrePop
-----------------------------------------------------------------------------*/
func TestHeap(t *testing.T) {
	input1 := []interface{}{6, 4, 9, 19, 15}
	expected1 := []interface{}{4, 6, 9, 15, 19}
	heapTestImpl(t, input1, expected1, true, nil)

	input2 := []interface{}{"benjamin", "alice", "john", "tom", "roy"}
	expected2 := []interface{}{"alice", "benjamin", "john", "roy", "tom"}
	heapTestImpl(t, input2, expected2, true, nil)
}

func TestHeapWithComparator(t *testing.T) {
	input1 := []interface{}{6, 4, 9, 19, 15}
	expected1 := []interface{}{19, 15, 9, 6, 4}
	heapTestImpl(t, input1, expected1, true, reverseInt{})

	input2 := []interface{}{"benjamin", "alice", "john", "tom", "roy"}
	expected2 := []interface{}{"tom", "roy", "john", "benjamin", "alice"}
	heapTestImpl(t, input2, expected2, true, reverseString{})
}

func TestMaxHeap(t *testing.T) {
	input1 := []interface{}{6, 4, 9, 19, 15}
	expected1 := []interface{}{19, 15, 9, 6, 4}
	heapTestImpl(t, input1, expected1, false, nil)

	input2 := []interface{}{"benjamin", "alice", "john", "tom", "roy"}
	expected2 := []interface{}{"tom", "roy", "john", "benjamin", "alice"}
	heapTestImpl(t, input2, expected2, false, nil)
}

func TestMaxHeapWithComparator(t *testing.T) {
	input1 := []interface{}{6, 4, 9, 19, 15}
	expected1 := []interface{}{4, 6, 9, 15, 19}
	heapTestImpl(t, input1, expected1, false, reverseInt{})

	input2 := []interface{}{"benjamin", "alice", "john", "tom", "roy"}
	expected2 := []interface{}{"alice", "benjamin", "john", "roy", "tom"}
	heapTestImpl(t, input2, expected2, false, reverseString{})
}

func heapTestImpl(t *testing.T, input []interface{}, expected []interface{}, isMinHeap bool, c Comparator) {
	// construct a heap
	heapSlice := []interface{}{}
	for _, v := range input {
		heapSlice = append(heapSlice, v)
		HeapPostPush(heapSlice, isMinHeap, c)
	}

	// Pop all elements from heap one by one
	for i := 0; i < len(expected); i++ {
		HeapPrePop(heapSlice, isMinHeap, c)
		ret := heapSlice[len(heapSlice)-1]
		heapSlice = heapSlice[:(len(heapSlice) - 1)]

		if ret != expected[i] {
			t.Errorf("Doesn't match, Pop(heap) = %v, expected[%d] = %v\n", ret, i, expected[i])
		}
	}

	if len(heapSlice) != 0 {
		t.Errorf("len(heapSlice) should be 0, but actual: %d\n", len(heapSlice))
	}
}

/*-----------------------------------------------------------------------------
// Test: HeapInit and HeapPreRemove
-----------------------------------------------------------------------------*/
func TestHeapRemove(t *testing.T) {
	input1 := []interface{}{6, 4, 9, 19, 15}
	expected1 := []interface{}{4, 6, 9, 19}
	heapRemoveTestImpl(t, input1, expected1, 15, true, nil)

	input2 := []interface{}{"benjamin", "alice", "john", "tom", "roy"}
	expected2 := []interface{}{"alice", "benjamin", "roy", "tom"}
	heapRemoveTestImpl(t, input2, expected2, "john", true, nil)
}

func TestHeapRemoveWithComparator(t *testing.T) {
	input1 := []interface{}{6, 4, 9, 19, 15}
	expected1 := []interface{}{19, 15, 6, 4}
	heapRemoveTestImpl(t, input1, expected1, 9, true, reverseInt{})

	input2 := []interface{}{"benjamin", "alice", "john", "tom", "roy"}
	expected2 := []interface{}{"tom", "john", "benjamin", "alice"}
	heapRemoveTestImpl(t, input2, expected2, "roy", true, reverseString{})
}

func TestMaxHeapRemove(t *testing.T) {
	input1 := []interface{}{6, 4, 9, 19, 15}
	expected1 := []interface{}{15, 9, 6, 4}
	heapRemoveTestImpl(t, input1, expected1, 19, false, nil)

	input2 := []interface{}{"benjamin", "alice", "john", "tom", "roy"}
	expected2 := []interface{}{"tom", "roy", "john", "benjamin"}
	heapRemoveTestImpl(t, input2, expected2, "alice", false, nil)
}

func TestMaxHeapRemoveWithComparator(t *testing.T) {
	input1 := []interface{}{6, 4, 9, 19, 15}
	expected1 := []interface{}{4, 6, 9, 15}
	heapRemoveTestImpl(t, input1, expected1, 19, false, reverseInt{})

	input2 := []interface{}{"benjamin", "alice", "john", "tom", "roy"}
	expected2 := []interface{}{"alice", "benjamin", "john", "roy"}
	heapRemoveTestImpl(t, input2, expected2, "tom", false, reverseString{})
}

func heapRemoveTestImpl(t *testing.T, input []interface{}, expected []interface{}, val interface{}, isMinHeap bool, c Comparator) {
	// establish a heap in-place
	HeapInit(input, isMinHeap, c)

	// find the index of the value to be removed
	index := 0
	for i := 0; i < len(input); i++ {
		if input[i] == val {
			index = i
			break
		}
	}

	// call HeapPreRemove
	HeapPreRemove(input, index, isMinHeap, c)
	valRemoved := input[len(input)-1]
	if valRemoved != val {
		t.Errorf("The removed value isn't expected, expected: %v, actual %v\n", val, valRemoved)
	}
	input = input[:(len(input) - 1)]

	// Pop all elements from heap one by one
	for i := 0; i < len(expected); i++ {
		HeapPrePop(input, isMinHeap, c)
		ret := input[len(input)-1]
		input = input[:(len(input) - 1)]

		if ret != expected[i] {
			t.Errorf("Doesn't match, Pop(heap) = %v, expected[%d] = %v\n", ret, i, expected[i])
		}
	}

	if len(input) != 0 {
		t.Errorf("len(input) should be 0, but actual: %d\n", len(input))
	}
}

/*-----------------------------------------------------------------------------
// Test: HeapInit and HeapPostUpdate
-----------------------------------------------------------------------------*/
func TestHeapFix(t *testing.T) {
	input1 := []interface{}{6, 4, 9, 19, 15}
	expected1 := []interface{}{4, 6, 9, 19, 25}
	heapFixTestImpl(t, input1, expected1, 15, 25, true, nil)

	input2 := []interface{}{"benjamin", "alice", "john", "tom", "roy"}
	expected2 := []interface{}{"alice", "benjamin", "ken", "roy", "tom"}
	heapFixTestImpl(t, input2, expected2, "john", "ken", true, nil)
}

func TestHeapFixWithComparator(t *testing.T) {
	input1 := []interface{}{6, 4, 9, 19, 15}
	expected1 := []interface{}{19, 15, 6, 4, 3}
	heapFixTestImpl(t, input1, expected1, 9, 3, true, reverseInt{})

	input2 := []interface{}{"benjamin", "alice", "john", "tom", "roy"}
	expected2 := []interface{}{"tom", "john", "benjamin", "alice", "ali"}
	heapFixTestImpl(t, input2, expected2, "roy", "ali", true, reverseString{})
}

func TestMaxHeapFix(t *testing.T) {
	input1 := []interface{}{6, 4, 9, 19, 15}
	expected1 := []interface{}{15, 13, 9, 6, 4}
	heapFixTestImpl(t, input1, expected1, 19, 13, false, nil)

	input2 := []interface{}{"benjamin", "alice", "john", "tom", "roy"}
	expected2 := []interface{}{"trevor", "tom", "roy", "john", "benjamin"}
	heapFixTestImpl(t, input2, expected2, "alice", "trevor", false, nil)
}

func TestMaxHeapFixWithComparator(t *testing.T) {
	input1 := []interface{}{6, 4, 9, 19, 15}
	expected1 := []interface{}{4, 6, 7, 9, 15}
	heapFixTestImpl(t, input1, expected1, 19, 7, false, reverseInt{})

	input2 := []interface{}{"benjamin", "alice", "john", "tom", "roy"}
	expected2 := []interface{}{"alice", "benjamin", "john", "roy", "zoo"}
	heapFixTestImpl(t, input2, expected2, "tom", "zoo", false, reverseString{})
}

func heapFixTestImpl(t *testing.T, input []interface{}, expected []interface{}, oldVal, newVal interface{}, isMinHeap bool, c Comparator) {
	// establish a heap in-place
	HeapInit(input, isMinHeap, c)

	// find the index of the value to be updated
	index := 0
	for i := 0; i < len(input); i++ {
		if input[i] == oldVal {
			index = i
			// update the value
			input[index] = newVal
			break
		}
	}

	// call HeapPostUpdate
	HeapPostUpdate(input, index, isMinHeap, c)

	// Pop all elements from heap one by one
	for i := 0; i < len(expected); i++ {
		HeapPrePop(input, isMinHeap, c)
		ret := input[len(input)-1]
		input = input[:(len(input) - 1)]

		if ret != expected[i] {
			t.Errorf("Doesn't match, Pop(heap) = %v, expected[%d] = %v\n", ret, i, expected[i])
		}
	}

	if len(input) != 0 {
		t.Errorf("len(input) should be 0, but actual: %d\n", len(input))
	}
}
