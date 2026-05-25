package main

import (
	"fmt"

	iterium "github.com/mowshon/iterium"
)

func main() {
	var values []string
	iterium.ProductStringInto("ab", 2, func(value []byte) bool {
		values = append(values, string(value))
		return true
	})
	fmt.Println(values)
}
