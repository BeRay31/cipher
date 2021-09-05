package full_test

import (
	"testing"

	"github.com/mkamadeus/cipher/cipher/vigenere/full"
)

func TestEncrypt(t *testing.T) {
	plain := "thisisplaintext"
	key := "sony"
	expected := "DPYFNWKHRDCABKO"

	encrypted := full.Encrypt(plain, key)
	if encrypted != expected {
		t.Fatalf("full vignere encryption failed, expected %v, found %v", expected, encrypted)
	}
}

func TestEncrypt2(t *testing.T) {
	plain := "t hi(*&"
	key := "so(*&^ny"
	expected := "DPY"

	encrypted := full.Encrypt(plain, key)
	if encrypted != expected {
		t.Fatalf("full vignere encryption failed, expected %v, found %v", expected, encrypted)
	}
}

func TestEncryptWithUpperCase(t *testing.T) {
	plain := "dCode Vigenere automatically!#!#!#!#"
	key := "KEY"
	expected := "YMKYAMPVWNAZVNXIPTHUQCNHET"

	encrypted := full.Encrypt(plain, key)
	if encrypted != expected {
		t.Fatalf("full vignere encryption failed, expected %v, found %v", expected, encrypted)
	}
}
