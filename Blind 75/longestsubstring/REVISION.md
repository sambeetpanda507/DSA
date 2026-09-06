# Longest Substring Without Repeating Characters

## Goal

Given a string `s`, return the length of its longest **contiguous substring** containing no repeated characters.

Example: `"abcabcbb"` returns `3` because `"abc"` is the longest valid substring.

## Approach: Sliding Window

`solution.go` maintains a window from `left` to `right` and a set containing the bytes currently in that window.

### Window invariant

At every point, the set contains only unique characters from `s[left:right]`.

### Steps

1. Inspect `s[right]`.
2. If it is not in the set, add it and advance `right` to expand the window.
3. If it is already present, the window has a duplicate:
   - record the current window length if it is the largest seen;
   - remove `s[left]` from the set;
   - advance `left`.
4. Keep shrinking one character at a time until the duplicate can be added on a later iteration.
5. Check the final window after the loop, then return the maximum length.

## Example Trace: `"pwwkew"`

| Window | Action | Best |
| --- | --- | --- |
| `p` | Add `p` | 0 |
| `pw` | Add `w` | 0 |
| `pw` + `w` | Duplicate: remove `p` | 2 |
| `w` + `w` | Duplicate: remove `w` | 2 |
| `wke` | Add `w`, `k`, and `e` | 2 |
| `wke` + `w` | Duplicate: remove the first `w`, then add `w` | 3 |
| Final check | Window length is `3` (`"kew"`) | 3 |

## Complexity

- Time: `O(n)` — each byte enters and leaves the set at most once.
- Space: `O(min(n, character set size))` for the window set.

## Important Details

- `right` advances only after adding a new character. On a duplicate, it stays in place while `left` shrinks the window.
- The final maximum check is required because the longest substring may end at the last byte.
- This implementation indexes strings as `byte`, so it is correct for ASCII / single-byte characters. For arbitrary Unicode characters, convert to `[]rune` and use `map[rune]bool`.

## Alternative Optimization

Store each character's latest index in `map[byte]int`. When a duplicate appears, jump `left` directly past its previous occurrence instead of deleting one character per iteration. It is still `O(n)`, but can be more direct to reason about.
