package iterium

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestMapInteger(t *testing.T) {
	newMap := Map(New(1, 2, 3, 4, 5), func(x int) int {
		return x * 2
	})

	if result, err := newMap.Slice(); assert.Nil(t, err) {
		assert.Exactly(t, []int{2, 4, 6, 8, 10}, result)
	}
}

func TestMapString(t *testing.T) {
	newMap := Map(New("a", "b", "c", "d"), func(x string) string {
		return strings.ToUpper(x)
	})

	if result, err := newMap.Slice(); assert.Nil(t, err) {
		assert.Exactly(t, []string{"A", "B", "C", "D"}, result)
	}
}
