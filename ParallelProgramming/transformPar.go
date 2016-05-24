package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"os"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

//Definition der benötigten globalen Variablen
var m *image.RGBA
var wg sync.WaitGroup
var barrier sync.WaitGroup
var countThreads int32
var differenceOfPixel [][]int32

//Objekte dieser Klasse besitzen Bild als Variable
type transformPar struct {
	pic image.Image
}

//Funktion, die Bild einliest
func (t transformPar) transformParallel(input, method string) bool {
	bounds := t.pic.Bounds()
	//Aufteilen des Inputs in Name und Dateiformat
	i := strings.Index(input, ".")
	name := input[:i]
	format := input[i:]
	//Erstellen des neuen Bildes mit entsprechendem Namen und Dateiformat
	newPic, _ = os.Create("pictures/" + name + "_" + method + "_par" + strconv.Itoa(getGOMAXPROCS()) + format)
	defer newPic.Close()
	m = image.NewRGBA(image.Rectangle{Min: image.Point{0, 0}, Max: image.Point{bounds.Max.X, bounds.Max.Y}})
	// 2-dimensionales Array, welches Differenzen der umliegenden Pixel speichert
	differenceOfPixel = make([][]int32, bounds.Max.Y)
	for element := range differenceOfPixel {
		differenceOfPixel[element] = make([]int32, bounds.Max.X)
	}
	//Channel, um Bearbeitungsreihenfolge der Threads einzuhalten
	order := make(chan int)
	go func() {
		order <- 0
	}()
	//WaitGroup wird auf Anzahl der Zeilen gesetzt, erst wenn alle Zeilen "Done" sind, endet Methode
	wg.Add(bounds.Max.Y)
	//Anzahl der laufenden Threads wird gezählt
	countThreads = 0
	//zwei for-Schleifen, um jeden Pixelwert auszulesen
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		//für jede Zeile eigene Go-Routine
		//erst wenn ersten drei Pixel transformiert wurden, darf nächste Zeile transformiert werden <------------------------- funktioniert nicht
		if y == <-order {
			//WaitGroup, die Zeilenbearbeitung synchronisiert
			barrier.Add(1)
			switch method {
			case "FloydSteinberg":
				go t.transformLineFloydSteinberg(y, bounds, order, &differenceOfPixel, &countThreads)
			case "Algorithm2":
				go t.transformLineAlgorithm2(y, bounds, order, &differenceOfPixel, &countThreads)
			case "Algorithm3":
				go t.transformLineAlgorithm3(y, bounds, order, &differenceOfPixel, &countThreads)
			case "Schwellwert":
				go t.transformLineSchwellwert(y, bounds, order)
			case "Graustufen":
				go t.transformLineGraustufen(y, bounds, order)
			}
		}
	}
	//falls alle Zeilen transformiert wurden
	wg.Wait()
	close(order)
	if format == ".png" {
		png.Encode(newPic, m)
	} else {
		jpeg.Encode(newPic, m, nil)
	}
	return true
}

//Funktion, die jeweils eine Zeile mit Floyd-Steinberg-Algorithmus transformiert
func (t transformPar) transformLineFloydSteinberg(y int, bounds image.Rectangle, order chan<- int, pDifferenceOfPixel *[][]int32, pCountThreads *int32) {
	atomic.AddInt32(pCountThreads, 1)
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		//Synchronisation über WaitGroup barrier, wenn Thread 1 bei 10 px, Thread 2 bei 7 px, Thread 3 bei 4 px und Thread 4 bei 1 px ist [usw]
		threadNr := int32(y + 1)
		if x == int((*pCountThreads-threadNr)*3+1) {
			if *pCountThreads != threadNr {
				fmt.Println("Sync") //<-----------------wird kaum aufgerufen!!!!!!!!!!!!!!!!!!!!!
			}
			barrier.Done() //zeigt an, dass bis zu bestimmtem Pixel bearbeitet wurde
			barrier.Wait() //wartet auf alle anderen Threads, dass diese bestimmte Pixelanzahl erreicht haben
			barrier.Add(1) //nachdem erfolgreich gewartet wurde, starten alle Threads wieder und Barriere wird erneut gesetzt
		}
		var difference int32
		value := color.GrayModel.Convert((t.pic).At(x, y)).(color.Gray).Y
		// Einberechnug der bereits errechneten Differenzen von umliegenden Pixel
		value = checkValueOfPixel(value, (*pDifferenceOfPixel)[y][x])
		//Setzen eines neuen Farbwertes für Pixel, abhängig von derzeitigem Wert
		if value >= 128 {
			m.Set(x, y, color.White)
			//verbleibende negative Differenz wird berechnet
			difference = int32(-(255 - int32(value)))
		} else {
			m.Set(x, y, color.Black)
			//verbleibende positive Differenz wird berechnet
			difference = int32(value)
		}
		//Aufteilung der Differenz nach Floyd-Steinberg auf umliegende Pixel
		if x < bounds.Max.X-1 {
			// x+1, y = 7/16
			atomic.AddInt32(&(*pDifferenceOfPixel)[y][x+1], difference*7/16)
		}
		if y < bounds.Max.Y-1 {
			if x < bounds.Max.X-1 {
				// x+1, y+1 = 1/16
				atomic.AddInt32(&(*pDifferenceOfPixel)[y+1][x+1], difference*1/16)
			}
			if x > 0 {
				// x-1, y+1 = 3/16
				atomic.AddInt32(&(*pDifferenceOfPixel)[y+1][x-1], difference*3/16)
			}
			// x, y+1 = 5/16
			atomic.AddInt32(&(*pDifferenceOfPixel)[y+1][x], difference*5/16)
		}
		if x == 3 && y < bounds.Max.Y-1 {
			order <- y + 1 //y wird erhöht, wenn drei Pixel der Zeile  durchlaufen wurden
		}
	}
	barrier.Done()
	wg.Done()
}

