package standard

import (
	"strings"

	"github.com/mkamadeus/cipher/common/stringutils"
)

func Decrypt(cipher string, key string) string {
	result := []rune{}
	i := 0
	key = strings.ReplaceAll(key, " ", "")
	for _, char := range cipher {
		keyEvaluated := key[i%len(key)]
		keyBase := stringutils.GetCharBase(rune(keyEvaluated))
		charBase := stringutils.GetCharBase(char)
		var toBeAppended rune
		// Ignore non alphabet and ignore space
		if charBase != -1 {
			toBeAppended = rune(((((int(char) - charBase) - (int(keyEvaluated) - keyBase)) + 26) % 26) + charBase)
			result = append(result, toBeAppended)
			i++
		}
	}

	return string(result)
}
