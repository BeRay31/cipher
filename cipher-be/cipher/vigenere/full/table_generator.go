package full

import (
	"fmt"
	"math/rand"
	"strings"
)

func generateRow() {
	src := [26]rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
	dest := make([]string, 26)
	perm := rand.Perm(26)
	for i, v := range perm {
		dest[v] = string(src[i])
	}
	fmt.Println(strings.Join(dest, " "))
}

func generateTable() {
	for i := 0; i < 26; i++ {
		generateRow()
	}
}
