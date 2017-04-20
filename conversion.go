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

// AlphaBlending performs alpha blending operation.
// All images must have the same width and height. Source and destination images must have the same format (8 bit per channel, for example GRAY8, BGR24 or BGRA32). Alpha must be 8-bit gray image.
// 	For every point:
// 		dst[i] = (src[i]*alpha[i] + dst[i]*(255 - alpha[i]))/255;
//This function is used for image drawing.
func AlphaBlending(src, alpha, dst gocvsimd.View) {

	gocvsimd.SimdSse2AlphaBlending(src, alpha, dst)
}
//BgraToGray converts 32-bit BGRA image to 8-bit gray image.
// All images must have the same width and height.
func BgraToGray(bgra gocvsimd.View) (gray gocvsimd.View) {

	gray = gocvsimd.View{}
	gray.Recreate(bgra.GetWidth(), bgra.GetHeight(), gocvsimd.GRAY8)

	gocvsimd.SimdSse2BgraToGray(bgra, gray)

	return gray
}

func bgratoyuv(bgra gocvsimd.View, f func(bgra, y, u, v gocvsimd.View)) (y, u, v gocvsimd.View) {

	y, u, v = gocvsimd.View{}, gocvsimd.View{}, gocvsimd.View{}
	y.Recreate(bgra.GetWidth(), bgra.GetHeight(), gocvsimd.GRAY8)
	u.Recreate(bgra.GetWidth(), bgra.GetHeight(), gocvsimd.GRAY8)
	v.Recreate(bgra.GetWidth(), bgra.GetHeight(), gocvsimd.GRAY8)

	gocvsimd.SimdSse2BgraToYuv420p(bgra, y, u, v)

	return y, u, v
}

// BgraToYuv420p converts 32-bit BGRA image to YUV420P.
// 	The input BGRA and output Y images must have the same width and height.
// The input U and V images must have the same width and height (half size relative to Y component).
func BgraToYuv420p(bgra gocvsimd.View) (y, u, v gocvsimd.View) {

	return bgratoyuv(bgra, gocvsimd.SimdSse2BgraToYuv420p)
}

// BgraToYuv422p converts 32-bit BGRA image to YUV422P.
// 	The input BGRA and output Y images must have the same width and height.
// The input U and V images must have the same width and height (their width is equal to half width of Y component).
func BgraToYuv422p(bgra gocvsimd.View) (y, u, v gocvsimd.View) {

	return bgratoyuv(bgra, gocvsimd.SimdSse2BgraToYuv422p)
}

// BgraToYuv444p converts 32-bit BGRA image to YUV444P.
// The input BGRA and output Y, U and V images must have the same width and height.
func BgraToYuv444p(bgra gocvsimd.View) (y, u, v gocvsimd.View) {

	return bgratoyuv(bgra, gocvsimd.SimdSse2BgraToYuv444p)
}
// GrayToBgra converts 8-bit gray image to 32-bit BGRA image.
// All images must have the same width and height.
func GrayToBgra(gray gocvsimd.View) (bgra gocvsimd.View) {

	bgra = gocvsimd.View{}
	bgra.Recreate(gray.GetWidth(), gray.GetHeight(), gocvsimd.BGRA32)

	gocvsimd.SimdSse2GrayToBgra(gray, bgra)

	return bgra
}

// DeinterleaveUv deinterleaves 16-bit UV interleaved image into separated 8-bit U and V planar images.
// All images must have the same width and height.
// This function used for NV12 to YUV420P conversion.
func DeinterleaveUv(uv gocvsimd.View) (u, v gocvsimd.View) {

	u, v = gocvsimd.View{}, gocvsimd.View{}
	u.Recreate(uv.GetWidth(), uv.GetHeight(), gocvsimd.GRAY8)
	v.Recreate(uv.GetWidth(), uv.GetHeight(), gocvsimd.GRAY8)

	gocvsimd.SimdSse2DeinterleaveUv(uv, u, v)

	return u, v
}
// FillBgr fills pixels data of 24-bit BGR image by given color(blue, green, red).
func FillBgr(dst gocvsimd.View, blue, green, red int) {

	gocvsimd.SimdSse2FillBgr(dst, blue, green, red)
}

// FillBgra fills pixels data of 32-bit BGRA image by given color(blue, green, red, alpha).
func FillBgra(dst gocvsimd.View, blue, green, red, alpha int) {

	gocvsimd.SimdSse2FillBgra(dst, blue, green, red, alpha)
}

// Int16ToGray converts 16-bit signed integer image to 8-bit gray image with saturation.
// All images must have the same width and height.
func Int16ToGray(src gocvsimd.View) (dst gocvsimd.View) {

	dst = gocvsimd.View{}
	dst.Recreate(dst.GetWidth(), dst.GetHeight(), gocvsimd.GRAY8)

	gocvsimd.SimdSse2Int16ToGray(src, dst)

	return dst
}

func yuvtobgra(y, u, v gocvsimd.View, alpha uint8, f func(y, u, v, bgra gocvsimd.View, alpha uint64)) (bgra gocvsimd.View) {

	bgra = gocvsimd.View{}
	bgra.Recreate(y.GetWidth(), y.GetHeight(), gocvsimd.BGRA32)

	f(y, u, v, bgra, uint64(alpha))

	return bgra
}

//Yuv420pToBgra converts YUV420P image to 32-bit BGRA image.
// 	The input Y and output BGRA images must have the same width and height.
// 	The input U and V images must have the same width and height (half size relative to Y component).
func Yuv420pToBgra(y, u, v gocvsimd.View, alpha uint8) (bgra gocvsimd.View) {

	return yuvtobgra(y, u, v, alpha, gocvsimd.SimdSse2Yuv420pToBgra)
}

// Yuv422pToBgra converts YUV422P image to 32-bit BGRA image.
// 	The input Y and output BGRA images must have the same width and height.
// 	The input U and V images must have the same width and height (their width is equal to half width of Y component).
func Yuv422pToBgra(y, u, v gocvsimd.View, alpha uint8) (bgra gocvsimd.View) {

	return yuvtobgra(y, u, v, alpha, gocvsimd.SimdSse2Yuv422pToBgra)
}
// Yuv444pToBgra converts YUV444P image to 32-bit BGRA image.
// 	The input Y, U, V and output BGRA images must have the same width and height.
func Yuv444pToBgra(y, u, v gocvsimd.View, alpha uint8) (bgra gocvsimd.View) {

	return yuvtobgra(y, u, v, alpha, gocvsimd.SimdSse2Yuv444pToBgra)
}

// TODO: Add Reorder16bit/32bit/64bit