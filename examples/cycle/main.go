package main

import (
	"fmt"

	iterium "github.com/mowshon/iterium"
)

func main() {
	values := iterium.SliceN(iterium.Cycle(iterium.New("A", "B", "C")), 8)
	fmt.Println(values)
}
