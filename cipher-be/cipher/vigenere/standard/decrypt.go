package standard

import "strings"

func Decrypt(cipher string, key string) string {
	cipher = strings.ToUpper(cipher)
	key = strings.ToUpper(key)

	result := []rune{}
	for i, char := range cipher {
		result = append(result, rune((((int(char)-65)-(int(key[i%len(key)])-65)+26)%26)+65))
	}

	return string(result)
}
