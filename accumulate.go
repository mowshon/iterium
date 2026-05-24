package iterium

import "iter"

// Accumulate yields accumulated results from applying operator.
func Accumulate[T any](seq iter.Seq[T], operator func(T, T) T) iter.Seq[T] {
	return func(yield func(T) bool) {
		var last T
		started := false

		for value := range seq {
			if !started {
				last = value
				started = true
			} else {
				last = operator(last, value)
			}

			if !yield(last) {
				return
			}
		}
	}
}
