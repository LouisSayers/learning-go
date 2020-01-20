package main

import (
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

type pixel struct {
	r, g, b, a uint32
}

//jpgImage Represents a JPEG image
type jpgImage struct {
	path   string
	pixels []pixel
	width  int
	height int
}

func retrieveImagePaths(path string) []string {
	var filePaths []string

	filepath.Walk("images", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("Got error: ", err.Error())
			return err
		} else if info.IsDir() {
			fmt.Println("Got a directory!")
		} else {
			filePaths = append(filePaths, path)
		}
		return nil
	})

	return filePaths
}

func loadImage(path string) image.Image {
	var img image.Image
	f, err := os.Open(path)
	defer f.Close()

	if err != nil {
		fmt.Println("Got Error: ", err)
	} else {
		img, err = jpeg.Decode(f)
		if err != nil {
			fmt.Println("Couldn't decode: ", err)
		}
	}

	return img
}

func newjpgImage(imgPath string) (*jpgImage, error) {
	var newImage jpgImage
	var err error
	img := loadImage(imgPath)

	if img == nil {
		err = errors.New("Could not load image")
	} else {
		imgNameParts := strings.Split(imgPath, "/")
		imgName := imgNameParts[len(imgNameParts)-1]

		bounds := img.Bounds()
		width := bounds.Dx()
		height := bounds.Dy()
		totalPixels := width * height
		pixels := make([]pixel, totalPixels)

		for h := 0; h < height; h++ {
			for w := 0; w < width; w++ {
				currentIndex := h*width + w
				currentPx := img.At(w, h)
				r, g, b, a := currentPx.RGBA()
				pixels[currentIndex] = pixel{r, g, b, a}
			}
		}
		newImage = jpgImage{imgName, pixels, width, height}
	}
	return &newImage, err
}

func getImages(imagesPath string) chan *jpgImage {
	imagePaths := retrieveImagePaths(imagesPath)
	ch := make(chan *jpgImage)
	var wg sync.WaitGroup

	for _, imgPath := range imagePaths {
		wg.Add(1)
		go func(path string) {
			defer wg.Done()
			newImage, err := newjpgImage(path)
			if err == nil {
				ch <- newImage
			}
		}(imgPath)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	return ch
}

func xImages(ch chan *jpgImage) []*jpgImage {
	var images []*jpgImage
	for img := range ch {
		images = append(images, img)
	}

	return images
}

func main() {
	start := time.Now()
	ch := getImages("images")
	images := xImages(ch)
	finish := time.Since(start)

	for _, img := range images {
		for _, needle := range images {
			same := compare(img, needle)
			fmt.Printf("Images are the same? %t\n", same)
		}
	}
	fmt.Println("Time taken: ", finish)
}
