package full_test

import (
	"testing"

	"github.com/mkamadeus/cipher/cipher/vigenere/full"
)

func TestDecrypt(t *testing.T) {
	cipher := "lvvqhzngfhrvl"
	key := "sony"
	expected := "thisplaintext"

	encrypted := full.Decrypt(cipher, key)
	if encrypted != expected {
		t.Fatalf("full vignere decryption failed, expected %v, found %v", expected, encrypted)
	}
}

func TestDecryptWithUpperCase(t *testing.T) {
	cipher := "nGmni Tskcxipo esdskkxgmejvc!#!#!#!#"
	key := "KEY"
	expected := "dCode Vigenere automatically!#!#!#!#"

	encrypted := full.Decrypt(cipher, key)
	if encrypted != expected {
		t.Fatalf("full vignere encryption failed, expected %v, found %v", expected, encrypted)
	}
}
