// Copyright (c) 2019, Benjamin Wang (benjamin_wang@aliyun.com). All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

package queue_test

import (
	"reflect"
	"testing"

	"github.com/ahrtr/gocontainer/queue"
)

func TestQueueSize(t *testing.T) {
	q := queue.New()

	q.Add(5, 6)

	if q.Size() != 2 {
		t.Errorf("The length isn't expected, expect: 2, actual: %d", q.Size())
	}
}

func TestQueuePeek(t *testing.T) {
	q := queue.New()

	q.Add(5, "hello")

	val1, ok := q.Peek().(int)
	if !ok {
		t.Errorf("The value type peeked from queue isn't expected, expect: int, actual: %s\n", reflect.TypeOf(val1).String())
	}

	if val1 != 5 {
		t.Errorf("The value peeked from queue isn't expected, expect: 5, actual: %d\n", val1)
	}

	val2, ok := q.Peek().(int)
	if !ok {
		t.Errorf("The value type peeked from queue isn't expected, expect: int, actual: %s\n", reflect.TypeOf(val1).String())
	}

	if val2 != 5 {
		t.Errorf("The value peeked from queue isn't expected, expect: 5, actual: %d\n", val2)
	}
}

func TestQueuePoll(t *testing.T) {
	q := queue.New()

	q.Add(5, "hello")

	val1, ok := q.Poll().(int)
	if !ok {
		t.Errorf("The value type polled from queue isn't expected, expect: int, actual: %s\n", reflect.TypeOf(val1).String())
	}

	if val1 != 5 {
		t.Errorf("The value polled from queue isn't expected, expect: 5, actual: %d\n", val1)
	}

	val2, ok := q.Poll().(string)
	if !ok {
		t.Errorf("The value type polled from queue isn't expected, expect: string, actual: %s\n", reflect.TypeOf(val1).String())
	}

	if val2 != "hello" {
		t.Errorf("The value polled from queue isn't expected, expect: hello, actual: %s\n", val2)
	}
}

func TestQueueIsEmpty(t *testing.T) {
	q := queue.New()

	q.Add(5, 6)

	if isEmpty1 := q.IsEmpty(); isEmpty1 {
		t.Errorf("The queue shouldn't be empty\n")
	}

	q.Clear()
	if isEmpty2 := q.IsEmpty(); !isEmpty2 {
		t.Errorf("The queue should be empty\n")
	}
}

func TestQueueClear(t *testing.T) {
	q := queue.New()

	q.Add(5, 6)
	q.Clear()

	if q.Size() != 0 {
		t.Errorf("The length isn't expected, expect: 0, actual: %d\n", q.Size())
	}
}
