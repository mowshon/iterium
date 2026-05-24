package iterium

import (
	"iter"
	"math"
)

// CombinationsCount returns n choose k, saturating on overflow.
func CombinationsCount(n, k int) int64 {
	count, ok := CombinationsCountOK(n, k)
	if !ok {
		return math.MaxInt64
	}
	return count
}

// CombinationsCountOK returns n choose k and reports overflow.
func CombinationsCountOK(n, k int) (int64, bool) {
	if n < 0 || k < 0 {
		return 0, false
	}
	if k > n {
		return 0, true
	}
	if k > n-k {
		k = n - k
	}

	result := int64(1)
	for i := 1; i <= k; i++ {
		numerator := int64(n - k + i)
		denominator := int64(i)

		gcd := gcd64(numerator, denominator)
		numerator /= gcd
		denominator /= gcd

		gcd = gcd64(result, denominator)
		result /= gcd
		denominator /= gcd

		if denominator != 1 {
			return math.MaxInt64, false
		}
		if numerator != 0 && result > math.MaxInt64/numerator {
			return math.MaxInt64, false
		}
		result *= numerator
	}
	return result, true
}

func gcd64(a, b int64) int64 {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Combinations returns r-length combinations in lexicographic index order.
// Each yielded slice is safe to keep.
func Combinations[T any](symbols []T, r int) iter.Seq[[]T] {
	return func(yield func([]T) bool) {
		CombinationsInto(symbols, r, func(value []T) bool {
			out := make([]T, r)
			copy(out, value)
			return yield(out)
		})
	}
}

// CombinationsInto generates combinations using a reused result buffer.
// The yielded slice is only valid until the next yield call.
func CombinationsInto[T any](symbols []T, r int, yield func([]T) bool) {
	n := len(symbols)
	if r < 0 || r > n {
		return
	}
	if r == 0 {
		yield([]T{})
		return
	}

	indices := make([]int, r)
	result := make([]T, r)
	for i := range indices {
		indices[i] = i
	}

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
			if indices[i] != i+n-r {
				indices[i]++
				for j := i + 1; j < r; j++ {
					indices[j] = indices[j-1] + 1
				}
				break
			}
		}
	}
}
