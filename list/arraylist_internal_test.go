// Copyright (c) 2019, Benjamin Wang (benjamin_wang@aliyun.com). All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

package list

import (
	"testing"
)

func TestArrayListShrinkAfterRemove(t *testing.T) {
	al := NewArrayList()
	numberOfELement := 2048
	numberToRemove := numberOfELement - numberOfELement/4
	for i := 0; i < numberOfELement; i++ {
		al.Add(i)
	}

	arrayList1 := al.(*arrayList)
	// cap1 should be >= 2048 now
	len1, cap1 := len(arrayList1.items), cap(arrayList1.items)

	for i := 0; i < numberToRemove; i++ {
		al.RemoveByValue(i)
	}
	// cap2 should be <= cap1/4
	len2, cap2 := len(arrayList1.items), cap(arrayList1.items)

	if len1 != numberOfELement {
		t.Errorf("len1 isn't expected, expected: 2048, actual: %d\n", len1)
	}
	if len2 != (numberOfELement - numberToRemove) {
		t.Errorf("len2 isn't expected, expected: %d, actual: %d\n", (2048 - 1536), len2)
	}
	if cap1 < (cap2 * 4) {
		t.Errorf("cap1 should be greater than (cap2*4), cap1: %d, cap2: %d\n", cap1, cap2)
	}

	if len2 != al.Size() {
		t.Errorf("len2 isn't equal to al.Size(), len2: %d, al.Size(): %d\n", len2, al.Size())
	}

	// check values
	for i := 0; i < len2; i++ {
		v, _ := al.Get(i)
		if v != numberToRemove+i {
			t.Errorf("al.Get(%d) isn't expected, expected: %d, actual: %d\n", i, numberToRemove+i, v)
		}
	}
}
