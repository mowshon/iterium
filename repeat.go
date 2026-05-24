package iterium

import "iter"

// Repeat returns a Go iterator sequence that repeats value n times.
// A negative n repeats forever.
func Repeat[T any](value T, n int) iter.Seq[T] {
	return func(yield func(T) bool) {
		if n < 0 {
			for {
				if !yield(value) {
					return
				}
			}
		}

		for i := 0; i < n; i++ {
			if !yield(value) {
				return
			}
		}
	}
}
