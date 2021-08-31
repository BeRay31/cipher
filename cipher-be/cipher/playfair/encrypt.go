package playfair

import "strings"

func Encrypt(plain, key string) (string, error) {
	plain = strings.ToUpper(plain)
	plain = strings.ReplaceAll(plain, "J", "I")
	table := GenerateTable(key)

	result := []rune{}
	for i := 0; i < len(plain); {
		var a, b rune
		if i == len(plain)-1 {
			a = rune(plain[i])
			b = 'X'
		} else {
			a, b = rune(plain[i]), rune(plain[i+1])
		}

		// check if bigram is a pair of same char
		if a == b {
			b = 'X'
			i++
		} else {
			i += 2
		}

		pos1, err := table.Find(a)
		if err != nil {
			return "", err
		}
		pos2, err := table.Find(b)
		if err != nil {
			return "", err
		}

		if pos1/5 == pos2/5 {
			// same row
			newPos1, newPos2 := table.ShiftHorizontal(pos1, pos2)
			encryptedBigram := []rune{rune(table[newPos1]), rune(table[newPos2])}
			result = append(result, encryptedBigram...)
		} else if pos1%5 == pos2%5 {
			// same col
			newPos1, newPos2 := table.ShiftVertical(pos1, pos2)
			encryptedBigram := []rune{rune(table[newPos1]), rune(table[newPos2])}
			result = append(result, encryptedBigram...)
		} else {
			// different row & col
			newPos1, newPos2 := table.ShiftCycle(pos1, pos2)
			encryptedBigram := []rune{rune(table[newPos1]), rune(table[newPos2])}
			result = append(result, encryptedBigram...)
		}
	}

	return string(result), nil
}
