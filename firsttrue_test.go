package iternium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFirstTrue(t *testing.T) {
	numbers := Range(10)

	firstTrue := FirstTrue(numbers, func(x int) bool {
		return x == 5
	})

	if value, err := firstTrue.Next(); assert.Nil(t, err) {
		assert.Exactly(t, 5, value)
	}

	if _, err := firstTrue.Next(); assert.NotNil(t, err) {
		assert.ErrorIs(t, stopIterationErr, err)
	}
}

func TestFirstTrueEmpty(t *testing.T) {
	firstTrue := FirstTrue(Range(10), func(x int) bool {
		return x > 100
	})

	if slice, err := firstTrue.Slice(); assert.Nil(t, err) {
		assert.Exactly(t, []int{}, slice)
	}
}
