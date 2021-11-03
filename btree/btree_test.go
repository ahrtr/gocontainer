// Copyright 2014 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package btree_test

import (
	"flag"
	"fmt"
	"github.com/ahrtr/gocontainer/btree"
	"math/rand"
	"reflect"
	"sort"
	"sync"
	"testing"
	"time"
)

func init() {
	seed := time.Now().Unix()
	fmt.Println(seed)
	rand.Seed(seed)
}

// perm returns a random permutation of n integers in the range [0, n).
func perm(n int) (out []interface{}) {
	for _, v := range rand.Perm(n) {
		out = append(out, v)
	}
	return
}

// rang returns an ordered list of integers in the range [0, n).
func rang(n int) (out []interface{}) {
	for i := 0; i < n; i++ {
		out = append(out, i)
	}
	return
}

// all extracts all interface{}s from a tree in order as a slice.
func all(t btree.Interface) (out []interface{}) {
	t.Ascend(func(a interface{}) bool {
		out = append(out, a)
		return true
	})
	return
}

// rangerev returns a reversed ordered list of integers in the range [0, n).
func rangrev(n int) (out []interface{}) {
	for i := n - 1; i >= 0; i-- {
		out = append(out, i)
	}
	return
}

// allrev extracts all interface{}s from a tree in reverse order as a slice.
func allrev(t btree.Interface) (out []interface{}) {
	t.Descend(func(a interface{}) bool {
		out = append(out, a)
		return true
	})
	return
}

var btreeDegree = flag.Int("degree", 32, "B-Tree degree")

func TestBTree(t *testing.T) {
	tr := btree.New(*btreeDegree)
	const treeSize = 10000
	for i := 0; i < 10; i++ {
		if min := tr.Min(); min != nil {
			t.Fatalf("empty min, got %+v", min)
		}
		if max := tr.Max(); max != nil {
			t.Fatalf("empty max, got %+v", max)
		}
		for _, item := range perm(treeSize) {
			if x := tr.ReplaceOrInsert(item); x != nil {
				t.Fatal("insert found interface{}", item)
			}
		}
		for _, item := range perm(treeSize) {
			if x := tr.ReplaceOrInsert(item); x == nil {
				t.Fatal("insert didn't find interface{}", item)
			}
		}
		if min, want := tr.Min(), interface{}(0); min != want {
			t.Fatalf("min: want %+v, got %+v", want, min)
		}
		if max, want := tr.Max(), interface{}(treeSize-1); max != want {
			t.Fatalf("max: want %+v, got %+v", want, max)
		}
		got := all(tr)
		want := rang(treeSize)
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("mismatch:\n got: %v\nwant: %v", got, want)
		}

		gotrev := allrev(tr)
		wantrev := rangrev(treeSize)
		if !reflect.DeepEqual(gotrev, wantrev) {
			t.Fatalf("mismatch:\n got: %v\nwant: %v", got, want)
		}

		for _, item := range perm(treeSize) {
			if x := tr.Delete(item); x == nil {
				t.Fatalf("didn't find %v", item)
			}
		}
		if got = all(tr); len(got) > 0 {
			t.Fatalf("some left!: %v", got)
		}
	}
}

//nolint
func ExampleBTree() {
	tr := btree.New(*btreeDegree)
	for i := 0; i < 10; i++ {
		tr.ReplaceOrInsert(i)
	}
	fmt.Println("len:       ", tr.Size())
	fmt.Println("get3:      ", tr.Get(3))
	fmt.Println("get100:    ", tr.Get(100))
	fmt.Println("del4:      ", tr.Delete(4))
	fmt.Println("del100:    ", tr.Delete(100))
	fmt.Println("replace5:  ", tr.ReplaceOrInsert(5))
	fmt.Println("replace100:", tr.ReplaceOrInsert(100))
	fmt.Println("min:       ", tr.Min())
	fmt.Println("delmin:    ", tr.DeleteMin())
	fmt.Println("max:       ", tr.Max())
	fmt.Println("delmax:    ", tr.DeleteMax())
	fmt.Println("len:       ", tr.Size())
	// Output:
	// len:        10
	// get3:       3
	// get100:     <nil>
	// del4:       4
	// del100:     <nil>
	// replace5:   5
	// replace100: <nil>
	// min:        0
	// delmin:     0
	// max:        100
	// delmax:     100
	// len:        8
}

