package iternium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func isLess(x int) bool {
	return x < 5
}

func startsWithJ(x string) bool {
	return x[0] == 'J'
}

func TestTakeWhileInteger(t *testing.T) {
	withIntegers := TakeWhile(Count(1, 1), isLess)
	if slice, err := withIntegers.Slice(); assert.Nil(t, err) {
		assert.Exactly(t, []int{1, 2, 3, 4}, slice)
	}
}

func TestTakeWhileString(t *testing.T) {
	withStrings := TakeWhile(New("James", "John", "David"), startsWithJ)
	if slice, err := withStrings.Slice(); assert.Nil(t, err) {
		assert.Exactly(t, []string{"James", "John"}, slice)
	}
}

func TestTakeWhileTillTheEnd(t *testing.T) {
	tillTheEnd := TakeWhile(New(2, 2, 2), isLess)
	if slice, err := tillTheEnd.Slice(); assert.Nil(t, err) {
		assert.Exactly(t, []int{2, 2, 2}, slice)
	}
}
