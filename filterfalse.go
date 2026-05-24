package iterium

import "iter"

// FilterFalse lazily filters values where predicate returns false.
func FilterFalse[T any](seq iter.Seq[T], predicate func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for value := range seq {
			if !predicate(value) && !yield(value) {
				return
			}
		}
	}
}
