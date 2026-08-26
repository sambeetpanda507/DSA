# Notes for `main.py`

`main.py` solves the **Longest Consecutive Sequence** problem.

## What the code does

The `Solution.longestConsecutive()` method:

1. Converts `nums` into a set for `O(1)` average membership checks.
2. Iterates over each unique number.
3. Treats a number as the start of a sequence only if `num - 1` is not in the set.
4. Walks forward from that start and counts how long the streak is.
5. Keeps the maximum streak length seen so far.

## Example

For `nums = [1, 0, 1, 2]`, the set becomes `{0, 1, 2}`.

The consecutive sequence is `0, 1, 2`, so the result is `3`.

## Complexity

- Time: `O(n)` average
- Space: `O(n)`

## Notes

- Duplicates do not affect the result because the input is converted to a set.
- The `main()` function is only a small local test harness that prints the result for one sample input.
