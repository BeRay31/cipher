package affine

import (
	"github.com/mkamadeus/cipher/common/stringutils"
	"github.com/pkg/errors"
)

var relativePrimes = []int{1, 3, 5, 7, 9, 11, 15, 17, 19, 21, 23}

func Encrypt(plain string, m, b int) (string, error) {
	isRelativePrime := false
	for _, rp := range relativePrimes {
		if isRelativePrime {
			break
		}
		if m == rp {
			isRelativePrime = true
		}
	}
	if !isRelativePrime {
		return "", errors.Errorf("variable m is not relative prime to 26, found %d", m)
	}

	normalized := stringutils.Normalize(plain)
	result := []rune{}
	for _, char := range normalized {
		ascii := int(char) - stringutils.GetCharBase(char)
		encrypted := (m*ascii + b) % 26
		result = append(result, rune(encrypted+65))
	}
	return string(result), nil
}
