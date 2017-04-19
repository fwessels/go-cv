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

	y, u, v := gocv.BgraToYuv444p(bgra)

	fmt.Println(y)
	fmt.Println(u)
	fmt.Println(v)
}
