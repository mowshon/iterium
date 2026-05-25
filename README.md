# Iterium

Iterium is a Go iterator toolkit inspired by Python's `itertools`, built on Go's official `iter` package.

It is useful when you want to build lazy, synchronous pipelines without creating channels, goroutines, or intermediate slices for every step. Iterium works best for finite data transformations, bounded search spaces, test-case generation, combinatorics, and small adapter boundaries where a slice or channel is needed at the edge.

## Benchmark: Iterium vs Python 3.14 `itertools`

These benchmarks compare Iterium best-practice streaming code against Python 3.14 `itertools` equivalents. The Go side uses `Into` APIs where the generated value is only inspected and not retained; this avoids per-item result allocation and is the recommended style for hot combinatoric loops.

Local run: Linux x86_64, Go 1.26.2, Python 3.14.5. Times and peak RSS are best of 3 runs, measured with `/usr/bin/time` after building the Go benchmark binary. RSS is process-level memory, so it includes the Go runtime or Python interpreter baseline. The percent columns are relative reductions versus Python: `(python - iterium) / python`.

| Workload | Results streamed | Iterium API | Python API | Iterium time | Python time | Speedup | Time reduction | Iterium RSS | Python RSS | RSS reduction |
| --- | ---: | --- | --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: |
| Range -> map -> filter | 10,000,000 range values | `Range` + `Map` + `Filter` | `range` + `map` + generator filter | 0.06 s | 0.94 s | 15.7x | 93.6% | 2.3 MiB | 13.1 MiB | 82.9% |
| Cartesian product count | 26^5 = 11,881,376 products | `ProductBytesInto` | `itertools.product` | 0.03 s | 0.42 s | 14.0x | 92.9% | 2.3 MiB | 13.1 MiB | 82.9% |
| MD5 brute force, worst-case `zzzzz` | 26^5 = 11,881,376 MD5 checks | `ProductBytesInto` + `crypto/md5` | `itertools.product` + `hashlib.md5` | 1.26 s | 7.53 s | 6.0x | 83.3% | 2.3 MiB | 13.1 MiB | 82.9% |
| Combinations count | C(35, 8) = 23,535,820 combinations | `CombinationsInto` | `itertools.combinations` | 0.08 s | 1.01 s | 12.6x | 92.1% | 2.1 MiB | 13.0 MiB | 83.7% |
| Combinations with replacement count | C(37, 8) = 38,608,020 combinations | `CombinationsWithReplacementInto` | `itertools.combinations_with_replacement` | 0.12 s | 1.60 s | 13.3x | 92.5% | 2.3 MiB | 13.1 MiB | 82.9% |
| Permutations count | P(12, 8) = 19,958,400 permutations | `PermutationsInto` | `itertools.permutations` | 0.09 s | 0.85 s | 9.4x | 89.4% | 2.3 MiB | 13.0 MiB | 82.7% |

Conclusion: on these streaming workloads, Iterium is roughly 6x to 16x faster than Python 3.14 `itertools` and uses about 83% less peak resident memory as a command-line process. The MD5 crack benchmark is the most realistic mixed workload here because it combines iterator overhead, candidate construction, early stopping, and hashing; Iterium completes the same worst-case 5-letter lowercase search in 1.26 s versus 7.53 s for Python.

For best results in Go, use the regular safe APIs such as `Product`, `Combinations`, and `Permutations` when you need to keep yielded slices. Use the `Into` APIs when you only inspect each value during iteration. In Python, the closest best practice is also to stream with `itertools` and avoid materializing large products or combinations.

Reproduce the comparison:

```bash
go build -o /tmp/iterium-compare ./benchmarks/compare/go
for case in range-map-filter product-repeat5 md5-repeat5 combinations combinations-with-replacement permutations; do
  /usr/bin/time -f "iterium $case: %e seconds, %M KB RSS" /tmp/iterium-compare "$case"
  /usr/bin/time -f "python  $case: %e seconds, %M KB RSS" /usr/bin/python3.14 benchmarks/compare/python_itertools.py "$case"
done
```

Iterium uses `iter.Seq` and `iter.Seq2` as the public API. That means most values are consumed with normal Go `for range` loops:

```go
for value := range iterium.Range(0, 10, 2) {
	fmt.Println(value)
}
```

## Installation

```bash
go get github.com/mowshon/iterium
```

