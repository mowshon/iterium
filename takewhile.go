package iterium

import "iter"

// TakeWhile yields values until predicate returns false.
func TakeWhile[T any](seq iter.Seq[T], predicate func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for value := range seq {
			if !predicate(value) {
				return
			}
			if !yield(value) {
				return
			}
		}
	}
}