//Funktion, die jeweils eine Zeile mit Algorithmus 2 transformiert
func (t transformPar) transformLineAlgorithm2(y int, bounds image.Rectangle, order chan<- int, pDifferenceOfPixel *[][]int32, pCountThreads *int32) {
	atomic.AddInt32(pCountThreads, 1)
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		//Synchronisation über WaitGroup barrier, wenn Thread 1 bei 10 px, Thread 2 bei 7 px, Thread 3 bei 4 px und Thread 4 bei 1 px ist [usw]
		threadNr := int32(y + 1)
		if x == int((*pCountThreads-threadNr)*3+1) {
			barrier.Done() //zeigt an, dass bis zu bestimmtem Pixel bearbeitet wurde
			barrier.Wait() //wartet auf alle anderen Threads, dass diese bestimmte Pixelanzahl erreicht haben
			barrier.Add(1) //nachdem erfolgreich gewartet wurde, starten alle Threads wieder und Barriere wird erneut gesetzt
		}
		var difference int32
		value := color.GrayModel.Convert((t.pic).At(x, y)).(color.Gray).Y
		// Einberechnug der bereits errechneten Differenzen von umliegenden Pixel
		value = checkValueOfPixel(value, (*pDifferenceOfPixel)[y][x])
		//Setzen eines neuen Farbwertes für Pixel, abhängig von derzeitigem Wert
		if value >= 128 {
			m.Set(x, y, color.White)
			//verbleibende negative Differenz wird berechnet
			difference = int32(-(255 - int32(value)))
		} else {
			m.Set(x, y, color.Black)
			//verbleibende positive Differenz wird berechnet
			difference = int32(value)
		}
		//Aufteilung der Differenz nach Algorithmus 2 auf umliegende Pixel
		if x < bounds.Max.X-1 {
			// x+1, y = 4/12 = 1/3
			atomic.AddInt32(&(*pDifferenceOfPixel)[y][x+1], difference/3)
		}
		if x < bounds.Max.X-2 {
			//x+2, y = 1/12
			atomic.AddInt32(&(*pDifferenceOfPixel)[y][x+2], difference/12)
		}
		if y < bounds.Max.Y-1 {
			if x < bounds.Max.X-1 {
				// x+1, y+1 = 1/12
				atomic.AddInt32(&(*pDifferenceOfPixel)[y+1][x+1], difference/12)
			}
			if x > 0 {
				// x-1, y+1 = 1/12
				atomic.AddInt32(&(*pDifferenceOfPixel)[y+1][x-1], difference/12)
			}
			// x, y+1 = 4/12 = 1/3
			atomic.AddInt32(&(*pDifferenceOfPixel)[y+1][x], difference/3)
		}
		if y < bounds.Max.Y-2 {
			// x, y+2 = 1/12
			atomic.AddInt32(&(*pDifferenceOfPixel)[y+2][x], difference/12)
		}
		if x == 3 && y < bounds.Max.Y-1 {
			order <- y + 1 //y wird erhöht, wenn drei Pixel der Zeile  durchlaufen wurden
		}
	}
	barrier.Done()
	wg.Done()
}

