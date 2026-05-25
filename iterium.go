package iterium

import (
	"golang.org/x/exp/constraints"
)

// Number is the type constraint which includes all numbers.
type Number interface {
	constraints.Integer | constraints.Float
}

// Signed is a type restriction on all numbers especially
// including negative numbers and floating point numbers.
type Signed interface {
	constraints.Signed | constraints.Float
}

// argsTrio takes the first three values from the slice and returns them as arguments.
func argsTrio[T any](args []T, first, second, third T) (T, T, T) {
	switch len(args) {
	case 0:
		return first, second, third
	case 1:
		return args[0], second, third
	case 2:
		return args[0], args[1], third
	default:
		return args[0], args[1], args[2]
	}
}
