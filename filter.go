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
