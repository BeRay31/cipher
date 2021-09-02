package full

import (
	"strings"

	"github.com/mkamadeus/cipher/common/stringutils"
)

func Encrypt(plain string, key string) string {
	trimmedPlainLen := len(strings.ReplaceAll(plain, " ", ""))
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
