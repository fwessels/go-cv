package gocv

import (
	"github.com/fwessels/go-cv-simd/sse2"
)

func filter(src gocvsimd.View, f func(src, dst gocvsimd.View)) (dst gocvsimd.View) {

	dst = gocvsimd.View{}
	dst.Recreate(src.GetWidth(), src.GetHeight(), src.GetFormat())

	gocvsimd.SimdSse2GaussianBlur3x3(src, dst)

	return dst
}

func GaussianBlur3x3(src gocvsimd.View) (dst gocvsimd.View) {

	return filter(src, gocvsimd.SimdSse2GaussianBlur3x3)
}

func MeanFilter3x3(src gocvsimd.View) (dst gocvsimd.View) {

	return filter(src, gocvsimd.SimdSse2MeanFilter3x3)
}

func MedianFilterRhomb3x3(src gocvsimd.View) (dst gocvsimd.View) {

	return filter(src, gocvsimd.SimdSse2MedianFilterRhomb3x3)
}

func MedianFilterSquare3x3(src gocvsimd.View) (dst gocvsimd.View) {

	return filter(src, gocvsimd.SimdSse2MedianFilterSquare3x3)
}

func MedianFilterRhomb5x5(src gocvsimd.View) (dst gocvsimd.View) {

	return filter(src, gocvsimd.SimdSse2MedianFilterRhomb5x5)
}

func MedianFilterSquare5x5(src gocvsimd.View) (dst gocvsimd.View) {

	return filter(src, gocvsimd.SimdSse2MedianFilterSquare5x5)
}

func Laplace(src gocvsimd.View) (dst gocvsimd.View) {

	return filter(src, gocvsimd.SimdSse2Laplace)
}

func SobelDx(src gocvsimd.View) (dst gocvsimd.View) {

	return filter(src, gocvsimd.SimdSse2SobelDx)
}

func SobelDy(src gocvsimd.View) (dst gocvsimd.View) {

	return filter(src, gocvsimd.SimdSse2SobelDy)
}

//
func ContourAnchors(src gocvsimd.View, step uint64, threshold int16, dst gocvsimd.View) {

	gocvsimd.SimdSse2ContourAnchors(src, step, uint64(threshold), dst)
}
