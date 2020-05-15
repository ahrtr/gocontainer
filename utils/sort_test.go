// Copyright (c) 2019, Benjamin Wang (benjamin_wang@aliyun.com). All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

package utils

import "testing"

func TestSort(t *testing.T) {
	input1 := []interface{}{6, 4, 9, 19, 15}
	expected1 := []interface{}{4, 6, 9, 15, 19}
	sortTestImpl(t, input1, expected1, false, nil)

	input2 := []interface{}{"benjamin", "alice", "john", "tom", "roy"}
	expected2 := []interface{}{"alice", "benjamin", "john", "roy", "tom"}
	sortTestImpl(t, input2, expected2, false, nil)
}

func TestSortWithComparator(t *testing.T) {
	input1 := []interface{}{6, 4, 9, 19, 15}
	expected1 := []interface{}{19, 15, 9, 6, 4}
	sortTestImpl(t, input1, expected1, false, reverseInt{})

	input2 := []interface{}{"benjamin", "alice", "john", "tom", "roy"}
	expected2 := []interface{}{"tom", "roy", "john", "benjamin", "alice"}
	sortTestImpl(t, input2, expected2, false, reverseString{})
}

func TestReverseSort(t *testing.T) {
	input1 := []interface{}{6, 4, 9, 19, 15}
	expected1 := []interface{}{19, 15, 9, 6, 4}
	sortTestImpl(t, input1, expected1, true, nil)

	input2 := []interface{}{"benjamin", "alice", "john", "tom", "roy"}
	expected2 := []interface{}{"tom", "roy", "john", "benjamin", "alice"}
	sortTestImpl(t, input2, expected2, true, nil)
}

func TestReverseSortWithComparator(t *testing.T) {
	input1 := []interface{}{6, 4, 9, 19, 15}
	expected1 := []interface{}{4, 6, 9, 15, 19}
	sortTestImpl(t, input1, expected1, true, reverseInt{})

	input2 := []interface{}{"benjamin", "alice", "john", "tom", "roy"}
	expected2 := []interface{}{"alice", "benjamin", "john", "roy", "tom"}
	sortTestImpl(t, input2, expected2, true, reverseString{})
}

func sortTestImpl(t *testing.T, input []interface{}, expected []interface{}, reverse bool, c Comparator) {
	// sort
	if reverse {
		ReverseSort(input, c)
	} else {
		Sort(input, c)
	}

	// check result
	for i := 0; i < len(input); i++ {
		if input[i] != expected[i] {
			t.Errorf("Doesn't match, input[%d] = %v, expected[%d] = %v\n", i, input[i], i, expected[i])
		}
	}
}

type reverseString struct{}

// Compare returns reverse order for string
func (i reverseString) Compare(v1, v2 interface{}) (int, error) {
	i1, i2 := v1.(string), v2.(string)

	if i1 < i2 {
		return 1, nil
	}
	if i1 > i2 {
		return -1, nil
	}
	return 0, nil
}

type reverseInt struct{}

// Compare returns reverse order for int
func (i reverseInt) Compare(v1, v2 interface{}) (int, error) {
	i1, i2 := v1.(int), v2.(int)

	if i1 < i2 {
		return 1, nil
	}
	if i1 > i2 {
		return -1, nil
	}
	return 0, nil
}
