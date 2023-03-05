package iterium

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func pow(a, b float64) float64 {
	return math.Pow(a, b)
}

func merge(a, b string) string {
	return fmt.Sprintf("%s-%s", a, b)
}

func TestStarmapInteger(t *testing.T) {
	starmap := StarMap(New([]float64{2, 5}, []float64{3, 2}, []float64{10, 3}), pow)
	if result, err := starmap.Slice(); assert.Nil(t, err) {
		assert.Exactly(t, []float64{32, 9, 1000}, result)
	}
}

func TestStarmapString(t *testing.T) {
	starmap := StarMap(New([]string{"a", "b"}, []string{"c", "d"}), merge)
	if result, err := starmap.Slice(); assert.Nil(t, err) {
		assert.Exactly(t, []string{"a-b", "c-d"}, result)
	}
}
