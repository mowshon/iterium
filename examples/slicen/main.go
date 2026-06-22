package main

import (
	"fmt"

	iterium "github.com/mowshon/iterium/v2"
)

func main() {
	values := iterium.SliceN(iterium.Count[int](), 5)
	fmt.Println(values)
}
