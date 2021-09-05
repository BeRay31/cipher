package auto

import (
	"github.com/mkamadeus/cipher/common/stringutils"
)

func Encrypt(plain string, key string) string {
	key = stringutils.Normalize(key)
	plain = stringutils.Normalize(plain)

	result := []rune{}

	initialKeyLen := len(key)
	for i, char := range plain {
		var keyEvaluated byte
		if i >= len(key) {
			keyEvaluated = byte(plain[i-initialKeyLen])
		} else {
			keyEvaluated = key[i]
		}
		keyBase := stringutils.GetCharBase(rune(keyEvaluated))
		charBase := stringutils.GetCharBase(char)
		y := int(keyEvaluated) - keyBase
		x := int(char) - charBase
		toBeAppended := rune(((x + y) % 26) + charBase)
		result = append(result, toBeAppended)
	}

	return string(result)
}
