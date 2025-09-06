package utils

import "testing"

func TestCountChars(t *testing.T) {
	cases := []struct {
		s string
		n int
	}{
		{"", 0},
		{"abc", 3},
		{"Привет", 12},
		{"Hello 👋", 10},
	}
	for _, c := range cases {
		if got := CountChars(c.s); got != c.n {
			t.Fatalf("%q: want %d, got %d", c.s, c.n, got)
		}
	}
}
