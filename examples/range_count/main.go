package main

import (
	"fmt"

	iterium "github.com/mowshon/iterium/v2"
)

func main() {
	fmt.Println(iterium.RangeCount(0, 10, 3))
	fmt.Println(iterium.RangeCount(10, 0, -4))
}
