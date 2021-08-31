package affine

import "github.com/mkamadeus/cipher/common/stringutils"

func Encrypt(plain string, m, b int) string {
	normalized := stringutils.Normalize(plain)
	result := []rune{}
	for _, char := range normalized {
		ascii := int(char) - 65
		encrypted := (m*ascii + b) % 26
		result = append(result, rune(encrypted+65))
	}
	return string(result)
}
