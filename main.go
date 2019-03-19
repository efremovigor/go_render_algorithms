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

type Rect struct {
	min image.Point
	max image.Point
}

type Circle struct {
	radius   int
	position image.Point
}

func RectsCollision(a Rect, b Rect) bool {
	if a.max.X < b.min.X || a.min.X > b.max.X {
		return false
	}
	if a.max.Y < b.min.Y || a.min.Y > b.max.Y {
		return false
	}

	return true
}

func CirclesCollision(a Circle, b Circle) bool {
	r := a.radius + b.radius + 1
	r *= r
	pow1 := math.Pow(float64(a.position.X-b.position.X), 2)
	pow2 := math.Pow(float64(a.position.Y-b.position.Y), 2)
	return r < (int(pow1 + pow2))
}

func main() {
	rect := image.Rect(0, 0, 100, 100)
	img := image.NewRGBA(rect)
	fillRect(img, rect, colornames.White)

	circle1 := Circle{position: image.Point{X: 5, Y: 20}, radius: 15}
	circle2 := Circle{position: image.Point{X: 20, Y: 5}, radius: 5}
	drawCircle(circle1, img, colornames.Green)
	drawCircle(circle2, img, colornames.Red)
	fmt.Println(CirclesCollision(circle1, circle2))
	//img = rotate90(img)
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
	errorDraw := math.Abs(dy) * 2
	errorDraw2 := 0.0
	y := y0

	for x := x0; x <= x1; x++ {
		if steep {
			img.Set(int(y), int(x), color) // if transposed, de-transpose
		} else {
			img.Set(int(x), int(y), color)
		}
		errorDraw2 += errorDraw

		if errorDraw2 > dx {
			if y1 > y0 {
				y += 1
			} else {
				y -= 1
			}
			errorDraw2 -= dx * 2
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
func drawRect(min image.Point, max image.Point, img *image.RGBA) (new *image.RGBA) {
	for x := min.X; x <= max.X; x++ {
		for y := min.Y; y <= max.Y; y++ {
			img.Set(x, y, colornames.Green)
		}
	}
	return
}

func drawCircle(circle Circle, img *image.RGBA, color color.Color) {
	x := 0
	y := circle.radius
	delta := 1 - 2*circle.radius
	errorDraw := 0
	for y >= 0 {
		img.Set(circle.position.X+x, circle.position.Y+y, color)
		img.Set(circle.position.X+x, circle.position.Y-y, color)
		img.Set(circle.position.X-x, circle.position.Y+y, color)
		img.Set(circle.position.X-x, circle.position.Y-y, color)
		errorDraw = 2*(delta+y) - 1
		if (delta < 0) && (errorDraw <= 0) {
			x += 1
			delta += 2*x + 1
			continue
		}

		if delta > 0 && errorDraw > 0 {
			y -= 1
			delta -= 2*y + 1
			continue
		}
		x += 1
		delta += 2 * (x - y)
		y -= 1
	}

}
