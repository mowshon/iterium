package iterium

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRange(t *testing.T) {
	assert.Exactly(t, []int{0, 1, 2, 3, 4}, Slice(Range(5)))
}

func TestRangeEmpty(t *testing.T) {
	assert.Exactly(t, []int{}, Slice(Range[int]()))
	assert.Exactly(t, []int{}, Slice(Range(50, 1)))
}

func TestRangeDefaultStep(t *testing.T) {
	assert.Exactly(t, []int{0, 1, 2, 3, 4}, Slice(Range(0, 5)))
}

func TestRangeSigned(t *testing.T) {
	assert.Exactly(t, []int{0, -1, -2, -3, -4}, Slice(Range(-5)))
}

func TestRangeWithStep(t *testing.T) {
	assert.Exactly(t, []int{0, 2, 4, 6, 8}, Slice(Range(0, 10, 2)))
}

func TestRangeZeroStep(t *testing.T) {
	assert.Exactly(t, []int{}, Slice(Range(0, 10, 0)))
}

func TestRangeFloat(t *testing.T) {
	assert.Exactly(t, []float64{0, 1.5, 3, 4.5, 6, 7.5, 9}, Slice(Range(0.0, 10.0, 1.5)))
}
