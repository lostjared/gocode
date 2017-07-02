package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"math/cmplx"
	"fmt"
	"strconv"
)

func mandlebrot(z complex128, r uint8, g uint8, b uint8) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{r-contrast*n, g-contrast*n, b-contrast*n, 255}
		}
	}
	return color.Black
}



func main() {
	if len(os.Args) != 5 {
		fmt.Fprintln(os.Stderr, "Error takes 5 arguments, filename r g b")
		os.Exit(1)
	}
	fn, err := os.Create(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error could not create file: ", os.Args[1])
		os.Exit(1)
	}


	r, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: ", err)
		os.Exit(1)
	}
	g, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: ", err)
		os.Exit(1)
	}

	b, err := strconv.Atoi(os.Args[4])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: ", err)
		os.Exit(1)
	}

	rx := uint8(r)
	gx := uint8(g)
	bx := uint8(b)

	fmt.Printf("RGB Values: %d %d %d\n", rx,gx,bx)

	const (
		xmin, ymin, xmax, ymax = -2,-2,+2,+2
		width = 1280
		height = 720
	)

	image := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x,y)

			image.Set(px, py, mandlebrot(z, rx, gx, bx))
		}
	}
	png.Encode(fn, image)
	fmt.Println("Wrote image to file: ", os.Args[1])

}
