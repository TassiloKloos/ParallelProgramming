package main

//Matrikelnummer Carsten Bieber:
//Matrikelnummer Tassilo Kloos: 2257414

import "testing"

func TestTransformWithoutFilterPNG(t *testing.T) {
	tr := transformSeq{readPicture("schwarz_weiss.png")}
	if !tr.transformWithoutFilter("schwarz_weiss.png") {
		t.Error("Fehler beim Transformieren des PNG-Bildes mit Schwellwert")
	}
}

func TestTransformWithoutFilterJPG(t *testing.T) {
	tr := transformSeq{readPicture("schwarz_weiss.png")}
	if !tr.transformWithoutFilter("schwarz_weiss.jpg") {
		t.Error("Fehler beim Transformieren des JPG-Bildes mit Schwellwert")
	}
}

func TestTransformWithFloydSteinbergPNG(t *testing.T) {
	tr := transformSeq{readPicture("schwarz_weiss.png")}
	if !tr.transformWithFloydSteinberg("schwarz_weiss.png") {
		t.Error("Fehler beim Transformieren des PNG-Bildes ohne Filter")
	}
}

func TestTransformWithFloydSteinbergJPG(t *testing.T) {
	tr := transformSeq{readPicture("schwarz_weiss.png")}
	if !tr.transformWithFloydSteinberg("schwarz_weiss.jpg") {
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
