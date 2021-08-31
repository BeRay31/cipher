package standard

import "strings"

func Encrypt(plain string, key string) string {
	plain = strings.ToUpper(plain)
	key = strings.ToUpper(key)

	result := []rune{}
	for i, char := range plain {
		result = append(result, rune((((int(char)-65)+(int(key[i%len(key)])-65))%26)+65))
	}

	return string(result)
}
