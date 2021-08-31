package standard

func Encrypt(plain []byte, key []byte) []byte {
	result := []byte{}
	for i, char := range plain {
		result = append(result, byte((int(char)+int(key[i%len(key)]))%256))
	}

	return result
}