func TestDeleteMin(t *testing.T) {
	tr := btree.New(3)
	for _, v := range perm(100) {
		tr.ReplaceOrInsert(v)
	}
	var got []interface{}
	for v := tr.DeleteMin(); v != nil; v = tr.DeleteMin() {
		got = append(got, v)
	}
	if want := rang(100); !reflect.DeepEqual(got, want) {
		t.Fatalf("ascendrange:\n got: %v\nwant: %v", got, want)
	}
}

func TestDeleteMax(t *testing.T) {
	tr := btree.New(3)
	for _, v := range perm(100) {
		tr.ReplaceOrInsert(v)
	}
	var got []interface{}
	for v := tr.DeleteMax(); v != nil; v = tr.DeleteMax() {
		got = append(got, v)
	}
	// Reverse our list.
	for i := 0; i < len(got)/2; i++ {
		got[i], got[len(got)-i-1] = got[len(got)-i-1], got[i]
	}
	if want := rang(100); !reflect.DeepEqual(got, want) {
		t.Fatalf("ascendrange:\n got: %v\nwant: %v", got, want)
	}
}

func TestAscendRange(t *testing.T) {
	tr := btree.New(2)
	for _, v := range perm(100) {
		tr.ReplaceOrInsert(v)
	}
	var got []interface{}
	tr.AscendRange(40, 60, func(a interface{}) bool {
		got = append(got, a)
		return true
	})
	if want := rang(100)[40:60]; !reflect.DeepEqual(got, want) {
		t.Fatalf("ascendrange:\n got: %v\nwant: %v", got, want)
	}
	got = got[:0]
	tr.AscendRange(40, 60, func(a interface{}) bool {
		if a.(int) > 50 {
			return false
		}
		got = append(got, a)
		return true
	})
	if want := rang(100)[40:51]; !reflect.DeepEqual(got, want) {
		t.Fatalf("ascendrange:\n got: %v\nwant: %v", got, want)
	}
}

func TestDescendRange(t *testing.T) {
	tr := btree.New(2)
	for _, v := range perm(100) {
		tr.ReplaceOrInsert(v)
	}
	var got []interface{}
	tr.DescendRange(60, 40, func(a interface{}) bool {
		got = append(got, a)
		return true
	})
	if want := rangrev(100)[39:59]; !reflect.DeepEqual(got, want) {
		t.Fatalf("descendrange:\n got: %v\nwant: %v", got, want)
	}
	got = got[:0]
	tr.DescendRange(60, 40, func(a interface{}) bool {
		if a.(int) < 50 {
			return false
		}
		got = append(got, a)
		return true
	})
	if want := rangrev(100)[39:50]; !reflect.DeepEqual(got, want) {
		t.Fatalf("descendrange:\n got: %v\nwant: %v", got, want)
	}
}
func TestAscendLessThan(t *testing.T) {
	tr := btree.New(*btreeDegree)
	for _, v := range perm(100) {
		tr.ReplaceOrInsert(v)
	}
	var got []interface{}
	tr.AscendLessThan(60, func(a interface{}) bool {
		got = append(got, a)
		return true
	})
	if want := rang(100)[:60]; !reflect.DeepEqual(got, want) {
		t.Fatalf("ascendrange:\n got: %v\nwant: %v", got, want)
	}
	got = got[:0]
	tr.AscendLessThan(60, func(a interface{}) bool {
		if a.(int) > 50 {
			return false
		}
		got = append(got, a)
		return true
	})
	if want := rang(100)[:51]; !reflect.DeepEqual(got, want) {
		t.Fatalf("ascendrange:\n got: %v\nwant: %v", got, want)
	}
}

func TestDescendLessOrEqual(t *testing.T) {
	tr := btree.New(*btreeDegree)
	for _, v := range perm(100) {
		tr.ReplaceOrInsert(v)
	}
	var got []interface{}
	tr.DescendLessOrEqual(40, func(a interface{}) bool {
		got = append(got, a)
		return true
	})
	if want := rangrev(100)[59:]; !reflect.DeepEqual(got, want) {
		t.Fatalf("descendlessorequal:\n got: %v\nwant: %v", got, want)
	}
	got = got[:0]
	tr.DescendLessOrEqual(60, func(a interface{}) bool {
		if a.(int) < 50 {
			return false
		}
		got = append(got, a)
		return true
	})
	if want := rangrev(100)[39:50]; !reflect.DeepEqual(got, want) {
		t.Fatalf("descendlessorequal:\n got: %v\nwant: %v", got, want)
	}
}
func TestAscendGreaterOrEqual(t *testing.T) {
	tr := btree.New(*btreeDegree)
	for _, v := range perm(100) {
		tr.ReplaceOrInsert(v)
	}
	var got []interface{}
	tr.AscendGreaterOrEqual(40, func(a interface{}) bool {
		got = append(got, a)
		return true
	})
	if want := rang(100)[40:]; !reflect.DeepEqual(got, want) {
		t.Fatalf("ascendrange:\n got: %v\nwant: %v", got, want)
	}
	got = got[:0]
	tr.AscendGreaterOrEqual(40, func(a interface{}) bool {
		if a.(int) > 50 {
			return false
		}
		got = append(got, a)
		return true
	})
	if want := rang(100)[40:51]; !reflect.DeepEqual(got, want) {
		t.Fatalf("ascendrange:\n got: %v\nwant: %v", got, want)
	}
}

