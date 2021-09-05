package standard

import (
	"github.com/mkamadeus/cipher/common/stringutils"
)

func Encrypt(plain string, key string) string {
	result := []rune{}
	i := 0
	key = stringutils.Normalize(key)
	plain = stringutils.Normalize(plain)

	for _, char := range plain {
		keyEvaluated := key[i%len(key)]
		keyBase := stringutils.GetCharBase(rune(keyEvaluated))
		charBase := stringutils.GetCharBase(char)
		y := int(keyEvaluated) - keyBase
		x := int(char) - charBase
		toBeAppended := rune(((x + y) % 26) + charBase)
		result = append(result, toBeAppended)
		i++
	}

	return string(result)
}
