# Iterium Examples

Each subdirectory is a small runnable Go program for one public Iterium function or a tightly related pair.

Run one example:

```bash
go run ./examples/range
```

Run all examples:

```bash
go test ./examples/...
```

## Layout

- Source iterators: `new`, `empty`, `range`, `count`, `repeat`
- Transforms and searches: `map`, `filter`, `filterfalse`, `accumulate`, `takewhile`, `dropwhile`, `firsttrue`, `firstfalse`, `starmap`, `cycle`
- Combinatorics: `product`, `product2`, `combinations`, `combinations_with_replacement`, `permutations`
- Count helpers: `range_count`, `product_count`, `count_helpers`
- Reused-buffer APIs: `product_into`, `product_bytes_into`, `product_string_into`, `product_runes_into`, `combinations_into`, `combinations_with_replacement_into`, `permutations_into`
- Adapters: `slice`, `slicen`, `sliceuntil`, `slice2`, `chan`, `chan2`
- Constants: `alphabets`
- Applied demos: `md5_crack`

The safe sequence APIs yield values that callers can keep. The `Into` examples copy the reused buffer before storing it.
