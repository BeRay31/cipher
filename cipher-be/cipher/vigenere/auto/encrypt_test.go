package auto_test

import (
	"testing"

	"github.com/mkamadeus/cipher/cipher/vigenere/auto"
)

func TestEncrypt(t *testing.T) {
	plain := "negarapenghasilminyak"
	key := "INDO"
	expected := "vrjoeeveegwefosmavjms"

	encrypted := auto.Encrypt(plain, key)
	if encrypted != expected {
		t.Fatalf("auto-key vignere encryption failed, expected %v, found %v", expected, encrypted)
	}
}

func TestEncryptWithUpperCase(t *testing.T) {
	plain := "dCode Vigenere automatically!#!#!#!#"
	key := "KEY"
	expected := "nGmggJlkzvkvrelxogthucttny"

	encrypted := auto.Encrypt(plain, key)
	if encrypted != expected {
		t.Fatalf("auto-key vignere encryption failed, expected %v, found %v", expected, encrypted)
	}
}

func TestEncryptIgnoreSymbol(t *testing.T) {
	plain := "dCode Vigenere !#!#!#!# automatically"
	key := "KEY"
	expected := "nGmggJlkzvkvrelxogthucttny"

	encrypted := auto.Encrypt(plain, key)
	if encrypted != expected {
		t.Fatalf("standard vignere encryption failed, expected %v, found %v", expected, encrypted)
	}
}
