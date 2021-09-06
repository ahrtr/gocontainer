package main

import (
	"fmt"
	"github.com/ahrtr/gocontainer/btree"
)

// example for basic usage
func btreeExample1() {
	items := []int {5, 9, 2, 4, 11, 6}
	tr := btree.New(2)

	fmt.Printf("tr.Size(): %d\n", tr.Size()) // should be 0 in the beginning

	// Insert values
	fmt.Printf("Inserting %d items: %v\n", len(items), items)
	for _, item := range items {
		tr.ReplaceOrInsert(item)
	}

	// Search values
	fmt.Printf("    tr.Size(): %d\n", tr.Size()) // should be len(items): 6 now
	fmt.Printf("    tr.Min(): %v\n", tr.Min()) // should be 2
	fmt.Printf("    tr.Max(): %v\n", tr.Max()) // should be 11
	fmt.Printf("    tr.Has(6): %t\n", tr.Has(6))  // true
	fmt.Printf("    tr.Get(6): %v\n", tr.Get(6))  // 6
	fmt.Printf("    tr.Has(7): %t\n", tr.Has(7))  // false
	fmt.Printf("    tr.Get(7): %v\n", tr.Get(7))  // nil

	// Delete values
	fmt.Println("Deleting items:")
	fmt.Printf("    tr.DeleteMin(): %v\n", tr.DeleteMin()) // 2 is deleted and returned
	fmt.Printf("    tr.Min(): %v\n", tr.Min()) // should be 4 now
	fmt.Printf("    tr.DeleteMax(): %v\n", tr.DeleteMax()) // 11 is deleted and returned
	fmt.Printf("    tr.Max(): %v\n", tr.Max()) // should be 9 now
	fmt.Printf("    tr.Delete(6): %v\n", tr.Delete(6)) // 6 is deleted and returned
	fmt.Printf("    tr.Delete(7): %v\n", tr.Delete(7)) // 7 doesn't exist, so nil is returned

	fmt.Printf("tr.Size(): %d\n", tr.Size()) // should be 3 now because 3 items have already been removed
}

// example for search in ascending or descending order
func btreeExample2() {
	items := []int {5, 9, 2, 4, 11, 6}
	tr := btree.New(2)

	// Insert values
	fmt.Printf("Inserting %d items: %v\n", len(items), items)
	for _, item := range items {
		tr.ReplaceOrInsert(item)
	}

	fmt.Println("Iterating in ascending order:")
	fmt.Printf("    tr.Ascend(x): ") // should be 2, 4, 5, 6, 9, 11
	tr.Ascend(func(item interface{}) bool {
		fmt.Printf("%v  ", item)
		return true
	})

	fmt.Printf("\n    tr.AscendRange(4, 9, x): ") // should be in [4, 9): 4, 5, 6
	tr.AscendRange(4, 9, func(item interface{}) bool {
		fmt.Printf("%v  ", item)
		return true
	})

	fmt.Printf("\n    tr.AscendLessThan(9, x): ") // should be in [2, 9): 2, 4, 5, 6
	tr.AscendLessThan(9, func(item interface{}) bool {
		fmt.Printf("%v  ", item)
		return true
	})

	fmt.Printf("\n    tr.AscendGreaterOrEqual(5, x): ") // should be in [5, 11]: 5, 6, 9, 11
	tr.AscendGreaterOrEqual(5, func(item interface{}) bool {
		fmt.Printf("%v  ", item)
		return true
	})

	fmt.Println("\nIterating in descending order:")
	fmt.Printf("    tr.Descend(x): ") // should be 11, 9, 6, 5, 4, 2
	tr.Descend(func(item interface{}) bool {
		fmt.Printf("%v  ", item)
		return true
	})

	fmt.Printf("\n    tr.DescendRange(9, 4, x): ") // should be in [9, 4): 6, 5, 4
	tr.DescendRange(9, 4, func(item interface{}) bool {
		fmt.Printf("%v  ", item)
		return true
	})

	fmt.Printf("\n    tr.DescendGreaterThan(5, x): ") // should be in [11, 5): 6, 9, 11
	tr.DescendGreaterThan(5, func(item interface{}) bool {
		fmt.Printf("%v  ", item)
		return true
	})

	fmt.Printf("\n    tr.DescendLessOrEqual(9, x): ") // should be in [9, 2]: 9, 6, 5, 4, 2
	tr.DescendLessOrEqual(9, func(item interface{}) bool {
		fmt.Printf("%v  ", item)
		return true
	})
	fmt.Println()
}

