package iterium

import "iter"

// StarMap applies a binary function to two-item slices from a sequence.
// It panics if any yielded slice has fewer than two elements.
func StarMap[T any](seq iter.Seq[[]T], apply func(T, T) T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for value := range seq {
			if !yield(apply(value[0], value[1])) {
				return
			}
		}
	}
}
