package main

import (
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

type transformSeq struct {
	pic image.Image
}

//Funktion, die Bild ohne Filter transformiert
func (t transformSeq) transformWithoutFilter(input string) bool {
	bounds := t.pic.Bounds()
	i := strings.Index(input, ".")
	name := input[:i]
	format := input[i:]
	newPic, _ = os.Create("pictures/" + name + "_schwellwert" + format)
	defer newPic.Close()
	m := image.NewRGBA(image.Rectangle{Min: image.Point{0, 0}, Max: image.Point{bounds.Max.X, bounds.Max.Y}})
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := t.pic.At(x, y).RGBA()
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
		jpeg.Encode(newPic, m, nil)
	}
	return true
}

//Funktion, die Bild nach Filter von Floyd / Steinberg transformiert
func (t transformSeq) transformWithFloydSteinberg(input string) bool {
	bounds := t.pic.Bounds()
	i := strings.Index(input, ".")
	name := input[:i]
	format := input[i:]
	newPic, _ = os.Create("pictures/" + name + "_floydsteinberg" + format)
	defer newPic.Close()
	m := image.NewRGBA(image.Rectangle{Min: image.Point{0, 0}, Max: image.Point{bounds.Max.X, bounds.Max.Y}})
	// Array, welches Differenzen der umliegenden Pixel speichert
	differenceOfPixel := make([][]float32, bounds.Max.Y)
	for element := range differenceOfPixel {
		differenceOfPixel[element] = make([]float32, bounds.Max.X)
	}
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			var difference float32
			r, g, b, _ := t.pic.At(x, y).RGBA()
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
				if x > 0 {
					// x-1, y+1 = 3/16
					differenceOfPixel[y+1][x-1] = differenceOfPixel[y+1][x-1] + difference*3/16
				}
				// x, y+1 = 5/16
				differenceOfPixel[y+1][x] = differenceOfPixel[y+1][x] + difference*5/16
			}
		}
	}
	if format == ".png" {
		png.Encode(newPic, m)
	} else {
		jpeg.Encode(newPic, m, nil)
	}
	return true
}

//Funktion, die Bild mit Algorithmus 2 transformiert
func (t transformSeq) transformWithAlgorithm2(input string) bool {
	bounds := t.pic.Bounds()
	i := strings.Index(input, ".")
	name := input[:i]
	format := input[i:]
	newPic, _ = os.Create("pictures/" + name + "_algorithm2" + format)
	defer newPic.Close()
	m := image.NewRGBA(image.Rectangle{Min: image.Point{0, 0}, Max: image.Point{bounds.Max.X, bounds.Max.Y}})
	// Array, welches Differenzen der umliegenden Pixel speichert
	differenceOfPixel := make([][]float32, bounds.Max.Y)
	for element := range differenceOfPixel {
		differenceOfPixel[element] = make([]float32, bounds.Max.X)
	}
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			var difference float32
			r, g, b, _ := t.pic.At(x, y).RGBA()
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
	}
	if format == ".png" {
		png.Encode(newPic, m)
	} else {
		jpeg.Encode(newPic, m, nil)
	}
	return true
}

//Funktion, die Bild mit Algorithmus 3 transformiert
func (t transformSeq) transformWithAlgorithm3(input string) bool {
	bounds := t.pic.Bounds()
	i := strings.Index(input, ".")
	name := input[:i]
	format := input[i:]
	newPic, _ = os.Create("pictures/" + name + "_algorithm3" + format)
	defer newPic.Close()
	m := image.NewRGBA(image.Rectangle{Min: image.Point{0, 0}, Max: image.Point{bounds.Max.X, bounds.Max.Y}})
	// Array, welches Differenzen der umliegenden Pixel speichert
	differenceOfPixel := make([][]float32, bounds.Max.Y)
	for element := range differenceOfPixel {
		differenceOfPixel[element] = make([]float32, bounds.Max.X)
	}
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			var difference float32
			r, g, b, _ := t.pic.At(x, y).RGBA()
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
	}
	if format == ".png" {
		png.Encode(newPic, m)
	} else {
		jpeg.Encode(newPic, m, nil)
	}
	return true
}
