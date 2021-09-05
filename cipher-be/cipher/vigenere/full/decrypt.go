package full

import (
	"github.com/mkamadeus/cipher/common/stringutils"
)

func lookUpY(table [26][26]rune, x int, target rune) int {
	for i, letter := range table[x] {
		if letter == target {
			return i
		}
	}
	return -1
}

func Decrypt(cipher string, key string) string {
	fullVigenereTable := FullVigenereTable

	cipher = stringutils.Normalize(cipher)
	key = stringutils.Normalize(key)

	result := []rune{}

	for i, char := range cipher {
		keyEvaluated := key[i%len(key)]
		keyBase := stringutils.GetCharBase(rune(keyEvaluated))
		x := (int(keyEvaluated) - keyBase)
		toBeAppended := rune(lookUpY(fullVigenereTable, x, char) + 65)
		result = append(result, toBeAppended)
	}

	return string(result)
}
