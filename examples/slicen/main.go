package main

import (
	"fmt"

	iterium "github.com/mowshon/iterium"
)

func main() {
	values := iterium.SliceN(iterium.Count[int](), 5)
	fmt.Println(values)
}
