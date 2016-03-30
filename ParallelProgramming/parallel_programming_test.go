package main

//Matrikelnummer Carsten Bieber:
//Matrikelnummer Tassilo Kloos: 2257414

import "testing"

//Testfall, ob ReadPicture richtig funktioniert
func TestReadPicture(t *testing.T) {
	pic := readPicture("schwarz_weiss")
	if pic == nil {
		t.Error("Fehler beim Einlesen des Bildes")
	}
}

//Testfall, ob ReadPicture bei falsche Bild richtig funktioniert
func TestReadPictureFalse(t *testing.T) {
	pic := readPicture("false")
	if pic != nil {
		t.Error("Fehler beim Einlesen eines nicht vorhandenen Bildes")
	}
}

func TestTransformWithoutFilter(t *testing.T) {
	if !transformWithoutFilter("schwarz_weiss") {
		t.Error("Fehler beim Transformieren des Bildes ohne Filter")
	}
}

func TestAnalyzePictureNormal(t *testing.T) {
	if !analyzePicture("schwarz_weiss", "normal") {
		t.Error("Fehler beim Analysieren des Bildes ohne Filter")
	}
}

func TestAnalyzePictureFloydSteinberg(t *testing.T) {
	if !analyzePicture("schwarz_weiss", "FloydSteinberg") {
		t.Error("Fehler beim Analysieren des Bildes mit FloydSteinberg")
	}
}

func TestTransformPicture(t *testing.T) {
	transformPicture("schwarz_weiss")
}

func TestMain(t *testing.T) {
	main()
}
