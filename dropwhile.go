package iterium

import "iter"

// DropWhile skips values until predicate returns false, then yields the rest.
func DropWhile[T any](seq iter.Seq[T], predicate func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		dropping := true
		for value := range seq {
			if dropping {
				if predicate(value) {
					continue
				}
				dropping = false
			}

			if !yield(value) {
				return
			}
		}
	}
}