```go
import "github.com/mowshon/iterium"
```

Iterium requires Go 1.23.4 or newer.

## Quick Example

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

The sequence above is lazy. `Range` does not build a slice of numbers, `Map` does not build a slice of mapped values, and `Filter` only evaluates values as the final consumer asks for them.

## How Iterium Fits Go

Go's standard iterator type is a function that pushes values into a `yield` callback. When the consumer breaks from a `for range` loop, `yield` returns false and the iterator stops early.

Use Iterium when:

- You want reusable sequence helpers with the normal `for value := range seq` syntax.
- You want Python-like iterator tools in Go without hiding goroutines or channels.
- You need combinatorics such as Cartesian products, combinations, and permutations.
- You want to stop early without generating the rest of a sequence.

Prefer a plain `for` loop when the logic is shorter and clearer as direct Go code. Iterators are strongest when they express a pipeline or reusable sequence.

## Sources

Sources create sequences.

| Function | Produces | Common use case |
| --- | --- | --- |
| `New(values...)` | The provided values | Turn a small list into a sequence for a pipeline |
| `Empty[T]()` | No values | Return an empty sequence from generic code |
| `Range(args...)` | A finite numeric range | Page numbers, IDs, offsets, bounded numeric work |
| `Count(args...)` | An unbounded numeric sequence | Generate indexes or ticks until another step stops |
| `Repeat(value, n)` | A repeated value | Fill defaults, pair constants with generated values |

### New

`New` is the smallest source. Use it when values are already known and you want to feed them through the same iterator pipeline as generated data.

```go
names := iterium.Filter(iterium.New("alice", "", "bob"), func(name string) bool {
	return name != ""
})

fmt.Println(iterium.Slice(names)) // [alice bob]
```

### Empty

`Empty` is useful in functions that return `iter.Seq[T]` but sometimes have nothing to yield.

```go
func activeIDs(enabled bool) iter.Seq[int] {
	if !enabled {
		return iterium.Empty[int]()
	}
	return iterium.Range(1, 4)
}
```

### Range

`Range` is finite and works like Python-style ranges:

```go
iterium.Slice(iterium.Range(5))        // [0 1 2 3 4]
iterium.Slice(iterium.Range(2, 8, 2))  // [2 4 6]
iterium.Slice(iterium.Range(-3))       // [0 -1 -2]
iterium.Slice(iterium.Range(10, 0, -3)) // [10 7 4 1]
```

A real use case is pagination or chunk offsets:

```go
const pageSize = 100

for offset := range iterium.Range(0, totalRows, pageSize) {
	rows := loadRows(offset, pageSize)
	process(rows)
}
```

Use `RangeCount(start, stop, step)` when you need to size or validate work before iterating.

### Count

`Count` is infinite. Always combine it with `SliceN`, `SliceUntil`, `TakeWhile`, `FirstTrue`, or a `break`.

```go
ids := iterium.SliceN(iterium.Count(1000, 5), 4)
fmt.Println(ids) // [1000 1005 1010 1015]
```

Use it for generated indexes:

```go
for attempt := range iterium.Count(1) {
	if send() == nil || attempt == 3 {
		break
	}
}
```

### Repeat

`Repeat(value, n)` yields exactly `n` values. A negative `n` repeats forever.

```go
fmt.Println(iterium.Slice(iterium.Repeat("pending", 3))) // [pending pending pending]
```

A common pattern is pairing a constant with generated values:

```go
for value := range iterium.Map(iterium.Range(3), func(id int) string {
	return fmt.Sprintf("tenant-a:%d", id)
}) {
	fmt.Println(value)
}
```

## Transforms And Search

Transforms take a sequence and return another sequence.

| Function | Use when |
| --- | --- |
| `Map(seq, fn)` | Each value should be converted to another value |
| `Filter(seq, pred)` | Keep values that pass a predicate |
| `FilterFalse(seq, pred)` | Keep values that fail a predicate |
| `Accumulate(seq, op)` | Build running totals or rolling state |
| `TakeWhile(seq, pred)` | Read until the first value that fails a condition |
| `DropWhile(seq, pred)` | Skip an initial prefix, then keep the rest |
| `StarMap(seq, fn)` | Apply a binary function to two-item slices |
| `Cycle(seq)` | Repeat a finite sequence forever |
| `FirstTrue(seq, pred)` | Find the first matching value |
| `FirstFalse(seq, pred)` | Find the first non-matching value |

