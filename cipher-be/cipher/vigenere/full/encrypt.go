package full

import (
	"github.com/mkamadeus/cipher/common/stringutils"
)

func Encrypt(plain string, key string) string {
	fullVigenereTable := FullVigenereTable

	plain = stringutils.Normalize(plain)
	key = stringutils.Normalize(key)

	result := []rune{}

	for i, char := range plain {
		keyEvaluated := key[i%len(key)]
		keyBase := stringutils.GetCharBase(rune(keyEvaluated))
		charBase := stringutils.GetCharBase(char)
		y := (int(keyEvaluated) - keyBase)
		x := (int(char) - charBase)
		toBeAppended := fullVigenereTable[y][x]
		result = append(result, toBeAppended)
	}

	return string(result)
}
