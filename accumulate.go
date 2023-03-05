package iterium

// Accumulate returns an iterator that sends the accumulated
// result from the binary function to the channel.
func Accumulate[N Number](iterable Iter[N], operator func(N, N) N) Iter[N] {
	iter := Instance[N](iterable.Count(), iterable.IsInfinite())

	go func() {
		defer iterRecover()
		defer iter.Close()

		var last N
		var start bool
		for true {
			next, err := iterable.Next()
			if err != nil {
				return
			}

			if !start {
				iter.Chan() <- next
				last = next
				start = true
				continue
			}

			result := operator(last, next)
			iter.Chan() <- result
			last = result
		}
	}()

	return iter
}
