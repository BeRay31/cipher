package affine_test

import (
	"testing"

	"github.com/mkamadeus/cipher/cipher/affine"
)

func TestDecrypt(t *testing.T) {
	cipher := "CZOLNE"
	expected := "KRIPTO"

	encrypted := affine.Decrypt(cipher, 7, 10)

	if encrypted != expected {
		t.Fatalf("affine encryption failed, expected %v, found %v", expected, encrypted)
	}
}
