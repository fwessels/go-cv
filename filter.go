/*
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

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

// GaussianBlur3x3 performs Gaussian blur filtration with window 3x3.
// 	For every point:
//		dst[x, y] = (src[x-1, y-1] + 2*src[x, y-1] + src[x+1, y-1] +
//		2*(src[x-1, y] + 2*src[x, y] + src[x+1, y]) +
//		src[x-1, y+1] + 2*src[x, y+1] + src[x+1, y+1] + 8) / 16;
// All images must have the same width, height and format (8-bit gray, 16-bit UV, 24-bit BGR or 32-bit BGRA).
func GaussianBlur3x3(src gocvsimd.View) (dst gocvsimd.View) {

	return filter(src, gocvsimd.SimdSse2GaussianBlur3x3)
}
// MeanFilter3x3 performs an averaging with window 3x3.
//	For every point:
//		dst[x, y] = (src[x-1, y-1] + src[x, y-1] + src[x+1, y-1] +
//			src[x-1, y] + src[x, y] + src[x+1, y] +
//			src[x-1, y+1] + src[x, y+1] + src[x+1, y+1] + 4) / 9;
// All images must have the same width, height and format (8-bit gray, 16-bit UV, 24-bit BGR or 32-bit BGRA).
func MeanFilter3x3(src gocvsimd.View) (dst gocvsimd.View) {

	return filter(src, gocvsimd.SimdSse2MeanFilter3x3)
}
// MedianFilterRhomb3x3 performs median filtration of input image (filter window is a rhomb 3x3).
// All images must have the same width, height and format (8-bit gray, 16-bit UV, 24-bit BGR or 32-bit BGRA).
func MedianFilterRhomb3x3(src gocvsimd.View) (dst gocvsimd.View) {

	return filter(src, gocvsimd.SimdSse2MedianFilterRhomb3x3)
}
// MedianFilterSquare3x3 performs median filtration of input image (filter window is a square 3x3).
// All images must have the same width, height and format (8-bit gray, 16-bit UV, 24-bit BGR or 32-bit BGRA).
func MedianFilterSquare3x3(src gocvsimd.View) (dst gocvsimd.View) {

	return filter(src, gocvsimd.SimdSse2MedianFilterSquare3x3)
}
// MedianFilterRhomb5x5 performs median filtration of input image (filter window is a rhomb 5x5).
// All images must have the same width, height and format (8-bit gray, 16-bit UV, 24-bit BGR or 32-bit BGRA).
func MedianFilterRhomb5x5(src gocvsimd.View) (dst gocvsimd.View) {

	return filter(src, gocvsimd.SimdSse2MedianFilterRhomb5x5)
}
// MedianFilterSquare5x5 performs median filtration of input image (filter window is a square 5x5).
// All images must have the same width, height and format (8-bit gray, 16-bit UV, 24-bit BGR or 32-bit BGRA).
func MedianFilterSquare5x5(src gocvsimd.View) (dst gocvsimd.View) {

	return filter(src, gocvsimd.SimdSse2MedianFilterSquare5x5)
}
// Laplace calculates Laplace's filter.
// All images must have the same width and height. Input image must has 8-bit gray format, output image must has 16-bit integer format.
func Laplace(src gocvsimd.View) (dst gocvsimd.View) {

	return filter(src, gocvsimd.SimdSse2Laplace)
}

// SobelDx calculates Sobel's filter along x axis.
// All images must have the same width and height. Input image must has 8-bit gray format, output image must has 16-bit integer format.
// 	For every point:
//		n dst[x, y] = (src[x+1,y-1] + 2*src[x+1, y] + src[x+1, y+1]) - (src[x-1,y-1] + 2*src[x-1, y] + src[x-1, y+1]).
func SobelDx(src gocvsimd.View) (dst gocvsimd.View) {

	return filter(src, gocvsimd.SimdSse2SobelDx)
}

// SobelDy calculates Sobel's filter along y axis.
// All images must have the same width and height. Input image must has 8-bit gray format, output image must have a 16-bit integer format.
// 	For every point:
//		dst[x, y] = (src[x-1,y+1] + 2*src[x, y+1] + src[x+1, y+1]) - (src[x-1,y-1] + 2*src[x, y-1] + src[x+1, y-1]);
func SobelDy(src gocvsimd.View) (dst gocvsimd.View) {

	return filter(src, gocvsimd.SimdSse2SobelDy)
}

//
func ContourAnchors(src gocvsimd.View, step uint64, threshold int16, dst gocvsimd.View) {

	gocvsimd.SimdSse2ContourAnchors(src, step, uint64(threshold), dst)
}
