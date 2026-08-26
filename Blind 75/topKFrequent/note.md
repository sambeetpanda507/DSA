# Notes for `main.py`

This solution returns the `k` most frequent elements in `nums`.

## Approach

1. Count how many times each number appears using a hash map.
2. Push each `(frequency, number)` pair into a heap.
3. Use negative frequencies so Python's `heapq` behaves like a max-heap.
4. Pop from the heap `k` times and collect the numbers.

## Why it works

- The hash map stores the exact frequency of every unique number.
- The heap keeps the highest frequency at the top because frequencies are inserted as negative values.
- Each pop returns the next most frequent number, so the first `k` pops produce the answer.

## Time and Space Complexity

- Time: `O(n + m log m + k log m)`
- Space: `O(m)`

Where:

- `n` = length of `nums`
- `m` = number of unique values in `nums`

## Example in `main()`

Input:

```python
nums = [1, 2, 1, 2, 1, 2, 3, 1, 3, 2]
k = 2
```

Frequencies:

- `1 -> 4`
- `2 -> 4`
- `3 -> 2`

So the result is the two most frequent values:

```python
[1, 2]
```

The order between values with the same frequency depends on heap ordering, but for this input the current code returns `[1, 2]`.
