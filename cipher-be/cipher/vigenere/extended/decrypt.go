package extended

import "github.com/mkamadeus/cipher/cipher/hill"

func Decrypt(cipher []byte, key []byte) []byte {
	result := []byte{}
	for i, char := range cipher {
		result = append(result, byte(hill.CorrectModulus((int(char-key[i%len(key)])), 256)))
	}

	return result
}
