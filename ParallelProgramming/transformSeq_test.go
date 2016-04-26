package main

//Matrikelnummer Carsten Bieber:
//Matrikelnummer Tassilo Kloos: 2257414

import "testing"

func TestTransformWithFloydSteinbergPNGSeq(t *testing.T) {
	tr := transformSeq{readPicture("schwarz_weiss.png")}
	if !tr.transformSequentiell("schwarz_weiss.png", "FloydSteinberg") {
		t.Error("Fehler beim Transformieren des PNG-Bildes ohne Filter")
	}
}

func TestTransformWithFloydSteinbergJPGSeq(t *testing.T) {
	tr := transformSeq{readPicture("schwarz_weiss.png")}
	if !tr.transformSequentiell("schwarz_weiss.jpg", "FloydSteinberg") {
		t.Error("Fehler beim Transformieren des JPG-Bildes ohne Filter")
	}
}

func TestAnalyzePictureSchwellwertPNGSeq(t *testing.T) {
	if !analyzePictureSeq("schwarz_weiss.png", "Schwellwert") {
		t.Error("Fehler beim sequentiellen Analysieren des PNG-Bildes mit Schwellwert-Verfahren")
	}
}

func TestAnalyzePictureSchwellwertJPGSeq(t *testing.T) {
	if !analyzePictureSeq("schwarz_weiss.jpg", "Schwellwert") {
		t.Error("Fehler beim sequentiellen Analysieren des JPG-Bildes mit Schwellwert-Verfahren")
	}
}

func TestAnalyzePictureFloydSteinbergPNGSeq(t *testing.T) {
	if !analyzePictureSeq("schwarz_weiss.png", "FloydSteinberg") {
		t.Error("Fehler beim sequentiellen Analysieren des PNG-Bildes mit FloydSteinberg")
	}
}

func TestAnalyzePictureFloydSteinbergJPGSeq(t *testing.T) {
	if !analyzePictureSeq("schwarz_weiss.jpg", "FloydSteinberg") {
		t.Error("Fehler beim sequentiellen Analysieren des JPG-Bildes mit FloydSteinberg")
	}
}

func TestAnalyzePictureAlgorithm2PNGSeq(t *testing.T) {
	if !analyzePictureSeq("schwarz_weiss.png", "Algorithm2") {
		t.Error("Fehler beim sequentiellen Analysieren des PNG-Bildes mit Algorithmus 2")
	}
}

func TestAnalyzePictureAlgorithm2JPGSeq(t *testing.T) {
	if !analyzePictureSeq("schwarz_weiss.jpg", "Algorithm2") {
		t.Error("Fehler beim sequentiellen Analysieren des JPG-Bildes mit Algorithmus 2")
	}
}

func TestAnalyzePictureAlgorithm3PNGSeq(t *testing.T) {
	if !analyzePictureSeq("schwarz_weiss.png", "Algorithm3") {
		t.Error("Fehler beim sequentiellen Analysieren des PNG-Bildes mit Algorithmus 3")
	}
}

func TestAnalyzePictureAlgorithm3JPGSeq(t *testing.T) {
	if !analyzePictureSeq("schwarz_weiss.jpg", "Algorithm3") {
		t.Error("Fehler beim sequentiellen Analysieren des JPG-Bildes mit Algorithmus 3")
	}
}

func TestAnalyzePictureGraustufenPNGSeq(t *testing.T) {
	if !analyzePictureSeq("schwarz_weiss.png", "Graustufen") {
		t.Error("Fehler beim sequentiellen Analysieren des PNG-Bildes mit Graustufen")
	}
}

func TestAnalyzePictureGraustufenJPGSeq(t *testing.T) {
	if !analyzePictureSeq("schwarz_weiss.jpg", "Graustufen") {
		t.Error("Fehler beim sequentiellen Analysieren des JPG-Bildes mit Graustufen")
	}
}
