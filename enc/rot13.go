package enc

/*
Rot13 cypher.
Call it like:
    input := "Do you have any cat pictures?"
    mapped := strings.Map(rot13, input)
 */
func Rot13(r rune) rune {
	if r >= 'a' && r <= 'z' {
		// Rotate lowercase letters 13 places.
		if r >= 'm' {
			return r - 13
		} else {
			return r + 13
		}
	} else if r >= 'A' && r <= 'Z' {
		// Rotate uppercase letters 13 places.
		if r >= 'M' {
			return r - 13
		} else {
			return r + 13
		}
	}
	return r
}
