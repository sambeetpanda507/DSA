# Group Anagrams Notes

## Overview

The `groupAnagrams` function groups words that are anagrams of each other.

Input example:

```go
[]string{"eat", "tea", "tan", "ate", "nat", "bat"}
```

Expected grouping:

```go
[["eat", "tea", "ate"], ["tan", "nat"], ["bat"]]
```

The order of groups can vary because Go maps do not preserve insertion order.

## Core Idea

Instead of sorting every string, the solution builds a frequency signature for each word.

- Create a fixed-size array of 26 integers.
- Count how many times each letter `a` to `z` appears.
- Convert that frequency array into a compact string key.
- Use the key in a map to collect all matching anagrams.

Words with the same character counts produce the same key, so they belong in the same group.

## How `main.go` Works

For each string:

1. Initialize `hash := [26]int{}`.
2. Loop through each character and increment its count.
3. Build a key like `1a1e1t` using `strings.Builder`.
4. Store the original string in `groups[key]`.

After all strings are processed:

- Iterate through the map values.
- Append each group to `result`.
- Return `result`.

## Why This Works

Anagrams contain the same letters with the same frequencies.

Examples:

- `"eat"` -> `1a1e1t`
- `"tea"` -> `1a1e1t`
- `"ate"` -> `1a1e1t`

Since all three generate the same key, they are grouped together.

## Time Complexity

Let:

- `n` = number of strings
- `k` = average length of a string

Complexity:

- Building counts for each string: `O(k)`
- Building the key: `O(26)`, which is constant
- Total: `O(n * k)`

## Space Complexity

- Map storage for grouped strings: `O(n * k)` in total output storage
- Temporary frequency array per word: `O(1)`

## Notes About This Approach

- This solution assumes all characters are lowercase English letters.
- It is more efficient than sorting each string, which would usually cost `O(k log k)` per word.
- Result ordering is not guaranteed because map iteration order in Go is not deterministic.