func TestDescendGreaterThan(t *testing.T) {
	tr := btree.New(*btreeDegree)
	for _, v := range perm(100) {
		tr.ReplaceOrInsert(v)
	}
	var got []interface{}
	tr.DescendGreaterThan(40, func(a interface{}) bool {
		got = append(got, a)
		return true
	})
	if want := rangrev(100)[:59]; !reflect.DeepEqual(got, want) {
		t.Fatalf("descendgreaterthan:\n got: %v\nwant: %v", got, want)
	}
	got = got[:0]
	tr.DescendGreaterThan(40, func(a interface{}) bool {
		if a.(int) < 50 {
			return false
		}
		got = append(got, a)
		return true
	})
	if want := rangrev(100)[:50]; !reflect.DeepEqual(got, want) {
		t.Fatalf("descendgreaterthan:\n got: %v\nwant: %v", got, want)
	}
}

const benchmarkTreeSize = 10000

func BenchmarkInsert(b *testing.B) {
	b.StopTimer()
	insertP := perm(benchmarkTreeSize)
	b.StartTimer()
	i := 0
	for i < b.N {
		tr := btree.New(*btreeDegree)
		for _, item := range insertP {
			tr.ReplaceOrInsert(item)
			i++
			if i >= b.N {
				return
			}
		}
	}
}

func BenchmarkSeek(b *testing.B) {
	b.StopTimer()
	size := 100000
	insertP := perm(size)
	tr := btree.New(*btreeDegree)
	for _, item := range insertP {
		tr.ReplaceOrInsert(item)
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		tr.AscendGreaterOrEqual(i%size, func(i interface{}) bool { return false })
	}
}

func BenchmarkDeleteInsert(b *testing.B) {
	b.StopTimer()
	insertP := perm(benchmarkTreeSize)
	tr := btree.New(*btreeDegree)
	for _, item := range insertP {
		tr.ReplaceOrInsert(item)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tr.Delete(insertP[i%benchmarkTreeSize])
		tr.ReplaceOrInsert(insertP[i%benchmarkTreeSize])
	}
}

func BenchmarkDeleteInsertCloneOnce(b *testing.B) {
	b.StopTimer()
	insertP := perm(benchmarkTreeSize)
	tr := btree.New(*btreeDegree)
	for _, item := range insertP {
		tr.ReplaceOrInsert(item)
	}
	tr = tr.Clone()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tr.Delete(insertP[i%benchmarkTreeSize])
		tr.ReplaceOrInsert(insertP[i%benchmarkTreeSize])
	}
}

func BenchmarkDeleteInsertCloneEachTime(b *testing.B) {
	b.StopTimer()
	insertP := perm(benchmarkTreeSize)
	tr := btree.New(*btreeDegree)
	for _, item := range insertP {
		tr.ReplaceOrInsert(item)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tr = tr.Clone()
		tr.Delete(insertP[i%benchmarkTreeSize])
		tr.ReplaceOrInsert(insertP[i%benchmarkTreeSize])
	}
}

func BenchmarkDelete(b *testing.B) {
	b.StopTimer()
	insertP := perm(benchmarkTreeSize)
	removeP := perm(benchmarkTreeSize)
	b.StartTimer()
	i := 0
	for i < b.N {
		b.StopTimer()
		tr := btree.New(*btreeDegree)
		for _, v := range insertP {
			tr.ReplaceOrInsert(v)
		}
		b.StartTimer()
		for _, item := range removeP {
			tr.Delete(item)
			i++
			if i >= b.N {
				return
			}
		}
		if tr.Size() > 0 {
			panic(tr.Size())
		}
	}
}

func BenchmarkGet(b *testing.B) {
	b.StopTimer()
	insertP := perm(benchmarkTreeSize)
	removeP := perm(benchmarkTreeSize)
	b.StartTimer()
	i := 0
	for i < b.N {
		b.StopTimer()
		tr := btree.New(*btreeDegree)
		for _, v := range insertP {
			tr.ReplaceOrInsert(v)
		}
		b.StartTimer()
		for _, item := range removeP {
			tr.Get(item)
			i++
			if i >= b.N {
				return
			}
		}
	}
}

