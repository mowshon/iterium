package iterium

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmpty(t *testing.T) {
	assert.Exactly(t, []int{}, Slice(Empty[int]()))
}

func TestSliceN(t *testing.T) {
	assert.Exactly(t, []int{0, 1, 2, 3}, SliceN(Count[int](), 4))
	assert.Exactly(t, []int{}, SliceN(Count[int](), 0))
}

func TestSliceUntil(t *testing.T) {
	values := SliceUntil(Count[int](), func(value int) bool {
		return value == 4
	})

	assert.Exactly(t, []int{0, 1, 2, 3}, values)
}

func TestChan(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var values []int
	for value := range Chan(ctx, New(1, 2, 3)) {
		values = append(values, value)
	}

	assert.Exactly(t, []int{1, 2, 3}, values)
}
