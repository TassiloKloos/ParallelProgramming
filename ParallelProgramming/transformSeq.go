package main

import (
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

//Objekte dieser Klasse besitzen Bild als Variable
type transformSeq struct {
	pic image.Image
}

//Funktion, die Bild einliest
func (t transformSeq) transformSequentiell(input, method string) bool {
	bounds := t.pic.Bounds()
	//Aufteilen des Inputs in Name und Dateiformat
	i := strings.Index(input, ".")
	name := input[:i]
	format := input[i:]
	//Erstellen des neuen Bildes mit entsprechendem Namen und Dateiformat
	newPic, _ = os.Create("pictures/" + name + "_" + method + "_seq" + format)
	defer newPic.Close()
	m := image.NewRGBA(image.Rectangle{Min: image.Point{0, 0}, Max: image.Point{bounds.Max.X, bounds.Max.Y}})
	// 2-dimensionales Array, welches Differenzen der umliegenden Pixel speichert
	differenceOfPixel := make([][]float32, bounds.Max.Y)
	for element := range differenceOfPixel {
		differenceOfPixel[element] = make([]float32, bounds.Max.X)
	}
	//zwei for-Schleifen, um jeden Pixelwert auszulesen
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		switch method {
		case "Schwellwert":
			m = t.transformLineSchwellwert(m, y, bounds, differenceOfPixel)
		case "FloydSteinberg":
			m = t.transformLineFloydSteinberg(m, y, bounds, differenceOfPixel)
		case "Algorithm2":
			m = t.transformLineAlgorithm2(m, y, bounds, differenceOfPixel)
		case "Algorithm3":
			m = t.transformLineAlgorithm3(m, y, bounds, differenceOfPixel)
		case "Graustufen":
			m = t.transformLineGraustufen(m, y, bounds, differenceOfPixel)
		}
	}
	if format == ".png" {
		png.Encode(newPic, m)
	} else {
		jpeg.Encode(newPic, m, nil)
	}
	return true
}

//Funktion, die jeweils eine Zeile mit Schwellwert transformiert
func (t transformSeq) transformLineSchwellwert(m *image.RGBA, y int, bounds image.Rectangle, differenceOfPixel [][]float32) *image.RGBA {
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		r, g, b, _ := t.pic.At(x, y).RGBA()
		// Einberechnug der bereits errechneten Differenzen von umliegenden Pixel
		r = checkValueOfPixel(r/256, differenceOfPixel[y][x])
		g = checkValueOfPixel(g/256, differenceOfPixel[y][x])
		b = checkValueOfPixel(b/256, differenceOfPixel[y][x])
		// Maximalwert: jeweils 256
		value := (r + g + b) / 3
		//Setzen eines neuen Farbwertes für Pixel, abhängig von derzeitigem Wert
		if value >= 128 {
			m.Set(x, y, color.RGBA{255, 255, 255, 255})
		} else {
			m.Set(x, y, color.RGBA{0, 0, 0, 255}) //wird in Test transformSeq_test nicht durchlaufen
		}
	}
	return m
}

