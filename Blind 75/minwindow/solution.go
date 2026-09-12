package main

import (
	"fmt"
	"math"
)

func minWindow(s string, t string) string {
	if len(t) == 0 {
		return ""
	}

	S := make(map[byte]int)
	T := make(map[byte]int)
	count := 0
	result := math.MaxInt
	window := [2]int{0, 0}
	for i := 0; i < len(t); i++ {
		T[t[i]]++
	}

	L := 0
	R := 0
	for R < len(s) {
		if _, ok := T[s[R]]; ok {
			S[s[R]]++

			if S[s[R]] == T[s[R]] {
				count++
			}
		}

		for count == len(T) {
			if R-L+1 < result {
				result = R - L + 1
				window[0] = L
				window[1] = L + result
			}

			if _, ok := T[s[L]]; ok {
				S[s[L]]--
			}

			if S[s[L]] < T[s[L]] {
				count--
			}

			L++
		}

		R++
	}

	// fmt.Printf("L = %d and R = %d and res = %d", L, R, result)
	return s[window[0]:window[1]]
}

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"
	fmt.Println(minWindow(s, t))
}
