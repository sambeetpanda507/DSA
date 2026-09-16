package main

import "testing"

func TestSearchRotated(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{name: "Example 1", nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 0, want: 4},
		{name: "Example 2", nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 3, want: -1},
		{name: "Example 3", nums: []int{1}, target: 0, want: -1},
		{name: "Example 4", nums: []int{1, 3}, target: 3, want: 1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out := search(test.nums, test.target)
			if out != test.want {
				t.Errorf("search(%d, %d) = %d", test.nums, test.target, out)
			}
		})
	}
}
