package standard_test

import (
	"testing"

	"github.com/mkamadeus/cipher/cipher/vigenere/standard"
)

func TestEncrypt(t *testing.T) {
	plain := "thisplaintext"
	key := "sony"
	expected := "LVVQHZNGFHRVL"

	encrypted := standard.Encrypt(plain, key)
	if encrypted != expected {
		t.Fatalf("standard vignere encryption failed, expected %v, found %v", expected, encrypted)
	}
}
