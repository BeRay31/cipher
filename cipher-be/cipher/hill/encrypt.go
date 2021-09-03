package hill

import (
	"github.com/mkamadeus/cipher/common/stringutils"
)

func Encrypt(plain string, key string) string {
	plain = stringutils.Normalize(plain)
	key = stringutils.Normalize(key)

	keyMatrix := BuildKeyMatrix(key)
	keyDim := len(keyMatrix[0])

	segmentedWord := BuildSegmentedPlainArr(BuildPlainString(plain, keyDim))
	resMatrix := transpose(multiply(keyMatrix, transpose(segmentedWord)))
	runeResMatrix := make([]rune, len(resMatrix)*len(resMatrix[0]))
	for i := 0; i < len(resMatrix); i++ {
		for j := 0; j < len(resMatrix[i]); j++ {
			runeResMatrix[i*keyDim+j] = rune(resMatrix[i][j]%26) + 65
		}
	}
	return string(runeResMatrix)
}
