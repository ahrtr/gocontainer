// Copyright (c) 2019, Benjamin Wang (benjamin_wang@aliyun.com). All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"reflect"
	"runtime"
)

func main() {
	funcs := []func(){
		stackExample1,

		queueExample1,

		setExample1,

		priorityqueueExample1,
		priorityqueueExample2,
		priorityqueueExample3,
		priorityqueueExample4,

		listArrayListExample1,
		listArrayListExample2,
		listArrayListExample3,
		listArrayListExample4,

		listLinkedListExample1,
		listLinkedListExample2,
		listLinkedListExample3,
		listLinkedListExample4,

		linkedmapExample1,
		linkedmapExample2,
		linkedmapExample3,

		btreeExample1,
		btreeExample2,
		btreeExample3,
		btreeExample4,
		btreeExample5,
	}

	for _, f := range funcs {
		fmt.Printf("\n%s starting...\n", runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name())
		f()
	}
	fmt.Printf("\n****** Total examples: %d\n", len(funcs))
}

// printFuncName prints each example function's name.
// This function isn't needed anymore.
func printFuncName() {
	pc, _, _, _ := runtime.Caller(1)
	fmt.Printf("\n%s starting...\n", runtime.FuncForPC(pc).Name())
}

// student implements sort.Comparator
type student struct {
	name string
	age  int
}

// Compare makes sure the students are printed in descending order by age.
func (s *student) Compare(v1, v2 interface{}) (int, error) {
	s1, s2 := v1.(student), v2.(student)
	if s1.age < s2.age {
		return 1, nil
	}
	if s1.age > s2.age {
		return -1, nil
	}
	return 0, nil
}

type reverseString struct{}

// Compare returns reverse order
func (i *reverseString) Compare(v1, v2 interface{}) (int, error) {
	i1, i2 := v1.(string), v2.(string)

	if i1 < i2 {
		return 1, nil
	}
	if i1 > i2 {
		return -1, nil
	}
	return 0, nil
}
