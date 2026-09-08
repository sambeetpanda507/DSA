package main

import "fmt"

func characterReplacement(s string, k int) int {
	freq := [26]int{}
	left := 0
	right := 0
	ans := 0
	for right < len(s) {
		curr := s[right]

		// Update the freq
		freq[curr-65]++

		// Is valid window
		maxFreq := 0
		for _, f := range freq {
			if f > maxFreq {
				maxFreq = f
			}
		}

		windowSize := right - left + 1
		if windowSize-maxFreq <= k {
			ans = windowSize
			right++
		} else {
			freq[s[left]-65]--
			left++
			right++
		}
	}

	return ans
}

func main() {
	s := "AAAB"
	k := 0
	ans := characterReplacement(s, k)
	fmt.Println(ans)
}
