package full_test

import (
	"testing"

	"github.com/mkamadeus/cipher/cipher/vigenere/full"
)

func TestDecryptt(t *testing.T) {
	cipher := "DPYFNWKHRDCABKO"
	key := "sony"
	expected := "THISISPLAINTEXT"

	decrypted := full.Decrypt(cipher, key)
	if decrypted != expected {
		t.Fatalf("full vignere Decryption failed, expected %v, found %v", expected, decrypted)
	}
}

func TestDecrypt2(t *testing.T) {
	cipher := "DPY"
	key := "son y"
	expected := "THI"

	decrypted := full.Decrypt(cipher, key)
	if decrypted != expected {
		t.Fatalf("full vignere decryption failed, expected %v, found %v", expected, decrypted)
	}
}

func TestDecryptionWithUpperCase(t *testing.T) {
	cipher := "YMKYAMPVWNAZVNXIPTHUQCNHET()*&^%$"
	key := "KEY"
	expected := "DCODEVIGENEREAUTOMATICALLY"

	decrypted := full.Decrypt(cipher, key)
	if decrypted != expected {
		t.Fatalf("full vignere decryption failed, expected %v, found %v", expected, decrypted)
	}
}
