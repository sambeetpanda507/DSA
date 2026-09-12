package main

import "testing"

func TestValidParentheses(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{name: "Example 1", input: "()", want: true},
		{name: "Example 2", input: "()[]{}", want: true},
		{name: "Example 3", input: "(]", want: false},
		{name: "Example 4", input: "([])", want: true},
		{name: "Example 5", input: "([)])", want: false},
		{name: "Example 6", input: "(", want: false},
		{name: "Example 7", input: "]", want: false},
		{name: "Example 8", input: "][]", want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := isValid(test.input)
			if got != test.want {
				t.Errorf("isValid(%s) = %v", test.input, test.want)
			}
		})
	}
}
