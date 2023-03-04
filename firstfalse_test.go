package iternium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFirstFalse(t *testing.T) {
	numbers := Range(10, 0, -1)

	firstFalse := FirstFalse(numbers, func(x int) bool {
		return x >= 2
	})

	if value, err := firstFalse.Next(); assert.Nil(t, err) {
		assert.Exactly(t, 1, value)
	}

	if _, err := firstFalse.Next(); assert.NotNil(t, err) {
		assert.ErrorIs(t, stopIterationErr, err)
	}
}

func TestFirstFalseEmpty(t *testing.T) {
	firstFalse := FirstFalse(Range(10), func(x int) bool {
		return x < 100
	})

	if slice, err := firstFalse.Slice(); assert.Nil(t, err) {
		assert.Exactly(t, []int{}, slice)
	}
}
