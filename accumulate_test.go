package iterium

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func mergeAcc(first, second string) string {
	return fmt.Sprintf("%s-%s", first, second)
}

func TestAccumulate(t *testing.T) {
	assert.Exactly(t, []int{1, 3, 6, 10, 15}, Slice(Accumulate(New(1, 2, 3, 4, 5), add)))
	assert.Exactly(t, []int{1, 2, 6, 24, 120}, Slice(Accumulate(New(1, 2, 3, 4, 5), mul)))
}

func TestAccumulateEmpty(t *testing.T) {
	assert.Exactly(t, []int{}, Slice(Accumulate(Empty[int](), add)))
}

func TestAccumulateOneValue(t *testing.T) {
	assert.Exactly(t, []int{7}, Slice(Accumulate(New(7), add)))
}

func TestAccumulateString(t *testing.T) {
	values := Accumulate(New("A", "B", "C", "D"), mergeAcc)

	assert.Exactly(t, []string{"A", "A-B", "A-B-C", "A-B-C-D"}, Slice(values))
}
