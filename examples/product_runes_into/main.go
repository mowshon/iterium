package main

import (
	"fmt"

	iterium "github.com/mowshon/iterium/v2"
)

func main() {
	var values []string
	iterium.ProductRunesInto([]rune("ab"), 2, func(value []rune) bool {
		values = append(values, string(value))
		return true
	})
	fmt.Println(values)
}
