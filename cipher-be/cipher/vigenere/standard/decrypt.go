package standard

import (
	"github.com/mkamadeus/cipher/common/stringutils"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func Decrypt(cipher string, key string) string {
	result := []rune{}
	i := 0
	key = stringutils.Normalize(key)
	cipher = stringutils.Normalize(cipher)

	for _, char := range cipher {
		keyEvaluated := key[i%len(key)]
		keyBase := stringutils.GetCharBase(rune(keyEvaluated))
		charBase := stringutils.GetCharBase(char)
		y := int(keyEvaluated) - keyBase
		x := int(char) - charBase
		toBeAppended := rune((((x - y) + 26) % 26) + charBase)
		result = append(result, toBeAppended)
		i++
	}

	return string(result)
}
