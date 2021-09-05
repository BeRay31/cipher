package extended

import "github.com/mkamadeus/cipher/cipher/hill"

func Encrypt(plain []byte, key []byte) []byte {
	result := []byte{}
	for i, char := range plain {
		result = append(result, byte(hill.CorrectModulus(int(char+key[i%len(key)]), 256)))
	}
	return result
}
