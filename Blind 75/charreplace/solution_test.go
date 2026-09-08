package main

import "testing"

func TestCharReplace(t *testing.T) {
	tests := []struct {
		name string
		s    string
		k    int
		want int
	}{
		{name: "empty string", s: "", k: 0, want: 0},
		{name: "one letter with k = 0", s: "A", k: 0, want: 1},
		{name: "one letter with k = 1", s: "A", k: 1, want: 1},
		{name: "one letter with k = 2", s: "A", k: 2, want: 1},
		{name: "ABAB", s: "ABAB", k: 2, want: 4},
		{name: "AABABBA", s: "AABABBA", k: 1, want: 4},
		{name: "k is 0", s: "AAAB", k: 0, want: 3},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := characterReplacement(test.s, test.k)
			if got != test.want {
				t.Errorf("characterReplacement(%s, %d) = %d", test.s, test.k, got)
			}
		})
	}
}