### Map

Use `Map` for conversions that should stay lazy.

```go
type User struct {
	ID   int
	Name string
}

names := iterium.Map(usersSeq, func(user User) string {
	return user.Name
})
```

### Filter And FilterFalse

Use `Filter` when the predicate describes the values you want. Use `FilterFalse` when the predicate describes values to reject.

```go
adults := iterium.Filter(usersSeq, func(user User) bool {
	return user.Age >= 18
})

nonEmpty := iterium.FilterFalse(iterium.New("", "api", "", "worker"), func(value string) bool {
	return value == ""
})

fmt.Println(iterium.Slice(nonEmpty)) // [api worker]
```

### Accumulate

Use `Accumulate` for running totals, balances, counters, or progressive state.

```go
dailyRevenue := iterium.New(120, 80, 100, -30)

runningRevenue := iterium.Accumulate(dailyRevenue, func(total, value int) int {
	return total + value
})

fmt.Println(iterium.Slice(runningRevenue)) // [120 200 300 270]
```

### TakeWhile

Use `TakeWhile` when the source is ordered and the first failing value means the rest is not needed.

```go
recent := iterium.TakeWhile(eventsByNewestFirst, func(event Event) bool {
	return event.CreatedAt.After(cutoff)
})

for event := range recent {
	index(event)
}
```

This is ideal for logs, sorted timestamps, increasing counters, or any source where early stop saves work.

### DropWhile

Use `DropWhile` to skip a header, warm-up period, or initial invalid prefix.

```go
measurements := iterium.DropWhile(sensorReadings, func(reading Reading) bool {
	return reading.Status == "warming-up"
})
```

After the predicate first returns false, all remaining values are yielded without checking the prefix again.

### FirstTrue And FirstFalse

Use these when you only need one value and want the upstream pipeline to stop immediately.

```go
user, ok := iterium.FirstTrue(usersSeq, func(user User) bool {
	return user.Email == targetEmail
})
if !ok {
	return errors.New("user not found")
}
```

### StarMap

`StarMap` expects each yielded slice to contain at least two values and applies a binary function to `value[0]` and `value[1]`.

```go
sums := iterium.StarMap(iterium.Product([]int{1, 2, 3}, 2), func(a, b int) int {
	return a + b
})

fmt.Println(iterium.Slice(sums)) // [2 3 4 3 4 5 4 5 6]
```

### Cycle

`Cycle` caches the first pass through a finite sequence, then replays the saved values forever.

```go
rotatingWorkers := iterium.Cycle(iterium.New("worker-a", "worker-b", "worker-c"))

for worker := range iterium.SliceN(rotatingWorkers, 5) {
	fmt.Println(worker)
}
```

Use it for round-robin assignment, repeating schedules, or test fixtures. Do not pass an infinite sequence to `Cycle`; it will keep caching forever.

## Combinatorics

This is where users often choose the wrong method. The main question is whether order matters and whether a value may be reused.

| Function | Order matters | Reuse allowed | Example output for `["A", "B"]`, `r=2` | Use case |
| --- | --- | --- | --- | --- |
| `Product(values, r)` | Yes | Yes | `AA AB BA BB` | All codes, all test-matrix choices, search spaces |
| `Combinations(values, r)` | No | No | `AB` | Pick unique teams, feature subsets, lottery-style selections |
| `CombinationsWithReplacement(values, r)` | No | Yes | `AA AB BB` | Multisets, repeated quantities, bundles |
| `Permutations(values, r)` | Yes | No | `AB BA` | Ordered plans, route/order testing, rank lists |

Stack Overflow questions about itertools often come down to this distinction: if you need "permutations with replacement", you usually want Cartesian product with `repeat`, not combinations.

### Product

`Product(symbols, repeat)` yields Cartesian products of one alphabet repeated `repeat` times.

```go
for code := range iterium.Product([]string{"A", "B", "C"}, 2) {
	fmt.Println(strings.Join(code, ""))
}
// AA AB AC BA BB BC CA CB CC
```

Real use cases:

- Generate all short coupon/code candidates from an alphabet.
- Run every combination in a test matrix.
- Explore all choices in a bounded search space.
- Build pairs/triples from a single homogeneous set.

