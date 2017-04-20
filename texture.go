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

// TextureBoostedSaturatedGradient calculates boosted saturated gradients for given input image.
// All images must have the same width, height and format (8-bit gray).
// 	For border pixels:
//		dx[x, y] = 0;
//		dy[x, y] = 0;
//	For other pixels:
//		dx[x, y] = (saturation + max(-saturation, min(saturation, (src[x + 1, y] - src[x - 1, y]))))*boost;
//		dy[x, y] = (saturation + max(-saturation, min(saturation, (src[x, y + 1] - src[x, y - 1]))))*boost;
func TextureBoostedSaturatedGradient(src gocvsimd.View, saturation, boost uint8, dx, dy gocvsimd.View) {

	gocvsimd.SimdSse2TextureBoostedSaturatedGradient(src, uint64(saturation), uint64(boost), dx, dy)
}

// TextureBoostedUv calculates boosted colorized texture feature of input image (actual for U and V components of YUV format).
// All images must have the same width, height and format (8-bit gray).
// 	For every pixel:
//		lo = 128 - (128/boost);
//		hi = 255 - lo;
//		dst[x, y] = max(lo, min(hi, src[i]))*boost;
func TextureBoostedUv(src gocvsimd.View, boost uint8, dst gocvsimd.View) {

	gocvsimd.SimdSse2TextureBoostedUv(src, uint64(boost), dst)
}

// TextureGetDifferenceSum calculates difference between current image and background.
// All images must have the same width, height and format (8-bit gray).
// 	For every pixel:
//		sum += current - average(lo[i], hi[i]);
func TextureGetDifferenceSum(src, lo, hi gocvsimd.View) int64 {

	return gocvsimd.SimdSse2TextureGetDifferenceSum(src, lo, hi)
}

// TexturePerformCompensation performs brightness compensation of input image.
// All images must have the same width, height and format (8-bit gray).
// 	For every pixel:
//		dst[i] = max(0, min(255, src[i] + shift));
func TexturePerformCompensation(src gocvsimd.View, shift int, dst gocvsimd.View) {

	gocvsimd.SimdSse2TexturePerformCompensation(src, shift, dst)
}