package main

import (
	"fmt"

	iterium "github.com/mowshon/iterium/v2"
)

func main() {
	values := iterium.Slice2(iterium.Product2([]int{1, 2}, []string{"a", "b"}))
	fmt.Println(values)
}
