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
func TestTransformWithoutFilterPNG(t *testing.T) {
	pic = readPicture("schwarz_weiss.png")
	if !transformWithoutFilter("schwarz_weiss.png") {
		t.Error("Fehler beim Transformieren des PNG-Bildes mit Schwellwert")
	}
}

func TestTransformWithoutFilterJPG(t *testing.T) {
	pic = readPicture("schwarz_weiss.jpg")
	if !transformWithoutFilter("schwarz_weiss.jpg") {
		t.Error("Fehler beim Transformieren des JPG-Bildes mit Schwellwert")
	}
}

func TestTransformWithFloydSteinbergPNG(t *testing.T) {
	pic = readPicture("schwarz_weiss.png")
	if !transformWithFloydSteinberg("schwarz_weiss.png") {
		t.Error("Fehler beim Transformieren des PNG-Bildes ohne Filter")
	}
}

func TestTransformWithFloydSteinbergJPG(t *testing.T) {
	pic = readPicture("schwarz_weiss.jpg")
	if !transformWithFloydSteinberg("schwarz_weiss.jpg") {
		t.Error("Fehler beim Transformieren des JPG-Bildes ohne Filter")
	}
}

func TestAnalyzePictureNormalPNG(t *testing.T) {
	if !analyzePicture("schwarz_weiss.png", "Schwellwert") {
		t.Error("Fehler beim Analysieren des PNG-Bildes ohne Filter")
	}
}

func TestAnalyzePictureNormalJPG(t *testing.T) {
	if !analyzePicture("schwarz_weiss.jpg", "Schwellwert") {
		t.Error("Fehler beim Analysieren des JPG-Bildes ohne Filter")
	}
}

func TestAnalyzePictureFloydSteinbergPNG(t *testing.T) {
	if !analyzePicture("schwarz_weiss.png", "FloydSteinberg") {
		t.Error("Fehler beim Analysieren des PNG-Bildes mit FloydSteinberg")
	}
}

func TestAnalyzePictureFloydSteinbergJPG(t *testing.T) {
	if !analyzePicture("schwarz_weiss.jpg", "FloydSteinberg") {
		t.Error("Fehler beim Analysieren des JPG-Bildes mit FloydSteinberg")
	}
}

func TestAnalyzePictureAlgorithm2PNG(t *testing.T) {
	if !analyzePicture("schwarz_weiss.png", "Algorithm2") {
		t.Error("Fehler beim Analysieren des PNG-Bildes mit Algorithmus 2")
	}
}

func TestAnalyzePictureAlgorithm2JPG(t *testing.T) {
	if !analyzePicture("schwarz_weiss.jpg", "Algorithm2") {
		t.Error("Fehler beim Analysieren des JPG-Bildes mit Algorithmus 2")
	}
}

func TestAnalyzePictureAlgorithm3PNG(t *testing.T) {
	if !analyzePicture("schwarz_weiss.png", "Algorithm3") {
		t.Error("Fehler beim Analysieren des PNG-Bildes mit Algorithmus 3")
	}
}

func TestAnalyzePictureAlgorithm3JPG(t *testing.T) {
	if !analyzePicture("schwarz_weiss.jpg", "Algorithm3") {
		t.Error("Fehler beim Analysieren des JPG-Bildes mit Algorithmus 3")
	}
}

func TestTransformPicture(t *testing.T) {
	transformPicture("schwarz_weiss.png")
}

func TestMain(t *testing.T) {
	main()
}
