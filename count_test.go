package iterium

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCount(t *testing.T) {
	assert.Exactly(t, []int{0, -1, -2, -3}, SliceN(Count(0, -1), 4))
}

func TestCountOneArg(t *testing.T) {
	assert.Exactly(t, []int{1, 2, 3, 4}, SliceN(Count(1), 4))
}

func TestCountNoArgs(t *testing.T) {
	assert.Exactly(t, []int{0, 1, 2, 3}, SliceN(Count[int](), 4))
}