//Funktion, die jeweils eine Zeile mit Floyd-Steinberg-Algorithmus transformiert
func (t transformSeq) transformLineFloydSteinberg(m *image.RGBA, y int, bounds image.Rectangle, differenceOfPixel [][]float32) *image.RGBA {
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		var difference float32
		r, g, b, _ := t.pic.At(x, y).RGBA()
		// Einberechnug der bereits errechneten Differenzen von umliegenden Pixel
		r = checkValueOfPixel(r/256, differenceOfPixel[y][x])
		g = checkValueOfPixel(g/256, differenceOfPixel[y][x])
		b = checkValueOfPixel(b/256, differenceOfPixel[y][x])
		// Maximalwert: jeweils 256
		value := (r + g + b) / 3
		//Setzen eines neuen Farbwertes für Pixel, abhängig von derzeitigem Wert
		if value >= 128 {
			m.Set(x, y, color.RGBA{255, 255, 255, 255})
			//verbleibende negative Differenz wird berechnet
			difference = float32(-(255 - float32(value)))
		} else {
			m.Set(x, y, color.RGBA{0, 0, 0, 255})
			//verbleibende positive Differenz wird berechnet
			difference = float32(value)
		}
		//Aufteilung der Differenz nach Floyd-Steinberg auf umliegende Pixel
		if x < bounds.Max.X-1 {
			// x+1, y = 7/16
			differenceOfPixel[y][x+1] = differenceOfPixel[y][x+1] + difference*7/16
		}
		if y < bounds.Max.Y-1 {
			if x < bounds.Max.X-1 {
				// x+1, y+1 = 1/16
				differenceOfPixel[y+1][x+1] = differenceOfPixel[y+1][x+1] + difference*1/16
			}
			if x > 0 {
				// x-1, y+1 = 3/16
				differenceOfPixel[y+1][x-1] = differenceOfPixel[y+1][x-1] + difference*3/16
			}
			// x, y+1 = 5/16
			differenceOfPixel[y+1][x] = differenceOfPixel[y+1][x] + difference*5/16
		}
	}
	return m
}

//Funktion, die jeweils eine Zeile mit Algorithmus 2 transformiert
func (t transformSeq) transformLineAlgorithm2(m *image.RGBA, y int, bounds image.Rectangle, differenceOfPixel [][]float32) *image.RGBA {
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		var difference float32
		r, g, b, _ := t.pic.At(x, y).RGBA()
		// Einberechnug der bereits errechneten Differenzen von umliegenden Pixel
		r = checkValueOfPixel(r/256, differenceOfPixel[y][x])
		g = checkValueOfPixel(g/256, differenceOfPixel[y][x])
		b = checkValueOfPixel(b/256, differenceOfPixel[y][x])
		// Maximalwert: jeweils 256
		value := (r + g + b) / 3
		//Setzen eines neuen Farbwertes für Pixel, abhängig von derzeitigem Wert
		if value >= 128 {
			m.Set(x, y, color.RGBA{255, 255, 255, 255})
			//verbleibende negative Differenz wird berechnet
			difference = float32(-(255 - float32(value)))
		} else {
			m.Set(x, y, color.RGBA{0, 0, 0, 255})
			//verbleibende positive Differenz wird berechnet
			difference = float32(value)
		}
		//Aufteilung der Differenz nach Algorithmus 2 auf umliegende Pixel
		if x < bounds.Max.X-1 {
			// x+1, y = 4/12 = 1/3
			differenceOfPixel[y][x+1] = differenceOfPixel[y][x+1] + difference/3
		}
		if x < bounds.Max.X-2 {
			//x+2, y = 1/12
			differenceOfPixel[y][x+2] = differenceOfPixel[y][x+2] + difference/12
		}
		if y < bounds.Max.Y-1 {
			if x < bounds.Max.X-1 {
				// x+1, y+1 = 1/12
				differenceOfPixel[y+1][x+1] = differenceOfPixel[y+1][x+1] + difference/12
			}
			if x > 0 {
				// x-1, y+1 = 1/12
				differenceOfPixel[y+1][x-1] = differenceOfPixel[y+1][x-1] + difference/12
			}
			// x, y+1 = 4/12 = 1/3
			differenceOfPixel[y+1][x] = differenceOfPixel[y+1][x] + difference/3
		}
		if y < bounds.Max.Y-2 {
			// x, y+2 = 1/12
			differenceOfPixel[y+2][x] = differenceOfPixel[y+2][x] + difference/12
		}
	}
	return m
}

