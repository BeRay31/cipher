package hill_test

import (
	"testing"

	"github.com/mkamadeus/cipher/cipher/hill"
)

func TestDecrypt(t *testing.T) {
	cipher := "CFIXOP"
	key := "bcef"
	decrypted, err := hill.Decrypt(cipher, key)
	expected := "ABCDEF"

	if err != nil {
		t.Fatalf("error on hill decryption: %v", err)
	} else if decrypted != expected {
		t.Fatalf("decryption failed, expected %v, found %v", expected, decrypted)
	}
}

func TestDecrypt2(t *testing.T) {
	cipher := "IJGAYNIIPSLB"
	key := "bcdabefga"
	decrypted, err := hill.Decrypt(cipher, key)
	expected := "ABCDEFBGHFJH"

	if err != nil {
		t.Fatalf("error on hill decryption: %v", err)
	} else if decrypted != expected {
		t.Fatalf("decryption failed, expected %v, found %v", expected, decrypted)
	}
}

func TestDecrypt3(t *testing.T) {
	cipher := "AFKPUZFEJZTJ"
	key := "f"
	decrypted, err := hill.Decrypt(cipher, key)
	expected := "ABCDEFBGHFJH"

	if err != nil {
		t.Fatalf("error on hill decryption: %v", err)
	} else if decrypted != expected {
		t.Fatalf("decryption failed, expected %v, found %v", expected, decrypted)
	}
}
