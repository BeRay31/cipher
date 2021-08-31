package stringutils

import "strings"

// Normalize normalizes a string to 26 alphabets only, all uppercase [A-Z]
func Normalize(s string) string {
	// set uppercase
	s = strings.ToUpper(s)

	// remove punctuation
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ReplaceAll(s, ",", "")
	s = strings.ReplaceAll(s, ".", "")
	s = strings.ReplaceAll(s, ":", "")
	s = strings.ReplaceAll(s, ";", "")
	s = strings.ReplaceAll(s, "\"", "")
	s = strings.ReplaceAll(s, "'", "")

	return s
}
