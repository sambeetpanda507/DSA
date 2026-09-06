package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	left := 0
	right := 0
	maxlen := 0
	set := make(map[byte]bool)
	for right < len(s) {
		curr := s[right]
		if _, ok := set[curr]; ok {
			setlen := len(set)
			if setlen > maxlen {
				maxlen = setlen
			}

			delete(set, s[left])
			left++
		} else {
			set[curr] = true
			right++
		}
	}

	setlen := len(set)
	if setlen > maxlen {
		maxlen = setlen
	}

	return maxlen
}

func main() {
	// s := "abcabcbb"
	s := "s"
	ans := lengthOfLongestSubstring(s)
	fmt.Println(ans)
}
