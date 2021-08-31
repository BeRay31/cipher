package standard

func getCharBase(char rune) int {
	intRep := int(char)
	if intRep >= 65 && intRep <= 90 { // Uppercase
		return 65
	} else if intRep >= 97 && intRep <= 122 {
		return 97
	} else {
		return -1
	}
}

func Encrypt(plain string, key string) string {
	result := []rune{}
	i := 0
	for _, char := range plain {
		keyEvaluated := key[i%len(key)]
		keyBase := getCharBase(rune(keyEvaluated))
		charBase := getCharBase(char)
		var toBeAppended rune
		// Ignore non alphabet
		if charBase == -1 {
			toBeAppended = char
		} else {
			toBeAppended = rune((((int(char) - charBase) + (int(keyEvaluated) - keyBase)) % 26) + charBase)
		}
		result = append(result, toBeAppended)

		if char != ' ' {
			i++
		}
	}

	return string(result)
}