//Funktion, die jeweils eine Zeile mit Algorithmus 3 transformiert
func (t transformSeq) transformLineAlgorithm3(m *image.RGBA, y int, bounds image.Rectangle, differenceOfPixel [][]float32) *image.RGBA {
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		var difference float32
		r, g, b, _ := t.pic.At(x, y).RGBA()
		// Einberechnug der bereits errechneten Differenzen von umliegenden Pixel
		r = checkValueOfPixel(r/256, differenceOfPixel[y][x])
		g = checkValueOfPixel(g/256, differenceOfPixel[y][x])
		b = checkValueOfPixel(b/256, differenceOfPixel[y][x])
		// Maximalwert: jeweils 256
		value := (r + g + b) / 3
		//Setzen eines neuen Farbwertes für Pixel, abhängig von derzeitigem Wert
		if value >= 128 {
			m.Set(x, y, color.RGBA{255, 255, 255, 255})
			//verbleibende negative Differenz wird berechnet
			difference = float32(-(255 - float32(value)))
		} else {
			m.Set(x, y, color.RGBA{0, 0, 0, 255})
			//verbleibende positive Differenz wird berechnet
			difference = float32(value)
		}
		//Aufteilung der Differenz nach Algorithmus 3 auf umliegende Pixel
		if x < bounds.Max.X-1 {
			// x+1, y = 8/42 = 4/21
			differenceOfPixel[y][x+1] = differenceOfPixel[y][x+1] + difference*4/21
		}
		if x < bounds.Max.X-2 {
			// x+2, y   = 4/42 = 2/21
			differenceOfPixel[y][x+2] = differenceOfPixel[y][x+2] + difference*2/21
		}
		if y < bounds.Max.Y-1 {
			if x < bounds.Max.X-1 {
				// x+1, y+1 = 4/42 = 2/21
				differenceOfPixel[y+1][x+1] = differenceOfPixel[y+1][x+1] + difference*2/21
			}
			if x < bounds.Max.X-2 {
				// x+2, y+1   = 2/42 = 1/21
				differenceOfPixel[y+1][x+2] = differenceOfPixel[y+1][x+2] + difference*1/21
			}
			if x > 0 {
				// x-1, y+1 = 4/42 = 2/21
				differenceOfPixel[y+1][x-1] = differenceOfPixel[y+1][x-1] + difference*2/21
			}
			if x > 1 {
				// x-2, y+1 = 2/42 = 1/21
				differenceOfPixel[y+1][x-2] = differenceOfPixel[y+1][x-2] + difference*1/21
			}
			// x, y+1 = 8/42 = 4/21
			differenceOfPixel[y+1][x] = differenceOfPixel[y+1][x] + difference*4/21
		}
		if y < bounds.Max.Y-2 {
			if x < bounds.Max.X-1 {
				// x+1, y+2 = 2/42 = 1/21
				differenceOfPixel[y+2][x+1] = differenceOfPixel[y+2][x+1] + difference*1/21
			}
			if x < bounds.Max.X-2 {
				// x+2, y+2   = 1/42
				differenceOfPixel[y+2][x+2] = differenceOfPixel[y+2][x+2] + difference*1/42
			}
			if x > 0 {
				// x-1, y+2 = 2/42 = 1/21
				differenceOfPixel[y+2][x-1] = differenceOfPixel[y+2][x-1] + difference*1/21
			}
			if x > 1 {
				// x-2, y+2 = 1/42
				differenceOfPixel[y+2][x-2] = differenceOfPixel[y+2][x-2] + difference*1/42
			}
			// x, y+2 = 4/42 = 2/21
			differenceOfPixel[y+2][x] = differenceOfPixel[y+2][x] + difference*2/21
		}
	}
	return m
}

//Funktion, die jeweils eine Zeile mit Graustufen transformiert
func (t transformSeq) transformLineGraustufen(m *image.RGBA, y int, bounds image.Rectangle, differenceOfPixel [][]float32) *image.RGBA {
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		r, g, b, _ := t.pic.At(x, y).RGBA()
		// Einberechnug der bereits errechneten Differenzen von umliegenden Pixel
		r = checkValueOfPixel(r/256, differenceOfPixel[y][x])
		g = checkValueOfPixel(g/256, differenceOfPixel[y][x])
		b = checkValueOfPixel(b/256, differenceOfPixel[y][x])
		// Maximalwert: jeweils 256
		value := uint8((r + g + b) / 3)
		//Setzen eines neuen Farbwertes für Pixel, abhängig von derzeitigem Wert
		m.Set(x, y, color.RGBA{value, value, value, 255})
	}
	return m
}
