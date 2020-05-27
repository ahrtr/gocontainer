// Copyright (c) 2019, Benjamin Wang (benjamin_wang@aliyun.com). All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

package linkedmap_test

import (
	"testing"

	"github.com/ahrtr/gocontainer/map/linkedmap"
)

func TestLinkedMapSize(t *testing.T) {
	lm := linkedmap.New()
	lm.Put(24, "benjamin")
	lm.Put(43, "alice")
	lm.Put(18, "john")

	if lm.Size() != 3 {
		t.Errorf("The length isn't expected, expect: 3, actual: %d\n", lm.Size())
	}

	lm.Remove(43)
	if lm.Size() != 2 {
		t.Errorf("The length isn't expected, expect: 2, actual: %d\n", lm.Size())
	}

	lm.Clear()
	if lm.Size() != 0 {
		t.Errorf("The length isn't expected, expect: 0, actual: %d\n", lm.Size())
	}
	if !lm.IsEmpty() {
		t.Error("The container should be empty")
	}
}

func TestLinkedMapValue(t *testing.T) {
	lm := linkedmap.New()
	keys := []int{24, 43, 18, 23, 35}
	values := []string{"benjamin", "alice", "john", "tom", "bill"}
	for i := 0; i < len(keys); i++ {
		lm.Put(keys[i], values[i])
	}

	// test ContainsKey & ContainsValue
	for _, k := range keys {
		if !lm.ContainsKey(k) {
			t.Errorf("The linkedMap should contain key: %d\n", k)
		}
	}

	for _, v := range values {
		if !lm.ContainsValue(v) {
			t.Errorf("The linkedMap should contain value: %s\n", v)
		}
	}

	// test Get & GetOrDefault
	for i, k := range keys {
		v := lm.Get(k)
		if v != values[i] {
			t.Errorf("The value associated with key: %v isn't expected, expect: %v, actual %v\n", k, values[i], v)
		}
	}

	v := lm.GetOrDefault(50, "defaultName")
	if v != "defaultName" {
		t.Errorf("The returned value should be the default value, but actual: %v\n", v)
	}

	// test Remove, RemoveFirstElement and RemoveLastElement
	v, ok := lm.Remove(43)
	if v != "alice" || !ok {
		t.Errorf("Failed to remove element with key 43, returned value: %v, success: %t\n", v, ok)
	}

	k, v, ok := lm.RemoveFirstElement()
	if k != 24 || v != "benjamin" || !ok {
		t.Errorf("Failed to remove the first element, key: %v, value: %v, success: %t\n", k, v, ok)
	}

	k, v, ok = lm.RemoveLastElement()
	if k != 35 || v != "bill" || !ok {
		t.Errorf("Failed to remove the last element, key: %v, value: %v, success: %t\n", k, v, ok)
	}

	if lm.Size() != 2 {
		t.Errorf("The length isn't expected, expect: 2, actual: %d\n", lm.Size())
	}
}

func TestLinkedMapIterate(t *testing.T) {
	lm := linkedmap.New()
	keys := []interface{}{24, 43, 18, 23, 35}
	values := []interface{}{"benjamin", "alice", "john", "tom", "bill"}
	for i := 0; i < len(keys); i++ {
		lm.Put(keys[i], values[i])
	}

	checkIterateResult(t, lm, keys, values)
	checkReverseIterateResult(t, lm, []interface{}{35, 23, 18, 43, 24}, []interface{}{"bill", "tom", "john", "alice", "benjamin"})
}

func TestLinkedMapAccessOrder(t *testing.T) {
	lm := linkedmap.New().WithAccessOrder(true)
	keys := []int{24, 43, 18, 23, 35}
	values := []string{"benjamin", "alice", "john", "tom", "bill"}
	for i := 0; i < len(keys); i++ {
		lm.Put(keys[i], values[i])
	}

	lm.Get(23)
	lm.Get(24)
	lm.Get(18)
	lm.Get(35)
	lm.Get(43)

	checkIterateResult(t, lm, []interface{}{23, 24, 18, 35, 43}, []interface{}{"tom", "benjamin", "john", "bill", "alice"})
	checkReverseIterateResult(t, lm, []interface{}{43, 35, 18, 24, 23}, []interface{}{"alice", "bill", "john", "benjamin", "tom"})
}

func checkIterateResult(t *testing.T, lm linkedmap.Interface, expectedKey, expectedValue []interface{}) {
	it, hasNext := lm.Iterator()
	var k, v interface{}

	for i := 0; i < len(expectedKey); i++ {
		if !hasNext {
			t.Error("Unexpectedly reaching the end")
			break
		}

		k, v, hasNext = it()
		if k != expectedKey[i] || v != expectedValue[i] {
			t.Errorf("Unexpected key or value, Iterate: %d, expect: (%v, %v), actual: (%v, %v)\n", i, expectedKey[i], expectedValue[i], k, v)
		}
	}

	if hasNext {
		t.Error("The iterator should have already reached the end")
	}
}

func checkReverseIterateResult(t *testing.T, lm linkedmap.Interface, expectedKey, expectedValue []interface{}) {
	it, hasPrev := lm.ReverseIterator()
	var k, v interface{}

	for i := 0; i < len(expectedKey); i++ {
		if !hasPrev {
			t.Error("Unexpectedly reaching the end")
			break
		}

		k, v, hasPrev = it()
		if k != expectedKey[i] || v != expectedValue[i] {
			t.Errorf("Unexpected key or value, Iterate: %d, expect: (%v, %v), actual: (%v, %v)\n", i, expectedKey[i], expectedValue[i], k, v)
		}
	}

	if hasPrev {
		t.Error("The iterator should have already reached the end")
	}
}
