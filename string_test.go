package iterium

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFastAlphabetConstants(t *testing.T) {
	assert.Exactly(t, []byte("abcdefghijklmnopqrstuvwxyz"), AsciiLowercaseBytes)
	assert.Exactly(t, []byte("0123456789abcdefABCDEF"), HexDigitsBytes)
	assert.Exactly(t, []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"), AsciiLettersRunes)
}

func TestPrintableMatchesPythonStringPrintableOrder(t *testing.T) {
	expected := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~ \t\n\r\x0b\x0c"

	assert.Exactly(t, expected, strings.Join(Printable, ""))
}
