package extended_test

import (
	"bytes"
	"testing"

	"github.com/mkamadeus/cipher/cipher/vigenere/extended"
)

func TestDecrypt(t *testing.T) {
	cipher := []byte{255, 1, 3, 5, 251, 253, 255, 0}
	key := []byte{255, 2, 1, 2}
	expected := []byte{0, 254, 2, 3, 252, 253, 254, 255}

	decrypted := extended.Decrypt(cipher, key)
	if bytes.Equal(expected, decrypted) {
		t.Fatalf("extended vigenere encryption failed, expected %v, found %v", expected, decrypted)
	}
}
