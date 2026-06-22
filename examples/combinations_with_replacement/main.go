package main

import (
	"fmt"

	iterium "github.com/mowshon/iterium/v2"
)

func main() {
	checked, ok := iterium.CombinationsWithReplacementCountOK(2, 3)
	fmt.Println("count:", checked, ok)
	for value := range iterium.CombinationsWithReplacement([]string{"A", "B"}, 3) {
		fmt.Println(value)
	}
}
