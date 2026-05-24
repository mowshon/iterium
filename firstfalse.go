package iterium

import "iter"

// FirstFalse returns the first value where predicate returns false.
func FirstFalse[T any](seq iter.Seq[T], predicate func(T) bool) (T, bool) {
	var zero T
	for value := range seq {
		if !predicate(value) {
			return value, true
		}
	}
	return zero, false
}
