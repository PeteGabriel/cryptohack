package xor

/**
 We can XOR integers by first converting the integer from decimal to binary.
 We can XOR strings by first converting each character to the integer representing
 the Unicode character.

 Xor the given string with a number. Returns the result after xor in string.
 */
func xorString(str string, pred uint8) string {
	inAscii := []byte(str) //get the ascii of each char
	xorResult := make([]byte, 0)

	for _, val := range inAscii {
		xorResult = append(xorResult, val ^ pred)
	}

	return string(xorResult)
}
