# Find Minimum in Rotated Sorted Array

## Problem

Given a sorted array that may have been rotated, return the minimum value.

Example:

```go
[]int{4, 5, 6, 7, 0, 1, 2}
```

The answer is:

```go
0
```

## Key Idea

A rotated sorted array is made of two sorted parts:

```text
4 5 6 7 | 0 1 2
high part | low part
```

The minimum is the first value in the low part.

Binary search works because we can look at the current window and decide which half still may contain the minimum.

## Current Approach in solution.go

The code keeps three things:

```go
left := 0
right := len(nums) - 1
ans := math.MaxInt
```

During each loop:

```go
mid := left + (right-left)/2
```

This is the safe midpoint formula. It avoids overflow in languages where `left + right` can become too large.

Then it updates the best minimum seen so far:

```go
if nums[mid] < ans {
	ans = nums[mid]
}
```

Important: compare `nums[mid]`, not `mid`.

`mid` is an index. `nums[mid]` is the actual value.

## Important Bug You Fixed

This kind of loop:

```go
for left <= right
```

must shrink the search space every time.

So when moving the right pointer, use:

```go
right = mid - 1
```

not:

```go
right = mid
```

Why?

If `left == right`, then `mid == left == right`.

So this would get stuck:

```go
right = mid
```

Because `right` does not move.

Example:

```text
left = 0
right = 0
mid = 0
right = mid // still 0
```

That creates an infinite loop.

## Loop Rule to Remember

There are two common binary search styles.

Style 1:

```go
for left <= right {
	// exclude mid after checking it
	left = mid + 1
	right = mid - 1
}
```

Use this when `mid` has already been processed.

Style 2:

```go
for left < right {
	// keep mid when it may still be the answer
	right = mid
}
```

Use this when `mid` might still be the answer.

Do not mix these accidentally. The loop condition and pointer updates must agree.

## Common Mistakes

1. Returning the index instead of the value.

Wrong:

```go
ans = mid
```

Correct:

```go
ans = nums[mid]
```

2. Using a midpoint formula with misplaced parentheses.

Wrong:

```go
mid := (left + (right-left)) / 2
```

This simplifies to:

```go
mid := right / 2
```

Correct:

```go
mid := left + (right-left)/2
```

3. Using `right = mid` inside a `left <= right` loop.

That may not shrink the window and can cause an infinite loop.

## Simpler Standard Version

A shorter version does not need `ans`:

```go
func findMin(nums []int) int {
	left := 0
	right := len(nums) - 1

	for left < right {
		mid := left + (right-left)/2

		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return nums[left]
}
```

Mental model:

If `nums[mid] > nums[right]`, the minimum is to the right of `mid`.

Otherwise, `mid` may be the minimum, so keep it by doing `right = mid`.

## Revision Summary

- Use `nums[mid]` when tracking the minimum value.
- Use `left + (right-left)/2` for midpoint.
- With `left <= right`, move past `mid` using `mid + 1` or `mid - 1`.
- With `left < right`, it is okay to keep `mid` using `right = mid`.
- The main danger in binary search is not shrinking the window.
