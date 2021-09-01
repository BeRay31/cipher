package stringutils

func GetCharBase(char rune) int {
	intRep := int(char)
	if intRep >= 65 && intRep <= 90 { // Uppercase
		return 65
	} else if intRep >= 97 && intRep <= 122 {
		return 97
	} else {
		return -1
	}
}
