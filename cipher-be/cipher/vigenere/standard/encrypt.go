package standard

import (
	"strings"

	"github.com/mkamadeus/cipher/common/stringutils"
)

func Encrypt(plain string, key string) string {
	result := []rune{}
	i := 0
	key = strings.ReplaceAll(key, " ", "")
	for _, char := range plain {
		keyEvaluated := key[i%len(key)]
		keyBase := stringutils.GetCharBase(rune(keyEvaluated))
		charBase := stringutils.GetCharBase(char)
		y := int(keyEvaluated) - keyBase
		x := int(char) - charBase
		// Ignore non alphabet and ignore space
		if charBase != -1 {
			toBeAppended := rune(((x + y) % 26) + charBase)
			result = append(result, toBeAppended)
			i++
		}
	}

	return string(result)
}
