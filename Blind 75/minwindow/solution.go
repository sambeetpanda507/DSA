package main

import (
	"fmt"
	"math"
)

func minWindow(s string, t string) string {
	left := 0
	right := 0
	n := len(t)
	tMap := make(map[rune]bool)
	min := math.MaxInt
	for _, r := range t {
		tMap[r] = true
	}

	sMap := make(map[rune]int)
	m := 0
	for right < len(s) {
		var curr rune = rune(s[right])
		isPresent := tMap[curr]
		if isPresent {
			m++
		}

		sMap[curr]++
		for m == n {
			windowSize := right - left + 1
			if windowSize < min {
				min = windowSize
			}

			lChar := rune(s[left])
			isPresent = tMap[lChar]
			if isPresent {
				m--
			}

			sMap[lChar]--
			left++
		}

		right++
	}

	return s[left : left+min]
}

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"
	fmt.Println(minWindow(s, t))
}
