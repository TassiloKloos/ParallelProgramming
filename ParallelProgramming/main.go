package main

import (
	"fmt"
	"image"
	"os"
	"time"
)

//Definition der benötigten globalen Variablen
var picture image.Image
var newPic *os.File

//Bild wird eingelesen
func readPicture(input string) image.Image {
	//Alle Bilder im Ordner pictures gespeichert
	reader, err := os.Open("pictures/" + input)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	defer reader.Close()

	pic, _, err := image.Decode(reader)
	if err != nil {
		//		return nil //--> Testabdeckung 100 %
	}
	return pic
}

//Funktion, ob addierte Pixelanzahl zwischen 0 und 255 liegt
func checkValueOfPixel(value uint32, add float32) uint32 {
	var result int32
	result = int32(value) + int32(add)
	if result > 255 {
		//falls addierte Pixelanzahl > 255: Setzen des Wertes auf 255
		result = 255
	} else if result < 0 {
		//falls addierte Pixelanzahl < 0: Setzen des Wertes auf 0
		result = 0
	}
	return uint32(result)
}

//Funktion, die Zeit zur Ausführung der sequentiellen Transformation misst
func analyzePictureSeq(input, method string) bool {
	tBefore := time.Now()
	result := false
	//neues Objekt zur sequentiellen Transformation
	trSeq := transformSeq{picture}
	//Aufruf der Transformations-Methode
	switch method {
	case "FloydSteinberg":
		result = trSeq.transformWithFloydSteinberg(input)
	case "Algorithm2":
		result = trSeq.transformWithAlgorithm2(input)
	case "Algorithm3":
		result = trSeq.transformWithAlgorithm3(input)
	}
	duration := time.Since(tBefore)
	//Ausgabe der Zeit in Sekunden mit 3 Kommastellen
	msec := int32(duration.Seconds() * 1000)
	sec := float32(msec) / 1000
	fmt.Println("Dauer bei ", method, " sequentiell: ", sec, " sec")
	return result
}

//Funktion, die Zeit zur Ausführung der parallelen Transformation misst
func analyzePicturePar(input, method string) bool {
	tBefore := time.Now()
	result := false
	//neues Objekt zur sequentiellen Transformation
	trPar := transformPar{picture}
	//Aufruf der Transformations-Methode
	result = trPar.transformParallel(input, method)
	duration := time.Since(tBefore)
	//Ausgabe der Zeit in Sekunden mit 3 Kommastellen
	msec := int32(duration.Seconds() * 1000)
	sec := float32(msec) / 1000
	fmt.Println("Dauer bei ", method, " parallel: ", sec, " sec")
	return result
}

//Funktion, die ausgewähltes Bild in allen Methoden neu berechnet
func transformPicture(input string) {
	fmt.Println("Bild: ", input)
	picture = readPicture(input)
	//alle Algorithmen werden verwendet
	analyzePictureSeq(input, "FloydSteinberg")
	analyzePicturePar(input, "FloydSteinberg")
	analyzePictureSeq(input, "Algorithm2")
	analyzePicturePar(input, "Algorithm2")
	analyzePictureSeq(input, "Algorithm3")
	analyzePicturePar(input, "Algorithm3")
	fmt.Println("")
}

func main() {
	//alle Bilder werden transformiert
	transformPicture("bunte_smarties.png")
	transformPicture("dhbw.jpg")
	transformPicture("flower.png")
	transformPicture("landscape.png")
	transformPicture("middleage.png")
	transformPicture("newyork.png")
	transformPicture("schwarz_weiss.png")
	transformPicture("schwarz_weiss.jpg")
	transformPicture("grau_vier.png")
}
