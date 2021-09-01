package auto

import (
	"strings"

	"github.com/mkamadeus/cipher/common/stringutils"
)

func Decrypt(cipher string, key string) string {
	result := []rune{}
	i := 0
	key = strings.ReplaceAll(key, " ", "")
	initialKeyLen := len(key)
	for _, char := range cipher {
		var keyEvaluated byte
		if i >= len(key) {
			keyEvaluated = byte(result[i-initialKeyLen])
		} else {
			keyEvaluated = key[i]
		}
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
