package iterium

import (
	"math"
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

func TestRangeCountIntegerPrecision(t *testing.T) {
	const start = int64(1 << 53)

	assert.Exactly(t, int64(3), RangeCount(start, start+3, int64(1)))
	assert.Exactly(t, int64(3), RangeCount(start+3, start, int64(-1)))
}

func TestRangeCountIntegerOverflowSaturates(t *testing.T) {
	assert.Exactly(t, int64(math.MaxInt64), RangeCount(int64(math.MinInt64), int64(math.MaxInt64), int64(1)))
	assert.Exactly(t, int64(math.MaxInt64), RangeCount(int64(math.MaxInt64), int64(math.MinInt64), int64(-1)))
}

func TestRangeCountUnsigned(t *testing.T) {
	assert.Exactly(t, int64(4), RangeCount(uint64(0), uint64(10), uint64(3)))
}

func TestRangeCountDefinedNumericTypes(t *testing.T) {
	type localInt int64
	type localFloat float64
	type localUint uint64

	assert.Exactly(t, int64(3), RangeCount(localInt(1<<53), localInt((1<<53)+3), localInt(1)))
	assert.Exactly(t, int64(4), RangeCount(localFloat(0), localFloat(1), localFloat(0.25)))
	assert.Exactly(t, int64(4), RangeCount(localUint(0), localUint(10), localUint(3)))
}
