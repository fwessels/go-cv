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

//AbsDifferenceSum gets sum of absolute difference of two gray 8-bit images.
// Both images must have the same width and height.
func AbsDifferenceSum(a, b gocvsimd.View) uint64 {

	return gocvsimd.SimdSse2AbsDifferenceSum(a, b)
}
// AbsDifferenceSumMasked gets sum of absolute difference of two gray 8-bit images based on gray 8-bit mask.
// Gets the absolute difference sum for all points when mask[i] == index.
// Both images and mask must have the same width and height.
func AbsDifferenceSumMasked(a, b, mask gocvsimd.View, index uint8) uint64 {

	return gocvsimd.SimdSse2AbsDifferenceSumMasked(a, b, mask, uint64(index))
}
// AbsDifferenceSums3x3 gets 9 sums of absolute difference of two gray 8-bit images with various relative shifts in neighborhood 3x3.
// Both images must have the same width and height. The image height and width must be equal or greater 3.
// The sums are calculated with central part (indent width = 1) of the current image and with part of the background image with corresponding shift.
// The shifts are lain in the range [-1, 1] for axis x and y.
func AbsDifferenceSums3x3(current, background gocvsimd.View) [9]uint64 {

	return gocvsimd.SimdSse2AbsDifferenceSums3x3(current, background)
}
// AbsDifferenceSums3x3Masked gets 9 sums of absolute difference of two gray 8-bit images with various relative shifts in neighborhood 3x3 based on gray 8-bit mask.
// Gets the absolute difference sums for all points when mask[i] == index.
// Both images and mask must have the same width and height. The image height and width must be equal or greater 3.
// The sums are calculated with central part (indent width = 1) of the current image and with part of the background image with the corresponding shift.
// The shifts are lain in the range [-1, 1] for axis x and y.
func AbsDifferenceSums3x3Masked(current, background, mask gocvsimd.View, index uint8) [9]uint64 {

	return gocvsimd.SimdSse2AbsDifferenceSums3x3Masked(current, background, mask, uint64(index))
}

// AbsGradientSaturatedSum puts to destination 8-bit gray image saturated sum of absolute gradient for every point of source 8-bit gray image.
// Both images must have the same width and height.
// 	For border pixels:
// 		dst[x, y] = 0;
// 	For other pixels:
// 		dx = abs(src[x + 1, y] - src[x - 1, y]);
// 		dy = abs(src[x, y + 1] - src[x, y - 1]);
// 		dst[x, y] = min(dx + dy, 255);
func AbsGradientSaturatedSum(src, dst gocvsimd.View) {

	gocvsimd.SimdSse2AbsGradientSaturatedSum(src, dst)
}
// AddFeatureDifference adds feature difference to common difference in sum.
// All images must have the same width, height and format (8-bit gray).
// 	For every point:
// 		excess = max(lo[i] - value[i], 0) + max(value[i] - hi[i], 0);
// 		difference[i] += (weight * excess*excess) >> 16;
// This function is used for difference estimation in algorithm of motion detection.
func AddFeatureDifference(value, lo, hi gocvsimd.View, weight uint16, difference gocvsimd.View) {

	gocvsimd.SimdSse2AddFeatureDifference(value, lo, hi, uint64(weight), difference)
}
// SquaredDifferenceSum calculates sum of squared differences for two 8-bit gray images.
// All images must have the same width and height.
// 	For every point:
//		sum += (a[i] - b[i])*(a[i] - b[i]);
func SquaredDifferenceSum(a, b gocvsimd.View) uint64 {

	return gocvsimd.SimdSse2SquaredDifferenceSum(a, b)
}

// SquaredDifferenceSumMasked calculates sum of squared differences for two images with using mask.
// All images must have the same width, height and format (8-bit gray).
// 	For every point:
//		if(mask[i] == index)
//			sum += (a[i] - b[i])*(a[i] - b[i]);
func SquaredDifferenceSumMasked(a, b, mask gocvsimd.View, index uint8) uint64 {

	return gocvsimd.SimdSse2SquaredDifferenceSumMasked(a, b, mask, uint64(index))
}
