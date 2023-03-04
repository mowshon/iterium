package iternium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCycle(t *testing.T) {
	cycle := Cycle(New(1, 2, 3))
	assert.Exactly(t, true, cycle.IsInfinite())

	var dump []int

	for {
		value, err := cycle.Next()
		if len(dump) >= 9 {
			assert.ErrorContains(t, err, "stop iteration")
			break
		}

		dump = append(dump, value)

		if len(dump) == 9 {
			cycle.Close()
		}
	}

	expected := []int{1, 2, 3, 1, 2, 3, 1, 2, 3}
	assert.Exactly(t, expected, dump)
}

func TestCycleWithInfiniteIter(t *testing.T) {
	countNumbers := Count(50, 1)
	cycle := Cycle(countNumbers)

	var dump []int
	for {
		value, err := cycle.Next()
		if len(dump) >= 5 {
			assert.ErrorContains(t, err, "stop iteration")
			break
		}

		dump = append(dump, value)

		if len(dump) == 5 {
			cycle.Close()
		}
	}

	expected := []int{50, 51, 52, 53, 54}
	assert.Exactly(t, expected, dump)
}

func TestCycleEmptySlice(t *testing.T) {
	repeat := Repeat(1, 5)
	repeat.Close()

	cycle := Cycle(repeat)
	_, err := cycle.Next()

	assert.Exactly(t, err, stopIterationErr)
}
