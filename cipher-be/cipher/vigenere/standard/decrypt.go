package standard

func Decrypt(cipher string, key string) string {
	result := []rune{}
	i := 0
	for _, char := range cipher {
		keyEvaluated := key[i%len(key)]
		keyBase := getCharBase(rune(keyEvaluated))
		charBase := getCharBase(char)
		var toBeAppended rune
		// Ignore non alphabet
		if charBase == -1 {
			toBeAppended = char
		} else {
			toBeAppended = rune(((((int(char) - charBase) - (int(keyEvaluated) - keyBase)) + 26) % 26) + charBase)
		}
		result = append(result, toBeAppended)
		// result = append(result, rune((((int(char)-65)-(int(key[i%len(key)])-65)+26)%26)+65))
		if char != ' ' {
			i++
		}
	}

	return string(result)
}
