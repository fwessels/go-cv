package main

import (
	"fmt"
	"github.com/fwessels/go-cv"
	"github.com/fwessels/go-cv-simd/sse2"
	_ "image/jpeg"
)

func main() {

	bgra, _ := gocvsimd.LoadImage("data/images/lena.jpg")

	fmt.Println(bgra)

	blurred := gocv.GaussianBlur3x3(bgra)

	fmt.Println(blurred)
}