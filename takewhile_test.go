package iterium

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func isLess(value int) bool {
	return value < 5
}

func startsWithJ(value string) bool {
	return value[0] == 'J'
}

func TestTakeWhileInteger(t *testing.T) {
	values := TakeWhile(Count(1, 1), isLess)

	assert.Exactly(t, []int{1, 2, 3, 4}, Slice(values))
}

func TestTakeWhileString(t *testing.T) {
	values := TakeWhile(New("James", "John", "David"), startsWithJ)

	assert.Exactly(t, []string{"James", "John"}, Slice(values))
}

func TestTakeWhileTillTheEnd(t *testing.T) {
	values := TakeWhile(New(2, 2, 2), isLess)

	assert.Exactly(t, []int{2, 2, 2}, Slice(values))
}
