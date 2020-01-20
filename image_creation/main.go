package main

import (
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"os"
)

func main() {
	min := image.Point{0, 0}
	max := image.Point{100, 100}
	img := image.NewRGBA(image.Rectangle{min, max})

	f, err := os.Create("TestFile.jpg")
	if err != nil {
		log.Fatal("Couldn't create a file...")
	}
	defer f.Close()

	for h := 0; h < 100; h++ {
		for x := 100 - h; x > 0; x-- {
			img.Set(x, h, color.RGBA{255, 0 , 0, 1})
		}
	}

	jpeg.Encode(f, img, &jpeg.Options{})
}
