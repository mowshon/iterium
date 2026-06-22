package main

import (
	"fmt"

	iterium "github.com/mowshon/iterium/v2"
)

func main() {
	fmt.Println("count:", iterium.CombinationsCount(4, 2))
	for value := range iterium.Combinations([]string{"A", "B", "C", "D"}, 2) {
		fmt.Println(value)
	}
}
