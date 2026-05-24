# Iterium

Iterium is a Go iterator toolkit inspired by Python's `itertools`, built on the official Go `iter` package.

Version 2 uses `iter.Seq` and `iter.Seq2` as the public API. Iterators are synchronous functions, so normal pipelines do not allocate channels or start goroutines.

## Installation

```bash
go get github.com/mowshon/iterium
```

```go
import "github.com/mowshon/iterium"
```

Iterium v2 requires Go 1.23.4 or newer.

## Basic Usage

```go
for value := range iterium.Range(0, 10, 2) {
	fmt.Println(value)
}
```

```go
values := iterium.Filter(
	iterium.Map(iterium.Range(0, 10), func(value int) int {
		return value + 1
	}),
	func(value int) bool {
		return value%2 == 0
	},
)

fmt.Println(iterium.Slice(values)) // [2 4 6 8 10]
```

## Sources

```go
iterium.New(1, 2, 3)
iterium.Empty[int]()
iterium.Range(5)
iterium.Count(10, 2)
iterium.Repeat("A", 3)
```

`Count` and `Repeat(value, -1)` are unbounded. Use `SliceN`, `SliceUntil`, or break from a `for range` loop.

```go
firstFive := iterium.SliceN(iterium.Count[int](), 5)
```

## Transforms

```go
iterium.Map(seq, fn)
iterium.Filter(seq, pred)
iterium.FilterFalse(seq, pred)
iterium.Accumulate(seq, op)
iterium.TakeWhile(seq, pred)
iterium.DropWhile(seq, pred)
iterium.StarMap(seq, binaryFn)
iterium.Cycle(seq)
```

`FirstTrue` and `FirstFalse` return a value and a boolean:

```go
value, ok := iterium.FirstTrue(iterium.Range(10), func(value int) bool {
	return value == 5
})
```

These helpers are Iterium extras, not CPython `itertools` functions: `Range`, `Map`, `Filter`, `FirstTrue`, and `FirstFalse`.

## Combinatorics

```go
for value := range iterium.Product([]string{"A", "B"}, 2) {
	fmt.Println(value)
}

for value := range iterium.Combinations([]string{"A", "B", "C"}, 2) {
	fmt.Println(value)
}

for value := range iterium.CombinationsWithReplacement([]string{"A", "B"}, 2) {
	fmt.Println(value)
}

for value := range iterium.Permutations([]string{"A", "B", "C"}, 2) {
	fmt.Println(value)
}
```

Two different element types can use `Product2`, which returns `iter.Seq2`:

```go
for number, letter := range iterium.Product2([]int{1, 2}, []string{"a", "b"}) {
	fmt.Println(number, letter)
}
```

Count helpers are available for checking sizes before iterating:

```go
iterium.ProductCount(26, 4)
iterium.CombinationsCount(26, 5)
iterium.CombinationsWithReplacementCount(26, 5)
iterium.PermutationCount(10, 5)
```

The `*CountOK` variants report integer overflow instead of saturating.

## Fast Reused-Buffer APIs

Safe sequence APIs yield slices that callers can keep. The `Into` APIs reuse a buffer for speed:

```go
iterium.ProductBytesInto(iterium.AsciiLowercaseBytes, 4, func(value []byte) bool {
	// Copy value here if it must live after this callback.
	return true
})
```

Byte and rune alphabet variables such as `AsciiLowercaseBytes`, `DigitsBytes`, and `HexDigitsRunes` are provided for these fast paths.

Available reused-buffer APIs:

```go
iterium.ProductInto(symbols, repeat, yield)
iterium.ProductBytesInto(symbols, repeat, yield)
iterium.ProductStringInto(symbols, repeat, yield)
iterium.ProductRunesInto(symbols, repeat, yield)
iterium.CombinationsInto(symbols, r, yield)
iterium.CombinationsWithReplacementInto(symbols, r, yield)
iterium.PermutationsInto(symbols, r, yield)
```

## Adapters

Use slices for finite sequences:

```go
all := iterium.Slice(iterium.Range(5))
some := iterium.SliceN(iterium.Count[int](), 5)
```

Use channels only at boundaries that need channels. Cancellation is required:

```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel()

for value := range iterium.Chan(ctx, iterium.Range(5)) {
	fmt.Println(value)
}
```

For pull-style iteration, use the standard library:

```go
next, stop := iter.Pull(iterium.Range(5))
defer stop()

for {
	value, ok := next()
	if !ok {
		break
	}
	fmt.Println(value)
}
```
