package auto

import (
	"github.com/mkamadeus/cipher/common/stringutils"
)

func Decrypt(cipher string, key string) string {
	key = stringutils.Normalize(key)
	cipher = stringutils.Normalize(cipher)

	result := []rune{}
	initialKeyLen := len(key)
	for i, char := range cipher {
		var keyEvaluated byte
		if i >= len(key) {
			keyEvaluated = byte(result[i-initialKeyLen])
		} else {
			keyEvaluated = key[i]
		}
		keyBase := stringutils.GetCharBase(rune(keyEvaluated))
		charBase := stringutils.GetCharBase(char)
		y := int(keyEvaluated) - keyBase
		x := int(char) - charBase
		toBeAppended := rune((((x - y) + 26) % 26) + charBase)
		result = append(result, toBeAppended)
	}

	return string(result)
}
