package full

import (
	"strings"

	"github.com/mkamadeus/cipher/common/stringutils"
)

func lookUpX(table [26][26]rune, y int, target rune) int {
	for i, letter := range table[y] {
		if letter == target {
			return i
		}
	}
	return -1
}

func Decrypt(cipher string, key string) string {
	fullVigenereTable := FullVigenereTable

	trimmedCipherLen := len(strings.ReplaceAll(cipher, " ", ""))
	key = strings.ReplaceAll(key, " ", "")

	var keyUsed string = key
	if len(key) > trimmedCipherLen {
		keyUsed = string(key[:trimmedCipherLen])
	}

	result := []rune{}

	i := 0
	for _, char := range cipher {
		keyEvaluated := keyUsed[i%len(keyUsed)]
		keyBase := stringutils.GetCharBase(rune(keyEvaluated))
		charBase := stringutils.GetCharBase(char)
		y := (int(keyEvaluated) - keyBase)
		// Ignore non alphabet
		if charBase != -1 {
			toBeAppended := rune(lookUpX(fullVigenereTable, y, char) + 65)
			result = append(result, toBeAppended)
			i++
		}
	}

	return string(result)
}
