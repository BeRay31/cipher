package extended_test

import (
	"bytes"
	"testing"

	"github.com/mkamadeus/cipher/cipher/vigenere/extended"
)

func TestEncrypt(t *testing.T) {
	plain := []byte{0, 1, 2, 3, 252, 253, 254, 255}
	key := []byte{255, 0, 1, 2}
	expected := []byte{255, 1, 3, 5, 251, 253, 255, 0}

	encrypted := extended.Encrypt(plain, key)
	if bytes.Equal(expected, encrypted) {
		t.Fatalf("extended vigenere encryption failed, expected %v, found %v", expected, encrypted)
	}
}
