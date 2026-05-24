# Migrating From Iterium v1 To v2

Iterium v2 replaces the old channel-backed `Iter[T]` API with Go's official `iter.Seq` and `iter.Seq2` APIs.

## Module Path

```go
import "github.com/mowshon/iterium/v2"
```

## Iteration

v1:

```go
values := iterium.Range(5)
for {
	value, err := values.Next()
	if err != nil {
		break
	}
	fmt.Println(value)
}
```

v2:

```go
for value := range iterium.Range(5) {
	fmt.Println(value)
}
```

For explicit pull-style control, use the standard library:

```go
next, stop := iter.Pull(iterium.Range(5))
defer stop()

value, ok := next()
```

## Channels

v1:

```go
for value := range iterium.Range(5).Chan() {
	fmt.Println(value)
}
```

v2:

```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel()

for value := range iterium.Chan(ctx, iterium.Range(5)) {
	fmt.Println(value)
}
```

The context is required so a channel producer can stop when the consumer exits early.

## Slices

v1:

```go
values, err := iterium.Range(5).Slice()
```

v2:

```go
values := iterium.Slice(iterium.Range(5))
```

For unbounded or unknown-length sequences:

```go
first := iterium.SliceN(iterium.Count[int](), 5)
until := iterium.SliceUntil(iterium.Count[int](), func(value int) bool {
	return value == 5
})
```

## Counts

Iterators no longer have `.Count()`.

```go
total := iterium.ProductCount(26, 4)
total, ok := iterium.CombinationsCountOK(26, 5)
```

The non-OK count helpers saturate at `math.MaxInt64` on overflow. The `*CountOK` helpers report overflow with `ok == false`.

## Infinite Metadata

`IsInfinite` and `SetInfinite` were removed. `iter.Seq` does not carry metadata, so callers own their bounds:

```go
values := iterium.SliceN(iterium.Repeat("A", -1), 3)
```

## Removed APIs

`Iter[T]`, `IterRecover`, `.Next()`, `.Chan()`, `.Slice()`, `.Close()`, `.Count()`, `IsInfinite`, and `SetInfinite` are not part of v2.

## Symbol Examples

```go
iterium.New(1, 2, 3)
iterium.Empty[int]()
iterium.Range(0, 10)
iterium.Count(0, 2)
iterium.Repeat("A", 3)
```

```go
iterium.Map(iterium.Range(3), func(value int) int { return value * 2 })
iterium.Filter(iterium.Range(5), func(value int) bool { return value%2 == 0 })
iterium.FilterFalse(iterium.Range(5), func(value int) bool { return value%2 == 0 })
iterium.Accumulate(iterium.New(1, 2, 3), func(a, b int) int { return a + b })
iterium.TakeWhile(iterium.Count[int](), func(value int) bool { return value < 3 })
iterium.DropWhile(iterium.Range(5), func(value int) bool { return value < 3 })
iterium.StarMap(iterium.Product([]int{1, 2}, 2), func(a, b int) int { return a + b })
```

```go
value, ok := iterium.FirstTrue(iterium.Range(5), func(value int) bool { return value == 3 })
value, ok = iterium.FirstFalse(iterium.Range(5), func(value int) bool { return value < 3 })
```

```go
iterium.Cycle(iterium.New(1, 2, 3))
iterium.Product([]string{"a", "b"}, 2)
iterium.Product2([]int{1, 2}, []string{"a", "b"})
iterium.Combinations([]string{"a", "b", "c"}, 2)
iterium.CombinationsWithReplacement([]string{"a", "b"}, 2)
iterium.Permutations([]string{"a", "b", "c"}, 2)
```

```go
iterium.Slice(iterium.Range(5))
iterium.Slice2(iterium.Product2([]int{1}, []string{"a"}))
iterium.Chan(ctx, iterium.Range(5))
iterium.Chan2(ctx, iterium.Product2([]int{1}, []string{"a"}))
```
