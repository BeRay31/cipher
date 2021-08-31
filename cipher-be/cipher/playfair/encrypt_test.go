package playfair_test

import (
	"testing"

	"github.com/mkamadeus/cipher/cipher/playfair"
)

func TestEncrypt(t *testing.T) {
	plain := "temuiibunantimalam"
	key := "JALANGANESHASEPULUH"
	cipher, err := playfair.Encrypt(plain, key)
	expected := "ZBRSFYKUPGLGRKVSNLQV"

	if err != nil {
		t.Fatalf("error on playfair encryption: %v", err)
	} else if cipher != expected {
		t.Fatalf("encryption failed, expected %v, found %v", expected, cipher)
	}
}
