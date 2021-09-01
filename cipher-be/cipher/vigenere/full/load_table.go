package full

import (
	"os"
	"path/filepath"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func LoadTable() [26][26]rune {
	dataPath, _ := filepath.Abs("cipher/vigenere/full/full_vigenere_table.txt")
	data, err := os.ReadFile(dataPath)
	check(err)

	var vigenereTable [26][26]rune

	splittedStrings := strings.Split(string(data), "\n")
	for i, sentence := range splittedStrings {
		splittedSentence := strings.Split(sentence, " ")
		for j, letter := range splittedSentence {
			vigenereTable[i][j] = rune(letter[0])
		}
	}

	return vigenereTable
}