func BenchmarkGetCloneEachTime(b *testing.B) {
	b.StopTimer()
	insertP := perm(benchmarkTreeSize)
	removeP := perm(benchmarkTreeSize)
	b.StartTimer()
	i := 0
	for i < b.N {
		b.StopTimer()
		tr := btree.New(*btreeDegree)
		for _, v := range insertP {
			tr.ReplaceOrInsert(v)
		}
		b.StartTimer()
		for _, item := range removeP {
			tr = tr.Clone()
			tr.Get(item)
			i++
			if i >= b.N {
				return
			}
		}
	}
}

type byInts []interface{}

func (a byInts) Len() int {
	return len(a)
}

func (a byInts) Less(i, j int) bool {
	return a[i].(int) < a[j].(int)
}

func (a byInts) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func BenchmarkAscend(b *testing.B) {
	arr := perm(benchmarkTreeSize)
	tr := btree.New(*btreeDegree)
	for _, v := range arr {
		tr.ReplaceOrInsert(v)
	}
	sort.Sort(byInts(arr))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		j := 0
		tr.Ascend(func(item interface{}) bool {
			if item.(int) != arr[j].(int) {
				b.Fatalf("mismatch: expected: %v, got %v", arr[j].(int), item.(int))
			}
			j++
			return true
		})
	}
}

func BenchmarkDescend(b *testing.B) {
	arr := perm(benchmarkTreeSize)
	tr := btree.New(*btreeDegree)
	for _, v := range arr {
		tr.ReplaceOrInsert(v)
	}
	sort.Sort(byInts(arr))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		j := len(arr) - 1
		tr.Descend(func(item interface{}) bool {
			if item.(int) != arr[j].(int) {
				b.Fatalf("mismatch: expected: %v, got %v", arr[j].(int), item.(int))
			}
			j--
			return true
		})
	}
}
func BenchmarkAscendRange(b *testing.B) {
	arr := perm(benchmarkTreeSize)
	tr := btree.New(*btreeDegree)
	for _, v := range arr {
		tr.ReplaceOrInsert(v)
	}
	sort.Sort(byInts(arr))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		j := 100
		tr.AscendRange(100, arr[len(arr)-100], func(item interface{}) bool {
			if item.(int) != arr[j].(int) {
				b.Fatalf("mismatch: expected: %v, got %v", arr[j].(int), item.(int))
			}
			j++
			return true
		})
		if j != len(arr)-100 {
			b.Fatalf("expected: %v, got %v", len(arr)-100, j)
		}
	}
}

func BenchmarkDescendRange(b *testing.B) {
	arr := perm(benchmarkTreeSize)
	tr := btree.New(*btreeDegree)
	for _, v := range arr {
		tr.ReplaceOrInsert(v)
	}
	sort.Sort(byInts(arr))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		j := len(arr) - 100
		tr.DescendRange(arr[len(arr)-100], 100, func(item interface{}) bool {
			if item.(int) != arr[j].(int) {
				b.Fatalf("mismatch: expected: %v, got %v", arr[j].(int), item.(int))
			}
			j--
			return true
		})
		if j != 100 {
			b.Fatalf("expected: %v, got %v", len(arr)-100, j)
		}
	}
}
func BenchmarkAscendGreaterOrEqual(b *testing.B) {
	arr := perm(benchmarkTreeSize)
	tr := btree.New(*btreeDegree)
	for _, v := range arr {
		tr.ReplaceOrInsert(v)
	}
	sort.Sort(byInts(arr))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		j := 100
		k := 0
		tr.AscendGreaterOrEqual(100, func(item interface{}) bool {
			if item.(int) != arr[j].(int) {
				b.Fatalf("mismatch: expected: %v, got %v", arr[j].(int), item.(int))
			}
			j++
			k++
			return true
		})
		if j != len(arr) {
			b.Fatalf("expected: %v, got %v", len(arr), j)
		}
		if k != len(arr)-100 {
			b.Fatalf("expected: %v, got %v", len(arr)-100, k)
		}
	}
}
func BenchmarkDescendLessOrEqual(b *testing.B) {
	arr := perm(benchmarkTreeSize)
	tr := btree.New(*btreeDegree)
	for _, v := range arr {
		tr.ReplaceOrInsert(v)
	}
	sort.Sort(byInts(arr))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		j := len(arr) - 100
		k := len(arr)
		tr.DescendLessOrEqual(arr[len(arr)-100], func(item interface{}) bool {
			if item.(int) != arr[j].(int) {
				b.Fatalf("mismatch: expected: %v, got %v", arr[j].(int), item.(int))
			}
			j--
			k--
			return true
		})
		if j != -1 {
			b.Fatalf("expected: %v, got %v", -1, j)
		}
		if k != 99 {
			b.Fatalf("expected: %v, got %v", 99, k)
		}
	}
}

