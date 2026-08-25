# Container With Most Water

## Problem

Given an array `height`, where each value is the height of a vertical line at that index, choose two lines that form a container with the x-axis. Return the maximum amount of water the container can hold.

The area between two lines is:

```text
width * min(left_height, right_height)
```

## Approach

Use two pointers:

- `start` begins at the left edge.
- `end` begins at the right edge.
- Compute the area formed by those two lines.
- Move the pointer with the smaller height inward.

Moving the taller line cannot help if the shorter line stays fixed, because the width gets smaller and the limiting height does not increase. The only useful move is to try to find a taller line on the shorter side.

## Algorithm

1. Set `start = 0` and `end = len(height) - 1`.
2. Track the best area in `result`.
3. While `start < end`:
   - Compute `width = end - start`.
   - Compute `area = width * min(height[start], height[end])`.
   - Update `result`.
   - Move the pointer with the smaller height.
4. Return `result`.

## Complexity

- Time: `O(n)` because each pointer moves at most `n` times.
- Space: `O(1)` because only a few variables are used.

## Example

```python
height = [1, 8, 6, 2, 5, 4, 8, 3, 7]
```

The best pair is at indexes `1` and `8`:

```text
width = 8 - 1 = 7
height = min(8, 7) = 7
area = 7 * 7 = 49
```

So the answer is `49`.
