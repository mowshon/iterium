package iterium

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func add(a, b int) int {
	return a + b
}

func mul(a, b int) int {
	return a * b
}

func TestAccumulate(t *testing.T) {
	accumulateAdd := Accumulate(New(1, 2, 3, 4, 5), add)
	if slice, err := accumulateAdd.Slice(); assert.Nil(t, err) {
		assert.Exactly(t, []int{1, 3, 6, 10, 15}, slice)
	}

	accumulateMul := Accumulate(New(1, 2, 3, 4, 5), mul)
	if slice, err := accumulateMul.Slice(); assert.Nil(t, err) {
		assert.Exactly(t, []int{1, 2, 6, 24, 120}, slice)
	}
}

func mergeAcc(first, second string) string {
	return fmt.Sprintf("%s-%s", first, second)
}

func TestAccumulateString(t *testing.T) {
	data := New("A", "B", "C", "D")
	accumulateString := Accumulate(data, mergeAcc)
	if slice, err := accumulateString.Slice(); assert.Nil(t, err) {
		assert.Exactly(t, []string{"A", "A-B", "A-B-C", "A-B-C-D"}, slice)
	}
}
