package main

import "testing"

func TestMinWindow(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want string
	}{
		{
			name: "example 1",
			s:    "ADOBECODEBANC",
			t:    "ABC",
			want: "BANC",
		},
		{
			name: "example 2",
			s:    "a",
			t:    "a",
			want: "a",
		},
		{
			name: "example 3",
			s:    "a",
			t:    "aa",
			want: "",
		},
		{
			name: "empty t",
			s:    "aa",
			t:    "",
			want: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := minWindow(test.s, test.t)
			if got != test.want {
				t.Errorf("minWindow(%s, %s) = %s", test.s, test.t, got)
			}
		})
	}
}
