package iterium

import (
	"context"
	"iter"
	"strings"
)

func add(a, b int) int {
	return a + b
}

func mul(a, b int) int {
	return a * b
}

func joinStrings(seq iter.Seq[[]string]) []string {
	return Slice(Map(seq, func(value []string) string {
		return strings.Join(value, "")
	}))
}

func testContext() (context.Context, context.CancelFunc) {
	return context.WithCancel(context.Background())
}
