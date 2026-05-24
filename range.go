package iterium

import (
	"iter"
	"math"
)

func RangeCount[N Number](start, stop, step N) int64 {
	if step == 0 {
		return 0
	}

	a, b, c := float64(start), float64(stop), float64(step)
	count := math.Ceil((b - a) / c)
	if count < 0 {
		return 0
	}
	return int64(count)
}

// Range returns a reusable Go iterator sequence of numbers.
func Range[S Signed](args ...S) iter.Seq[S] {
	var start, stop, step S
	var total int64

	switch len(args) {
	case 0:
		return Empty[S]()
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
		return Empty[S]()
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
