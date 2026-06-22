package main

import (
	"fmt"

	iterium "github.com/mowshon/iterium/v2"
)

func main() {
	count := iterium.ProductCount(26, 4)
	checked, ok := iterium.ProductCountOK(26, 4)
	fmt.Println(count)
	fmt.Println(checked, ok)
}
