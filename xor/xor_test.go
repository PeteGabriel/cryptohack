package xor

import (
	"testing"
	is2 "github.com/matryer/is"
)

func TestStringToAsci(t *testing.T) {
	is := is2.New(t)
	testCases := []struct {
		input    string
		expected string
		num      int
	}{
		{
			input:    "label",
			expected: "aloha",
			num:      13,
		},
		{
			input:    "peter",
			expected: "yl}l{",
			num:      9,
		},
	}
	for _, tc := range testCases {
		is.Equal(xorString(tc.input, uint8(tc.num)), tc.expected)
	}
}
