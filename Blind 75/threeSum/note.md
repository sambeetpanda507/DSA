# 3Sum Notes

## Problem

Given an integer array `nums`, return all unique triplets `[a, b, c]` such that:

```text
a + b + c = 0
```

The solution must avoid duplicate triplets.

## Approach

The implementation in `main.py` uses sorting plus a two-pointer scan.

1. Sort `nums`.
2. Fix one value at index `i`.
3. Search the remaining sorted range with two pointers:
   - `j = i + 1`
   - `k = len(nums) - 1`
4. Compute `nums[i] + nums[j] + nums[k]`.
5. Move pointers based on the sum:
   - If the sum is less than `0`, move `j` right to increase the sum.
   - If the sum is greater than `0`, move `k` left to decrease the sum.
   - If the sum is `0`, store the triplet and move both pointers.

## Duplicate Handling

Because the array is sorted, duplicates are adjacent.

After finding a valid triplet:

```python
while j < len(nums) and nums[j] == nums[j - 1]:
    j += 1

while k >= 0 and nums[k] == nums[k + 1]:
    k -= 1
```

This skips duplicate values for the second and third positions.

After finishing one fixed `i` value:

```python
while i < len(nums) - 2 and nums[i] == nums[i - 1]:
    i += 1
```

This skips duplicate starting values so the same triplet is not produced again.

## Example

Input:

```python
nums = [-1, 0, 1, 2, -1, -4]
```

Sorted:

```python
[-4, -1, -1, 0, 1, 2]
```

Output:

```python
[[-1, -1, 2], [-1, 0, 1]]
```

## Complexity

Sorting costs `O(n log n)`.

The outer loop runs `O(n)` times, and each two-pointer scan is `O(n)`, so the main search is:

```text
O(n^2)
```

Overall time complexity:

```text
O(n^2)
```

Extra space complexity, excluding the returned result:

```text
O(1)
```

Python's sort may use additional internal space depending on implementation details.
