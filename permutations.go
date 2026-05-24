package iterium

import (
	"iter"
	"math"
)

// PermutationCount returns the total number of possible permutations
// of k elements from a sequence of n elements.
func PermutationCount(countOfSymbols, limit int) int64 {
	count, ok := PermutationCountOK(countOfSymbols, limit)
	if !ok {
		return math.MaxInt64
	}
	return count
}

// PermutationCountOK returns permutation count and reports overflow.
func PermutationCountOK(countOfSymbols, limit int) (int64, bool) {
	if countOfSymbols < 0 || limit < 0 {
		return 0, false
	}
	if limit > countOfSymbols {
		return 0, true
	}

	result := int64(1)
	for i := countOfSymbols - limit + 1; i <= countOfSymbols; i++ {
		factor := int64(i)
		if factor != 0 && result > math.MaxInt64/factor {
			return math.MaxInt64, false
		}
		result *= factor
	}
	return result, true
}

// Permutations returns r-length permutations in CPython itertools order.
// Each yielded slice is safe to keep.
func Permutations[T any](symbols []T, r int) iter.Seq[[]T] {
	return func(yield func([]T) bool) {
		PermutationsInto(symbols, r, func(value []T) bool {
			out := make([]T, r)
			copy(out, value)
			return yield(out)
		})
	}
}

// PermutationsInto generates permutations using a reused result buffer.
// The yielded slice is only valid until the next yield call.
func PermutationsInto[T any](symbols []T, r int, yield func([]T) bool) {
	n := len(symbols)
	if r < 0 || r > n {
		return
	}

	indices := make([]int, n)
	cycles := make([]int, r)
	result := make([]T, r)
	for i := 0; i < n; i++ {
		indices[i] = i
	}
	for i := 0; i < r; i++ {
		cycles[i] = n - i
	}

	for i := 0; i < r; i++ {
		result[i] = symbols[indices[i]]
	}
	if !yield(result) {
		return
	}

	if n == 0 {
		return
	}

	for {
		advanced := false
		for i := r - 1; i >= 0; i-- {
			cycles[i]--
			if cycles[i] == 0 {
				index := indices[i]
				copy(indices[i:], indices[i+1:])
				indices[n-1] = index
				cycles[i] = n - i
				continue
			}

			j := cycles[i]
			indices[i], indices[n-j] = indices[n-j], indices[i]
			for k := i; k < r; k++ {
				result[k] = symbols[indices[k]]
			}
			if !yield(result) {
				return
			}
			advanced = true
			break
		}
		if !advanced {
			return
		}
	}
}
