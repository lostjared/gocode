package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"fmt"
	"strconv"
)

func procPixel(col color.Color, alpha float32) color.Color {
	r,g,b,a := col.RGBA()
	rf := float32(r)
	gf := float32(g)
	bf := float32(b)
	rf += (alpha*rf)
	gf += (alpha*gf)
	bf += (alpha*bf)
	return color.RGBA{uint8(rf), uint8(gf), uint8(bf), uint8(a)}

}

func FilterImage(im image.Image, alpha float32) image.Image {
	s := im.Bounds()
	nw := image.NewRGBA(image.Rect(0,0,s.Max.X, s.Max.Y))

	for x := 0; x < s.Max.X; x++ {
		for y := 0; y < s.Max.Y; y++ {
			col := im.At(x, y)
			value := procPixel(col, alpha)
			nw.Set(x,y, value)
		}
	}
	return nw
}


func main() {
	fmt.Println("Test Filter v1.0\nWritten by Jared Bruni")
	if (len(os.Args) != 4) {
		fmt.Fprintln(os.Stderr, "Program requires three arguments\ninput output iterations")
		os.Exit(1)
	}
	fn, err := os.Create(os.Args[2])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error could not create file: ", err)
		os.Exit(1)
	}
	ifn, err := os.Open(os.Args[1])

	if (err != nil) {
		fmt.Fprintln(os.Stderr, "Error could not open file: ", err)
	}

	alpha, err := strconv.Atoi(os.Args[3])

	if (err != nil) {
		fmt.Fprintln(os.Stderr, "Error converting to float: ", err)
		os.Exit(1)
	}

	alpha_value := float32(alpha)

	img, err := png.Decode(ifn)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error could not open file: ", err)
		os.Exit(1)
	}
	filtered_image := FilterImage(img,alpha_value)
	png.Encode(fn, filtered_image)
	fmt.Println("Image outputted to: ", os.Args[2])
}