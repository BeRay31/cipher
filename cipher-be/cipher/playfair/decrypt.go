package playfair

import "strings"

func Decrypt(cipher, key string) (string, error) {
	table := GenerateTable(key)

	var result []rune
	for i := 0; i < len(cipher); i += 2 {
		a, b := rune(cipher[i]), rune(cipher[i+1])

		// shifts characters
		pos1, err := table.Find(a)
		if err != nil {
			return "", err
		}
		pos2, err := table.Find(b)
		if err != nil {
			return "", err
		}
		var encryptedBigram []rune
		var newPos1, newPos2 int
		if pos1/5 == pos2/5 {
			// same row
			newPos1, newPos2 = table.ShiftHorizontal(pos1, pos2, -1)
		} else if pos1%5 == pos2%5 {
			// same col
			newPos1, newPos2 = table.ShiftVertical(pos1, pos2, -1)
		} else {
			// different row & col
			newPos1, newPos2 = table.ShiftCycle(pos1, pos2)
		}
		encryptedBigram = []rune{rune(table[newPos1]), rune(table[newPos2])}
		result = append(result, encryptedBigram...)
	}

	decrypted := strings.ReplaceAll(string(result), "X", "")
	return decrypted, nil
}
