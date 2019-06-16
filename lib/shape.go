package lib

import (
	"fmt"
	"golang.org/x/image/colornames"
	"image"
	"image/color"
	"math"
)

type Rect struct {
	Min image.Point
	Max image.Point
}

type Circle struct {
	Radius   int
	Position image.Point
}

func RectsCollision(a Rect, b Rect) bool {
	if a.Max.X < b.Min.X || a.Min.X > b.Max.X {
		return false
	}
	if a.Max.Y < b.Min.Y || a.Min.Y > b.Max.Y {
		return false
	}

	return true
}

func CirclesCollision(a Circle, b Circle) bool {
	r := a.Radius + b.Radius + 1
	r *= r
	pow1 := math.Pow(float64(a.Position.X-b.Position.X), 2)
	pow2 := math.Pow(float64(a.Position.Y-b.Position.Y), 2)
	return r < (int(pow1 + pow2))
}

func Line(x0 float64, y0 float64, x1 float64, y1 float64, img *image.RGBA, color color.Color) {
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

func FillRect(img *image.RGBA, rect image.Rectangle, color color.Color) {
	for x := rect.Min.X; x < rect.Max.X; x++ {
		for y := rect.Min.Y; y < rect.Max.Y; y++ {
			img.Set(x, y, color)
		}
	}
}

func Rotate90(img *image.RGBA) (new *image.RGBA) {
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

func Rotate180(img *image.RGBA) (new *image.RGBA) {
	new = image.NewRGBA(image.Rect(0, 0, img.Bounds().Max.X, img.Bounds().Max.Y))
	for x := 0; x < img.Bounds().Max.X; x++ {
		for y := 0; y < img.Bounds().Max.Y; y++ {
			new.Set(x, y, img.At(img.Bounds().Max.X-1-x, img.Bounds().Max.Y-1-y))
		}
	}
	return
}

func Rotate270(img *image.RGBA) (new *image.RGBA) {
	new = image.NewRGBA(image.Rect(0, 0, img.Bounds().Max.X, img.Bounds().Max.Y))
	for x := 0; x < img.Bounds().Max.X; x++ {
		for y := 0; y < img.Bounds().Max.Y; y++ {
			new.Set(x, y, img.At(img.Bounds().Max.X-1-x, y))
		}
	}
	return
}
func DrawRect(min image.Point, max image.Point, img *image.RGBA) (new *image.RGBA) {
	for x := min.X; x <= max.X; x++ {
		for y := min.Y; y <= max.Y; y++ {
			img.Set(x, y, colornames.Green)
		}
	}
	return
}

func DrawCircle(circle Circle, img *image.RGBA, color color.Color) {
	x := 0
	y := circle.Radius
	delta := 1 - 2*circle.Radius
	errorDraw := 0
	for y >= 0 {
		img.Set(circle.Position.X+x, circle.Position.Y+y, color)
		img.Set(circle.Position.X+x, circle.Position.Y-y, color)
		img.Set(circle.Position.X-x, circle.Position.Y+y, color)
		img.Set(circle.Position.X-x, circle.Position.Y-y, color)
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
