package main

import (
	"fmt"
	"game/lib"
	"golang.org/x/image/colornames"
	"image"
	"image/png"
	"os"
)

func main() {
	rect := image.Rect(0, 0, 100, 100)
	img := image.NewRGBA(rect)
	lib.FillRect(img, rect, colornames.White)

	circle1 := lib.Circle{Position: image.Point{X: 5, Y: 20}, Radius: 15}
	circle2 := lib.Circle{Position: image.Point{X: 20, Y: 5}, Radius: 5}
	lib.DrawCircle(circle1, img, colornames.Green)
	lib.DrawCircle(circle2, img, colornames.Red)
	fmt.Println(lib.CirclesCollision(circle1, circle2))
	img = lib.Rotate90(img)
	f, _ := os.Create("image.png")
	png.Encode(f, img)
}