For large products, use `ProductCount` or `ProductCountOK` before iterating:

```go
count, ok := iterium.ProductCountOK(len(iterium.AsciiLowercase), 6)
if !ok || count > 10_000_000 {
	return errors.New("search space is too large")
}
```

### Product2

Use `Product2` when you have two different element types and want Go to keep both types.

```go
for region, tier := range iterium.Product2(
	[]string{"us-east", "eu-west"},
	[]int{1, 2, 3},
) {
	fmt.Printf("%s tier %d\n", region, tier)
}
```

`Product2` returns `iter.Seq2[A, B]`, so consume it with `for first, second := range ...`.

### Combinations

`Combinations(symbols, r)` yields unique selections where order does not matter.

```go
members := []string{"Ana", "Ben", "Cy", "Dee"}

for team := range iterium.Combinations(members, 2) {
	fmt.Println(team)
}
```

Real use cases:

- Select review pairs from a team.
- Generate unique feature subsets.
- Choose sample groups where `AB` and `BA` mean the same thing.

### CombinationsWithReplacement

Use `CombinationsWithReplacement` when values may repeat, but order still does not matter.

```go
flavors := []string{"vanilla", "chocolate", "mint"}

for scoops := range iterium.CombinationsWithReplacement(flavors, 2) {
	fmt.Println(scoops)
}
```

Real use cases:

- Choose product bundles where duplicate items are allowed.
- Model quantities without caring about order.
- Generate multiset fixtures for tests.

### Permutations

`Permutations(symbols, r)` yields ordered arrangements without reusing a value.

```go
steps := []string{"build", "test", "deploy"}

for order := range iterium.Permutations(steps, 3) {
	fmt.Println(order)
}
```

Real use cases:

- Try possible task orders.
- Test rank lists.
- Generate route/order candidates where `AB` and `BA` are different.

## Count Helpers

Count helpers let you estimate work before doing it.

```go
iterium.ProductCount(26, 4)
iterium.CombinationsCount(26, 5)
iterium.CombinationsWithReplacementCount(26, 5)
iterium.PermutationCount(10, 5)
```

The non-OK variants saturate at `math.MaxInt64` on overflow. The `*CountOK` variants return `(count, ok)`:

```go
count, ok := iterium.CombinationsCountOK(60, 30)
if !ok {
	return errors.New("too many combinations")
}
fmt.Println(count)
```

## Safe Slices vs Reused Buffers

The normal combinatoric sequence APIs yield slices that are safe to keep:

```go
var saved [][]string
for value := range iterium.Product([]string{"A", "B"}, 2) {
	saved = append(saved, value)
}
```

That safety may allocate one slice per yielded value. For hot paths, use the `Into` APIs. They reuse one result buffer and are much faster, but the value is only valid until the next callback call.

```go
var saved [][]string

iterium.ProductInto([]string{"A", "B"}, 2, func(value []string) bool {
	saved = append(saved, append([]string(nil), value...))
	return true
})
```

If you only inspect each value inside the callback, do not copy:

```go
iterium.ProductBytesInto(iterium.AsciiLowercaseBytes, 4, func(value []byte) bool {
	if bytes.Equal(value, []byte("test")) {
		fmt.Println("found")
		return false
	}
	return true
})
```

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

## Alphabet Constants

String alphabets are convenient with the safe APIs:

```go
iterium.AsciiLowercase
iterium.AsciiUppercase
iterium.AsciiLetters
iterium.Digits
iterium.HexDigits
iterium.OctDigits
iterium.Punctuation
iterium.Whitespace
iterium.Printable
```

Byte and rune alphabets are better for allocation-sensitive code:

```go
iterium.AsciiLowercaseBytes
iterium.AsciiUppercaseBytes
iterium.AsciiLettersBytes
iterium.DigitsBytes
iterium.HexDigitsBytes
iterium.OctDigitsBytes

iterium.AsciiLowercaseRunes
iterium.AsciiUppercaseRunes
iterium.AsciiLettersRunes
iterium.DigitsRunes
iterium.HexDigitsRunes
iterium.OctDigitsRunes
```

## Adapters

Adapters collect or bridge sequences at application boundaries.

