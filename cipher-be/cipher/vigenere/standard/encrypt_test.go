package standard_test

import (
	"testing"

	"github.com/mkamadeus/cipher/cipher/vigenere/standard"
)

func TestEncrypt(t *testing.T) {
	plain := "thisplaintext"
	key := "sony"
	expected := "lvvqhzngfhrvl"

	encrypted := standard.Encrypt(plain, key)
	if encrypted != expected {
		t.Fatalf("standard vignere encryption failed, expected %v, found %v", expected, encrypted)
	}
}

func TestEncryptWithUpperCase(t *testing.T) {
	plain := "dCode Vigenere automatically!#!#!#!#"
	key := "KEY"
	expected := "nGmni Tskcxipo esdskkxgmejvc!#!#!#!#"

	encrypted := standard.Encrypt(plain, key)
	if encrypted != expected {
		t.Fatalf("standard vignere encryption failed, expected %v, found %v", expected, encrypted)
	}
}
