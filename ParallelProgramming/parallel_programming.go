package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	_ "image/png"
	"os"
	"time"
)

//Definition der benötigten Variablen
var pic image.Image
var newPic *os.File
var percent float32

func readPicture(input string) image.Image {
	reader, err := os.Open("pictures/" + input + ".png")
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	defer reader.Close()

	pic, _, err := image.Decode(reader)
	if err != nil {
		return nil
	}
	return pic
}

//Funktion, die Bild ohne Filter transformiert
func transformWithoutFilter(input string) bool {
	pic := readPicture(input)
	bounds := pic.Bounds()
	newPic, _ = os.Create("pictures/" + input + "_neu.png")
	defer newPic.Close()
	m := image.NewRGBA(image.Rectangle{Min: image.Point{0, 0}, Max: image.Point{bounds.Max.X, bounds.Max.Y}})
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := pic.At(x, y).RGBA()
			// Maximalwert: 65535, bzw. 196605 --> Hälfte: 98302,5
			wert := r + g + b
			percent = float32(wert*10000/196605) / 100 //Runden auf 2 Kommazahlen
			if percent >= 50 {
				m.Set(x, y, color.RGBA{255, 255, 255, 255})
			} else {
				m.Set(x, y, color.RGBA{0, 0, 0, 255})
			}
		}
	}
	png.Encode(newPic, m)
	return true
}

//Funktion, die Bild nach Filter von Floyd / Steinberg transformiert
func transformWithFloydSteinberg(input string) bool {
	pic := readPicture(input)
	bounds := pic.Bounds()
	newPic, _ = os.Create("pictures/" + input + "_floydsteinberg.png")
	defer newPic.Close()
	m := image.NewRGBA(image.Rectangle{Min: image.Point{0, 0}, Max: image.Point{bounds.Max.X, bounds.Max.Y}})
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := pic.At(x, y).RGBA()
			// Maximalwert: 65535, bzw. 196605 --> Hälfte: 98302,5
			wert := r + g + b
			percent = float32(wert*10000/196605) / 100 //Runden auf 2 Kommazahlen
			if percent >= 50 {
				m.Set(x, y, color.RGBA{255, 255, 255, 255})
			} else {
				m.Set(x, y, color.RGBA{0, 0, 0, 255})
			}
			//TODO:Verteilen der prozentual fehlenden Punkte
			//falls nicht letzte Y-Reihe (dort keine Verteilung auf untere Reihe möglich)
			if y < bounds.Max.Y-1 {
				//Verteilung auf Reihe darunter: 3/16, 5/16, 1/16
			}
			//falls nicht letzte X-Reihe (dort keine Verteilung nach links / rechts möglich)
			if x < bounds.Max.X-1 {
				//Verteilung auf rechts: 7/16
			}
		}
	}
	png.Encode(newPic, m)
	return true
}

//Funktion, die Zeit zur Ausführung der Transformation misst
func analyzePicture(input, method string) bool {
	tBefore := time.Now()
	result := false
	//Aufruf der Transformations-Methode
	switch method {
	case "normal":
		result = transformWithoutFilter(input)
	case "FloydSteinberg":
		result = transformWithFloydSteinberg(input)
	}
	//TODO: aufteilen in verschiedene Filter und Sequentiell / Parallel
	duration := time.Since(tBefore)
	min := int32(duration.Minutes())
	msec := int32(duration.Seconds() * 1000)
	sec := float32(msec) / 1000
	overSixty := int32(sec / 60)
	if overSixty > 0 {
		sec = sec - float32(overSixty*60)
	}
	fmt.Println("Dauer bei ", method, ": ", min, " min, ", sec, " sec")
	return result
}

//Funktion, die ausgewähltes Bild in allen Methoden neu berechnet
func transformPicture(input string) {
	fmt.Println("Bild: ", input)
	analyzePicture(input, "normal")
	analyzePicture(input, "FloydSteinberg")
	fmt.Println("")
}

func testPicture() {
	//	transformPicture("landscape")
	//	transformPicture("bunte_smarties")
	//	transformPicture("flower")
	//	transformPicture("newyork")
	//	transformPicture("middleage")
	transformPicture("schwarz_weiss")
}

func main() {
	testPicture()
}
