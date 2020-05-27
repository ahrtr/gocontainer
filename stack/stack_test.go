// Copyright (c) 2019, Benjamin Wang (benjamin_wang@aliyun.com). All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

package stack_test

import (
	"reflect"
	"testing"

	"github.com/ahrtr/gocontainer/stack"
)

func TestStackSize(t *testing.T) {
	s := stack.New()

	s.Push(5)
	s.Push(6)

	if s.Size() != 2 {
		t.Errorf("The length isn't expected, expect: 2, actual: %d", s.Size())
	}
}

func TestStackValue(t *testing.T) {
	s := stack.New()

	s.Push(5)
	s.Push("hello")

	// Peek "hello"
	val1, ok := s.Peek().(string)
	if !ok {
		t.Errorf("The value type peeked from stack isn't expected, expect: string, actual: %s\n", reflect.TypeOf(val1).String())
	}

	if val1 != "hello" {
		t.Errorf("The value peeked from stack isn't expected, expect: hello, actual: %s\n", val1)
	}

	// Pop "hello"
	val2, ok := s.Pop().(string)
	if !ok {
		t.Errorf("The value type popped from stack isn't expected, expect: string, actual: %s\n", reflect.TypeOf(val2).String())
	}

	if val2 != "hello" {
		t.Errorf("The value popped from stack isn't expected, expect: hello, actual: %s\n", val2)
	}

	// Peek 5
	val3, ok := s.Peek().(int)
	if !ok {
		t.Errorf("The value type popped from stack isn't expected, expect: int, actual: %s\n", reflect.TypeOf(val3).String())
	}

	if val3 != 5 {
		t.Errorf("The value popped from stack isn't expected, expect: 5, actual: %d\n", val3)
	}

	// Pop 5
	val4, ok := s.Pop().(int)
	if !ok {
		t.Errorf("The value type popped from stack isn't expected, expect: int, actual: %s\n", reflect.TypeOf(val4).String())
	}

	if val4 != 5 {
		t.Errorf("The value popped from stack isn't expected, expect: 5, actual: %d\n", val4)
	}
}

func TestStackIsEmpty(t *testing.T) {
	s := stack.New()

	s.Push(5)
	s.Push(6)

	if isEmpty1 := s.IsEmpty(); isEmpty1 {
		t.Errorf("The stack shouldn't be empty\n")
	}

	s.Clear()
	if isEmpty2 := s.IsEmpty(); !isEmpty2 {
		t.Errorf("The stack should be empty\n")
	}
}

func TestStackClear(t *testing.T) {
	s := stack.New()

	s.Push(5)
	s.Push(6)
	s.Clear()

	if s.Size() != 0 {
		t.Errorf("The length isn't expected, expect: 0, actual: %d", s.Size())
	}
}
