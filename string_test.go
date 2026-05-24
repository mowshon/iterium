package iterium

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFastAlphabetConstants(t *testing.T) {
	assert.Exactly(t, []byte("abcdefghijklmnopqrstuvwxyz"), AsciiLowercaseBytes)
	assert.Exactly(t, []byte("0123456789abcdefABCDEF"), HexDigitsBytes)
	assert.Exactly(t, []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"), AsciiLettersRunes)
}