//Funktion, die jeweils eine Zeile mit Algorithmus 3 transformiert
func (t transformPar) transformLineAlgorithm3(y int, bounds image.Rectangle, order chan<- int, pDifferenceOfPixel *[][]int32, pCountThreads *int32) {
	atomic.AddInt32(pCountThreads, 1)
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		//Synchronisation über WaitGroup barrier, wenn Thread 1 bei 10 px, Thread 2 bei 7 px, Thread 3 bei 4 px und Thread 4 bei 1 px ist [usw]
		threadNr := int32(y + 1)
		if x == int((*pCountThreads-threadNr)*3+1) {
			barrier.Done() //zeigt an, dass bis zu bestimmtem Pixel bearbeitet wurde
			barrier.Wait() //wartet auf alle anderen Threads, dass diese bestimmte Pixelanzahl erreicht haben
			barrier.Add(1) //nachdem erfolgreich gewartet wurde, starten alle Threads wieder und Barriere wird erneut gesetzt
		}
		var difference int32
		value := color.GrayModel.Convert((t.pic).At(x, y)).(color.Gray).Y
		// Einberechnug der bereits errechneten Differenzen von umliegenden Pixel
		value = checkValueOfPixel(value, (*pDifferenceOfPixel)[y][x])
		//Setzen eines neuen Farbwertes für Pixel, abhängig von derzeitigem Wert
		if value >= 128 {
			m.Set(x, y, color.White)
			//verbleibende negative Differenz wird berechnet
			difference = int32(-(255 - int32(value)))
		} else {
			m.Set(x, y, color.Black)
			//verbleibende positive Differenz wird berechnet
			difference = int32(value)
		}
		//Aufteilung der Differenz nach Algorithmus 3 auf umliegende Pixel
		if x < bounds.Max.X-1 {
			// x+1, y = 8/42 = 4/21
			atomic.AddInt32(&(*pDifferenceOfPixel)[y][x+1], difference*4/21)
		}
		if x < bounds.Max.X-2 {
			// x+2, y   = 4/42 = 2/21
			atomic.AddInt32(&(*pDifferenceOfPixel)[y][x+2], difference*2/21)
		}
		if y < bounds.Max.Y-1 {
			if x < bounds.Max.X-1 {
				// x+1, y+1 = 4/42 = 2/21
				atomic.AddInt32(&(*pDifferenceOfPixel)[y+1][x+1], difference*2/21)
			}
			if x < bounds.Max.X-2 {
				// x+2, y+1   = 2/42 = 1/21
				atomic.AddInt32(&(*pDifferenceOfPixel)[y+1][x+2], difference*1/21)
			}
			if x > 0 {
				// x-1, y+1 = 4/42 = 2/21
				atomic.AddInt32(&(*pDifferenceOfPixel)[y+1][x-1], difference*2/21)
			}
			if x > 1 {
				// x-2, y+1 = 2/42 = 1/21
				atomic.AddInt32(&(*pDifferenceOfPixel)[y+1][x-2], difference*1/21)
			}
			// x, y+1 = 8/42 = 4/21
			atomic.AddInt32(&(*pDifferenceOfPixel)[y+1][x], difference*4/21)
		}
		if y < bounds.Max.Y-2 {
			if x < bounds.Max.X-1 {
				// x+1, y+2 = 2/42 = 1/21
				atomic.AddInt32(&(*pDifferenceOfPixel)[y+2][x+1], difference*1/21)
			}
			if x < bounds.Max.X-2 {
				// x+2, y+2   = 1/42
				atomic.AddInt32(&(*pDifferenceOfPixel)[y+2][x+2], difference*1/42)
			}
			if x > 0 {
				// x-1, y+2 = 2/42 = 1/21
				atomic.AddInt32(&(*pDifferenceOfPixel)[y+2][x-1], difference*1/21)
			}
			if x > 1 {
				// x-2, y+2 = 1/42
				atomic.AddInt32(&(*pDifferenceOfPixel)[y+2][x-2], difference*1/42)
			}
			// x, y+2 = 4/42 = 2/21
			atomic.AddInt32(&(*pDifferenceOfPixel)[y+2][x], difference*2/21)
		}
		if x == 3 && y < bounds.Max.Y-1 {
			order <- y + 1 //y wird erhöht, wenn drei Pixel der  Zeile  durchlaufen wurden
		}
	}
	barrier.Done()
	wg.Done()
}

//Funktion, die jeweils eine Zeile mit Schwellwert transformiert
func (t transformPar) transformLineSchwellwert(y int, bounds image.Rectangle, order chan<- int) {
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		value := color.GrayModel.Convert((t.pic).At(x, y)).(color.Gray).Y
		//Setzen eines neuen Farbwertes für Pixel, abhängig von derzeitigem Wert
		if value >= 128 {
			m.Set(x, y, color.White)
		} else {
			m.Set(x, y, color.Black)
		}
		if x == 0 && y < bounds.Max.Y-1 {
			order <- y + 1 //y wird erhöht, wenn erster Pixel der Zeile  durchlaufen wurden
		}
	}
	wg.Done()
}

//Funktion, die jeweils eine Zeile mit Graustufen transformiert
func (t transformPar) transformLineGraustufen(y int, bounds image.Rectangle, order chan<- int) {
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		value := color.GrayModel.Convert((t.pic).At(x, y)).(color.Gray).Y
		//Setzen eines neuen Farbwertes für Pixel, abhängig von derzeitigem Wert
		m.Set(x, y, color.RGBA{value, value, value, 255})
		if x == 0 && y < bounds.Max.Y-1 {
			order <- y + 1 //y wird erhöht, wenn erster Pixel der Zeile  durchlaufen wurden
		}
	}
	wg.Done()
}
