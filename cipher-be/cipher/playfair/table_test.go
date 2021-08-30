package playfair_test

import (
	"testing"

	"github.com/mkamadeus/cipher/cipher/playfair"
)

func TestGenerateTable(t *testing.T) {
	key := "JALANGANESHASEPULUH"
	table := playfair.GenerateTable(key)
	result := [5][5]rune{{'A', 'L', 'N', 'G', 'E'}, {'S', 'H', 'P', 'U', 'B'}, {'C', 'D', 'F', 'I', 'K'}, {'M', 'O', 'Q', 'R', 'T'}, {'V', 'W', 'X', 'Y', 'Z'}}

	for i, r := range result {
		for j, value := range r {
			if table.Table[i][j] != value {
				t.Fatalf("Table not equal")
			}
		}
	}
}

func TestGenerateTableNormalize(t *testing.T) {
	key := "J ALAN GAN.E,SHA SEPU LUH"
	table := playfair.GenerateTable(key)
	result := [5][5]rune{{'A', 'L', 'N', 'G', 'E'}, {'S', 'H', 'P', 'U', 'B'}, {'C', 'D', 'F', 'I', 'K'}, {'M', 'O', 'Q', 'R', 'T'}, {'V', 'W', 'X', 'Y', 'Z'}}

	for i, r := range result {
		for j, value := range r {
			if table.Table[i][j] != value {
				t.Fatalf("Table not equal")
			}
		}
	}
}
