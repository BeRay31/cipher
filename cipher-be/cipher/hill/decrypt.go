package hill

import (
	"log"

	"github.com/mkamadeus/cipher/common/stringutils"
	"gonum.org/v1/gonum/mat"
)

func Decrypt(cipher string, key string) string {
	cipher = stringutils.Normalize(cipher)
	key = stringutils.Normalize(key)

	keyMatrix := BuildKeyMatrix(key)
	keyDim := len(keyMatrix[0])
	keyDenseMatrix := mat.NewDense(keyDim, keyDim, MatToFloatArr(keyMatrix))

	var keyInv mat.Dense
	err := keyInv.Inverse(keyDenseMatrix)
	if err != nil {
		log.Fatalf("A is not invertible: %v", err)
	}
	segmentedWord := BuildSegmentedPlainArr(BuildPlainString(cipher, keyDim))
	transposedSegment := transpose(segmentedWord)
	denseTransposedSegment := mat.NewDense(len(transposedSegment), len(transposedSegment[0]), MatToFloatArr(transposedSegment))
	var resMatrix mat.Dense
	resMatrix.Mul(keyDenseMatrix, denseTransposedSegment)
	r, c := resMatrix.Dims()
	runeResMatrix := make([]rune, r*c)
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			runeResMatrix[i*c+j] = rune(resMatrix.At(i, j))
		}
	}

	return string(runeResMatrix)
}
