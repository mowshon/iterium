package main

import (
	"fmt"

	iterium "github.com/mowshon/iterium/v2"
)

func main() {
	fmt.Println("product:", iterium.ProductCount(3, 2))
	fmt.Println("combinations:", iterium.CombinationsCount(4, 2))
	fmt.Println("combinations with replacement:", iterium.CombinationsWithReplacementCount(2, 3))
	fmt.Println("permutations:", iterium.PermutationCount(4, 2))

	count, ok := iterium.CombinationsCountOK(60, 30)
	fmt.Println("checked combinations:", count, ok)
}
