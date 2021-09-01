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
		t.Fatalf("Table not equal, expected %v, found %v", expected, table)
	}
}

func TestGenerateTableUnique(t *testing.T) {
	key := "PISANG"
	table := playfair.GenerateTable(key)
	expected := "PISANGBCDEFHKLMOQRTUVWXYZ"

	t.Log(table)

	if expected != string(table) {
		t.Fatalf("Table not equal, expected %v, found %v", expected, table)
	}
}

func TestGenerateTableNormalize(t *testing.T) {
	key := "J ALAN GAN.E,SHA SEPU LUH"
	table := playfair.GenerateTable(key)
	expected := "ALNGESHPUBCDFIKMOQRTVWXYZ"

	if expected != string(table) {
		t.Fatalf("Table not equal, expected %v, found %v", expected, table)
	}
}

func TestShiftHorizontal(t *testing.T) {
	key := "JALANGANESHASEPULUH"
	table := playfair.GenerateTable(key)
	expected := "ALNGESHPUBCDFIKMOQRTVWXYZ"

	if expected != string(table) {
		t.Fatalf("Table not equal, expected %v, found %v", expected, table)
	}

	row, col := 0, 4
	expectedRow, expectedCol := 1, 0
	newRow, newCol := table.ShiftHorizontal(row, col, 1)
	if newRow != expectedRow || newCol != expectedCol {
		t.Fatalf("Horizontal shift fails")
	}

}
