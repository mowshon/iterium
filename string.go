package iterium

// concatMultipleSlices merge more than two slices at once.
func concatMultipleSlices[T any](slices ...[]T) (result []T) {
	for _, s := range slices {
		result = append(result, s...)
	}

	return result
}

// AsciiLowercase represents lower case letters.
var AsciiLowercase = []string{
	"a", "b", "c", "d", "e", "f", "g",
	"h", "i", "j", "k", "l", "m", "n",
	"o", "p", "q", "r", "s", "t", "u",
	"v", "w", "x", "y", "z",
}

// AsciiUppercase represents upper case letters.
var AsciiUppercase = []string{
	"A", "B", "C", "D", "E", "F", "G",
	"H", "I", "J", "K", "L", "M", "N",
	"O", "P", "Q", "R", "S", "T", "U",
	"V", "W", "X", "Y", "Z",
}

// AsciiLetters is a concatenation of AsciiLowercase and AsciiUppercase.
var AsciiLetters = concatMultipleSlices(AsciiLowercase, AsciiUppercase)

// AsciiLowercaseBytes represents lower case letters as bytes.
var AsciiLowercaseBytes = []byte("abcdefghijklmnopqrstuvwxyz")

// AsciiUppercaseBytes represents upper case letters as bytes.
var AsciiUppercaseBytes = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

// AsciiLettersBytes is a concatenation of AsciiLowercaseBytes and AsciiUppercaseBytes.
var AsciiLettersBytes = concatMultipleSlices(AsciiLowercaseBytes, AsciiUppercaseBytes)

// Digits is a slice of the digits in the string type.
var Digits = []string{
	"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
}

// DigitsBytes is a slice of ASCII digit bytes.
var DigitsBytes = []byte("0123456789")

// HexDigits represents hexadecimal letters.
var HexDigits = []string{
	"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
	"a", "b", "c", "d", "e", "f", "A", "B", "C", "D",
	"E", "F",
}

// HexDigitsBytes represents hexadecimal bytes.
var HexDigitsBytes = []byte("0123456789abcdefABCDEF")

// OctDigits represents octadecimal letters.
var OctDigits = []string{
	"0", "1", "2", "3", "4", "5", "6", "7",
}

// OctDigitsBytes represents octal digit bytes.
var OctDigitsBytes = []byte("01234567")

// AsciiLowercaseRunes represents lower case letters as runes.
var AsciiLowercaseRunes = []rune("abcdefghijklmnopqrstuvwxyz")

// AsciiUppercaseRunes represents upper case letters as runes.
var AsciiUppercaseRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

// AsciiLettersRunes is a concatenation of AsciiLowercaseRunes and AsciiUppercaseRunes.
var AsciiLettersRunes = concatMultipleSlices(AsciiLowercaseRunes, AsciiUppercaseRunes)

// DigitsRunes is a slice of ASCII digit runes.
var DigitsRunes = []rune("0123456789")

// HexDigitsRunes represents hexadecimal runes.
var HexDigitsRunes = []rune("0123456789abcdefABCDEF")

// OctDigitsRunes represents octal digit runes.
var OctDigitsRunes = []rune("01234567")

// Punctuation is a slice of ASCII characters that
// are considered punctuation marks in the C locale
var Punctuation = []string{
	"!", "\"", "#", "$", "%", "&", "'", "(",
	")", "*", "+", ",", "-", ".", "/", ":",
	";", "<", "=", ">", "?", "@", "[", "\\",
	"]", "^", "_", "`", "{", "|", "}", "~",
}

// Whitespace contains all ASCII characters that are considered whitespace
var Whitespace = []string{
	" ", "\t", "\n", "\r", "\x0b", "\x0c",
}

// Printable is a slice of ASCII characters which are considered printable.
var Printable = concatMultipleSlices(
	AsciiLetters, Digits, Punctuation, Whitespace,
)
