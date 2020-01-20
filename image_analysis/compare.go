package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"math"
	"os"
)

const threshold = 100

func compare(haystack, needle *jpgImage) bool {
	if haystack.height < needle.height {
		return false
	}
	if haystack.width < needle.width {
		return false
	}
	if haystack.path == needle.path {
		return false
	}

	return compareImages(haystack, needle)
}

func compareImages(haystack, needle *jpgImage) bool {
	hPxls := haystack.pixels
	nPxls := needle.pixels
	maxYToSearch := haystack.height - needle.height + 1
	maxXToSearch := haystack.width - needle.width + 1

	for hy := 0; hy < maxYToSearch; hy++ {
		for hx := 0; hx < maxXToSearch; hx++ {
			location := hy*haystack.width + hx
			diff := comparePxls(nPxls[0], hPxls[location])

			if diff < threshold {
				result := comparisonSearch(haystack, needle, location)
				fmt.Printf("Result at %dx%d is %v\n", hx, hy, result)
			}
		}
	}

	return false
}

type imgComparisonResult struct {
	haystack, needle *jpgImage
	startX, startY   int
	found            bool
	diff             float64
}

func comparisonSearch(haystack *jpgImage, needle *jpgImage, start int) *imgComparisonResult {
	hPxls := haystack.pixels
	nPxls := needle.pixels
	startX := start % haystack.width
	startY := start / haystack.width
	endY := startY + needle.height
	endX := startX + needle.width
	var rowDiff float64
	var totalDiff float64

	imgName := fmt.Sprintf("imgComp(%vx%v)-%v", startX, startY, haystack.path)
	f, err := os.Create(imgName)
	defer f.Close()
	if err != nil {
		log.Fatal("Could not create file: ", imgName)
	}

	min := image.Point{0, 0}
	max := image.Point{needle.width, needle.height * 2}
	img := image.NewRGBA(image.Rectangle{min, max})

	for hy := startY; hy < endY; hy++ {
		for hx := startX; hx < endX; hx++ {
			hLoc := hy*haystack.width + hx
			nLoc := (hy-startY)*needle.width + (hx - startX)
			compRes := comparePxls(nPxls[nLoc], hPxls[hLoc])
			if startY == 250 {
				compRes = comparePxls(nPxls[nLoc], hPxls[hLoc])
			}
			rowDiff += compRes

			nColor := color.RGBA64{
				R: uint16(nPxls[nLoc].r),
				G: uint16(nPxls[nLoc].g),
				B: uint16(nPxls[nLoc].b),
				A: uint16(nPxls[nLoc].a),
			}
			hColor := color.RGBA64{
				R: uint16(hPxls[hLoc].r),
				G: uint16(hPxls[hLoc].g),
				B: uint16(hPxls[hLoc].b),
				A: uint16(hPxls[hLoc].a),
			}
			img.Set(hx-startX, hy-startY, nColor)
			img.Set(hx-startX, needle.height + hy-startY, hColor)
		}

		totalDiff += rowDiff
		rowDiff = 0
	}

	jpeg.Encode(f, img, &jpeg.Options{Quality: 100})

	if err != nil {
		fmt.Println("Got error encoding image: ", err)
	}

	avgDiff := math.Round(totalDiff / float64(needle.width * needle.height))

	return &imgComparisonResult{
		haystack: haystack,
		needle:   needle,
		startX:   startX,
		startY:   startY,
		found:    true,
		diff:     avgDiff,
	}
}

func comparePxls(p1 pixel, p2 pixel) float64 {
	var diff float64

	diff += math.Abs(float64(int(p1.r) - int(p2.r)))
	diff += math.Abs(float64(int(p1.g) - int(p2.g)))
	diff += math.Abs(float64(int(p1.b) - int(p2.b)))
	diff += math.Abs(float64(int(p1.a) - int(p2.a)))

	return diff
}
