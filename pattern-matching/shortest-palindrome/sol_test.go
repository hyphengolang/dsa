package shortestpalindrome

import (
	"testing"
	"unicode/utf8"
)

func TestAlgo(t *testing.T) {
	n := algo("123432")
	_ = n
}

func solution(input, output string) {
	//read from input and output
}

// this is deemed to be the most efficient-way to reverse a string in Go
func rev(s string) string {
	size := len(s)
	buf := make([]byte, size)
	for start := 0; start < size; {
		r, n := utf8.DecodeRuneInString(s[start:])
		start += n
		utf8.EncodeRune(buf[size-start:], r)
	}
	return string(buf)
}

func algo(s string) (lp string) {
	for i, j, rs := 0, len(s), rev(s); i < j; i, j = i+1, j-1 {
		if a, b := j-1, i+1; s[a:] == rs[:b] {
			lp = rs[b:]
		}
	}
	return s + lp
}
