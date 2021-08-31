package affine_test

import (
	"testing"

	"github.com/mkamadeus/cipher/cipher/affine"
)

func TestEncrypt(t *testing.T) {
	plain := "kripto"
	expected := "CZOLNE"

	encrypted := affine.Encrypt(plain, 7, 10)

	if encrypted != expected {
		t.Fatalf("affine encryption failed, expected %v, found %v", expected, encrypted)
	}
}
