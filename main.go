package main

import (
	"fmt"
	"golang.org/x/image/colornames"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

func main() {
	rect := image.Rect(0, 0, 100, 100)
	img := image.NewRGBA(rect)
	fillRect(img, rect, colornames.Black)
	line(13, 20, 80, 40, img, colornames.White)
	line(20, 13, 40, 80, img, colornames.Red)
	line(80, 40, 13, 20, img, colornames.Red)

	img = rotate90(img)

	f, _ := os.Create("image.png")
	png.Encode(f, img)
}

func line(x0 float64, y0 float64, x1 float64, y1 float64, img *image.RGBA, color color.Color) {
	steep := false
	if math.Abs(x0-x1) < math.Abs(y0-y1) {
		x0, y0 = y0, x0
		x1, y1 = y1, x1
		steep = true
	}
	if x0 > x1 { // make it left-to-right
		x0, x1 = x1, x0
		y0, y1 = y1, y0
	}

	dx := x1 - x0
	dy := y1 - y0
	derror2 := math.Abs(dy)*2
	error2 := 0.0
	y := y0

	for x := x0; x <= x1; x++ {
		if steep {
			img.Set(int(y), int(x), color) // if transposed, de-transpose
		} else {
			img.Set(int(x), int(y), color)
		}
		error2 += derror2

		if error2 > dx {
			if y1>y0 {
				y += 1
			} else {
				y -= 1
			}
			error2 -= dx*2
		}
	}
}

func fillRect(img *image.RGBA, rect image.Rectangle, color color.Color) {
	for x := rect.Min.X; x < rect.Max.X; x++ {
		for y := rect.Min.Y; y < rect.Max.Y; y++ {
			img.Set(x, y, color)
		}
	}
}

func rotate90(img *image.RGBA) (new *image.RGBA) {
	new = image.NewRGBA(image.Rect(0, 0, img.Bounds().Max.X, img.Bounds().Max.Y))
	for x := 0; x <= img.Bounds().Max.X; x++ {
		for y := 0; y <= img.Bounds().Max.Y; y++ {
			new.Set(x, y, img.At(x, img.Bounds().Max.Y-1-y))
		}

	}
	fmt.Println(img.Bounds().Min.X)
	fmt.Println(img.Bounds().Max.Y)
	return
}

func rotate180(img *image.RGBA) (new *image.RGBA) {
	new = image.NewRGBA(image.Rect(0, 0, img.Bounds().Max.X, img.Bounds().Max.Y))
	for x := 0; x < img.Bounds().Max.X; x++ {
		for y := 0; y < img.Bounds().Max.Y; y++ {
			new.Set(x, y, img.At(img.Bounds().Max.X-1-x, img.Bounds().Max.Y-1-y))
		}
	}
	return
}

func rotate270(img *image.RGBA) (new *image.RGBA) {
	new = image.NewRGBA(image.Rect(0, 0, img.Bounds().Max.X, img.Bounds().Max.Y))
	for x := 0; x < img.Bounds().Max.X; x++ {
		for y := 0; y < img.Bounds().Max.Y; y++ {
			new.Set(x, y, img.At(img.Bounds().Max.X-1-x, y))
		}
	}
	return
}
