package iterium

import "iter"

// Cycle yields values from seq, lazily caching the first pass, then replays the cache forever.
func Cycle[T any](seq iter.Seq[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		saved := make([]T, 0)

		for value := range seq {
			saved = append(saved, value)
			if !yield(value) {
				return
			}
		}

		if len(saved) == 0 {
			return
		}

		for {
			for _, value := range saved {
				if !yield(value) {
					return
				}
			}
		}
	}
}
