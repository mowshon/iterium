package iterium

import "iter"

// ProductSeq returns a reusable Go iterator sequence for Cartesian products.
// Each yielded slice is safe to keep.
func ProductSeq[T any](symbols []T, repeat int) iter.Seq[[]T] {
	return func(yield func([]T) bool) {
		ProductInto(symbols, repeat, func(value []T) bool {
			out := make([]T, len(value))
			copy(out, value)
			return yield(out)
		})
	}
}

// ProductInto generates Cartesian products using a reused result buffer.
// The yielded slice is only valid until the next yield call.
func ProductInto[T any](symbols []T, repeat int, yield func([]T) bool) {
	if repeat < 0 || len(symbols) == 0 {
		return
	}
	if repeat == 0 {
		yield([]T{})
		return
	}

	indices := make([]int, repeat)
	result := make([]T, repeat)

	for {
		for i, index := range indices {
			result[i] = symbols[index]
		}
		if !yield(result) {
			return
		}

		for i := repeat - 1; i >= 0; i-- {
			indices[i]++
			if indices[i] < len(symbols) {
				break
			}
			if i == 0 {
				return
			}
			indices[i] = 0
		}
	}
}

// ProductBytesInto is a specialized zero-allocation product generator for byte alphabets.
// The yielded slice is reused and must be copied by callers that keep it.
func ProductBytesInto(symbols []byte, repeat int, yield func([]byte) bool) {
	if repeat < 0 || len(symbols) == 0 {
		return
	}
	if repeat == 0 {
		yield([]byte{})
		return
	}

	indices := make([]int, repeat)
	result := make([]byte, repeat)

	for {
		for i, index := range indices {
			result[i] = symbols[index]
		}
		if !yield(result) {
			return
		}

		for i := repeat - 1; i >= 0; i-- {
			indices[i]++
			if indices[i] < len(symbols) {
				break
			}
			if i == 0 {
				return
			}
			indices[i] = 0
		}
	}
}

// Product2Seq returns a reusable Go iterator sequence for a two-slice Cartesian product.
func Product2Seq[A, B any](first []A, second []B) iter.Seq2[A, B] {
	return func(yield func(A, B) bool) {
		for _, a := range first {
			for _, b := range second {
				if !yield(a, b) {
					return
				}
			}
		}
	}
}

// CombinationsSeq returns r-length combinations in lexicographic index order.
// Each yielded slice is safe to keep.
func CombinationsSeq[T any](symbols []T, r int) iter.Seq[[]T] {
	return func(yield func([]T) bool) {
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
			out := make([]T, r)
			copy(out, result)
			if !yield(out) {
				return
			}

			i := r - 1
			for ; i >= 0 && indices[i] == i+n-r; i-- {
			}
			if i < 0 {
				return
			}
			indices[i]++
			for j := i + 1; j < r; j++ {
				indices[j] = indices[j-1] + 1
			}
		}
	}
}

// CombinationsWithReplacementSeq returns r-length combinations with replacement.
// Each yielded slice is safe to keep.
func CombinationsWithReplacementSeq[T any](symbols []T, r int) iter.Seq[[]T] {
	return func(yield func([]T) bool) {
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
			out := make([]T, r)
			copy(out, result)
			if !yield(out) {
				return
			}

			i := r - 1
			for ; i >= 0 && indices[i] == n-1; i-- {
			}
			if i < 0 {
				return
			}
			indices[i]++
			for j := i + 1; j < r; j++ {
				indices[j] = indices[i]
			}
		}
	}
}

// PermutationsSeq returns r-length permutations in CPython itertools order.
// Each yielded slice is safe to keep.
func PermutationsSeq[T any](symbols []T, r int) iter.Seq[[]T] {
	return func(yield func([]T) bool) {
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
		out := make([]T, r)
		copy(out, result)
		if !yield(out) {
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
				out := make([]T, r)
				copy(out, result)
				if !yield(out) {
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
}
