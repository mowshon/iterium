package main

import (
	"fmt"

	iterium "github.com/mowshon/iterium"
)

func main() {
	checked, ok := iterium.PermutationCountOK(3, 2)
	fmt.Println("count:", checked, ok)
	for value := range iterium.Permutations([]string{"A", "B", "C"}, 2) {
		fmt.Println(value)
	}
}
