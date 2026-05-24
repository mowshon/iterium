package iterium

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCycle(t *testing.T) {
	assert.Exactly(t, []int{1, 2, 3, 1, 2, 3, 1, 2}, SliceN(Cycle(New(1, 2, 3)), 8))
}

func TestCycleEmptySlice(t *testing.T) {
	assert.Exactly(t, []int{}, Slice(Cycle(Empty[int]())))
}
