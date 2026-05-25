package main

import (
	"fmt"

	iterium "github.com/mowshon/iterium"
)

func main() {
	fmt.Println(iterium.Slice(iterium.Repeat("go", 3)))
	fmt.Println(iterium.SliceN(iterium.Repeat("forever", -1), 2))
}
