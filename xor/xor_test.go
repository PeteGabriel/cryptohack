package xor

import (
	"github.com/matryer/is"
	"testing"
)

func TestStringToAsci(t *testing.T){
	is := is.New(t)
	input := "label"
	expected := "aloha"
	is.Equal(xorString(input, 13), expected)
}