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

func TestTransformPicture(t *testing.T) {
	transformPicture("schwarz_weiss.png")
	//geeignete Prüfung einbauen
}

func TestTransformProcessor(t *testing.T) {
	transformProcessor("schwarz_weiss.png", "FloydSteinberg")
	//geeignete Prüfung einbauen
}

func TestGetterAndSetter(t *testing.T) {
	setGOMAXPROCS(4)
	if getGOMAXPROCS() != 4 {
		t.Error("Fehler beim Setzen der maximal benötigten Prozesse")
	}
}

func TestMain(t *testing.T) {
	main()
	//geeignete Prüfung einbauen
}
