package main

import (
	"context"
	"fmt"

	iterium "github.com/mowshon/iterium/v2"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for value := range iterium.Chan2(ctx, iterium.Product2([]int{1}, []string{"a", "b"})) {
		fmt.Println(value)
	}
}
