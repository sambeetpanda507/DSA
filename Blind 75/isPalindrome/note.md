# Valid Palindrome

`main.py` checks whether a string is a palindrome after ignoring non-alphanumeric
characters and letter case.

## Approach

1. Remove every character that is not a letter or digit.
2. Convert the remaining string to lowercase.
3. Use two pointers:
   - `start` begins at the left side.
   - `end` begins at the right side.
4. Move both pointers inward while characters match.
5. Return `False` immediately when a mismatch is found.
6. If all mirrored characters match, return `True`.

## Complexity

- Time: `O(n)`, where `n` is the length of the input string.
- Space: `O(n)`, because the cleaned string is built before checking.

## Example

```python
s = "race a car"
```

After cleanup:

```text
"raceacar"
```

This is not a palindrome, so the program prints:

```text
False
```
