package main

import (
	"fmt"

	iterium "github.com/mowshon/iterium"
)

func main() {
	fmt.Println(string(iterium.AsciiLowercaseBytes[:3]))
	fmt.Println(string(iterium.DigitsRunes[:4]))
	fmt.Println(iterium.HexDigits[:6])
}