const cloneTestSize = 10000

func cloneTest(t *testing.T, b btree.Interface, start int, p []interface{}, wg *sync.WaitGroup, trees *[]btree.Interface, lock *sync.Mutex) {
	t.Logf("Starting btree.New clone at %v", start)
	lock.Lock()
	*trees = append(*trees, b)
	lock.Unlock()
	for i := start; i < cloneTestSize; i++ {
		b.ReplaceOrInsert(p[i])
		if i%(cloneTestSize/5) == 0 {
			wg.Add(1)
			go cloneTest(t, b.Clone(), i+1, p, wg, trees, lock)
		}
	}
	wg.Done()
}

func TestCloneConcurrentOperations(t *testing.T) {
	b := btree.New(*btreeDegree)
	trees := []btree.Interface{}
	p := perm(cloneTestSize)
	var wg sync.WaitGroup
	wg.Add(1)
	go cloneTest(t, b, 0, p, &wg, &trees, &sync.Mutex{})
	wg.Wait()
	want := rang(cloneTestSize)
	t.Logf("Starting equality checks on %d trees", len(trees))
	for i, tree := range trees {
		if !reflect.DeepEqual(want, all(tree)) {
			t.Errorf("tree %v mismatch", i)
		}
	}
	t.Log("Removing half from first half")
	toRemove := rang(cloneTestSize)[cloneTestSize/2:]
	for i := 0; i < len(trees)/2; i++ {
		tree := trees[i]
		wg.Add(1)
		go func() {
			for _, item := range toRemove {
				tree.Delete(item)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	t.Log("Checking all values again")
	for i, tree := range trees {
		var wantpart []interface{}
		if i < len(trees)/2 {
			wantpart = want[:cloneTestSize/2]
		} else {
			wantpart = want
		}
		if got := all(tree); !reflect.DeepEqual(wantpart, got) {
			t.Errorf("tree %v mismatch, want %v got %v", i, len(want), len(got))
		}
	}
}

func BenchmarkDeleteAndRestore(b *testing.B) {
	items := perm(16392)
	b.ResetTimer()
	b.Run(`CopyBigFreeList`, func(b *testing.B) {
		fl := btree.NewFreeList(16392)
		tr := btree.NewWithFreeList(*btreeDegree, fl)
		for _, v := range items {
			tr.ReplaceOrInsert(v)
		}
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			dels := make([]interface{}, 0, tr.Size())
			tr.Ascend(btree.ItemIterator(func(b interface{}) bool {
				dels = append(dels, b)
				return true
			}))
			for _, del := range dels {
				tr.Delete(del)
			}
			// tr is now empty, we make a btree.New empty copy of it.
			tr = btree.NewWithFreeList(*btreeDegree, fl)
			for _, v := range items {
				tr.ReplaceOrInsert(v)
			}
		}
	})
	b.Run(`Copy`, func(b *testing.B) {
		tr := btree.New(*btreeDegree)
		for _, v := range items {
			tr.ReplaceOrInsert(v)
		}
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			dels := make([]interface{}, 0, tr.Size())
			tr.Ascend(btree.ItemIterator(func(b interface{}) bool {
				dels = append(dels, b)
				return true
			}))
			for _, del := range dels {
				tr.Delete(del)
			}
			// tr is now empty, we make a btree.New empty copy of it.
			tr = btree.New(*btreeDegree)
			for _, v := range items {
				tr.ReplaceOrInsert(v)
			}
		}
	})
	b.Run(`ClearBigFreelist`, func(b *testing.B) {
		fl := btree.NewFreeList(16392)
		tr := btree.NewWithFreeList(*btreeDegree, fl)
		for _, v := range items {
			tr.ReplaceOrInsert(v)
		}
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			tr.Clear()
			for _, v := range items {
				tr.ReplaceOrInsert(v)
			}
		}
	})
	b.Run(`Clear`, func(b *testing.B) {
		tr := btree.New(*btreeDegree)
		for _, v := range items {
			tr.ReplaceOrInsert(v)
		}
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			tr.Clear()
			for _, v := range items {
				tr.ReplaceOrInsert(v)
			}
		}
	})
}
