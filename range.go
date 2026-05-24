package iterium

import (
	"iter"
	"math"
	"reflect"
)

func RangeCount[N Number](start, stop, step N) int64 {
	if step == 0 {
		return 0
	}

	switch reflect.TypeOf(start).Kind() {
	case reflect.Float32, reflect.Float64:
		a, b, c := float64(start), float64(stop), float64(step)
		count := math.Ceil((b - a) / c)
		if count < 0 {
			return 0
		}
		if count > math.MaxInt64 {
			return math.MaxInt64
		}
		return int64(count)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return rangeCountUint(uint64(start), uint64(stop), uint64(step))
	default:
		return rangeCountInt(int64(start), int64(stop), int64(step))
	}
}

// Range returns a reusable Go iterator sequence of numbers.
func Range[S Signed](args ...S) iter.Seq[S] {
	var start, stop, step S
	var total int64

	switch len(args) {
	case 0:
		return Empty[S]()
	case 1:
		stop = args[0]
		if args[0] < 0 {
			step = -1
		} else {
			step = 1
		}
	default:
		start = args[0]
		stop = args[1]
		if len(args) >= 3 {
			step = args[2]
		} else {
			step = 1
		}
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

func rangeCountUint(start, stop, step uint64) int64 {
	if stop <= start {
		return 0
	}
	return ceilDivUint64(stop-start, step)
}

func rangeCountInt(start, stop, step int64) int64 {
	if step > 0 {
		if stop <= start {
			return 0
		}
		return ceilDivUint64(uint64(stop)-uint64(start), uint64(step))
	}
	if stop >= start {
		return 0
	}
	return ceilDivUint64(uint64(start)-uint64(stop), int64Magnitude(step))
}

func int64Magnitude(value int64) uint64 {
	if value >= 0 {
		return uint64(value)
	}
	return uint64(-(value + 1)) + 1
}

func ceilDivUint64(value, divisor uint64) int64 {
	quotient := value / divisor
	if value%divisor != 0 {
		quotient++
	}
	if quotient > uint64(math.MaxInt64) {
		return math.MaxInt64
	}
	return int64(quotient)
}
