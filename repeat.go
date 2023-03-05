package iterium

// Repeat returns a channel from which a value can be retrieved n-number of times.
func Repeat[T any](value T, n int) Iter[T] {
	// Initialisation of a new channel.
	iter := Instance[T](int64(n), false)

	go func() {
		defer iterRecover()
		defer iter.Close()

		for step := 0; step < n; step++ {
			iter.Chan() <- value
		}
	}()

	return iter
}
