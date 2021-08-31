package affine

import "github.com/mkamadeus/cipher/common/stringutils"

func Decrypt(cipher string, m, b int) string {
	cipher = stringutils.Normalize(cipher)

	var x int
	for x = 0; (m*x)%26 != 1; x++ {
		if x == 26 {
			return ""
		}
	}

	result := []rune{}
	for _, char := range cipher {
		ascii := int(char) - 65
		decrypted := (x * (ascii - b + 26)) % 26
		result = append(result, rune(decrypted+65))
	}

	return string(result)

}
