package iterium

import (
	"iter"
	"math"
)

// CombinationsWithReplacementCount calculates the total number of combinations with replacement
// for a given set of n elements and a combination length of k.
func CombinationsWithReplacementCount(n, k int) int64 {
	count, ok := CombinationsWithReplacementCountOK(n, k)
	if !ok {
		return math.MaxInt64
	}
	return count
}

// CombinationsWithReplacementCountOK returns replacement combination count and reports overflow.
func CombinationsWithReplacementCountOK(n, k int) (int64, bool) {
	if n < 0 || k < 0 {
		return 0, false
	}
	if k == 0 {
		return 1, true
	}
	if n == 0 && k > 0 {
		return 0, true
	}
	maxInt := int(^uint(0) >> 1)
	if n > maxInt-k+1 {
		return math.MaxInt64, false
	}
	return CombinationsCountOK(n+k-1, k)
}

// CombinationsWithReplacement returns r-length combinations with replacement.
// Each yielded slice is safe to keep.
func CombinationsWithReplacement[T any](symbols []T, r int) iter.Seq[[]T] {
	return func(yield func([]T) bool) {
		CombinationsWithReplacementInto(symbols, r, func(value []T) bool {
			out := make([]T, r)
			copy(out, value)
			return yield(out)
		})
	}
}

// CombinationsWithReplacementInto generates replacement combinations using a reused result buffer.
// The yielded slice is only valid until the next yield call.
func CombinationsWithReplacementInto[T any](symbols []T, r int, yield func([]T) bool) {
	n := len(symbols)
	if r < 0 || (n == 0 && r > 0) {
		return
	}
	if r == 0 {
		yield([]T{})
		return
	}

	indices := make([]int, r)
	result := make([]T, r)

	for {
		for i, index := range indices {
			result[i] = symbols[index]
		}
		if !yield(result) {
			return
		}

		for i := r - 1; ; i-- {
			if i < 0 {
				return
			}
			if indices[i] != n-1 {
				indices[i]++
				for j := i + 1; j < r; j++ {
					indices[j] = indices[i]
				}
				break
			}
		}
	}
}
