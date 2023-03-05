package iterium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCount(t *testing.T) {
	count := Count(0, -1)

	var storage []int
	for i := 0; i <= 10; i++ {
		next, err := count.Next()
		assert.Nil(t, err)
		storage = append(storage, next)
	}

	assert.Exactly(t, []int{0, -1, -2, -3, -4, -5, -6, -7, -8, -9, -10}, storage)
}

func TestCountInfinite(t *testing.T) {
	count := Count(0, -2)
	slice, err := count.Slice()
	assert.ErrorIs(t, err, infiniteIteratorErr)
	assert.Nil(t, slice)
}

func TestCountOneArg(t *testing.T) {
	count := Count(1)

	var storage []int
	for i := 0; i <= 3; i++ {
		next, err := count.Next()
		assert.Nil(t, err)
		storage = append(storage, next)
	}

	assert.Exactly(t, []int{1, 2, 3, 4}, storage)
}

func TestCountNoArgs(t *testing.T) {
	count := Count[int]()

	var storage []int
	for i := 0; i <= 3; i++ {
		next, err := count.Next()
		assert.Nil(t, err)
		storage = append(storage, next)
	}

	assert.Exactly(t, []int{0, 1, 2, 3}, storage)
}
