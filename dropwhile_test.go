package iternium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDropWhile(t *testing.T) {
	dropwhile := DropWhile(New(1, 4, 6, 4, 1), func(x int) bool {
		return x < 5
	})

	if slice, err := dropwhile.Slice(); assert.Nil(t, err) {
		assert.Exactly(t, []int{6, 4, 1}, slice)
	}
}
