package iterium

import "iter"

// Count returns an infinite Go iterator sequence.
func Count[N Number](args ...N) iter.Seq[N] {
	start, step, _ := argsTrio[N](args, 0, 1, 0)

	return func(yield func(N) bool) {
		next := start
		for {
			if !yield(next) {
				return
			}
			next += step
		}
	}
}
