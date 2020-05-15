// Copyright (c) 2019, Benjamin Wang (benjamin_wang@aliyun.com). All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

package set

import (
	"testing"
)

func TestSetSize(t *testing.T) {
	s := New()

	s.Add(5, 6)

	if s.Size() != 2 {
		t.Errorf("The length isn't expected, expect: 2, actual: %d", s.Size())
	}
	if s.IsEmpty() {
		t.Errorf("The set shouldn't be empty\n")
	}

	s.Clear()
	if !s.IsEmpty() {
		t.Errorf("The set should be empty\n")
	}
}

func TestSetValue(t *testing.T) {
	s := New()

	s.Add(5, "hello")

	if !s.Contains(5) {
		t.Error("The value 5 isn't found in this set")
	}

	if !s.Contains("hello") {
		t.Error("The value 'hello' isn't found in this set")
	}

	if s.Contains(6) {
		t.Error("The value 6 isn't supposed to be in this set")
	}

	if s.Contains("world") {
		t.Error("The value 'world' isn't supposed to be in this set")
	}

	if !s.Remove(5) {
		t.Error("Failed to remove the value 5 in this set")
	}
	if s.Contains(5) {
		t.Error("The value 5 isn't supposed to be in this set after being deleted")
	}

	if !s.Remove("hello") {
		t.Error("Failed to remove the value 'hello' in this set")
	}
	if s.Contains("hello") {
		t.Error("The value 'hello' isn't supposed to be in this set after being deleted")
	}
}

func TestSetIterate(t *testing.T) {
	s := New()

	s.Add(5, "hello")

	s.Iterate(func(v interface{}) bool {
		switch tt := v.(type) {
		case int:
			if v.(int) != 5 {
				t.Errorf("The int value isn't expected, expect: 5, actual: %d\n", v)
			}
		case string:
			if v.(string) != "hello" {
				t.Errorf("The string value isn't expected, expect: hello, actual: %s\n", v)
			}
		default:
			t.Errorf("Unexpected type %T\n", tt)
		}

		return true
	})
}
