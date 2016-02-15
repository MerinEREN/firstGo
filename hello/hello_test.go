package main

import "testing"

func TestFactorial(t *testing.T) {
	cases := []struct {
		in, want int
	}{
		{5, 120},
		{6, 720},
	}
	for _, c := range cases {
		got := Factorial(c.in)
		if c.want != got {
			t.Errorf("Factorial(%q) = %q, want %q", c.in, got,
				c.want)
		}
	}
}
