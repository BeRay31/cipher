package playfair_test

import (
	"testing"

	"github.com/mkamadeus/cipher/cipher/playfair"
)

func TestGenerateTable(t *testing.T) {
	key := "JALANGANESHASEPULUH"
	table := playfair.GenerateTable(key)
	expected := "ALNGESHPUBCDFIKMOQRTVWXYZ"

	if expected != string(table) {
		t.Fatalf("Table not equal")
	}
}

func TestGenerateTableNormalize(t *testing.T) {
	key := "J ALAN GAN.E,SHA SEPU LUH"
	table := playfair.GenerateTable(key)
	expected := "ALNGESHPUBCDFIKMOQRTVWXYZ"

	if expected != string(table) {
		t.Fatalf("Table not equal")
	}
}

func TestShiftHorizontal(t *testing.T) {
	key := "JALANGANESHASEPULUH"
	table := playfair.GenerateTable(key)
	expected := "ALNGESHPUBCDFIKMOQRTVWXYZ"

	if expected != string(table) {
		t.Fatalf("Table not equal")
	}

	row, col := 0, 4
	expectedRow, expectedCol := 1, 0
	newRow, newCol := table.ShiftHorizontal(row, col)
	if newRow != expectedRow || newCol != expectedCol {
		t.Fatalf("Horizontal shift fails")
	}

}
