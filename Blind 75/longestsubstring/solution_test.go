package main

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name string
		input string
		want  int
	}{
		{name: "empty string", input: "", want: 0},
		{name: "single character", input: "a", want: 1},
		{name: "all unique", input: "abcdef", want: 6},
		{name: "all repeated", input: "aaaa", want: 1},
		{name: "repeating prefix", input: "abcabcbb", want: 3},
		{name: "window shrinks repeatedly", input: "pwwkew", want: 3},
		{name: "duplicate after long window", input: "dvdf", want: 3},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := lengthOfLongestSubstring(test.input)
			if got != test.want {
				t.Errorf("lengthOfLongestSubstring(%q) = %d, want %d", test.input, got, test.want)
			}
		})
	}
}
