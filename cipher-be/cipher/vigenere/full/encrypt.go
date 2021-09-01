package full

import (
	"strings"

	"github.com/mkamadeus/cipher/common/stringutils"
)

func Encrypt(plain string, key string) string {
	fullVigenereTable := LoadTable()

	trimmedPlainLen := len(strings.ReplaceAll(plain, " ", ""))
	key = strings.ReplaceAll(key, " ", "")

	var keyUsed string = key
	if len(key) > trimmedPlainLen {
		keyUsed = string(key[:trimmedPlainLen])
	}

	result := []rune{}

	i := 0
	for _, char := range plain {
		keyEvaluated := keyUsed[i%len(keyUsed)]
		keyBase := stringutils.GetCharBase(rune(keyEvaluated))
		charBase := stringutils.GetCharBase(char)
		y := (int(keyEvaluated) - keyBase)
		x := (int(char) - charBase)
		// Ignore non alphabet
		if charBase != -1 {
			toBeAppended := fullVigenereTable[y][x]
			result = append(result, toBeAppended)
			i++
		}

	}

	return string(result)
}
