package hill_test

import (
	"testing"

	"github.com/mkamadeus/cipher/cipher/hill"
)

func TestEncrypt(t *testing.T) {
	cipher := "ABCDEF"
	key := "bcef"
	encrypted, err := hill.Encrypt(cipher, key)
	expected := "CFIXOP"

	if err != nil {
		t.Fatalf("error on hill encryption: %v", err)
	} else if encrypted != expected {
		t.Fatalf("encryption failed, expected %v, found %v", expected, encrypted)
	}
}

func TestEncrypt2(t *testing.T) {
	cipher := "ABCDEFBGHFJH"
	key := "bcdabefga"
	encrypted, err := hill.Encrypt(cipher, key)
	expected := "IJGAYNIIPSLB"

	if err != nil {
		t.Fatalf("error on hill encryption: %v", err)
	} else if encrypted != expected {
		t.Fatalf("encryption failed, expected %v, found %v", expected, encrypted)
	}
}

func TestEncrypt3(t *testing.T) {
	cipher := "ABCDEFBGHFJH"
	key := "f"
	encrypted, err := hill.Encrypt(cipher, key)
	expected := "AFKPUZFEJZTJ"

	if err != nil {
		t.Fatalf("error on hill encryption: %v", err)
	} else if encrypted != expected {
		t.Fatalf("encryption failed, expected %v, found %v", expected, encrypted)
	}
}
