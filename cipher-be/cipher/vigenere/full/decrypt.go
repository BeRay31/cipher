package full

import (
	"strings"

	"github.com/mkamadeus/cipher/common/stringutils"
)

func Decrypt(cipher string, key string) string {
	trimmedCipherLen := len(strings.ReplaceAll(cipher, " ", ""))
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
		var toBeAppended rune
		// Ignore non alphabet
		if charBase == -1 {
			toBeAppended = char
		} else {
			toBeAppended = rune(((((int(char) - charBase) - (int(keyEvaluated) - keyBase)) + 26) % 26) + charBase)
		}
		result = append(result, toBeAppended)
		if char != ' ' {
			i++
		}
	}

	return string(result)
}
