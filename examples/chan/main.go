package main

import (
	"context"
	"fmt"

	iterium "github.com/mowshon/iterium/v2"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for value := range iterium.Chan(ctx, iterium.Range(3)) {
		fmt.Println(value)
	}
}
