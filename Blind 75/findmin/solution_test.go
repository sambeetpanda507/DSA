package main

import "testing"

func TestFindMin(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{name: "Example 1", input: []int{3, 4, 5, 1, 2}, want: 1},
		{name: "Example 2", input: []int{4, 5, 6, 7, 0, 1, 2}, want: 0},
		{name: "Example 3", input: []int{11, 13, 15, 17}, want: 11},
		{name: "Example 4", input: []int{2, 1}, want: 1},
		{name: "Example 4", input: []int{1, 2}, want: 1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out := findMin(test.input)
			if out != test.want {
				t.Errorf("findMin(%d) = %d", test.input, out)
			}
		})
	}
}
