package standard_test

import (
	"testing"

	"github.com/mkamadeus/cipher/cipher/vigenere/standard"
)

func TestDecrypt(t *testing.T) {
	cipher := "LVVQHZNGFHRVL"
	key := "sony"
	expected := "THISPLAINTEXT"

	encrypted := standard.Decrypt(cipher, key)
	if encrypted != expected {
		t.Fatalf("standard vignere decryption failed, expected %v, found %v", expected, encrypted)
	}
}
