package hill

import (
	"math"

	"github.com/mkamadeus/cipher/common/stringutils"
)

func isQuadratic(num int) bool {
	sqrt := int(math.Sqrt(float64(num)))
	return sqrt*sqrt == num
}

func multiply(x, y [][]int) [][]int {

	out := make([][]int, len(x))
	for i := 0; i < len(x); i++ {
		out[i] = make([]int, len(y[0]))
		for j := 0; j < len(y[0]); j++ {
			for k := 0; k < len(y); k++ {
				out[i][j] += x[i][k] * y[k][j]
			}
		}
	}
	return out
}

func transpose(x [][]int) [][]int {
	out := make([][]int, len(x[0]))
	for i := 0; i < len(x); i += 1 {
		for j := 0; j < len(x[0]); j += 1 {
			out[j] = append(out[j], x[i][j])
		}
	}
	return out
}

func BuildKeyMatrix(key string) [][]int {
	keyDimension := int(math.Sqrt(float64(len(key))))
	keyMatrix := make([][]int, keyDimension)
	for i := 0; i < keyDimension; i++ {
		keyMatrix[i] = make([]int, keyDimension)
		for j := 0; j < keyDimension; j++ {
			keyMatrix[i][j] = int(key[i*keyDimension+j]) - stringutils.GetCharBase(rune(key[i*keyDimension+j]))
		}
	}
	return keyMatrix
}

func BuildPlainString(plain string, dim int) ([]string, int) {
	arrayLen := len(plain) / dim
	if arrayLen*dim != len(plain) {
		arrayLen++
	}
	strArr := make([]string, arrayLen)
	count := 0
	offset := 0
	for i := 0; i < len(plain); i += dim {
		if i+dim > len(plain) {
			strArr[count] = string(plain[i:len(plain)])
			offset = i + dim - len(plain) + 1
		} else {
			strArr[count] = string(plain[i : i+dim])
		}
		count++
	}
	return strArr, offset
}

func BuildSegmentedPlainArr(segmentedPlain []string, offset int) [][]int {
	segmentedMatrix := make([][]int, len(segmentedPlain))
	dim := len(segmentedPlain[0])
	for i, segmented := range segmentedPlain {
		segmentedMatrix[i] = make([]int, dim)
		for j, letter := range segmented {
			segmentedMatrix[i][j] = int(letter) - stringutils.GetCharBase(letter)
			if i == len(segmentedPlain)-1 && j == len(segmented)-1 {
				for o := 0; o < offset; o++ {
					segmentedMatrix[i][j+o] = 0
				}
			}
		}
	}
	return segmentedMatrix
}

func MatToFloatArr(mat [][]int) []float64 {
	floatArray := make([]float64, len(mat)*len(mat[0]))
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			floatArray[i*len(mat)+j] = float64(mat[i][j])
		}
	}
	return floatArray
}
