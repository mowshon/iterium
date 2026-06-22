package main

import (
	"fmt"

	iterium "github.com/mowshon/iterium/v2"
)

func main() {
	values := iterium.StarMap(iterium.Product([]int{2, 3}, 2), func(left, right int) int {
		return left * right
	})
	fmt.Println(iterium.Slice(values))
}
