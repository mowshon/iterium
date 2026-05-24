package iterium

import (
	"iter"
	"math"
)

// ProductCount calculates the number of Cartesian products with repeat.
func ProductCount(countOfSymbols, repeat int) int64 {
	count, ok := ProductCountOK(countOfSymbols, repeat)
	if !ok {
		return math.MaxInt64
	}
	return count
}

// ProductCountOK calculates Cartesian product count and reports overflow.
func ProductCountOK(countOfSymbols, repeat int) (int64, bool) {
	if countOfSymbols < 0 || repeat < 0 {
		return 0, false
	}
	if repeat == 0 {
		return 1, true
	}

	result := int64(1)
	base := int64(countOfSymbols)
	for i := 0; i < repeat; i++ {
		if base != 0 && result > math.MaxInt64/base {
			return math.MaxInt64, false
		}
		result *= base
	}
	return result, true
}

// Product returns a reusable Go iterator sequence for Cartesian products.
// Each yielded slice is safe to keep.
func Product[T any](symbols []T, repeat int) iter.Seq[[]T] {
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
	if repeat < 0 {
		return
	}
	if repeat == 0 {
		yield([]T{})
		return
	}
	if len(symbols) == 0 {
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
	if repeat < 0 {
		return
	}
	if repeat == 0 {
		yield([]byte{})
		return
	}
	if len(symbols) == 0 {
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

// ProductStringInto is a specialized product generator for string alphabets.
// The yielded byte slice is reused and must be copied by callers that keep it.
func ProductStringInto(symbols string, repeat int, yield func([]byte) bool) {
	ProductBytesInto([]byte(symbols), repeat, yield)
}

// ProductRunesInto is a specialized product generator for rune alphabets.
// The yielded slice is reused and must be copied by callers that keep it.
func ProductRunesInto(symbols []rune, repeat int, yield func([]rune) bool) {
	ProductInto(symbols, repeat, yield)
}

// Product2 returns a reusable Go iterator sequence for a two-slice Cartesian product.
func Product2[A, B any](first []A, second []B) iter.Seq2[A, B] {
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
