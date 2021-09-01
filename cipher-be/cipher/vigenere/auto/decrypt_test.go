package auto_test

import (
	"testing"

	"github.com/mkamadeus/cipher/cipher/vigenere/auto"
)

func TestDecrypt(t *testing.T) {
	cipher := "vrjoeeveegwefosmavjms"
	key := "INDO"
	expected := "negarapenghasilminyak"

	decrypted := auto.Decrypt(cipher, key)
	if decrypted != expected {
		t.Fatalf("standard vignere encryption failed, expected %v, found %v", expected, decrypted)
	}
}

func TestDecryptWithUpperCase(t *testing.T) {
	cipher := "nGmggJlkzvkvrelxogthucttny"
	key := "KEY  "
	expected := "dCodeVigenereautomatically"

	decrypted := auto.Decrypt(cipher, key)
	if decrypted != expected {
		t.Fatalf("standard vignere encryption failed, expected %v, found %v", expected, decrypted)
	}
}
