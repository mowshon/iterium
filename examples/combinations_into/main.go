package main

import (
	"fmt"

	iterium "github.com/mowshon/iterium"
)

func main() {
	var kept [][]string
	iterium.CombinationsInto([]string{"A", "B", "C"}, 2, func(value []string) bool {
		kept = append(kept, append([]string(nil), value...))
		return true
	})
	fmt.Println(kept)
}
