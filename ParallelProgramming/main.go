package main

import (
	"fmt"
	"image"
	"os"
	"runtime"
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

//Auslesen der maximal benutzten Prozessoren
func getGOMAXPROCS() int {
	return runtime.GOMAXPROCS(0)
}

//Setzen der maximal benutzten Prozessoren
func setGOMAXPROCS(threads int) {
	runtime.GOMAXPROCS(threads)
}

//Funktion, ob addierte Pixelanzahl zwischen 0 und 255 liegt
func checkValueOfPixel(value uint8, add int32) uint8 {
	result := int32(value) + add
	if result > 255 {
		//falls addierte Pixelanzahl > 255: Setzen des Wertes auf 255
		result = 255
	} else if result < 0 {
		//falls addierte Pixelanzahl < 0: Setzen des Wertes auf 0
		result = 0
	}
	return uint8(result)
}

//Funktion, die Zeit zur Ausführung der sequentiellen Transformation misst
func analyzePictureSeq(input, method string) bool {
	tBefore := time.Now()
	result := false
	//neues Objekt zur sequentiellen Transformation
	trSeq := transformSeq{picture}
	//Aufruf der Transformations-Methode
	result = trSeq.transformSequentiell(input, method)
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
	//neues Objekt zur parallelen Transformation
	trPar := transformPar{picture}
	//Aufruf der Transformations-Methode
	result = trPar.transformParallel(input, method)
	duration := time.Since(tBefore)
	//Ausgabe der Zeit in Sekunden mit 3 Kommastellen
	msec := int32(duration.Seconds() * 1000)
	sec := float32(msec) / 1000
	fmt.Println("Dauer bei ", method, " parallel(", getGOMAXPROCS(), " Thread(s)): ", sec, " sec")
	return result
}

//Funktion, die ausgewähltes Bild sequentiell wie auch parallel neu berechnet
func transformProcessor(input, method string) {
	//1-4 Prozessoren werden verwendet
	picture = readPicture(input)
	analyzePictureSeq(input, method)
	setGOMAXPROCS(1)
	analyzePicturePar(input, method)
	setGOMAXPROCS(2)
	analyzePicturePar(input, method)
	setGOMAXPROCS(3)
	analyzePicturePar(input, method)
	setGOMAXPROCS(4)
	analyzePicturePar(input, method)
}

func main() {
	//	validInputName := false
	//	var name string
	//	for validInputName == false {
	//		validInputName = true
	//		reader := bufio.NewReader(os.Stdin)
	//		fmt.Print("Bild: ")
	//		name, _ = reader.ReadString('\n')
	//		name = strings.Trim(name, "\r")
	//		name = strings.Replace(name, "\r", "", -1)
	//		name = strings.Replace(name, "\n", "", -1)
	//		if _, err := os.Stat("pictures/" + name); os.IsNotExist(err) {
	//			validInputName = false
	//			fmt.Println("Falsche Eingabe!")
	//		}
	//	}
	//	validInputMethod := false
	//	var method string
	//	for validInputMethod == false {
	//		validInputMethod = true
	//		reader := bufio.NewReader(os.Stdin)
	//		fmt.Print("Methode: FloydSteinberg (1), Algorithmus 2 (2), Algorithmus 3 (3), Schwellwert (4), Graustufen (5): ")
	//		inputMethod, _ := reader.ReadString('\n')
	//		inputMethod = strings.Trim(inputMethod, "\r")
	//		inputMethod = strings.Replace(inputMethod, "\r", "", -1)
	//		inputMethod = strings.Replace(inputMethod, "\n", "", -1)
	//		switch inputMethod {
	//		case "1":
	//			method = "FloydSteinberg"
	//		case "2":
	//			method = "Algorithm2"
	//		case "3":
	//			method = "Algorithm3"
	//		case "4":
	//			method = "Schwellwert"
	//		case "5":
	//			method = "Graustufen"
	//		default:
	//			validInputMethod = false
	//			fmt.Println("Falsche Eingabe!")
	//		}
	//	}
	//	transformProcessor(name, method)
	transformProcessor("dhbw.jpg", "FloydSteinberg") //<-- zu Testzwecken, um Eingabe zu überspringen
}
