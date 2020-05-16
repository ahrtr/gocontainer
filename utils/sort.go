// Copyright (c) 2019, Benjamin Wang (benjamin_wang@aliyun.com). All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

package utils

import (
	"sort"
)

type sortableContainer struct {
	items []interface{}
	cmp   Comparator
}

type reverseSortableContainer struct {
	*sortableContainer
}

// Sort sorts values into ascending sequence according to their natural ordering, or according to the provided comparator.
func Sort(values []interface{}, c Comparator) {
	sort.Sort(&sortableContainer{values, c})
}

// ReverseSort sorts the values into opposite ordering to Sort
func ReverseSort(values []interface{}, c Comparator) {
	sort.Sort(&reverseSortableContainer{&sortableContainer{values, c}})
}

func (sc *sortableContainer) Len() int {
	return len(sc.items)
}
func (sc *sortableContainer) Swap(i, j int) {
	sc.items[i], sc.items[j] = sc.items[j], sc.items[i]
}
func (sc *sortableContainer) Less(i, j int) bool {
	var cmpRet int
	var err error
	if nil != sc.cmp {
		cmpRet, err = sc.cmp.Compare(sc.items[i], sc.items[j])
	} else {
		cmpRet, err = Compare(sc.items[i], sc.items[j])
	}
	if err != nil {
		panic(err)
	}
	return cmpRet < 0
}

// Less returns the opposite of the embedded implementation's Less method.
func (sc *reverseSortableContainer) Less(i, j int) bool {
	return sc.sortableContainer.Less(j, i)
}
