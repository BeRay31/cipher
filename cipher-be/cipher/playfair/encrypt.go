package playfair

import "strings"

func Encrypt(plain, key string) string {
	normalized := strings.ReplaceAll(plain, "J", "I")
	table := GenerateTable(key)

	for i := 0; i < len(normalized); i += 2 {
		a, b := normalized[i], normalized[i+1]
		if table.IsSameRow(rune(a), rune(b)) {
			
		}
		else if table.IsSameColumn(rune(a), rune(b)) {

		}
	}
}
