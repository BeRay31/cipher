package playfair_test

import (
	"testing"

	"github.com/mkamadeus/cipher/cipher/playfair"
)

func TestDecrypt(t *testing.T) {
	cipher := "ZBRSFYKUPGLGRKVSNLQV"
	key := "JALANGANESHASEPULUH"
	decrypted, err := playfair.Decrypt(cipher, key)
	expected := "TEMUIIBUNANTIMALAM"

	if err != nil {
		t.Fatalf("error on playfair encryption: %v", err)
	} else if decrypted != expected {
		t.Fatalf("encryption failed, expected %v, found %v", expected, decrypted)
	}
}