| Function | Use when |
| --- | --- |
| `Slice(seq)` | The sequence is finite and you need a `[]T` |
| `SliceN(seq, n)` | The sequence may be infinite or you only need the first `n` values |
| `SliceUntil(seq, stop)` | Collect until a stop condition is reached |
| `Slice2(seq2)` | Collect `iter.Seq2[A, B]` into `[]Pair[A, B]` |
| `Chan(ctx, seq)` | A channel API is required at a boundary |
| `Chan2(ctx, seq2)` | A channel of `Pair[A, B]` is required |

### Slice, SliceN, SliceUntil

```go
all := iterium.Slice(iterium.Range(5))
some := iterium.SliceN(iterium.Count[int](), 5)
until := iterium.SliceUntil(iterium.Count[int](), func(value int) bool {
	return value == 5
})

fmt.Println(all)   // [0 1 2 3 4]
fmt.Println(some)  // [0 1 2 3 4]
fmt.Println(until) // [0 1 2 3 4]
```

Do not call `Slice` on `Count`, `Repeat(value, -1)`, or `Cycle` unless another transform makes the sequence finite.

### Slice2

```go
pairs := iterium.Slice2(iterium.Product2(
	[]int{1, 2},
	[]string{"small", "large"},
))

fmt.Println(pairs)
```

### Chan And Chan2

Use channels only when the rest of your program already speaks channels. For normal pipelines, range over the sequence directly.

`Chan` requires a context so the producer can stop if the consumer exits early.

```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel()

for value := range iterium.Chan(ctx, iterium.Range(5)) {
	if value == 3 {
		break
	}
	fmt.Println(value)
}
```

## Pull-Style Iteration

When a `for range` loop is not the right shape, use the standard library `iter.Pull`.

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

This is useful when you need to manually coordinate multiple sequences.

## Real Use Cases

### Test Matrix Generation

```go
browsers := []string{"chrome", "firefox"}
locales := []string{"en", "fr", "es"}

for browser, locale := range iterium.Product2(browsers, locales) {
	t.Run(browser+"-"+locale, func(t *testing.T) {
		runUITest(t, browser, locale)
	})
}
```

Use `Product2` for two different dimensions. Use `Product` when every position comes from the same alphabet or symbol list.

### Early-Stop Search

```go
code, ok := iterium.FirstTrue(
	iterium.Product(iterium.Digits, 4),
	func(parts []string) bool {
		return strings.Join(parts, "") == "0427"
	},
)
if ok {
	fmt.Println("found", strings.Join(code, ""))
}
```

Only candidates up to the matching value are generated. For a hot path, prefer `ProductBytesInto` or `ProductStringInto`.

### Running Business Totals

```go
balances := iterium.Accumulate(transactions, func(balance, tx int64) int64 {
	return balance + tx
})

for balance := range balances {
	if balance < 0 {
		alert(balance)
		break
	}
}
```

### Round-Robin Assignment

```go
workers := iterium.Cycle(iterium.New("a", "b", "c"))
tasks := []Task{task1, task2, task3, task4}

nextWorker, stop := iter.Pull(workers)
defer stop()

for _, task := range tasks {
	worker, _ := nextWorker()
	assign(task, worker)
}
```

## Practical Rules

- Use `for range` as the default consumer.
- Use `Slice`, `SliceN`, and `SliceUntil` only at boundaries.
- Use `FirstTrue`, `FirstFalse`, `TakeWhile`, or `break` to stop expensive pipelines early.
- Use `Product` for order-sensitive choices with replacement.
- Use `Combinations` when order does not matter and values cannot repeat.
- Use `CombinationsWithReplacement` when order does not matter and values can repeat.
- Use `Permutations` when order matters and values cannot repeat.
- Use `Into` APIs for performance-sensitive combinatorics.
- Copy values from `Into` callbacks only when you need to keep them.
- Do not pass infinite sequences to `Slice` or `Cycle`.

## References

- Go `iter` package documentation: https://pkg.go.dev/iter
- Go blog, "Range Over Function Types": https://go.dev/blog/range-functions
- Python `itertools` documentation: https://docs.python.org/3/library/itertools.html
- Stack Overflow example on choosing `product` when order matters and replacement is needed: https://stackoverflow.com/questions/68665316/python-use-of-itertools-to-find-all-combinations-permutations-with-replacemen
- Stack Overflow example showing why large Cartesian products should be streamed instead of materialized: https://stackoverflow.com/questions/72471146/memory-problem-when-using-itertools-and-product-python-brute-force-script
