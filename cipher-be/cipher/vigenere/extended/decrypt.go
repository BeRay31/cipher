package extended

func Decrypt(cipher []byte, key []byte) []byte {
	result := []byte{}
	for i, char := range cipher {
		result = append(result, byte((int(char-key[i%len(key)])+256)%256))
	}

	return result
}
