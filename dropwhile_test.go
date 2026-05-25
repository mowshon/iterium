package iterium

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDropWhile(t *testing.T) {
	values := DropWhile(New(1, 4, 6, 4, 1), func(value int) bool {
		return value < 5
	})

	assert.Exactly(t, []int{6, 4, 1}, Slice(values))
}

func TestDropWhileEmptyWhenPredicateNeverFalse(t *testing.T) {
	values := DropWhile(New(1, 2, 3), func(value int) bool {
		return value < 5
	})

	assert.Exactly(t, []int{}, Slice(values))
}
