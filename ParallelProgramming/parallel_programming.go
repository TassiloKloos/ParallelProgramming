package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	_ "image/png"
	"os"
)

//Definition der benötigten Variablen
var pic image.Image
var newPic *os.File
var countPixel int
var percent float32

func readPicture(input string) (image.Image, string) {
	fmt.Println("            ", input)
	reader, err := os.Open("pictures/" + input + ".png")
	if err != nil {
		fmt.Println(err.Error())
		return nil, input
	}
	defer reader.Close()

	pic, _, err := image.Decode(reader)
	if err != nil {
		return nil, input
	}
	return pic, input
}

func analyzePicture(pic image.Image, input string) bool {
	//aufteilen in verschiedene Filter und Sequentiell / Parallel
	bounds := pic.Bounds()
	newPic, _ = os.Create("pictures/" + input + "_neu.png")
	defer newPic.Close()
	m := image.NewRGBA(image.Rectangle{Min: image.Point{0, 0}, Max: image.Point{bounds.Max.X, bounds.Max.Y}})
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := pic.At(x, y).RGBA()
			// Maximalwert: 65535, bzw. 196605
			// Hälfte: 98302,5
			wert := r + g + b
			percent = float32(wert*10000/196605) / 100 //Runden auf 2 Kommazahlen
			var wert2 int
			if wert > 98302 {
				wert2 = 196605
				m.Set(x, y, color.RGBA{255, 255, 255, 255})
			} else {
				wert2 = 0
				m.Set(x, y, color.RGBA{0, 0, 0, 255})
			}
			fmt.Println("Wert: ", wert, " (", percent, "%) --> ", wert2)
			countPixel = countPixel + 1
		}
	}
	png.Encode(newPic, m)
	fmt.Println("Pixelanzahl: ", countPixel)
	fmt.Println("")
	return true
}

func testPicture() {
	analyzePicture(readPicture("landscape"))
	//	analyzePicture(readPicture("landscapehd"))
	//	analyzePicture(readPicture("bunte_smarties"))
	//	analyzePicture(readPicture("flower"))
	//	analyzePicture(readPicture("weiss_klein"))
	//	analyzePicture(readPicture("schwarz_klein"))
	//	analyzePicture(readPicture("weiss_eins"))
	//	analyzePicture(readPicture("schwarz_eins"))
}

func main() {
	testPicture()
}
