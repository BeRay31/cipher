package auto_test

import (
	"testing"

	"github.com/mkamadeus/cipher/cipher/vigenere/auto"
)

func TestDecrypt(t *testing.T) {
	cipher := "vrjoeeveegwefosmavjms"
	key := "INDO"
	expected := "NEGARAPENGHASILMINYAK"

	decrypted := auto.Decrypt(cipher, key)
	if decrypted != expected {
		t.Fatalf("auto-key vignere decryption failed, expected %v, found %v", expected, decrypted)
	}
}

func TestDecryptWithUpperCase(t *testing.T) {
	cipher := "nGmgg(*&%^&Jlkzvkvrelxogthucttny"
	key := "KEY  "
	expected := "DCODEVIGENEREAUTOMATICALLY"

	decrypted := auto.Decrypt(cipher, key)
	if decrypted != expected {
		t.Fatalf("auto-ke vignere decryption failed, expected %v, found %v", expected, decrypted)
	}
}
