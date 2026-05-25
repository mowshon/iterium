package iterium

// Pair stores two values produced by two-value iterator adapters.
type Pair[A, B any] struct {
	First  A
	Second B
}
