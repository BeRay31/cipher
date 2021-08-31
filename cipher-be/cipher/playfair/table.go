package playfair

import (
	"errors"
	"sort"
	"unicode"

	"github.com/mkamadeus/cipher/common/stringutils"
)

type Table string

func GenerateTable(key string) Table {
	// normalize key
	key = stringutils.Normalize(key)

	// prepare map
	exists := map[rune]bool{}
	for c := 'A'; c <= 'Z'; c++ {
		exists[c] = false
	}
	delete(exists, 'J')

	// put words from key
	result := []rune{}
	for _, c := range key {
		if c != 'J' && !exists[c] {
			exists[c] = true
			result = append(result, c)
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
			result = append(result, k)
		}
	}

	return Table(result)
}

func (table Table) Find(c rune) (int, error) {
	c = unicode.ToUpper(c)
	for i, ch := range table {
		if ch == c {
			return i, nil
		}
	}
	return -1, errors.New("character not found")
}

func (table Table) IsSameRow(i, j rune) (bool, error) {
	pos1, err := table.Find(i)
	if err != nil {
		return false, err
	}
	pos2, err := table.Find(j)
	if err != nil {
		return false, err
	}

	return pos1/5 == pos2/5, nil
}

func (table Table) IsSameColumn(i, j rune) (bool, error) {
	pos1, err := table.Find(i)
	if err != nil {
		return false, err
	}
	pos2, err := table.Find(j)
	if err != nil {
		return false, err
	}

	return pos1/5 == pos2/5, nil
}

func (table Table) ShiftHorizontal(i, j, offset int) (int, int) {
	ri, ci := i/5, i%5
	rj, cj := j/5, j%5

	return ri*5 + (ci+offset)%5, rj*5 + (cj+offset)%5
}

func (table Table) ShiftVertical(i, j, offset int) (int, int) {
	ri, ci := i/5, i%5
	rj, cj := j/5, j%5

	return ((ri+offset)%5)*5 + ci, ((rj+offset)%5)*5 + cj
}

func (table Table) ShiftCycle(i, j int) (int, int) {
	ri, ci := i/5, i%5
	rj, cj := j/5, j%5

	return ri*5 + cj, rj*5 + ci
}
