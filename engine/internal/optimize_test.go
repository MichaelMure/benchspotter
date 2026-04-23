// Package tuning provides example benchmarks for the benchspotter optimizer.
//
// Run the optimizer against these benchmarks with:
//
//	benchspotter optimize --bench BenchmarkHybridSort --strategy sa
//	benchspotter optimize --bench BenchmarkMapOps     --strategy sa
package internal_test

import (
	"math/rand"
	"testing"

	"benchspotter/benchinput"
)

// ── BenchmarkHybridSort ───────────────────────────────────────────────────────
//
// Tunes the insertion-sort cutoff of a hybrid quicksort.
//
// Below the threshold, the algorithm switches from quicksort to insertion sort.
// Insertion sort has lower constant factors for tiny slices (no recursion,
// better branch prediction), but its O(n²) cost dominates for larger ones.
//
// The optimal threshold is typically 8–24 elements depending on CPU/cache.
// Too small → pure quicksort with excessive recursion overhead.
// Too large → insertion sort on large partitions (quadratic blowup).
//
// This is a classic single-parameter float optimisation with a clear interior
// minimum, ideal for demonstrating the SA strategy.

var SortThreshold = benchinput.Float("threshold", 16, 1, 128)

const sortN = 10_000

var sortSeed = rand.New(rand.NewSource(42))

func makeData(n int) []int {
	data := make([]int, n)
	r := rand.New(rand.NewSource(42))
	for i := range data {
		data[i] = r.Int()
	}
	return data
}

var sortData = makeData(sortN)

func BenchmarkHybridSort(b *testing.B) {
	threshold := int(SortThreshold)
	if threshold < 1 {
		threshold = 1
	}

	buf := make([]int, sortN)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copy(buf, sortData)
		hybridSort(buf, threshold)
	}
}

func insertionSort(a []int) {
	for i := 1; i < len(a); i++ {
		key := a[i]
		j := i - 1
		for j >= 0 && a[j] > key {
			a[j+1] = a[j]
			j--
		}
		a[j+1] = key
	}
}

func hybridSort(a []int, threshold int) {
	if len(a) <= threshold {
		insertionSort(a)
		return
	}
	// Median-of-three pivot selection.
	mid := len(a) / 2
	if a[0] > a[mid] {
		a[0], a[mid] = a[mid], a[0]
	}
	if a[mid] > a[len(a)-1] {
		a[mid], a[len(a)-1] = a[len(a)-1], a[mid]
	}
	if a[0] > a[mid] {
		a[0], a[mid] = a[mid], a[0]
	}
	pivot := a[mid]
	lo, hi := 0, len(a)-1
	for lo <= hi {
		for a[lo] < pivot {
			lo++
		}
		for a[hi] > pivot {
			hi--
		}
		if lo <= hi {
			a[lo], a[hi] = a[hi], a[lo]
			lo++
			hi--
		}
	}
	hybridSort(a[:hi+1], threshold)
	hybridSort(a[lo:], threshold)
}

// ── BenchmarkMapOps ───────────────────────────────────────────────────────────
//
// Measures the cost of allocating and populating a map.
// Two float parameters produce a 2-D optimisation landscape:
//
//   - capacityFactor: initial capacity hint as a multiple of the insert count.
//     Too small → rehashing overhead. Too large → wasted allocation and GC.
//     Sweet spot ≈ 1.0–1.5 (confirmed empirically on this benchmark).
//
//   - fillFraction: fraction of baseN elements to insert.
//     Scales work roughly linearly → strong positive correlation with ns/op.
//     Lets the correlation matrix clearly distinguish the two parameters.

var (
	CapacityFactor = benchinput.Float("capacityFactor", 1.5, 0.25, 6.0)
	FillFraction   = benchinput.Float("fillFraction", 0.75, 0.1, 1.0)
)

const baseN = 10_000

var mapSink int

func BenchmarkMapOps(b *testing.B) {
	cf := CapacityFactor
	ff := FillFraction

	insertN := int(float64(baseN) * ff)
	if insertN < 1 {
		insertN = 1
	}
	hint := int(float64(insertN) * cf)
	if hint < 1 {
		hint = 1
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m := make(map[int]struct{}, hint)
		for j := 0; j < insertN; j++ {
			m[j] = struct{}{}
		}
		mapSink = len(m)
	}
}
