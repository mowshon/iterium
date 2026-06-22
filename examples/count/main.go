package main

import (
	"fmt"

	iterium "github.com/mowshon/iterium/v2"
)

func main() {
	fmt.Println(iterium.SliceN(iterium.Count(10, 5), 5))
}
