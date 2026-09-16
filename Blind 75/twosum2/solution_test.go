package main

import "testing"

func TestTwoSum2(t *testing.T) {
	tests := []struct {
		name    string
		numbers []int
		target  int
		want    []int
	}{
		{name: "Example 1", numbers: []int{2, 7, 11, 15}, target: 9, want: []int{1, 2}},
		{name: "Example 2", numbers: []int{2, 3, 4}, target: 6, want: []int{1, 3}},
		{name: "Example 3", numbers: []int{-1, 0}, target: -1, want: []int{1, 2}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out := twoSum(test.numbers, test.target)
			if len(out) != len(test.want) {
				t.Errorf("twoSum(%d, %d) = %d\n", test.numbers, test.target, out)
				return
			}

			for i := range out {
				if out[i] != test.want[i] {
					t.Errorf("twoSum(%d, %d) = %d\n", test.numbers, test.target, out)
					return
				}
			}
		})
	}
}
