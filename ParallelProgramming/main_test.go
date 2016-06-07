package main

//Matrikelnummer Carsten Bieber: 4346441
//Matrikelnummer Tassilo Kloos: 2257414

import (
	"bufio"
	"bytes"
	"os"
	"testing"
)

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

//Test, ob CheckValueOfPixel bei Zahlen > 255 funktoniert
func TestCheckValueOfPixel255(t *testing.T) {
	if checkValueOfPixel(255, 20) > 255 {
		t.Error("Fehler beim Berechnen des Pixelwertes über 255")
	}
}

//Test, ob CheckValueOfPixel bei Zahlen < 0 funktoniert
func TestCheckValueOfPixel0(t *testing.T) {
	if checkValueOfPixel(0, -20) < 0 {
		t.Error("Fehler beim Berechnen des Pixelwertes unter 0")
	}
}

//Test, ob paralleles und sequentielles Verfahren das gleiche Ergebnis liefern
func TestParallelSequentiell(t *testing.T) {
	transformProcessor("schwarz_weiss.png", "FloydSteinberg")
	file1, err := os.Open("pictures/schwarz_weiss_FloydSteinberg_seq.png")
	if err != nil {
		t.Error("Fehler beim Öffnen des Bildes schwarz_weiss_FloydSteinberg_seq.png")
	}
	file2, err := os.Open("pictures/schwarz_weiss_FloydSteinberg_par4.png")
	if err != nil {
		t.Error("Fehler beim Öffnen des Bildes schwarz_weiss_FloydSteinberg_par4.png")
	}
	scan1 := bufio.NewScanner(file1)
	scan2 := bufio.NewScanner(file2)

	for scan1.Scan() {
		scan2.Scan()
		if !bytes.Equal(scan1.Bytes(), scan2.Bytes()) {
			t.Error("Fehler beim Vergleich der sequentiellen und parallelen Lösung")
		}
	}
}

//Test, ob Getter und Setter der maximal verwendeten Prozessoren funktioniert
func TestGetterAndSetter(t *testing.T) {
	setGOMAXPROCS(4)
	if getGOMAXPROCS() != 4 {
		t.Error("Fehler beim Setzen der maximal benötigten Prozesse")
	}
}
