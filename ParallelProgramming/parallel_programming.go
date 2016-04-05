package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	_ "image/png"
	"os"
	"strings"
	"time"
)

//Definition der benötigten globalen Variablen
var pic image.Image
var newPic *os.File

//Bild wird eingelesen
func readPicture(input string) image.Image {
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
		result = 255
	} else if result < 0 {
		result = 0
	}
	return uint32(result)
}

//Funktion, die Bild ohne Filter transformiert
func transformWithoutFilter(input string) bool {
	bounds := pic.Bounds()
	i := strings.Index(input, ".")
	name := input[:i]
	format := input[i:]
	newPic, _ = os.Create("pictures/" + name + "_schwellwert" + format)
	defer newPic.Close()
	m := image.NewRGBA(image.Rectangle{Min: image.Point{0, 0}, Max: image.Point{bounds.Max.X, bounds.Max.Y}})
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := pic.At(x, y).RGBA()
			// Maximalwert: 65535, bzw. 196605 --> neues Maximum: 256
			wert := (r + g + b) / 256 / 3
			if wert >= 128 {
				m.Set(x, y, color.RGBA{255, 255, 255, 255})
			} else {
				m.Set(x, y, color.RGBA{0, 0, 0, 255})
			}
		}
	}
	if format == ".png" {
		png.Encode(newPic, m)
	} else {
		//TODO Speichern als JPG
	}
	return true
}

//Funktion, die Bild nach Filter von Floyd / Steinberg transformiert
func transformWithFloydSteinberg(input string) bool {
	bounds := pic.Bounds()
	i := strings.Index(input, ".")
	name := input[:i]
	format := input[i:]
	newPic, _ = os.Create("pictures/" + name + "_floydsteinberg" + format)
	defer newPic.Close()
	m := image.NewRGBA(image.Rectangle{Min: image.Point{0, 0}, Max: image.Point{bounds.Max.X, bounds.Max.Y}})
	// Array, welches Differenzen der umliegenden Pixel speichert
	differenceOfPixel := make([][]float32, bounds.Max.Y+1)
	for element := range differenceOfPixel {
		differenceOfPixel[element] = make([]float32, bounds.Max.X+1)
	}
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			var difference float32
			r, g, b, _ := pic.At(x, y).RGBA()

			//			//zu Testzwecken, Ausgabe der Differenzen-Matrix
			//			for by := bounds.Min.Y; by < bounds.Max.Y; by++ {
			//				for bx := bounds.Min.X; bx < bounds.Max.X; bx++ {
			//					fmt.Printf("[%v] ", int32(differenceOfPixel[by][bx]))
			//				}
			//				fmt.Println("")
			//			}
			//			fmt.Println("")

			r = checkValueOfPixel(r/256, differenceOfPixel[y][x])
			g = checkValueOfPixel(g/256, differenceOfPixel[y][x])
			b = checkValueOfPixel(b/256, differenceOfPixel[y][x])
			// Maximalwert: 65535, --> neuer Maximalwert: 256
			value := (r + g + b) / 3
			if value >= 128 {
				m.Set(x, y, color.RGBA{255, 255, 255, 255})
				//verbleibende Differenz wird berechnet
				difference = float32(-(255 - float32(value)))
			} else {
				m.Set(x, y, color.RGBA{0, 0, 0, 255})
				//verbleibende Differenz wird berechnet
				difference = float32(value)
			}
			if x < bounds.Max.X-1 {
				// x+1, y = 7/16
				differenceOfPixel[y][x+1] = differenceOfPixel[y][x+1] + difference*7/16
			}
			if y < bounds.Max.Y-1 {
				if x < bounds.Max.X-1 {
					// x+1, y+1 = 1/16
					differenceOfPixel[y+1][x+1] = differenceOfPixel[y+1][x+1] + difference*1/16
				}
			}
			if x > 0 {
				// x-1, y+1 = 3/16
				differenceOfPixel[y+1][x-1] = differenceOfPixel[y+1][x-1] + difference*3/16
			}
			// x, y+1 = 5/16
			differenceOfPixel[y+1][x] = differenceOfPixel[y+1][x] + difference*5/16
		}
	}
	if format == ".png" {
		png.Encode(newPic, m)
	} else {
		//TODO Speichern als JPG
	}
	return true
}

//Funktion, die Bild mit Algorithmus 2 transformiert
func transformWithAlgorithm2(input string) bool {
	bounds := pic.Bounds()
	i := strings.Index(input, ".")
	name := input[:i]
	format := input[i:]
	newPic, _ = os.Create("pictures/" + name + "_algorithm2" + format)
	defer newPic.Close()
	m := image.NewRGBA(image.Rectangle{Min: image.Point{0, 0}, Max: image.Point{bounds.Max.X, bounds.Max.Y}})
	// Array, welches Differenzen der umliegenden Pixel speichert
	differenceOfPixel := make([][]float32, bounds.Max.Y+1)
	for element := range differenceOfPixel {
		differenceOfPixel[element] = make([]float32, bounds.Max.X+1)
	}
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			var difference float32
			r, g, b, _ := pic.At(x, y).RGBA()
			r = checkValueOfPixel(r/256, differenceOfPixel[y][x])
			g = checkValueOfPixel(g/256, differenceOfPixel[y][x])
			b = checkValueOfPixel(b/256, differenceOfPixel[y][x])
			// Maximalwert: 65535, --> neuer Maximalwert: 256
			value := (r + g + b) / 3
			if value >= 128 {
				m.Set(x, y, color.RGBA{255, 255, 255, 255})
				//verbleibende Differenz wird berechnet
				difference = float32(-(255 - float32(value)))
			} else {
				m.Set(x, y, color.RGBA{0, 0, 0, 255})
				//verbleibende Differenz wird berechnet
				difference = float32(value)
			}
			if x < bounds.Max.X-1 {
				// x+1, y = 4/12 = 1/3
				differenceOfPixel[y][x+1] = differenceOfPixel[y][x+1] + difference/3
			}
			if y < bounds.Max.Y-1 {
				if x < bounds.Max.X-1 {
					// x+1, y+1 = 1/12
					differenceOfPixel[y+1][x+1] = differenceOfPixel[y+1][x+1] + difference/12
				}
			}
			if x > 0 {
				// x-1, y+1 = 1/12
				differenceOfPixel[y+1][x-1] = differenceOfPixel[y+1][x-1] + difference/12
			}
			// x, y+1 = 4/12 = 1/3
			differenceOfPixel[y+1][x] = differenceOfPixel[y+1][x] + difference/3
		}
		//TODO 	x+2, y = 1/12
		//		x, y+2 = 1/12
	}
	if format == ".png" {
		png.Encode(newPic, m)
	} else {
		//TODO Speichern als JPG
	}
	return true
}

