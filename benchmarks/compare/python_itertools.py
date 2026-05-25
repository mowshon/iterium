#!/usr/bin/env python3.14

import hashlib
import itertools
import string
import sys


RANGE_LIMIT = 10_000_000
PRODUCT_REPEAT = 5
MD5_PASSWORD = "zzzzz"
COMBINATION_N = 35
COMBINATION_R = 8
PERMUTATION_N = 12
PERMUTATION_R = 8
REPLACEMENT_N = 30
REPLACEMENT_R = 8
EXPECTED_RANGE = RANGE_LIMIT // 2
EXPECTED_PRODUCT = 11_881_376


def main() -> None:
    if len(sys.argv) != 2:
        print(
            "usage: python_itertools.py "
            "<range-map-filter|product-repeat5|md5-repeat5|combinations|combinations-with-replacement|permutations>",
            file=sys.stderr,
        )
        raise SystemExit(2)

    benchmark = sys.argv[1]
    if benchmark == "range-map-filter":
        count = sum(1 for value in map(lambda v: v + 1, range(RANGE_LIMIT)) if value % 2 == 0)
        must_equal(count, EXPECTED_RANGE)
        print(count)
    elif benchmark == "product-repeat5":
        count = sum(1 for _ in itertools.product(string.ascii_lowercase, repeat=PRODUCT_REPEAT))
        must_equal(count, EXPECTED_PRODUCT)
        print(count)
    elif benchmark == "md5-repeat5":
        target = hashlib.md5(MD5_PASSWORD.encode()).digest()
        found = None
        for candidate in itertools.product(string.ascii_lowercase.encode(), repeat=PRODUCT_REPEAT):
            raw = bytes(candidate)
            if hashlib.md5(raw).digest() == target:
                found = raw.decode()
                break
        must_equal(found, MD5_PASSWORD)
        print(found)
    elif benchmark == "combinations":
        count = sum(1 for _ in itertools.combinations(range(COMBINATION_N), COMBINATION_R))
        must_equal(count, 23_535_820)
        print(count)
    elif benchmark == "combinations-with-replacement":
        count = sum(1 for _ in itertools.combinations_with_replacement(range(REPLACEMENT_N), REPLACEMENT_R))
        must_equal(count, 38_608_020)
        print(count)
    elif benchmark == "permutations":
        count = sum(1 for _ in itertools.permutations(range(PERMUTATION_N), PERMUTATION_R))
        must_equal(count, 19_958_400)
        print(count)
    else:
        print(f"unknown benchmark: {benchmark}", file=sys.stderr)
        raise SystemExit(2)


def must_equal(got, want) -> None:
    if got != want:
        print(f"unexpected result: got {got!r}, want {want!r}", file=sys.stderr)
        raise SystemExit(1)


if __name__ == "__main__":
    main()
