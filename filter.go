package iterium

import "iter"

// Filter lazily filters a sequence.
func Filter[T any](seq iter.Seq[T], predicate func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for value := range seq {
			if predicate(value) && !yield(value) {
				return
			}
		}
	}
}
