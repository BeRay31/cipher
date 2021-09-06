package extended_test

import (
	"bytes"
	"testing"

	"github.com/mkamadeus/cipher/cipher/vigenere/extended"
)

func TestEncrypt(t *testing.T) {
	plain := "thisisplaintext"
	bytePlain := []byte{}
	for _, charPlain := range plain {
		bytePlain = append(bytePlain, byte(charPlain))
	}
	key := []byte{0}
	expected := []byte{115, 104, 106, 117, 104, 115, 113, 110, 96, 105, 111, 118, 100, 120, 117}

	encrypted := extended.Encrypt(bytePlain, key)
	if !bytes.Equal(expected, encrypted) {
		t.Fatalf("extended vigenere encryption failed, expected %v, found %v", expected, encrypted)
	}
}
