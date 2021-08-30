package playfair

import (
	"log"
	"sort"
	"strings"
)

type Table struct {
	Table [5][5]rune
}

func GenerateTable(key string) *Table {
	// normalize key
	key = strings.ToUpper(key)
	key = strings.ReplaceAll(key, " ", "")
	key = strings.ReplaceAll(key, ",", "")
	key = strings.ReplaceAll(key, ".", "")
	key = strings.ReplaceAll(key, ":", "")
	key = strings.ReplaceAll(key, ";", "")

	log.Println(key)

	// prepare map
	exists := map[rune]bool{}
	for c := 'A'; c <= 'Z'; c++ {
		exists[c] = false
	}
	delete(exists, 'J')
	cell := 0

	// put words from key
	var result [5][5]rune
	for _, c := range key {
		if c != 'J' && !exists[c] {
			exists[c] = true
			result[cell/5][cell%5] = c
			cell++
		}
	}

	// put rest not in playfair array
	keys := make([]rune, 0, len(exists))
	for k := range exists {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	for _, k := range keys {
		if !exists[k] {
			exists[k] = true
			result[cell/5][cell%5] = k
			cell++
		}
	}

	return &Table{
		Table: result,
	}
}

func (t Table) IsSameRow(a, b rune) bool {
	for _, row := range t.Table {
		for i := 0; i < len(row); i++ {
			for j := i + 1; j < len(row); j++ {
				if row[i] == row[j] {
					return true
				}
			}
		}
	}

	return false
}

func (t Table) IsSameColumn(a, b rune) bool {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			for k := j + 1; k < 5; k++ {
				if t.Table[j][i] == t.Table[j][i] {
					return true
				}
			}
		}
	}

	return false
}
