package standard_test

import (
	"testing"

	"github.com/mkamadeus/cipher/cipher/vigenere/standard"
)

func TestDecrypt(t *testing.T) {
	cipher := "LVVQHZNGFHRVL"
	key := "SONY"
	expected := "THISPLAINTEXT"

	decrypted := standard.Decrypt(cipher, key)
	if decrypted != expected {
		t.Fatalf("affine encryption failed, expected %v, found %v", expected, decrypted)
	}
}
