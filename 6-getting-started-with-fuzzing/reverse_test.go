package main

import (
	"testing"
	"unicode/utf8"
)

func FuzzReverse(f *testing.F) {
	testcases := []string{"Hello, world", " ", "!12345"}

	for _, tc := range testcases {
		f.Add(tc) // use f.Add to provide a seed corpus
	}

	f.Fuzz(func(t *testing.T, origin string) {
		rev := Reverse(origin)
		doubleRev := Reverse(rev)

		t.Logf("Number of runes: orgin=%d, rev=%d, doubleRev=%d", utf8.RuneCountInString(origin), utf8.RuneCountInString(rev), utf8.RuneCountInString(doubleRev))

		if origin != doubleRev {
			t.Errorf("Before: %q, after: %q", origin, doubleRev)
		}

		if utf8.ValidString(origin) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}

/*
func TestReverse(t *testing.T) {
	testcases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{" ", " "},
		{"!12345", "54321!"},
	}

	for _, tc := range testcases {
		rev := Reverse(tc.in)
		if rev != tc.want {
			t.Errorf("Reverse: %q, want %q", rev, tc.want)
		}
	}
}
*/
