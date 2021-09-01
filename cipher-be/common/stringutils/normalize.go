package stringutils

import "strings"

// Normalize normalizes a string to 26 alphabets only, all uppercase [A-Z]
func Normalize(s string) string {
	// set uppercase
	normalizedChar := []rune{}
	for _, char := range s {
		charBase := GetCharBase(char)
		if charBase != -1 {
			normalizedChar = append(normalizedChar, char)
		}
	}
	return strings.ToUpper(string(normalizedChar))
}
