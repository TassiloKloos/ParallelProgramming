package main

//Matrikelnummer Carsten Bieber:
//Matrikelnummer Tassilo Kloos: 2257414

import "testing"

//Testfall, ob ReadPicture richtig funktioniert
func TestReadPicture(t *testing.T) {
	pic := readPicture("schwarz_weiss.png")
	if pic == nil {
		t.Error("Fehler beim Einlesen des Bildes")
	}
}

//Testfall, ob ReadPicture bei falsche Bild richtig funktioniert
func TestReadPictureFalse(t *testing.T) {
	pic := readPicture("false.png")
	if pic != nil {
		t.Error("Fehler beim Einlesen eines nicht vorhandenen Bildes")
	}
}

func TestCheckValueOfPixel255(t *testing.T) {
	if checkValueOfPixel(255, 20) > 255 {
		t.Error("Fehler beim Berechnen des Pixelwertes über 255")
	}
}

func TestCheckValueOfPixel0(t *testing.T) {
	if checkValueOfPixel(0, -20) < 0 {
		t.Error("Fehler beim Berechnen des Pixelwertes unter 0")
	}
}

//hier jeweils auch jpg testen
func TestTransformWithoutFilter(t *testing.T) {
	pic = readPicture("schwarz_weiss.png")
	if !transformWithoutFilter("schwarz_weiss") {
		t.Error("Fehler beim Transformieren des Bildes mit Schwellwert")
	}
}

func TestTtansformWithFloydSteinberg(t *testing.T) {
	pic = readPicture("schwarz_weiss.png")
	if !transformWithFloydSteinberg("schwarz_weiss") {
		t.Error("Fehler beim Transformieren des Bildes ohne Filter")
	}
}

func TestAnalyzePictureNormal(t *testing.T) {
	if !analyzePicture("schwarz_weiss.png", "Schwellwert") {
		t.Error("Fehler beim Analysieren des Bildes ohne Filter")
	}
}

func TestAnalyzePictureFloydSteinberg(t *testing.T) {
	if !analyzePicture("schwarz_weiss.png", "FloydSteinberg") {
		t.Error("Fehler beim Analysieren des Bildes mit FloydSteinberg")
	}
}

func TestAnalyzePictureAlgorithm2(t *testing.T) {
	if !analyzePicture("schwarz_weiss.png", "Algorithm2") {
		t.Error("Fehler beim Analysieren des Bildes mit Algorithmus 2")
	}
}

func TestAnalyzePictureAlgorithm3(t *testing.T) {
	if !analyzePicture("schwarz_weiss.png", "Algorithm3") {
		t.Error("Fehler beim Analysieren des Bildes mit Algorithmus 3")
	}
}

func TestTransformPicture(t *testing.T) {
	transformPicture("schwarz_weiss.png")
}

func TestMain(t *testing.T) {
	main()
}
