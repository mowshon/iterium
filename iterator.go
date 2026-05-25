package iterium

import (
	"context"
	"iter"
)

// New returns a reusable Go iterator sequence over values.
func New[T any](values ...T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, value := range values {
			if !yield(value) {
				return
			}
		}
	}
}

// Empty returns a reusable empty Go iterator sequence.
func Empty[T any]() iter.Seq[T] {
	return func(func(T) bool) {}
}

// Chan adapts an official Go iterator sequence to a channel.
// Prefer ranging over the sequence directly in hot paths.
func Chan[T any](ctx context.Context, seq iter.Seq[T]) <-chan T {
	ch := make(chan T)
	go func() {
		defer close(ch)
		for value := range seq {
			select {
			case <-ctx.Done():
				return
			case ch <- value:
			}
		}
	}()
	return ch
}

// Chan2 adapts a two-value Go iterator sequence to a channel of pairs.
// Prefer ranging over the sequence directly in hot paths.
func Chan2[A, B any](ctx context.Context, seq iter.Seq2[A, B]) <-chan Pair[A, B] {
	ch := make(chan Pair[A, B])
	go func() {
		defer close(ch)
		for first, second := range seq {
			select {
			case <-ctx.Done():
				return
			case ch <- Pair[A, B]{First: first, Second: second}:
			}
		}
	}()
	return ch
}

// Slice collects a finite sequence into a slice.
func Slice[T any](seq iter.Seq[T]) []T {
	result := make([]T, 0)
	for value := range seq {
		result = append(result, value)
	}
	return result
}

// SliceN collects up to n values from a sequence.
func SliceN[T any](seq iter.Seq[T], n int) []T {
	if n <= 0 {
		return []T{}
	}

	result := make([]T, 0, n)
	for value := range seq {
		result = append(result, value)
		if len(result) == n {
			break
		}
	}
	return result
}

// SliceUntil collects values until stop returns true for a value.
func SliceUntil[T any](seq iter.Seq[T], stop func(T) bool) []T {
	result := make([]T, 0)
	for value := range seq {
		if stop(value) {
			break
		}
		result = append(result, value)
	}
	return result
}

// Slice2 collects a finite two-value sequence into a slice of pairs.
func Slice2[A, B any](seq iter.Seq2[A, B]) []Pair[A, B] {
	result := make([]Pair[A, B], 0)
	for first, second := range seq {
		result = append(result, Pair[A, B]{First: first, Second: second})
	}
	return result
}
