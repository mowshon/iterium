package iternium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRange(t *testing.T) {
	values := Range(5)
	if slice, err := values.Slice(); assert.Nil(t, err) {
		assert.Exactly(t, []int{0, 1, 2, 3, 4}, slice)
	}
}

func TestRangeEmpty(t *testing.T) {
	values := Range[int]()
	if slice, err := values.Slice(); assert.Nil(t, err) {
		assert.Exactly(t, []int{}, slice)
	}

	wrongArgs := Range(50, 1)
	if slice, err := wrongArgs.Slice(); assert.Nil(t, err) {
		assert.Exactly(t, []int{}, slice)
	}
}

func TestRangeDefaultStep(t *testing.T) {
	values := Range(0, 5)
	if slice, err := values.Slice(); assert.Nil(t, err) {
		assert.Exactly(t, []int{0, 1, 2, 3, 4}, slice)
	}
}

func TestRangeSigned(t *testing.T) {
	values := Range(-5)
	if slice, err := values.Slice(); assert.Nil(t, err) {
		assert.Exactly(t, []int{0, -1, -2, -3, -4}, slice)
	}
}

func TestRangeWithStep(t *testing.T) {
	values := Range(0, 10, 2)
	if slice, err := values.Slice(); assert.Nil(t, err) {
		assert.Exactly(t, []int{0, 2, 4, 6, 8}, slice)
	}
}

func TestRangeFloat(t *testing.T) {
	values := Range(0.0, 10.0, 1.5)
	if slice, err := values.Slice(); assert.Nil(t, err) {
		assert.Exactly(t, []float64{0, 1.5, 3, 4.5, 6, 7.5, 9}, slice)
	}
}
