package main

//Matrikelnummer Carsten Bieber:
//Matrikelnummer Tassilo Kloos: 2257414

import "testing"

//Testfall, ob ReadPicture richtig funktioniert
func TestReadPicture(t *testing.T) {
	pic, _ := readPicture("schwarz_klein")
	if pic == nil {
		t.Error("Fehler beim Einlesen des Bildes")
	}
}

//Testfall, ob ReadPicture bei falsche Bild richtig funktioniert
func TestReadPictureFalse(t *testing.T) {
	pic, _ := readPicture("false")
	if pic != nil {
		t.Error("Fehler beim Einlesen eines nicht vorhandenen Bildes")
	}
}

func TestAnalyzePicture(t *testing.T) {
	result := analyzePicture(readPicture("schwarz_klein"))
	if result != true {
		t.Error("Fehler beim Analysieren des Bildes")
	}
}

func TestMain(t *testing.T) {
	main()
}
