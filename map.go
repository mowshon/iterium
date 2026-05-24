package iterium

import "iter"

// Map lazily maps a sequence.
func Map[T, W any](seq iter.Seq[T], apply func(T) W) iter.Seq[W] {
	return func(yield func(W) bool) {
		for value := range seq {
			if !yield(apply(value)) {
				return
			}
		}
	}
}