// example of Clone(). Once a btree is cloned, and the following writes to both the old and new btree use copy-on-write logic.
func btreeExample3() {
	items := []int {5, 9, 2, 4, 11, 6}
	tr := btree.New(2)

	// Insert values
	fmt.Printf("Inserting %d items: %v\n", len(items), items)
	for _, item := range items {
		tr.ReplaceOrInsert(item)
	}

	// In the beginning, both tr and  tr2 share exactly the same structure and data
	fmt.Println("tr2 := tr.Clone()")
	tr2 := tr.Clone()
	fmt.Printf("    tr.Size(): %d, tr2.Size(): %d\n", tr.Size(), tr2.Size())
	fmt.Printf("    tr.Ascend(x): ") // should be 2, 4, 5, 6, 9, 11
	tr.Ascend(func(item interface{}) bool {
		fmt.Printf("%v  ", item)
		return true
	})

	fmt.Printf("\n    tr2.Ascend(x): ") // should be 2, 4, 5, 6, 9, 11
	tr2.Ascend(func(item interface{}) bool {
		fmt.Printf("%v  ", item)
		return true
	})

	fmt.Printf("\ntr.ReplaceOrInsert(7): %v\n",tr.ReplaceOrInsert(7))
	fmt.Printf("tr.ReplaceOrInsert(13): %v\n", tr.ReplaceOrInsert(13))
	fmt.Printf("tr.Delete(9): %v\n", tr.Delete(9))

	fmt.Printf("tr2.ReplaceOrInsert(20): %v\n", tr2.ReplaceOrInsert(20))
	fmt.Printf("tr2.DeleteMin(): %v\n", tr2.DeleteMin())


	fmt.Printf("    tr.Size(): %d, tr2.Size(): %d\n", tr.Size(), tr2.Size())
	fmt.Printf("    tr.Ascend(x): ") // should be 2, 4, 5, 6, 9, 11
	tr.Ascend(func(item interface{}) bool {
		fmt.Printf("%v  ", item)
		return true
	})

	fmt.Printf("\n    tr2.Ascend(x): ") // should be 2, 4, 5, 6, 9, 11
	tr2.Ascend(func(item interface{}) bool {
		fmt.Printf("%v  ", item)
		return true
	})

	fmt.Println()
}

// example on how to use utils.Comparator
func btreeExample4() {
	values := []student{
		{name: "benjamin", age: 28},
		{name: "alice", age: 42},
		{name: "john", age: 35},
		{name: "tom", age: 18},
		{name: "bill", age: 25},
	}

	tr := btree.New(2).WithComparator(&student{})
	for _, v := range values {
		tr.ReplaceOrInsert(v)
	}

	fmt.Printf("tr.Size(): %d\n", tr.Size())
	fmt.Println("tr.Ascend(x):")
	tr.Ascend(func(item interface{}) bool {
		fmt.Printf("    %v\n", item)
		return true
	})
}


// example on how to use FreeList, which useful if you want to share a FreeList in multiple btree instances.
func btreeExample5() {
	freeList := btree.NewFreeList(32)
	tr := btree.NewWithFreeList(2, freeList)

	items := []int {5, 9, 2, 4, 11, 6}

	// Insert values
	fmt.Printf("Inserting %d items: %v\n", len(items), items)
	for _, item := range items {
		tr.ReplaceOrInsert(item)
	}
}