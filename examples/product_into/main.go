package main

import (
	"fmt"

	iterium "github.com/mowshon/iterium/v2"
)

func main() {
	var kept [][]string
	iterium.ProductInto([]string{"A", "B"}, 2, func(value []string) bool {
		kept = append(kept, append([]string(nil), value...))
		return true
	})
	fmt.Println(kept)
}
