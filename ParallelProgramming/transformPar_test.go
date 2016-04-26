package main

//Matrikelnummer Carsten Bieber:
//Matrikelnummer Tassilo Kloos: 2257414

import "testing"

func TestTransformWithFloydSteinbergPNGPar(t *testing.T) {
	tr := transformPar{readPicture("schwarz_weiss.png")}
	if !tr.transformParallel("schwarz_weiss.png", "FloydSteinberg") {
		t.Error("Fehler beim Transformieren des PNG-Bildes ohne Filter")
	}
}

func TestTransformWithFloydSteinbergJPGPar(t *testing.T) {
	tr := transformPar{readPicture("schwarz_weiss.png")}
	if !tr.transformParallel("schwarz_weiss.jpg", "FloydSteinberg") {
		t.Error("Fehler beim Transformieren des JPG-Bildes ohne Filter")
	}
}

func TestAnalyzePictureSchwellwertPNGPar(t *testing.T) {
	if !analyzePicturePar("schwarz_weiss.png", "Schwellwert") {
		t.Error("Fehler beim parallelen Analysieren des PNG-Bildes mit Schwellwert-Verfahren")
	}
}

func TestAnalyzePictureSchwellwertJPGPar(t *testing.T) {
	if !analyzePicturePar("schwarz_weiss.jpg", "Schwellwert") {
		t.Error("Fehler beim parallelen Analysieren des JPG-Bildes mit Schwellwert-Verfahren")
	}
}

func TestAnalyzePictureFloydSteinbergPNGPar(t *testing.T) {
	if !analyzePicturePar("schwarz_weiss.png", "FloydSteinberg") {
		t.Error("Fehler beim parallelen Analysieren des PNG-Bildes mit FloydSteinberg")
	}
}

func TestAnalyzePictureFloydSteinbergJPGPar(t *testing.T) {
	if !analyzePicturePar("schwarz_weiss.jpg", "FloydSteinberg") {
		t.Error("Fehler beim parallelen Analysieren des JPG-Bildes mit FloydSteinberg")
	}
}

func TestAnalyzePictureAlgorithm2PNGPar(t *testing.T) {
	if !analyzePicturePar("schwarz_weiss.png", "Algorithm2") {
		t.Error("Fehler beim parallelen Analysieren des PNG-Bildes mit Algorithmus 2")
	}
}

func TestAnalyzePictureAlgorithm2JPGPar(t *testing.T) {
	if !analyzePicturePar("schwarz_weiss.jpg", "Algorithm2") {
		t.Error("Fehler beim parallelen Analysieren des JPG-Bildes mit Algorithmus 2")
	}
}

func TestAnalyzePictureAlgorithm3PNGPar(t *testing.T) {
	if !analyzePicturePar("schwarz_weiss.png", "Algorithm3") {
		t.Error("Fehler beim parallelen Analysieren des PNG-Bildes mit Algorithmus 3")
	}
}

func TestAnalyzePictureAlgorithm3JPGPar(t *testing.T) {
	if !analyzePicturePar("schwarz_weiss.jpg", "Algorithm3") {
		t.Error("Fehler beim parallelen Analysieren des JPG-Bildes mit Algorithmus 3")
	}
}

func TestAnalyzePictureGraustufenPNGPar(t *testing.T) {
	if !analyzePicturePar("schwarz_weiss.png", "Graustufen") {
		t.Error("Fehler beim parallelen Analysieren des PNG-Bildes mit Graustufen")
	}
}

func TestAnalyzePictureGraustufenJPGPar(t *testing.T) {
	if !analyzePicturePar("schwarz_weiss.jpg", "Graustufen") {
		t.Error("Fehler beim parallelen Analysieren des JPG-Bildes mit Graustufen")
	}
}
