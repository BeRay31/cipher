package affine_test

import (
	"testing"

	"github.com/mkamadeus/cipher/cipher/affine"
)

func TestEncrypt(t *testing.T) {
	plain := "kripto"
	expected := "CZOLNE"

	encrypted, err := affine.Encrypt(plain, 7, 10)

	if err != nil || encrypted != expected {
		t.Fatalf("affine encryption failed, expected %v, found %v", expected, encrypted)
	}
}

func TestNotRelativePrime(t *testing.T) {
	plain := "kripto"

	_, err := affine.Encrypt(plain, 2, 10)

	if err == nil {
		t.Fatalf("affine encryption should fail, m is expected to be relative prime to 26")
	}
}
