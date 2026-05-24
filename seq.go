package iterium

import (
	"context"
	"iter"
)

// NewSeq returns a reusable Go iterator sequence over values.
func NewSeq[T any](values ...T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, value := range values {
			if !yield(value) {
				return
			}
		}
	}
}

// EmptySeq returns a reusable empty Go iterator sequence.
func EmptySeq[T any]() iter.Seq[T] {
	return func(func(T) bool) {}
}

// ChanSeq adapts an official Go iterator sequence to a channel.
// Prefer ranging over the sequence directly in hot paths.
func ChanSeq[T any](ctx context.Context, seq iter.Seq[T]) <-chan T {
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
		seq(func(first A, second B) bool {
			select {
			case <-ctx.Done():
				return false
			case ch <- Pair[A, B]{First: first, Second: second}:
				return true
			}
		})
	}()
	return ch
}

// SliceSeq collects a finite sequence into a slice.
func SliceSeq[T any](seq iter.Seq[T]) []T {
	result := make([]T, 0)
	for value := range seq {
		result = append(result, value)
	}
	return result
}

// Slice2 collects a finite two-value sequence into a slice of pairs.
func Slice2[A, B any](seq iter.Seq2[A, B]) []Pair[A, B] {
	result := make([]Pair[A, B], 0)
	seq(func(first A, second B) bool {
		result = append(result, Pair[A, B]{First: first, Second: second})
		return true
	})
	return result
}

// RangeSeq returns a reusable Go iterator sequence of numbers.
func RangeSeq[S Signed](args ...S) iter.Seq[S] {
	var start, stop, step S
	var total int64

	switch len(args) {
	case 0:
		return func(func(S) bool) {}
	case 1:
		stop, start, step = argsTrio(args, 0, 0, 1)
		if args[0] < 0 {
			step = -1
		}
	default:
		start, stop, step = argsTrio(args, 0, 0, 1)
	}

	total = RangeCount(start, stop, step)
	if total <= 0 {
		return func(func(S) bool) {}
	}

	return func(yield func(S) bool) {
		next := start
		for i := int64(0); i < total; i++ {
			if !yield(next) {
				return
			}
			next += step
		}
	}
}

// CountSeq returns an infinite Go iterator sequence.
func CountSeq[N Number](args ...N) iter.Seq[N] {
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

// RepeatSeq returns a Go iterator sequence that repeats value n times.
// A negative n repeats forever.
func RepeatSeq[T any](value T, n int) iter.Seq[T] {
	return func(yield func(T) bool) {
		if n < 0 {
			for {
				if !yield(value) {
					return
				}
			}
			// Unreachable; documents that the infinite branch never falls through.
			return
		}

		for i := 0; i < n; i++ {
			if !yield(value) {
				return
			}
		}
	}
}

// MapSeq lazily maps a sequence.
func MapSeq[T, W any](seq iter.Seq[T], apply func(T) W) iter.Seq[W] {
	return func(yield func(W) bool) {
		for value := range seq {
			if !yield(apply(value)) {
				return
			}
		}
	}
}

// StarMapSeq applies a binary function to two-item slices from a sequence.
func StarMapSeq[T any](seq iter.Seq[[]T], apply func(T, T) T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for value := range seq {
			if !yield(apply(value[0], value[1])) {
				return
			}
		}
	}
}

// FilterSeq lazily filters a sequence.
func FilterSeq[T any](seq iter.Seq[T], predicate func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for value := range seq {
			if predicate(value) && !yield(value) {
				return
			}
		}
	}
}

// FilterFalseSeq lazily filters values where predicate returns false.
func FilterFalseSeq[T any](seq iter.Seq[T], predicate func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for value := range seq {
			if !predicate(value) && !yield(value) {
				return
			}
		}
	}
}

// AccumulateSeq yields accumulated results from applying operator.
func AccumulateSeq[T any](seq iter.Seq[T], operator func(T, T) T) iter.Seq[T] {
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

// TakeWhileSeq yields values until predicate returns false.
func TakeWhileSeq[T any](seq iter.Seq[T], predicate func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for value := range seq {
			if !predicate(value) {
				return
			}
			if !yield(value) {
				return
			}
		}
	}
}

// DropWhileSeq skips values until predicate returns false, then yields the rest.
func DropWhileSeq[T any](seq iter.Seq[T], predicate func(T) bool) iter.Seq[T] {
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

// CycleSeq yields values from seq, lazily caching the first pass, then replays the cache forever.
func CycleSeq[T any](seq iter.Seq[T]) iter.Seq[T] {
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

// FirstTrueSeq returns the first value where predicate returns true.
func FirstTrueSeq[T any](seq iter.Seq[T], predicate func(T) bool) (T, bool) {
	var zero T
	for value := range seq {
		if predicate(value) {
			return value, true
		}
	}
	return zero, false
}

// FirstFalseSeq returns the first value where predicate returns false.
func FirstFalseSeq[T any](seq iter.Seq[T], predicate func(T) bool) (T, bool) {
	var zero T
	for value := range seq {
		if !predicate(value) {
			return value, true
		}
	}
	return zero, false
}
