# Minimum Window Substring - Revision Notes

## Problem

Given two strings `s` and `t`, return the smallest substring of `s` that contains every character from `t`, including duplicate characters.

Example:

```txt
s = "ADOBECODEBANC"
t = "ABC"

answer = "BANC"
```

## Core Idea

Use a sliding window.

We expand the right pointer `R` until the current window contains all required characters from `t`.

Once the window is valid, we move the left pointer `L` forward to shrink the window as much as possible while keeping it valid.

## Data Structures

```go
T := make(map[byte]int)
S := make(map[byte]int)
```

`T` stores the required frequency of each character in `t`.

`S` stores the frequency of relevant characters inside the current window.

Example:

```txt
t = "AABC"

T = {
  A: 2,
  B: 1,
  C: 1,
}
```

## Important Variables

```go
count := 0
```

`count` tracks how many unique characters currently meet their required frequency.

It does not count total characters. It counts satisfied character types.

Example:

```txt
t = "AABC"

Need:
A -> 2
B -> 1
C -> 1

If window has:
A -> 2
B -> 1
C -> 0

count = 2
```

The window is valid when:

```go
count == len(T)
```

## Window Expansion

Move `R` from left to right.

For every character in `s`:

1. If it is needed by `t`, add it to `S`.
2. If its frequency now exactly matches the required frequency, increment `count`.

```go
if _, ok := T[s[R]]; ok {
	S[s[R]]++

	if S[s[R]] == T[s[R]] {
		count++
	}
}
```

## Window Shrinking

When all required characters are satisfied:

```go
for count == len(T) {
```

The current window is valid, so:

1. Save it if it is smaller than the previous best.
2. Remove `s[L]` from the window.
3. If removing it makes the window invalid, decrement `count`.
4. Move `L` forward.

```go
if R-L+1 < result {
	result = R - L + 1
	window[0] = L
	window[1] = L + result
}
```

The end index is stored as exclusive:

```go
window[1] = L + result
```

So the final answer can be returned using:

```go
return s[window[0]:window[1]]
```

## Why This Works

The right pointer finds valid windows.

The left pointer removes unnecessary characters.

Together, each character is visited at most twice:

- once by `R`
- once by `L`

So the solution is efficient.

## Complexity

Time complexity:

```txt
O(n + m)
```

Where:

- `n = len(s)`
- `m = len(t)`

Space complexity:

```txt
O(k)
```

Where `k` is the number of distinct characters in `t`.

## Common Mistakes

### 1. Counting total matched characters instead of satisfied character types

This solution uses `count` for satisfied unique characters, not total matches.

That is why validity is checked with:

```go
count == len(T)
```

### 2. Forgetting duplicates in `t`

For `t = "aa"`, one `a` is not enough.

The frequency map handles this correctly.

### 3. Shrinking the window too early

Only shrink when the window is already valid.

### 4. Off-by-one errors

`R-L+1` is the current window length.

The result slice uses an exclusive end index:

```go
s[start:end]
```

So `end = start + length`.

## Mental Model

Think of the window as a flexible frame over `s`.

```txt
s = A D O B E C O D E B A N C
    L           R
```

Move `R` to collect enough required characters.

Then move `L` to throw away anything unnecessary.

Keep the smallest valid frame seen so far.