//Funktion, die Bild mit Algorithmus 3 transformiert
func transformWithAlgorithm3(input string) bool {
	bounds := pic.Bounds()
	i := strings.Index(input, ".")
	name := input[:i]
	format := input[i:]
	newPic, _ = os.Create("pictures/" + name + "_algorithm3" + format)
	defer newPic.Close()
	m := image.NewRGBA(image.Rectangle{Min: image.Point{0, 0}, Max: image.Point{bounds.Max.X, bounds.Max.Y}})
	// Array, welches Differenzen der umliegenden Pixel speichert
	differenceOfPixel := make([][]float32, bounds.Max.Y+1)
	for element := range differenceOfPixel {
		differenceOfPixel[element] = make([]float32, bounds.Max.X+1)
	}
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			var difference float32
			r, g, b, _ := pic.At(x, y).RGBA()
			r = checkValueOfPixel(r/256, differenceOfPixel[y][x])
			g = checkValueOfPixel(g/256, differenceOfPixel[y][x])
			b = checkValueOfPixel(b/256, differenceOfPixel[y][x])
			// Maximalwert: 65535, --> neuer Maximalwert: 256
			value := (r + g + b) / 3
			if value >= 128 {
				m.Set(x, y, color.RGBA{255, 255, 255, 255})
				//verbleibende Differenz wird berechnet
				difference = float32(-(255 - float32(value)))
			} else {
				m.Set(x, y, color.RGBA{0, 0, 0, 255})
				//verbleibende Differenz wird berechnet
				difference = float32(value)
			}
			if x < bounds.Max.X-1 {
				// x+1, y = 8/42 = 4/21
				differenceOfPixel[y][x+1] = differenceOfPixel[y][x+1] + difference*4/21
			}
			if y < bounds.Max.Y-1 {
				if x < bounds.Max.X-1 {
					// x+1, y+1 = 4/42 = 2/21
					differenceOfPixel[y+1][x+1] = differenceOfPixel[y+1][x+1] + difference*2/21
				}
			}
			if x > 0 {
				// x-1, y+1 = 4/42 = 2/21
				differenceOfPixel[y+1][x-1] = differenceOfPixel[y+1][x-1] + difference*2/21
			}
			// x, y+1 = 8/42 = 4/21
			differenceOfPixel[y+1][x] = differenceOfPixel[y+1][x] + difference*4/21
		}
		//TODO	x+2, y   = 4/42
		//		x-2, y+1 = 2/42
		//		x+2, y+1 = 2/42
		//		x-2, y+2 = 1/42
		//		x-1, y+2 = 2/42
		//		x  , y+2 = 4/42
		//		x+1, y+2 = 2/42
		//		x+2, y+2 = 1/42
	}
	if format == ".png" {
		png.Encode(newPic, m)
	} else {
		//TODO Speichern als JPG
	}
	return true
}

//Funktion, die Zeit zur Ausführung der Transformation misst
func analyzePicture(input, method string) bool {
	tBefore := time.Now()
	result := false
	//Aufruf der Transformations-Methode
	switch method {
	case "Schwellwert":
		result = transformWithoutFilter(input)
	case "FloydSteinberg":
		result = transformWithFloydSteinberg(input)
	case "Algorithm2":
		result = transformWithAlgorithm2(input)
	case "Algorithm3":
		result = transformWithAlgorithm3(input)
	}
	duration := time.Since(tBefore)
	msec := int32(duration.Seconds() * 1000)
	sec := float32(msec) / 1000
	fmt.Println("Dauer bei ", method, ": ", sec, " sec")
	return result
}

//Funktion, die ausgewähltes Bild in allen Methoden neu berechnet
func transformPicture(input string) {
	fmt.Println("Bild: ", input)
	pic = readPicture(input)
	analyzePicture(input, "Schwellwert")
	analyzePicture(input, "FloydSteinberg")
	analyzePicture(input, "Algorithm2")
	analyzePicture(input, "Algorithm3")
	fmt.Println("")
}

func main() {
	transformPicture("landscape.png")
	transformPicture("bunte_smarties.png")
	transformPicture("flower.png")
	transformPicture("newyork.png")
	transformPicture("middleage.png")
	transformPicture("schwarz_weiss.png")
	transformPicture("grau_vier.png")
	//	transformPicture("dhbw.jpg")
}
