package iterium

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func pow(a, b float64) float64 {
	return math.Pow(a, b)
}

func merge(a, b string) string {
	return fmt.Sprintf("%s-%s", a, b)
}

func TestStarmapInteger(t *testing.T) {
	values := StarMap(New([]float64{2, 5}, []float64{3, 2}, []float64{10, 3}), pow)

	assert.Exactly(t, []float64{32, 9, 1000}, Slice(values))
}

func TestStarmapString(t *testing.T) {
	values := StarMap(New([]string{"a", "b"}, []string{"c", "d"}), merge)

	assert.Exactly(t, []string{"a-b", "c-d"}, Slice(values))
}
